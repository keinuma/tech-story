package repository

import "github.com/keinuma/go-graphql/api/domain/entity"

type StoryRepository interface {
	GetStories(limit, offset int) (*entity.Stories, error)
	CreateStory(story entity.Story) (*entity.Story, error)
}
