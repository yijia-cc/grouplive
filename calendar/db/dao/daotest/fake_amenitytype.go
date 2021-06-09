package daotest

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

var _ dao.AmenityType = (*FakeAmenityType)(nil)

type FakeAmenityType struct {
	amenityTypes []entity.AmenityType
}

func (f FakeAmenityType) GetAmenityType(transaction tx.Transaction, ID graphql.ID) (entity.AmenityType, error) {
	panic("implement me")
}

func (f FakeAmenityType) GetAllAmenityTypes(_ tx.Transaction) ([]entity.AmenityType, error) {
	return f.amenityTypes, nil
}

func NewFakeAmenityType(amenityTypesFixture []entity.AmenityType) FakeAmenityType {
	return FakeAmenityType{amenityTypes: amenityTypesFixture}
}
