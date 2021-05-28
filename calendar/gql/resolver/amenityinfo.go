package resolver

import "github.com/yijia-cc/grouplive/calendar/entity"

type AmenityInfo struct {
	amenityInfo entity.AmenityInfo
}

func newAmenityInfo(info entity.AmenityInfo) AmenityInfo {
	return AmenityInfo{
		amenityInfo: info,
	}
}
