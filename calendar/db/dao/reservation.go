package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Reservation interface {
	GetAllReservations(transaction tx.Transaction) ([]entity.Reservation, error)
	GetWeekReservations(transaction tx.Transaction, week entity.Week) ([]entity.Reservation, error)
}

var _ Reservation = (*ReservationSQL) (nil)

type ReservationSQL struct {
	db *sql.DB
}

func (r ReservationSQL) GetAllReservations(transaction tx.Transaction) ([]entity.Reservation, error) {
	rows, err := transaction.DBTransaction.Query(`
SELECT id, amenity_id, time_slot_id, owner_id, week_id
FROM reservation ;
`)
	if err != nil {
		return nil, err
	}
	//rows, err := stmt.Query()
	//if err != nil {
	//	return nil, err
	//}

	reservations := make([]entity.Reservation, 0)

	for rows.Next() {
		reservation := entity.Reservation{}
		err = rows.Scan(
			&reservation.ID,
			&reservation.Amenity.ID,
			&reservation.TimeSlot.ID,
			&reservation.UserID, &reservation.WeekID)
		if err != nil {
			continue
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func (r ReservationSQL) GetWeekReservations(transaction tx.Transaction, week entity.Week) ([]entity.Reservation, error) {
	stmt, err := transaction.DBTransaction.Prepare(`
SELECT id, amenity_id, time_slot_id, owner_id, week_id
FROM reservation 
WHERE week_id = ?;
`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(week.ID)
	if err != nil {
		return nil, err
	}

	reservations := make([]entity.Reservation, 0)

	for rows.Next() {
		reservation := entity.Reservation{}
		err = rows.Scan(
			&reservation.ID,
			&reservation.Amenity.ID,
			&reservation.TimeSlot.ID,
			&reservation.UserID,
			&reservation.WeekID)
		if err != nil {
			continue
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func NewReservationSQL(db *sql.DB) ReservationSQL {
	return ReservationSQL{db: db}
}


