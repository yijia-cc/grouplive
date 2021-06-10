package entity

type ReservationStatus int

const (
	Upcoming ReservationStatus = iota
	Ongoing
	Past
)
