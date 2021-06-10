package repo

import "github.com/yijia-cc/grouplive/calendar/db/dao"

type Reservation struct {
	reservationDao dao.Reservation
	amenityDao dao.Amenity
	timeSlotDao dao.TimeSlot
}

