package model

import (
	"time"
)

type UserReaction struct {
	ReactionId int       `json:"reaction_id"`
	UserName   string    `json:"user_name"`
	EventId    int       `json:"event_id"`
	Attend     bool      `json:"attend"`
	Comment    string    `json:"comment,omitempty"`
	LastUpdate time.Time `json:"last_update"`
}
