package repository

import "github.com/keinuma/go-graphql/api/domain/entity"

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
}
