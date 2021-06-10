package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type Schedule struct{
	schedule entity.Schedule
}

func (s Schedule) WeekID() graphql.ID {
	return graphql.ID(s.schedule.WeekID)
}

func (s Schedule) PreviousWeekID() *graphql.ID {
	return (*graphql.ID)(&s.schedule.PreviousWeekID)
}

func (s Schedule) NextWeekID() *graphql.ID {
	return (*graphql.ID)(&s.schedule.NextWeekID)
}

func (s Schedule) TimeSlots() []TimeSlot {
	timeSlots := make([]TimeSlot, 0)
	for _, timeSlot := range s.schedule.TimeSlots {
		timeSlots = append(timeSlots, newTimeSlot(timeSlot))
	}
	return timeSlots
}

func (s Schedule) Reservations() []Reservation {
	reservations := make([]Reservation, 0)
	for _, reservation := range s.schedule.Reservations {
		reservations = append(reservations, newReservation(reservation))
	}
	return reservations
}

func newSchedule(schedule entity.Schedule) Schedule {
	return Schedule{
		schedule: schedule,
	}
}
