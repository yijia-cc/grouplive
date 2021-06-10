package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/service"
)

type Reservation struct {
	reservation entity.Reservation
	calendarService service.Calendar
}

func (r Reservation) ID() graphql.ID {
	return graphql.ID(r.reservation.ID)
}

func (r Reservation) Amenity() Amenity {
	return newAmenity(r.reservation.Amenity)
}

func (r Reservation) HoldDuration() *Duration {
	return &Duration{duration: r.reservation.HoldDuration}
}

func (r Reservation) TimeSlot() TimeSlot {
	return newTimeSlot(r.reservation.TimeSlot)
}

func (r Reservation) WeekID() graphql.ID {
	return graphql.ID(r.reservation.WeekID)
}


func (r Reservation) Status() ReservationStatus {
	return reservationStatusMap[r.reservation.Status]
}

func newReservation(reservation entity.Reservation) Reservation {
	return Reservation {
		reservation: reservation,
	}
}