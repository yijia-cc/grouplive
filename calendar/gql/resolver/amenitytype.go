package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/entity"
)

type AmenityType struct {
	amenityType entity.AmenityType
}

func (a AmenityType) AmenityInfoList() []AmenityInfo {
	gqlInfo := make([]AmenityInfo, 0)
	for _, info := range a.amenityType.AmenityInfoList {
		gqlInfo = append(gqlInfo, newAmenityInfo(info))
	}
	return gqlInfo
}

func (a AmenityType) Id() graphql.ID {
	return graphql.ID(a.amenityType.ID)
}

func (a AmenityType) Title() *string {
	return &a.amenityType.Title
}

func (a AmenityType) Description() *string {
	return &a.amenityType.Description
}

func (a AmenityType) ThumbnailUrl() *string {
	return &a.amenityType.ThumbnailURL
}

func newAmenityType(amenityType entity.AmenityType) AmenityType {
	return AmenityType{
		amenityType: amenityType,
	}
}
