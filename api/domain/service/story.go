package service

import (
	"github.com/keinuma/go-graphql/api/domain/entity"
	"github.com/keinuma/go-graphql/api/domain/repository"
)

type StoryService interface {
	GetStories(limit, offset int) (*entity.Stories, error)
	CreateStory(story entity.Story) (*entity.Story, error)
}

type Story struct {
	storyRepository repository.StoryRepository
}

func NewStory(storyRepository repository.StoryRepository) *Story {
	return &Story{
		storyRepository: storyRepository,
	}
}

func (s *Story) GetStories(limit, offset int) (*entity.Stories, error) {
	stories, err := s.storyRepository.GetStories(limit, offset)
	if err != nil {
		return nil, err
	}
	return stories, err
}

func (s *Story) CreateStory(input entity.Story) (*entity.Story, error) {
	story, err := s.storyRepository.CreateStory(input)
	if err != nil {
		return nil, err
	}
	return story, err
}
