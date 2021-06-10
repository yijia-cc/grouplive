package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type Amenity struct {
	amenity         entity.Amenity
}

func (a Amenity) ID() graphql.ID {
	return graphql.ID(a.amenity.ID)
}

func (a Amenity) Name() *string {
	return &a.amenity.Name
}

func (a Amenity) Type() AmenityType {
	return newAmenityType(a.amenity.Type)
}

func (a Amenity) OperationalHours() []TimeRange {
	hours := make([]TimeRange, 0)
	for _, timeRange := range a.amenity.OperationalHours {
		hours = append(hours, newTimeRange(timeRange))
	}
	return hours
}

func newAmenity(amenity entity.Amenity) Amenity {
	return Amenity{
		amenity: amenity,
	}
}
