package dao

import (
	"encoding/json"
	"github.com/keinuma/tech-story/domain/model"
)

type Users []*User

func (u *Users) ToEntity() (model.Users, error) {
	var entityUsers model.Users
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityUsers)
	if err != nil {
		return nil, err
	}
	return entityUsers, nil
}

type User struct {
	ID          int     `json:"id"`
	UID         string  `json:"uid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (u *User) ToDAO(user model.User) User {
	daoUser := User{
		UID:         user.UID,
		Name:        user.Name,
		Description: user.Description,
	}
	return daoUser
}

func (u *User) ToEntity() (*model.User, error) {
	var entityUser model.User
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityUser)
	if err != nil {
		return nil, err
	}
	return &entityUser, nil
}
