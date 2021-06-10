package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
	"time"
)

type FindWeekQuery struct {
	year int
	week int
}

type Week interface {
	FindWeek(transaction tx.Transaction, weekStart time.Time) (entity.Week, error)
}

var _ Week = (*WeekSQL)(nil)

type WeekSQL struct {
	db *sql.DB
}

func (w WeekSQL) FindWeek(transaction tx.Transaction, weekStart time.Time) (entity.Week, error) {
	row := transaction.DBTransaction.QueryRow(`
SELECT id, week_start 
FROM week
WHERE week_start = ?;
`, weekStart)

	week := entity.Week{}
	err := row.Scan(&week.ID, &week.StartDate)
	return week, err
}

func NewWeekSQL(db *sql.DB) WeekSQL {
	return WeekSQL{db: db}
}

