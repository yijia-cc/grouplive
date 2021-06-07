package model

import (
	"time"
)

type EventCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Event struct {
	EventId      int       `json:"event_id"`
	CategoryId   int       `json:"category_id"`
	UserName     int       `json:"user_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	RsvpRequired bool      `json:"rsvp_required"`
	CreatedAt    time.Time `json:"created_at"`
}
