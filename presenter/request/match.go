package request

import "time"

type NewMatch struct {
	StoryID   int       `json:"storyId"`
	Date      time.Time `json:"date"`
	Attendees []int     `json:"attendees"`
}
