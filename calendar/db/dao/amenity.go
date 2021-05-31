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
	stmt, err := transaction.DBTransaction.Prepare(
		`
SELECT id, name, type_id 
FROM amenity 
WHERE type_id = "?";
`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(amenityTypeId)
	if err != nil {
		return nil, err
	}

	amenityInfos := make([]entity.AmenityInfo, 0)
	for rows.Next() {
		amenityInfo := entity.AmenityInfo{}
		err = rows.Scan(&amenityInfo.ID, &amenityInfo.Name, &amenityInfo.AmenityTypeID)
		if err != nil {
			continue
		}
		amenityInfos = append(amenityInfos, amenityInfo)
	}
	return amenityInfos, nil
}

func NewAmenitySQL(db *sql.DB) AmenitySQL {
	return AmenitySQL{db: db}
}
