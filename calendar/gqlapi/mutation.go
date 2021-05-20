package gqlapi

import "github.com/graph-gophers/graphql-go"

type mutation struct {
}

func (mutation) MakeReservation(args struct{
	Reservation ReservationInput
}) Reservation {
	return Reservation{}
}

func (mutation) ConfirmReservation(args struct{
	Id graphql.ID
}) *Void {
	return nil
}

func (mutation) CancelReservation(args struct{
	Id graphql.ID
}) *Void {
	return nil
}

func (mutation) UpdateReservation(args struct{
	Reservation ReservationInput
}) Reservation {
	return Reservation{}
}

func (mutation) AddAmenityType(args struct{
	AmenityType AmenityTypeInput
}) *Void {
	return nil
}

func (mutation) DeleteAmenityType(args struct{
	Id graphql.ID
}) *Void {
	return nil
}

func (mutation) UpdateAmenityType(args struct{
	AmenityType AmenityTypeInput
}) *Void {
	return nil
}

func (mutation) AddAmenity(args struct{
	Amenity AmenityInput
}) *Void {
	return nil
}

func (mutation) DeleteAmenity(args struct{
	Id graphql.ID
}) *Void {
	return nil
}

func (mutation) UpdateAmenity(args struct{
	Amenity AmenityInput
}) *Void {
	return nil
}