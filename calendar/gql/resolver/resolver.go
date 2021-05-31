package resolver

import (
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/service"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

type Resolver struct {
	query
	mutation
	subscription
}

func NewResolver(
	authorizer auth.Authorizer,
	transactionFactory tx.TransactionFactory,
	amenityDao dao.Amenity,
	amenityTypeDao dao.AmenityType,
) *Resolver {

	calendarService := service.NewCalendar(authorizer, transactionFactory, amenityDao, amenityTypeDao)
	return &Resolver{
		query: newQuery(calendarService),
	}
}
