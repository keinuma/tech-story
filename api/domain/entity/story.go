package entity

type Stories []*Story

type Story struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Owner *User  `json:"owner"`
}
