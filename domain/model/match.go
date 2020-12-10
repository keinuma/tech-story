package model

import "time"

type Match struct {
	ID        int        `json:"id"`
	Story     Story      `json:"story"`
	Date      time.Time  `json:"date"`
	Attendees Users      `json:"attendees"`
	Comments  []*Comment `json:"comments"`
}
