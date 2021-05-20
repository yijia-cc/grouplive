package gqlapi

import "github.com/graph-gophers/graphql-go"

type Reservation struct {
}

func (Reservation) Id() graphql.ID {
	return ""
}

func (Reservation) Amenity() Amenity {
	return Amenity{}
}

func (Reservation) HoldDuration() *Duration {
	return nil
}

func (Reservation) TimeSlot() TimeSlot {
	return TimeSlot{}
}
