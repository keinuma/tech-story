package dao

import (
	"encoding/json"
	"github.com/keinuma/go-graphql/api/domain/entity"
)

type Stories []*Story

func (s *Stories) ToEntity() (entity.Stories, error) {
	var entityStories entity.Stories
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityStories)
	if err != nil {
		return nil, err
	}
	return entityStories, nil
}

type Story struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	User  User   `gorm:"foreignKey:User" json:"user"`
}

func (s *Story) ToDAO(story entity.Story) Story {
	daoStory := Story{
		Title: story.Title,
		User: User{
			ID: story.User.ID,
		},
	}
	return daoStory
}

func (s *Story) ToEntity() (*entity.Story, error) {
	var entityStory entity.Story
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityStory)
	if err != nil {
		return nil, err
	}
	return &entityStory, nil
}
