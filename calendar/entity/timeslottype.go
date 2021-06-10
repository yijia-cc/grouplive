package entity

type TimeSlotType int

const (
	Available TimeSlotType = iota
	BookedByOthers
	BookedByMe
)
