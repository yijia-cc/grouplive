package resolver

type TimeSlot struct{}

func (TimeSlot) Type() TimeSlotType {
	return available
}

func (TimeSlot) TimeRange() TimeRange {
	return TimeRange{}
}
