package entity

import (
	"time"
)

type Event struct {
	Id           int64     `json:"id"`
	Type         *Type     `json:"type"`
	User         *User     `json:"user"`
	Title        string    `json:"title"`
	MediaList    []*Media  `json:"media_list"`
	Description  string    `json:"description"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	RsvpRequired bool      `json:"rsvp_required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Active       bool      `json:"active"`
}

type SearchKey struct {
	EventId      int64     `json:"event_id"`
	TypeId       int64     `json:"type_id"`
	CategoryId   int64     `json:"category_id"`
	UserName     string    `json:"username"`
	Title        string    `json:"title"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	RsvpRequired string    `json:"rsvp_required"` // "TRUE", "FALSE", "", where "" means both ture and false rsvp are searched
}

type SearchType int
const (
	Mixed SearchType = iota + 1
	Grouped
	Dashboard
	UserInfo
	UserConfirmation
)