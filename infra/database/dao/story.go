package dao

import (
	"encoding/json"
	"github.com/keinuma/tech-story/domain/model"
)

type Stories []*Story

func (s *Stories) ToEntity() (model.Stories, error) {
	var entityStories model.Stories
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

func (s *Story) ToDAO(story model.Story) Story {
	daoStory := Story{
		Title: story.Title,
		User: User{
			ID: story.User.ID,
		},
	}
	return daoStory
}

func (s *Story) ToEntity() (*model.Story, error) {
	var entityStory model.Story
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
