package service

import (
	"errors"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/auth/permission"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/repo"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Calendar struct {
	authorizer auth.Authorizer
	transactionFactory tx.TransactionFactory
	amenityTypeRepo    repo.AmenityType
}

func (c Calendar) ListAmenityTypes(user *entity.User) ([]entity.AmenityType, error) {
	if !c.authorizer.HasPermission(user, permission.ViewAmenityTypes, nil) {
		return nil, errors.New("user is not allowed to view amenity types")
	}

	transaction, err := c.transactionFactory.NewTransaction()
	if err != nil {
		return nil, err
	}
	return c.amenityTypeRepo.GetAllAmenityTypes(transaction)
}

func NewCalendar(
	authorizer auth.Authorizer,
	transactionFactory tx.TransactionFactory,
	amenityDao dao.Amenity,
	amenityTypeDao dao.AmenityType,
	) Calendar {
	return Calendar{
		authorizer: authorizer,
		transactionFactory: transactionFactory,
		amenityTypeRepo: repo.NewAmenityType(amenityDao, amenityTypeDao),
	}
}

