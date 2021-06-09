package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/service"
)

type query struct {
	calendarService service.Calendar
}

func (q query) AmenityTypes(ctx context.Context) ([]AmenityType, error) {
	//user, err := auth.UserFromContext(ctx)
	//if err != nil {
	//	return nil, errors.New("user is unauthorized")
	//}

	//types, err := q.calendarService.ListAmenityTypes(user)
	types, err := q.calendarService.ListAmenityTypes(nil)
	if err != nil {
		return nil, err
	}
	gqlAmenityTypes := make([]AmenityType, 0)
	for _, amenityType := range types {
		gqlAmenityTypes = append(gqlAmenityTypes, newAmenityType(amenityType))
	}
	return gqlAmenityTypes, nil
}

func (q query) AmenityType(ctx context.Context, args struct{graphql.ID}) (AmenityType, error) {
	//user, err := auth.UserFromContext(ctx)
	//if err != nil {
	//	return AmenityType{}, errors.New("user is unauthorized")
	//}

	amenityType, err := q.calendarService.GetAmenityType(nil, args.ID)
	if err != nil {
		return AmenityType{}, err
	}
	gqlAmenityType := newAmenityType(amenityType)
	return gqlAmenityType, nil
}

func (query) MyCalendar(ctx context.Context, args struct {
	WeekStart graphql.Time
}) Schedule {
	return Schedule{}
}

func (q query) Reservations() []Reservation {
	return nil
}

func newQuery(calendarService service.Calendar) query {
	return query{calendarService: calendarService}
}
