package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/service"
)

type query struct {
	calendarService service.Calendar
}

func (q query) AmenityTypes(ctx context.Context) ([]AmenityType, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	types, err := q.calendarService.ListAmenityTypes(user)
	if err != nil {
		return nil, err
	}
	gqlAmenityTypes := make([]AmenityType, 0)
	for _, amenityType := range types {
		gqlAmenityTypes = append(gqlAmenityTypes, newAmenityType(amenityType))
	}
	return gqlAmenityTypes, nil
}

func (query) AmenityType(args struct {
	ID graphql.ID
}) AmenityType {
	return AmenityType{}
}

func (query) MyCalendar() Schedule {
	return Schedule{}
}

func (query) Reservations() []Reservation {
	return nil
}

func newQuery(calendarService service.Calendar) query {
	return query{calendarService: calendarService}
}
