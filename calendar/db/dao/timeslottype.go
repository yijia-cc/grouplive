package dao

import "database/sql"

type TimeSlotType interface {

}

var _ TimeSlotType = (*TimeSlotTypeSQL)(nil)

type TimeSlotTypeSQL struct {
	db *sql.DB
}
