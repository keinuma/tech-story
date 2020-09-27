package service

import (
	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/repository"
)

type UserService interface {
	CreateUser(story model.User) (*model.User, error)
	GetUsersByIDs(userIDs []int) ([]*model.User, error)
	GetUserByUID(uid string) (*model.User, error)
}

type User struct {
	storyRepository repository.UserRepository
}

func NewUser(storyRepository repository.UserRepository) *User {
	return &User{
		storyRepository: storyRepository,
	}
}

func (s *User) CreateUser(input model.User) (*model.User, error) {
	story, err := s.storyRepository.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return story, err
}

func (s *User) GetUsersByIDs(userIDs []int) ([]*model.User, error) {
	users, err := s.storyRepository.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *User) GetUsersByUID(uid string) (*model.User, error) {
	user, err := s.storyRepository.GetUserByUID(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
