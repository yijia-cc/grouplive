package resolver

import "github.com/graph-gophers/graphql-go"

type Schedule struct{}

func (Schedule) Id() graphql.ID {
	return "id"
}

func (Schedule) WeekStart() graphql.Time {
	return graphql.Time{}
}

func (Schedule) Previous() *Schedule {
	return nil
}

func (Schedule) Next() *Schedule {
	return nil
}

func (Schedule) TimeSlots() []TimeSlot {
	return nil
}

func (Schedule) Reservations() []Reservation {
	return nil
}
