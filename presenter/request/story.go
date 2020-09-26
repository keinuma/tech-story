package request

type NewStory struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}
