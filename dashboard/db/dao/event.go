package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/config"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"time"
)

type EventDao interface {
	CreatEventTx(tx *sql.Tx, event *entity.Event) (int64, error)
	UpdateEventTx(tx *sql.Tx, event *entity.Event) (int64, error)
	DeleteByEidTx(tx *sql.Tx, id int64) (int64, error)
	ReadDashboard() (map[string][]*entity.Event, error)
	Search(searchType entity.SearchType, searchKeys *entity.SearchKey) ([]byte, error)
	MetaMapping() ([]byte, error)
}

type EventDaoImpl struct {
	db *sql.DB
}

func (e EventDaoImpl) MetaMapping() ([]byte, error) {
	query := `SELECT category_id, category.name AS category_name, type.id AS type_id, type.name AS type_name 
              FROM type join category on category_id = category.id
              ORDER BY category_id, type_id`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type Meta struct {
		CategoryId   int	`json:"category_id"`
		CategoryName string `json:"category_name"`
		TypeId       int	`json:"type-id"`
		TypeName     string	`json:"type-name"`
	}

	var metaList []*Meta

	for rows.Next() {
		meta := Meta{}
		if err := rows.Scan(&meta.CategoryId, &meta.CategoryName, &meta.TypeId, &meta.TypeName); err != nil {
			return nil, err
		}
		metaList = append(metaList, &meta)
	}

	return json.Marshal(metaList)
}

func (e EventDaoImpl) Search(searchType entity.SearchType, searchKeys *entity.SearchKey) ([]byte, error) {
	query := e.buildQuery(searchType, searchKeys)

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events interface{}
	if searchType == entity.Mixed {
		events = e.scanMixed(rows)
	} else {
		events = e.scanGrouped(rows)
	}
	js, err := e.toJson(events)

	return js, err
}


func (e EventDaoImpl) buildQuery(searchType entity.SearchType, searchKeys *entity.SearchKey) string {
	sql := `SELECT e.id AS eid, username, title, description, start_time, end_time, rsvp_required, 
			e.created_at AS e_created_at, e.updated_at AS e_updated_at, e.active AS e_active, 
			m.id AS mid, media_name, media_url, m.created_at AS m_created_at, m.updated_at AS m_updated_at,
			m.active AS m_active, t.id AS tid, t.name AS t_name, c.id AS cid, c.name AS c_name 
			FROM event e
			    JOIN media m ON e.id = m.event_id
				JOIN type t ON e.type_id = t.id
				JOIN category c on t.category_id = c.id 
			WHERE e.active = TRUE AND m.active = TRUE
			ORDER BY start_time`

	if searchType == entity.Dashboard {
		return fmt.Sprintf("%s AND start_time >= '%s'", sql, time.Now().Format(config.Cfg.App.LocalDatetimeFormat))
	}

	if searchKeys.EventId > 0 {
		sql = fmt.Sprintf("%s AND e.id = %d", sql, searchKeys.EventId)
	}
	if searchKeys.TypeId > 0 {
		sql = fmt.Sprintf("%s AND t.id = %d", sql, searchKeys.TypeId)
	}
	if searchKeys.CategoryId > 0 {
		sql = fmt.Sprintf("%s AND c.id = %d", sql, searchKeys.CategoryId)
	}
	if searchKeys.UserName != "" {
		sql = fmt.Sprintf("%s AND username = '%s'", sql, searchKeys.UserName)
	}
	if searchKeys.Title != "" {
		sql = fmt.Sprintf("%s AND title LIKE '%s%s%s'", sql, "%", searchKeys.Title, "%")
	}
	if searchKeys.StartTime != (time.Time{}) {
		fmt.Println("searchKeys.StartTime: ", searchKeys.StartTime)
		sql = fmt.Sprintf("%s AND start_time >= '%s'", sql, searchKeys.StartTime.Format(config.Cfg.App.LocalDatetimeFormat))
	}
	if searchKeys.EndTime != (time.Time{}) {
		fmt.Println("searchKeys.EndTime: ", searchKeys.EndTime)
		sql = fmt.Sprintf("%s AND end_time <= '%s'", sql, searchKeys.EndTime.Format(config.Cfg.App.LocalDatetimeFormat))
	}
	// "TRUE", "FALSE", "", where "" means both ture and false rsvp are searched
	if searchKeys.RsvpRequired != "" {
		sql = fmt.Sprintf("%s AND rsvp_required = %s", sql, searchKeys.RsvpRequired)
	}

	return sql
}

