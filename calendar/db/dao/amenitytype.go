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
	rows, err := transaction.DBTransaction.Query(`
SELECT id, title, description, thumbnail_url 
FROM amenity_type;
`)
	if err != nil {
		return nil, err
	}

	amenityTypes := make([]entity.AmenityType, 0)

	for rows.Next() {
		amenityType := entity.AmenityType{}
		err = rows.Scan(&amenityType.ID, &amenityType.Title, &amenityType.Description, &amenityType.ThumbnailURL)
		if err != nil {
			continue
		}
		amenityTypes = append(amenityTypes, amenityType)
	}
	return amenityTypes, nil
}

func NewAmenityTypeSQL(db *sql.DB) AmenityTypeSQL {
	return AmenityTypeSQL{db: db}
}
