package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type TimeRange struct {
	timeRange entity.TimeRange
}

func (t TimeRange) Start() graphql.Time {
	return graphql.Time{Time: t.timeRange.Start}
}

func (t TimeRange) End() graphql.Time {
	return graphql.Time{Time: t.timeRange.End}
}

func newTimeRange(timeRange entity.TimeRange) TimeRange {
	return TimeRange{
		timeRange: timeRange,
	}
}
