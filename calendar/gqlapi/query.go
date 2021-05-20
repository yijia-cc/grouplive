package gqlapi

import "github.com/graph-gophers/graphql-go"

type query struct {
}

func (query) AmenityTypes() []AmenityType {
	return nil
}

func (query) AmenityType(args struct {
	ID graphql.ID
}) AmenityType {
	return AmenityType{}
}

func (query) Reservations() []Reservation {
	return nil
}
