package model

type Users []*User

type User struct {
	ID          int     `json:"id"`
	UID         string  `json:"uid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
