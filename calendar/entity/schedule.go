package entity

type Schedule struct {
	WeekID ID
	Reservations []Reservation
	TimeSlots []TimeSlot
	PreviousWeekID ID
	NextWeekID ID
}

