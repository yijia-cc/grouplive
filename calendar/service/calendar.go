package service

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/obs"
	"github.com/yijia-cc/grouplive/calendar/repo"
	"github.com/yijia-cc/grouplive/calendar/tx"
	"time"
)

type Calendar struct {
	logger             obs.Logger
	authorizer         auth.Authorizer
	transactionFactory tx.TransactionFactory
	amenityTypeRepo    repo.AmenityType
	scheduleRepo       repo.Schedule
}

func (c Calendar) ListAmenityTypes(user *entity.User) ([]entity.AmenityType, error) {
	//if !c.authorizer.HasPermission(user, permission.ViewAmenityTypes()) {
	//	return nil, errors.New("user is not allowed to view amenity types")
	//}

	transaction, err := c.transactionFactory.NewTransaction()
	if err != nil {
		return nil, err
	}
	return c.amenityTypeRepo.GetAllAmenityTypes(transaction)
}

func (c Calendar) GetAmenityType(user *entity.User, typeID graphql.ID) (entity.AmenityType, error) {
	//if !c.authorizer.HasPermission(user, permission.)
	transaction, err := c.transactionFactory.NewTransaction()
	if err != nil {
		return entity.AmenityType{}, err
	}

	return c.amenityTypeRepo.GetAmenityType(transaction, typeID)
}

func (c Calendar) GetWeekSchedule(user *entity.User, week *entity.Week) (entity.Schedule, error) {
	//if !c.authorizer.HasPermission(user, permission.)
	transaction, err := c.transactionFactory.NewTransaction()
	if err != nil {
		return entity.Schedule{}, err
	}

	if week == nil {
		year, weekNumber := time.Now().ISOWeek()
		week = &entity.Week{
			Year:   year,
			Number: weekNumber,
		}
	}

	week.StartDate = getWeekStartDate(week.Year, week.Number)
	return c.scheduleRepo.GetWeekSchedule(transaction, *week)
}

func (c Calendar) GetAmenitySchedule(amenity entity.Amenity, week *entity.Week) (entity.Schedule, error) {
	return entity.Schedule{}, nil
}

func getWeekStartDate(year int, weekNumber int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	for {
		date = date.AddDate(0, 0, 7)

		yNum, wNum := date.ISOWeek()
		if yNum < year || wNum < weekNumber {
			continue
		}
		return date.AddDate(0, 0, -1)
	}
}

func NewCalendar(
	logger obs.Logger,
	authorizer auth.Authorizer,
	transactionFactory tx.TransactionFactory,
	amenityDao dao.Amenity,
	amenityTypeDao dao.AmenityType,
	timeSlotDao dao.TimeSlot,
	reservationDao dao.Reservation,
	weekDao dao.Week,
) Calendar {
	return Calendar{
		logger:             logger,
		authorizer:         authorizer,
		transactionFactory: transactionFactory,
		amenityTypeRepo:    repo.NewAmenityType(amenityDao, amenityTypeDao),
		scheduleRepo:       repo.NewSchedule(timeSlotDao, reservationDao, weekDao),
	}
}
