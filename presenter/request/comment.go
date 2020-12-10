package request

import "time"

type NewComment struct {
	MatchID  int       `json:"matchID"`
	Text     string    `json:"text"`
	PostedAt time.Time `json:"postedAt"`
}
