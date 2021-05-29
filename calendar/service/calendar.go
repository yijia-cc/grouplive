package service

import (
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/repo"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Calendar struct {
	transactionFactory tx.TransactionFactory
	amenityTypeRepo    repo.AmenityType
}

func (c Calendar) ListAmenityTypes() ([]entity.AmenityType, error) {
	transaction, err := c.transactionFactory.NewTransaction()
	if err != nil {
		return nil, err
	}
	return c.amenityTypeRepo.GetAllAmenityTypes(transaction)
}

func NewCalendar(transactionFactory tx.TransactionFactory, amenityDao dao.Amenity, amenityTypeDao dao.AmenityType) Calendar {
	return Calendar{
		transactionFactory: transactionFactory,
		amenityTypeRepo: repo.NewAmenityType(amenityDao, amenityTypeDao),
	}
}
