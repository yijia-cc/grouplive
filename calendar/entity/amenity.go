package entity

type Amenity struct {
	ID ID
	Name string
	Type AmenityType
	OperationalHours []TimeRange
	Schedule Schedule
}
