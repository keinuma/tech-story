package repository

import "github.com/keinuma/tech-story/domain/model"

type StoryRepository interface {
	GetStories(limit, offset int) (*model.Stories, error)
	CreateStory(story model.Story) (*model.Story, error)
}
