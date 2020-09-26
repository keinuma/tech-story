package service

import (
	"github.com/keinuma/go-graphql/api/domain/model"
	"github.com/keinuma/go-graphql/api/domain/repository"
)

type StoryService interface {
	GetStories(limit, offset int) (*model.Stories, error)
	CreateStory(story model.Story) (*model.Story, error)
}

type Story struct {
	storyRepository repository.StoryRepository
}

func NewStory(storyRepository repository.StoryRepository) *Story {
	return &Story{
		storyRepository: storyRepository,
	}
}

func (s *Story) GetStories(limit, offset int) (*model.Stories, error) {
	stories, err := s.storyRepository.GetStories(limit, offset)
	if err != nil {
		return nil, err
	}
	return stories, err
}

func (s *Story) CreateStory(input model.Story) (*model.Story, error) {
	story, err := s.storyRepository.CreateStory(input)
	if err != nil {
		return nil, err
	}
	return story, err
}
