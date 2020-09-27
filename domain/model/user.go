package model

type Users []*User

type User struct {
	ID          int     `json:"id"`
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
