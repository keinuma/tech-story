package presenter

import (
	"encoding/json"

	"github.com/keinuma/go-graphql/api/graph/model"
	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/service"
)

type StoryHandler interface {
	GetStories(limit, offset int) ([]*model.Story, error)
	CreateStory(input model.NewStory) (*model.Story, error)
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
	var storyModels []*model.Story
	bytes, err := json.Marshal(stories)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &storyModels)
	if err != nil {
		return nil, err
	}
	return storyModels, err
}

func (s *Story) CreateStory(input model.NewStory) (*model.Story, error) {
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
	var storyModel model.Story
	bytes, err := json.Marshal(story)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &storyModel)
	if err != nil {
		return nil, err
	}
	return &storyModel, err
}
