package util

import (
	"fmt"
	"github.com/pborman/uuid"
	"github.com/yijia-cc/grouplive/dashboard/config"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func ParseMultipartRequest(r *http.Request) (*entity.Event, []*entity.Media, int, error) {
	if r.MultipartForm == nil {
		err := r.ParseMultipartForm(config.Cfg.App.FileUploadSizeLimit)
		if err != nil {
			return nil, nil, http.StatusBadRequest, err
		}
	}

	event, status, err := parseEvent(r.MultipartForm.Value)
	if err != nil {
		return nil, nil, status, err
	}

	mediaList, status, err := uploadFiles(r.MultipartForm.File)
	if err != nil {
		return nil, nil, status, err
	}

	return event, mediaList, 0, nil
}

func uploadFiles(files map[string][]*multipart.FileHeader) ([]*entity.Media, int, error) {
	var mediaList []*entity.Media

	for _, fhs := range files {
		for _, fh := range fhs {
			f, err := fh.Open()
			if err != nil {
				return nil, http.StatusBadRequest, fmt.Errorf("media file is invalid or not available: %v", err)
			}

			if _, err := os.Stat(config.Cfg.App.StaticMediaDir); os.IsNotExist(err) {
				err := os.MkdirAll(config.Cfg.App.StaticMediaDir, 0755)
				if err != nil {
					return nil, http.StatusInternalServerError, fmt.Errorf("unable to create server folder: %v", err)
				}
			}

			relativePath := filepath.Join(config.Cfg.App.StaticMediaDir, fmt.Sprintf("%s%s", uuid.New(), filepath.Ext(fh.Filename)))
			out, err := os.Create(relativePath)
			if err != nil {
				return nil, http.StatusInternalServerError, fmt.Errorf("unable to create server file: %v", err)
			}
			defer out.Close()

			if _, err := io.Copy(out, f); err != nil {
				return nil, http.StatusInternalServerError, fmt.Errorf("unable to copy file to server folder: %v", err)
			}

			media := entity.Media{
				MediaURL: relativePath,
				MediaName: fh.Filename,
			}
			mediaList = append(mediaList, &media)
		}
	}

	return mediaList, 0, nil
}

func parseEvent(m map[string][]string) (*entity.Event, int, error) {
	var err error
	var event entity.Event
	for key, values := range m {
		switch key {
		case "id":
			event.Id, err = strconv.ParseInt(values[0], 10, 64)
		case "type_id":
			event.Type = &entity.Type{}
			event.Type.Id, err = strconv.ParseInt(values[0], 10, 64)
		/*case "username":
			event.User = &entity.User{Username: values[0]}*/
		case "title":
			event.Title = values[0]
		case "description":
			event.Description = values[0]
		case "start_time":
			event.StartTime, err = time.Parse(config.Cfg.App.LocalDatetimeFormat, values[0])
		case "end_time":
			event.EndTime, err = time.Parse(config.Cfg.App.LocalDatetimeFormat, values[0])
		case "rsvp_required":
			event.RsvpRequired = values[0] == "true"
		/*case "active":
			event.Active = values[0] == "true"*/
		}

		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	}

	return &event, 0, nil
}


