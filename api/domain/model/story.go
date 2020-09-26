package model

type Stories []*Story

type Story struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	User  *User  `json:"user"`
}
