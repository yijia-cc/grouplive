package daotest

import (
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

var _ dao.Amenity = (*FakeAmenity)(nil)

type FakeAmenity struct {
	amenityInfos []entity.AmenityInfo
}

func (f FakeAmenity) FindAmenityInfo(_ tx.Transaction, amenityTypeId entity.ID) ([]entity.AmenityInfo, error) {
	amenityInfos := make([]entity.AmenityInfo, 0)
	for _, amenityInfo := range f.amenityInfos {
		if amenityInfo.AmenityTypeID == amenityTypeId {
			amenityInfos = append(amenityInfos, amenityInfo)
		}
	}
	return amenityInfos, nil
}

func NewFakeAmenity(amenityInfosFixture []entity.AmenityInfo) FakeAmenity  {
	return FakeAmenity{amenityInfos: amenityInfosFixture}
}