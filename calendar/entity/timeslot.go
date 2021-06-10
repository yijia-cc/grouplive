package entity

type TimeSlot struct {
	ID ID
	Type TimeSlotType
	Range TimeRange
	WeekId ID
}