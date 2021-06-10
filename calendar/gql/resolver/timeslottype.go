package resolver

import "github.com/yijia-cc/grouplive/calendar/entity"

type TimeSlotType string

const (
	available      TimeSlotType = "AVAILABLE"
	bookedByOthers              = "BOOKED_BY_OTHERS"
	bookedByMe                  = "BOOKED_BY_ME"
)

var timeSlotTypeMap = map[entity.TimeSlotType]TimeSlotType{
	entity.Available:      available,
	entity.BookedByMe:     bookedByMe,
	entity.BookedByOthers: bookedByOthers,
}
