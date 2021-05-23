package resolver

import "github.com/graph-gophers/graphql-go"

type AmenityType struct {}

func (AmenityType) Amenities(args struct {
	Filter *AmenityFilterInput
}) []Amenity {
	return nil
}

func (AmenityType) Id() graphql.ID {
	return "id"
}

func (AmenityType) Title() *string {
	return nil
}

func (AmenityType) Description() *string {
	return nil
}

func (AmenityType) ThumbnailUrl() *string {
	return nil
}
