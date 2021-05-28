package repo

import (
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type AmenityType struct {
	amenityDao     dao.Amenity
	amenityTypeDao dao.AmenityType
}

func (a AmenityType) GetAllAmenityTypes(transaction tx.Transaction) ([]entity.AmenityType, error) {
	types, err := a.amenityTypeDao.GetAllAmenityTypes(transaction)
	if err != nil {
		return nil, err
	}

	allAmenityTypes := make([]entity.AmenityType, 0)
	for _, amenityType := range types {
		infoList, err := a.amenityDao.FindAmenityInfo(transaction, amenityType.ID)
		if err != nil {
			return allAmenityTypes, err
		}
		amenityType.AmenityInfoList = infoList
		allAmenityTypes = append(allAmenityTypes, amenityType)
	}
	return allAmenityTypes, nil
}

func NewAmenityType(amenityDao dao.Amenity, amenityTypeDao dao.AmenityType) AmenityType {
	return AmenityType{
		amenityDao:     amenityDao,
		amenityTypeDao: amenityTypeDao,
	}
}
