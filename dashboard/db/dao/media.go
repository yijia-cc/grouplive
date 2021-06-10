package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/dashboard/entity"
)

type MediaDao interface {
	CreatMediaTx(tx *sql.Tx, media *entity.Media) (int64, error)
	//UpdateMediaTx(tx *sql.Tx, media *entity.Media) (int64, error)
	DeleteByEidTx(tx *sql.Tx, eid int64) (int64, error)
}

type MediaDaoImpl struct{
	db *sql.DB
}

func (MediaDaoImpl) DeleteByEidTx(tx *sql.Tx, eid int64) (int64, error) {
	sql := "UPDATE media SET active = FALSE WHERE event_id = ?"

	res, err := tx.Exec(sql, eid)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func (m MediaDaoImpl) CreatMediaTx(tx *sql.Tx, media *entity.Media) (int64, error) {
	sql := "INSERT INTO media (event_id, media_name, media_url, created_at, updated_at) values(?,?,?,?,?)"
	res, err := tx.Exec(sql, media.Event.Id, media.MediaName, media.MediaURL, media.CreatedAt, media.UpdatedAt)
	if err != nil {
		return -1, err
	}

	mediaId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return mediaId, err
}

func NewMediaDao(db *sql.DB) MediaDao {
	return MediaDaoImpl{db}
}


