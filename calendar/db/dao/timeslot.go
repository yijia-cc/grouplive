package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type TimeSlot interface {
	GetWeekTimeSlots(transaction tx.Transaction, week entity.Week) ([]entity.TimeSlot, error)
}

var _ TimeSlot = (*TimeSlotSQL) (nil)

type TimeSlotSQL struct {
	db *sql.DB
}

func (t TimeSlotSQL) GetWeekTimeSlots(transaction tx.Transaction, week entity.Week) ([]entity.TimeSlot, error) {
	rows, err := transaction.DBTransaction.Query(`
SELECT id, type_id, start, end, week_id
FROM time_slot 
WHERE week_id = ?;
`, week.ID)
	if err != nil {
		return nil, err
	}

	timeSlots := make([]entity.TimeSlot, 0)
	for rows.Next() {
		timeSlot := entity.TimeSlot{}
		err = rows.Scan(&timeSlot.ID, &timeSlot.Type, &timeSlot.Range.Start, &timeSlot.Range.End, &timeSlot.WeekId)
		if err != nil {
			return nil, err
		}
		timeSlots = append(timeSlots, timeSlot)
	}
	return timeSlots, nil
}

func NewTimeSlotSQL(db *sql.DB) TimeSlotSQL {
	return TimeSlotSQL{db: db}
}