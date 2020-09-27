package repository

import "github.com/keinuma/tech-story/domain/model"

type UserRepository interface {
	CreateUser(user model.User) (*model.User, error)
	GetUsersByIDs(userIDs []int) (model.Users, error)
}
