package resolver

import "github.com/graph-gophers/graphql-go"

type AmenityFilterInput struct {
	TimeRange *TimeRangeInput
}

type TimeRangeInput struct {
	Start graphql.Time
	End graphql.Time
}

type ReservationInput struct {
	ReservationID *graphql.ID
	AmenityID     graphql.ID
	TimeRange     TimeRangeInput
}

type AmenityTypeInput struct {
	Id *graphql.ID
	Title *string
	Description *string
	ThumbnailUrl *string
}

type AmenityInput struct {
	Id               *graphql.ID
	Name             *string
	Type             AmenityTypeInput
	OperationalHours []TimeRangeInput
}
