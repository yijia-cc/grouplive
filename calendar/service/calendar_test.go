package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yijia-cc/grouplive/calendar/auth/authtest"
	"github.com/yijia-cc/grouplive/calendar/db/dao/daotest"
	"github.com/yijia-cc/grouplive/calendar/entity"
	"github.com/yijia-cc/grouplive/calendar/tx/txtest"
)

func TestCalendar_ListAmenityTypes(t *testing.T) {
	poolInfo := entity.AmenityInfo{
		ID:            "302",
		Name:          "Sky pool",
		AmenityTypeID: "002",
	}
	dreamGymInfo := entity.AmenityInfo{
		ID:            "100",
		Name:          "Dream Gym",
		AmenityTypeID: "001",
	}
	miniGymInfo := entity.AmenityInfo{
		ID:            "200",
		Name:          "Mini Gym",
		AmenityTypeID: "001",
	}
	gymInfoList := []entity.AmenityInfo{
		dreamGymInfo,
		miniGymInfo,
	}
	testCases := []struct {
		name                 string
		amenityTypesFixture  []entity.AmenityType
		amenityInfosFixture  []entity.AmenityInfo
		inputUser            *entity.User
		expectedHasError     bool
		expectedAmenityTypes []entity.AmenityType
	}{
		{
			name: "one amenity type",
			amenityTypesFixture: []entity.AmenityType{
				{
					ID:           "001",
					Title:        "Gym",
					Description:  "Indoor gym 24/7",
					ThumbnailURL: "https://www.google.com",
				},
			},
			amenityInfosFixture: []entity.AmenityInfo{
				dreamGymInfo,
				miniGymInfo,
			},
			inputUser:        &entity.User{ID: "tester1"},
			expectedHasError: false,
			expectedAmenityTypes: []entity.AmenityType{
				{
					ID:              "001",
					Title:           "Gym",
					Description:     "Indoor gym 24/7",
					ThumbnailURL:    "https://www.google.com",
					AmenityInfoList: gymInfoList,
				},
			},
		},
		{
			name: "multiple amenity types",
			amenityTypesFixture: []entity.AmenityType{
				{
					ID:           "001",
					Title:        "Gym",
					Description:  "Indoor gym 24/7",
					ThumbnailURL: "https://www.google.com",
				},
				{
					ID:           "002",
					Title:        "Pool",
					Description:  "Pool with sunshine",
					ThumbnailURL: "https://www.yahoo.com",
				},
			},
			amenityInfosFixture: []entity.AmenityInfo{
				dreamGymInfo,
				miniGymInfo,
				poolInfo,
			},
			inputUser:        &entity.User{ID: "tester2"},
			expectedHasError: false,
			expectedAmenityTypes: []entity.AmenityType{
				{
					ID:              "001",
					Title:           "Gym",
					Description:     "Indoor gym 24/7",
					ThumbnailURL:    "https://www.google.com",
					AmenityInfoList: gymInfoList,
				},
				{
					ID:           "002",
					Title:        "Pool",
					Description:  "Pool with sunshine",
					ThumbnailURL: "https://www.yahoo.com",
					AmenityInfoList: []entity.AmenityInfo{
						poolInfo,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			fakeAmenityType := daotest.NewFakeAmenityType(testCase.amenityTypesFixture)
			fakeAmenity := daotest.NewFakeAmenity(testCase.amenityInfosFixture)
			fakeTransactionFactory := txtest.NewFakeTransactionFactory()
			stubAuthorizer := authtest.NewStubAuthorizer()
			calendarService := NewCalendar(stubAuthorizer, fakeTransactionFactory, fakeAmenity, fakeAmenityType)

			actual, err := calendarService.ListAmenityTypes(testCase.inputUser)
			if testCase.expectedHasError {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, testCase.expectedAmenityTypes, actual)
		})
	}
}
