package dao

import (
	"database/sql"
	"github.com/graph-gophers/graphql-go"

	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type AmenityType interface {
	GetAllAmenityTypes(transaction tx.Transaction) ([]entity.AmenityType, error)
	GetAmenityType(transaction tx.Transaction, ID graphql.ID) (entity.AmenityType, error)
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

func (a AmenityTypeSQL) GetAmenityType(transaction tx.Transaction, ID graphql.ID) (entity.AmenityType, error) {
	rows, err := transaction.DBTransaction.Query(`
SELECT id, title, description, thumbnail_url 
FROM amenity_type
WHERE id = ?;
`, ID)

	amenityType := entity.AmenityType{}
	for rows.Next() {
		err = rows.Scan(&amenityType.ID, &amenityType.Title, &amenityType.Description, &amenityType.ThumbnailURL)
		if err != nil {
			continue
		}
	}

	return amenityType, nil
}



func NewAmenityTypeSQL(db *sql.DB) AmenityTypeSQL {
	return AmenityTypeSQL{db: db}
}
