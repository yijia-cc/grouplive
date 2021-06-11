package dao

import (
	"database/sql"
	"encoding/json"
	"github.com/yijia-cc/grouplive/dashboard/db"
	"github.com/yijia-cc/grouplive/dashboard/entity"
	"time"
)

type ReactionDao interface {
	Creat(re *entity.Reaction) (int64, error)
	Update(re *entity.Reaction) (int64, error)
	ToggleById(rid int64) (int64, error)
	GetReactionId(username string, eid int64) (int64, error)
	GetReaction(username string, eid int64) (*entity.Reaction, error)
	GetConfirmedEvtTIDs(username string) ([]int64, error)
	GetConfirmations(username string) ([]byte, error)
	DeleteByEidTx(tx *sql.Tx, eid int64) (int64, error)
}


type ReactionDaoImpl struct {
	db *sql.DB
}

func (r ReactionDaoImpl) Creat(re *entity.Reaction) (int64, error) {
	query := `INSERT INTO 
    		  reaction(username, event_id, attend, created_at, updated_at, active) 
			  values(?, ?, ?, ?, ?, ?)`

	res, err := r.db.Exec(query, re.Username, re.EventId, re.Attend, re.CreatedAt, re.UpdatedAt, re.Active)
	if err != nil {
		return -1, err
	}

	rid, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return rid, err
}

func (r ReactionDaoImpl) Update(re *entity.Reaction) (int64, error) {
	query := `UPDATE reaction
			  SET attend = ?, updated_at = ?, active = ?
			  WHERE id = ?`

	res, err := r.db.Exec(query, re.Attend, re.UpdatedAt, re.Active, re.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r ReactionDaoImpl) ToggleById(rid int64) (int64, error) {
	query := `UPDATE reaction
			  SET attend = !attend, updated_at = ?
			  WHERE id = ?`

	res, err := r.db.Exec(query, time.Now(), rid)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r ReactionDaoImpl) GetReaction(username string, eid int64) (*entity.Reaction, error) {
	query := `SELECT id, attend, created_at, updated_at, active 
    		  FROM reaction WHERE username = ? and event_id = ? AND active = TRUE`

	re := entity.Reaction{
		Username: username,
		EventId:  eid,
	}

	err := r.db.QueryRow(query, username, eid).Scan(&re.Id, &re.Attend, &re.CreatedAt, &re.UpdatedAt, &re.Active)
	if err != nil {
		return nil, err
	}
	return &re, nil
}

func (r ReactionDaoImpl) GetReactionId(username string, eid int64) (int64, error) {
	var rid int64
	query := `SELECT id FROM reaction WHERE username = ? AND event_id = ? AND active = TRUE`

	// If no row was found by QueryRow, Scan() returns an ErrNoRows error;
	err := r.db.QueryRow(query, username, eid).Scan(&rid)
	if err == sql.ErrNoRows {
		return -1, nil
	}
	if err != nil {
		return -1, err
	}
	return rid, nil
}


// GetConfirmation check whether the given user has confirmed or unconfirmed the given event
func (r ReactionDaoImpl) GetConfirmation(username string, eid int64) (bool, error) {
	var confirmed bool
	query := `SELECT attend FROM reaction WHERE username = ? AND event_id = ? AND active = TRUE`

	// If no row was found by QueryRow, Scan() returns an ErrNoRows error;
	err := r.db.QueryRow(query, username, eid).Scan(&confirmed)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return confirmed, nil
}

// GetConfirmations return all events confirmed by a given user
func (r ReactionDaoImpl) GetConfirmations(username string) ([]byte, error) {
	type confirmation struct {
		ReactionId  int64     `json:"reaction_id"`
		EventPoster string    `json:"event_poster"`
		EventId     int64     `json:"event_id"`
		TypeId      int64     `json:"type_id"`
		Attend      bool      `json:"attend"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	query := `SELECT r.id, e.username, e.id, e.type_id, r.attend, r.created_at, r.updated_at
			  FROM reaction r JOIN event e ON r.event_id = e.id
			  WHERE r.username = ? AND r.attend = TRUE AND r.active = TRUE`

	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cnfms []*confirmation
	for rows.Next() {
		cnfm := confirmation{}
		err := rows.Scan(&cnfm.ReactionId, &cnfm.EventPoster, &cnfm.EventId, &cnfm.TypeId, &cnfm.Attend, &cnfm.CreatedAt, &cnfm.UpdatedAt);
		if err != nil {
			return nil, err
		}
		cnfms = append(cnfms, &cnfm)
	}

	return json.Marshal(cnfms)
}


// GetConfirmedEvtTIDs return all confirmed events' type ids for a given user
func (ReactionDaoImpl) GetConfirmedEvtTIDs(username string) ([]int64, error) {
	query := `SELECT e.type_Id
			  FROM reaction r 
			  	JOIN event e ON r.event_id = e.id
			  WHERE r.username = ? AND r.attend = TRUE AND r.active = TRUE`

	rows, err := db.DashDB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tids []int64
	for rows.Next() {
		var tid int64
		if err := rows.Scan(&tid); err != nil {
			return nil, err
		}
		tids = append(tids, tid)
	}

	return tids, nil
}


func (ReactionDaoImpl) DeleteByEidTx(tx *sql.Tx, eid int64) (int64, error) {
	query := "UPDATE reaction SET active = FALSE, updated_at = ? WHERE event_id = ?"

	res, err := tx.Exec(query, time.Now(), eid)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func NewReactionDao(db *sql.DB) ReactionDao {
	return ReactionDaoImpl{db}
}
