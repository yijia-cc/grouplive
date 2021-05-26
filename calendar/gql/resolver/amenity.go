package resolver

import "github.com/graph-gophers/graphql-go"

type Amenity struct{}

func (Amenity) Schedule(args struct {
	WeekStart *graphql.Time
}) *Schedule {
	return nil
}

func (Amenity) Id() graphql.ID {
	return "id"
}

func (Amenity) Name() *string {
	return nil
}

func (Amenity) Type() AmenityType {
	return AmenityType{}
}

func (Amenity) OperationalHours() []TimeRange {
	return nil
}
