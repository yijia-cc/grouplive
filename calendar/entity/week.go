package entity

import "time"

type Week struct {
	ID ID
	StartDate time.Time
	Year int
	Number int
}
