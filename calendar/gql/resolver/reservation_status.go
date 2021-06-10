package resolver

import "github.com/yijia-cc/grouplive/calendar/entity"

type ReservationStatus string

const (
	upcoming ReservationStatus = "UPCOMING"
	ongoing                    = "ONGOING"
	past                       = "PAST"
)

var reservationStatusMap = map[entity.ReservationStatus]ReservationStatus {
	entity.Upcoming: upcoming,
	entity.Ongoing: ongoing,
	entity.Past: past,
}
