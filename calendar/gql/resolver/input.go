package resolver

import "github.com/graph-gophers/graphql-go"

type AmenityFilterInput struct {
	TimeRange *TimeRangeInput
}

type TimeRangeInput struct {
	Start graphql.Time
	End   graphql.Time
}

type ReservationInput struct {
	ReservationID *graphql.ID
	AmenityID     graphql.ID
	TimeRange     TimeRangeInput
}

type AmenityTypeInput struct {
	ID           *graphql.ID
	Title        *string
	Description  *string
	ThumbnailUrl *string
}

type AmenityInput struct {
	ID              *graphql.ID
	Name             *string
	Type             AmenityTypeInput
	OperationalHours []TimeRangeInput
}

type ScheduleUpdateSubscribeInput struct {
	AmenityID graphql.ID
	WeekID graphql.ID
}