func (e EventDaoImpl) toJson(events interface{}) ([]byte, error) {
	var js []byte
	var err error

	switch evts := events.(type) {
	case []*entity.Event:
		js, err = json.Marshal(evts)
	case map[string][]*entity.Event:
		js, err = json.Marshal(evts)
	default:
		err = fmt.Errorf("invliad events type")
	}

	return js, err
}

func (e EventDaoImpl) scanMixed(rows *sql.Rows) interface{} {
	//events := map[string][]*entity.Event{}
	events := []*entity.Event{}

	var lastEvent *entity.Event
	for rows.Next() {
		m := &entity.Media{}
		u := &entity.User{}
		c := &entity.Category{}
		t := &entity.Type{
			Category: c,
		}
		e := &entity.Event{
			Type: t,
			User: u,
			MediaList: []*entity.Media{},
		}

		rows.Scan(&e.Id, &e.User.Username, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.RsvpRequired,
			&e.CreatedAt, &e.UpdatedAt, &e.Active, &m.Id, &m.MediaName, &m.MediaURL, &m.CreatedAt, &m.UpdatedAt,
			&m.Active, &t.Id, &t.Name, &c.Id, &c.Name)

		/*if _, ok := events[c.Name]; !ok {
			events[c.Name] = []*entity.Event{}
		}*/

		// Encountered a new Event
		if lastEvent == nil || lastEvent.Id != e.Id {
			//events[c.Name] = append(events[c.Name], e)
			events = append(events, e)
			lastEvent = e
		}

		lastEvent.MediaList = append(lastEvent.MediaList, m)
	}

	return events
}

func (e EventDaoImpl) scanGrouped(rows *sql.Rows) interface{} {
	events := map[string][]*entity.Event{}

	var lastEvent *entity.Event
	for rows.Next() {
		m := &entity.Media{}
		u := &entity.User{}
		c := &entity.Category{}
		t := &entity.Type{
			Category: c,
		}
		e := &entity.Event{
			Type: t,
			User: u,
			MediaList: []*entity.Media{},
		}

		rows.Scan(&e.Id, &e.User.Username, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.RsvpRequired,
			&e.CreatedAt, &e.UpdatedAt, &e.Active, &m.Id, &m.MediaName, &m.MediaURL, &m.CreatedAt, &m.UpdatedAt,
			&m.Active, &t.Id, &t.Name, &c.Id, &c.Name)

		/*fmt.Println(e.ID, e.User.Username, e.Title, e.Description, e.StartTime, e.EndTime, e.RsvpRequired,
		e.CreatedAt, e.UpdatedAt, e.Active, m.ID, m.MediaName, m.MediaURL, m.CreatedAt, m.UpdatedAt,
		m.Active, t.ID, t.Name, c.ID, c.Name)*/


		if _, ok := events[c.Name]; !ok {
			events[c.Name] = []*entity.Event{}
		}

		// Encountered a new Event
		if lastEvent == nil || lastEvent.Id != e.Id {
			events[c.Name] = append(events[c.Name], e)
			lastEvent = e
		}

		lastEvent.MediaList = append(lastEvent.MediaList, m)
	}

	return events
}

