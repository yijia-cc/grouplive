package entity

import "time"

type Media struct {
	Id        int64     `json:"id"`
	Event     *Event    `json:"event,omitempty"`
	MediaName string    `json:"media_name"`
	MediaURL  string    `json:"media_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    bool      `json:"active"`
}
