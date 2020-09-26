package repository

import "github.com/keinuma/go-graphql/api/domain/model"

type UserRepository interface {
	CreateUser(user model.User) (*model.User, error)
}
