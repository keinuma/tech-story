package model

type Stories []*Story

type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	User  User   `json:"user"`
}

type NewStory struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}
