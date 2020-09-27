package presenter

import (
	"github.com/keinuma/tech-story/presenter/request"

	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/service"
)

type StoryHandler interface {
	GetStories(limit, offset int) ([]*model.Story, error)
	CreateStory(input request.NewStory) (*model.Story, error)
}

type Story struct {
	storyService service.Story
}

func NewStory(storyService service.Story) *Story {
	return &Story{
		storyService: storyService,
	}
}

func (s *Story) GetStories(limit, offset int) ([]*model.Story, error) {
	stories, err := s.storyService.GetStories(limit, offset)
	if err != nil {
		return nil, err
	}
	return *stories, err
}

func (s *Story) CreateStory(input request.NewStory) (*model.Story, error) {
	storyEntity := model.Story{
		Title: input.Text,
		User: &model.User{
			ID: input.UserID,
		},
	}
	story, err := s.storyService.CreateStory(storyEntity)
	if err != nil {
		return nil, err
	}
	return story, err
}
