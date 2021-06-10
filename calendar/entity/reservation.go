package entity

import (
	"time"
)

type Reservation struct {
	ID ID
	Amenity Amenity
	HoldDuration time.Duration
	TimeSlot TimeSlot
	Status ReservationStatus
	UserID ID
	WeekID ID
}

