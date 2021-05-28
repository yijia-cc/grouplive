package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Amenity interface {
	FindAmenityInfo(transaction tx.Transaction, amenityTypeId entity.ID) ([]entity.AmenityInfo, error)
}

var _ Amenity = (*AmenitySQL)(nil)

type AmenitySQL struct {
	db *sql.DB
}

func (a AmenitySQL) FindAmenityInfo(transaction tx.Transaction, amenityTypeId entity.ID) ([]entity.AmenityInfo, error) {
	panic("implement me")
}

func NewAmenitySQL(db *sql.DB) AmenitySQL {
	return AmenitySQL{db: db}
}
