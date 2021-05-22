package resolver

type TimeSlotType string

const(
	available      TimeSlotType = "AVAILABLE"
	bookedByOthers              = "BOOKED_BY_OTHERS"
	bookedByMe                  = "BOOKED_BY_ME"
)
