package dao

import (
	"encoding/json"
	"github.com/keinuma/go-graphql/api/domain/entity"
)

type User struct {
	ID          int     `json:"id"`
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (u *User) ToDAO(user entity.User) User {
	daoUser := User{
		UUID:        user.UUID,
		Name:        user.Name,
		Description: user.Description,
	}
	return daoUser
}

func (u *User) ToEntity() (*entity.User, error) {
	var entityUser entity.User
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
