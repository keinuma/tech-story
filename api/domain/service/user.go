package service

import (
	"github.com/keinuma/go-graphql/api/domain/entity"
	"github.com/keinuma/go-graphql/api/domain/repository"
)

type UserService interface {
	CreateUser(story entity.User) (*entity.User, error)
}

type User struct {
	storyRepository repository.UserRepository
}

func NewUser(storyRepository repository.UserRepository) *User {
	return &User{
		storyRepository: storyRepository,
	}
}

func (s *User) CreateUser(input entity.User) (*entity.User, error) {
	story, err := s.storyRepository.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return story, err
}
