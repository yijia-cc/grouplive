package resolver

type ReservationStatus string

const (
	upcoming ReservationStatus = "UPCOMING"
	ongoing                    = "ONGOING"
	past                       = "PAST"
)
