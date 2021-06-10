package entity

import (
	"time"
)

type Reaction struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	EventId   int64     `json:"event_id"`
	Attend    bool      `json:"attend"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}