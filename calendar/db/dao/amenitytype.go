package dao

import (
	"database/sql"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type AmenityType interface {
	GetAllAmenityTypes(transaction tx.Transaction) ([]entity.AmenityType, error)
}

var _ AmenityType = (*AmenityTypeSQL)(nil)

type AmenityTypeSQL struct {
	db *sql.DB
}

func (a AmenityTypeSQL) GetAllAmenityTypes(transaction tx.Transaction) ([]entity.AmenityType, error) {
	panic("implement me")
}

func NewAmenityTypeSQL(db *sql.DB) AmenityTypeSQL {
	return AmenityTypeSQL{db: db}
}
