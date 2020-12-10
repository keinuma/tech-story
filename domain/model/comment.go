package model

import "time"

type Comment struct {
	ID       int       `json:"id"`
	Match    Match     `json:"match"`
	Text     string    `json:"text"`
	PostedAt time.Time `json:"postedAt"`
}
