package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type AmenityInfo struct {
	amenityInfo entity.AmenityInfo
}

func (a AmenityInfo) ID() graphql.ID {
	return graphql.ID(a.amenityInfo.ID)
}

func (a AmenityInfo) Name() *string {
	return &a.amenityInfo.Name
}

func (a AmenityInfo) AmenityTypeID() graphql.ID {
	return graphql.ID(a.amenityInfo.AmenityTypeID)
}

func newAmenityInfo(info entity.AmenityInfo) AmenityInfo {
	return AmenityInfo{
		amenityInfo: info,
	}
}
