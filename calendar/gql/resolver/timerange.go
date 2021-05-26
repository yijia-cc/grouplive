package resolver

import "github.com/graph-gophers/graphql-go"

type TimeRange struct {
	start graphql.Time
	end   graphql.Time
}

func (TimeRange) Start() graphql.Time {
	return graphql.Time{}
}

func (TimeRange) End() graphql.Time {
	return graphql.Time{}
}