func (e EventDaoImpl) ReadDashboard() (map[string][]*entity.Event, error) {
	sql := `SELECT e.id AS eid, username, title, description, start_time, end_time, rsvp_required, 
			e.created_at AS e_created_at, e.updated_at AS e_updated_at, e.active AS e_active, 
			m.id AS mid, media_name, media_url, m.created_at AS m_created_at, m.updated_at AS m_updated_at,
			m.active AS m_active, t.id AS tid, t.name AS t_name, c.id AS cid, c.name AS c_name 
			FROM event e
			    JOIN media m ON e.id = m.event_id
				JOIN type t ON e.type_id = t.id
				JOIN category c on t.category_id = c.id 
			WHERE e.active = TRUE and m.active = TRUE and e.start_time >= ?
			`
	stmt, err := db.DashDB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	now := time.Now()
	localNow := now.Format(config.Cfg.App.LocalDatetimeFormat)    // "yyyy-MM-dd hh:mm:ss"
	rows, err := stmt.Query(localNow)
	//rows, err := stmt.Query(now)   // Invalid!!
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := map[string][]*entity.Event{}
	var lastEvent *entity.Event
	for rows.Next() {
		m := &entity.Media{}
		u := &entity.User{}
		c := &entity.Category{}
		t := &entity.Type{
			Category: c,
		}
		e := &entity.Event{
			Type: t,
			User: u,
			MediaList: []*entity.Media{},
		}

		rows.Scan(&e.Id, &e.User.Username, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.RsvpRequired,
			&e.CreatedAt, &e.UpdatedAt, &e.Active, &m.Id, &m.MediaName, &m.MediaURL, &m.CreatedAt, &m.UpdatedAt,
			&m.Active, &t.Id, &t.Name, &c.Id, &c.Name)

		fmt.Println(e.Id, e.User.Username, e.Title, e.Description, e.StartTime, e.EndTime, e.RsvpRequired,
		e.CreatedAt, e.UpdatedAt, e.Active, m.Id, m.MediaName, m.MediaURL, m.CreatedAt, m.UpdatedAt,
		m.Active, t.Id, t.Name, c.Id, c.Name)


		if _, ok := events[c.Name]; !ok {
			events[c.Name] = []*entity.Event{}
		}

		// Encountered a new Event
		if lastEvent == nil || lastEvent.Id != e.Id {
			events[c.Name] = append(events[c.Name], e)
			lastEvent = e
		}

		lastEvent.MediaList = append(lastEvent.MediaList, m)
	}

	return events, nil
}


func (e EventDaoImpl) DeleteByEidTx(tx *sql.Tx, id int64) (int64, error) {
	sql := "UPDATE event SET active = FALSE, updated_at = ? WHERE id = ?"

	res, err := tx.Exec(sql, time.Now(), id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func (e EventDaoImpl) CreatEventTx(tx *sql.Tx, event *entity.Event) (int64, error) {
	sql := `INSERT INTO 
    		event(type_id, username, title, description, start_time, end_time, rsvp_required, created_at, updated_at) 
			values(?,?,?,?,?,?,?,?,?)`

	res, err := tx.Exec(
		sql,
		event.Type.Id,
		event.User.Username,
		event.Title,
		event.Description,
		event.StartTime,
		event.EndTime,
		event.RsvpRequired,
		event.CreatedAt,
		event.UpdatedAt,
	)

	if err != nil {
		return -1, err
	}

	eventId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return eventId, err
}

func (e EventDaoImpl) UpdateEventTx(tx *sql.Tx, event *entity.Event) (int64, error) {
	sql := `UPDATE event
			SET type_id=?, username=?, title=?, description=?, start_time=?, end_time=?, rsvp_required=?, updated_at=?
			WHERE id=?`

	res, err := tx.Exec(
		sql,
		event.Type.Id,
		event.User.Username,
		event.Title,
		event.Description,
		event.StartTime,
		event.EndTime,
		event.RsvpRequired,
		event.UpdatedAt,
		event.Id,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func NewEventDao(db *sql.DB) EventDao {
	return EventDaoImpl{db}
}

