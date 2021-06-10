package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type TimeSlot struct{
	timeSlot entity.TimeSlot
}

func (t TimeSlot) ID() graphql.ID {
	return graphql.ID(t.timeSlot.ID)
}

func (t TimeSlot) Type() TimeSlotType {
	return timeSlotTypeMap[t.timeSlot.Type]
}

func (t TimeSlot) TimeRange() TimeRange {
	return newTimeRange(t.timeSlot.Range)
}

func newTimeSlot(timeSlot entity.TimeSlot) TimeSlot{
	return TimeSlot{
		timeSlot: timeSlot,
	}
}