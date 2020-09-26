package gateway

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/keinuma/go-graphql/api/domain/model"
	"github.com/keinuma/go-graphql/api/infra/database/dao"
)

type Story struct {
	ctx context.Context
	tx  *gorm.DB
}

func NewStory(ctx context.Context, tx *gorm.DB) *Story {
	return &Story{
		ctx: ctx,
		tx:  tx,
	}
}

func (s *Story) GetStories(limit, offset int) (*model.Stories, error) {
	var daoStories dao.Stories
	if err := s.tx.Limit(limit).Offset(offset).Find(&daoStories).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("[gateway.GetStories] failed get stories")
	}
	entityStory, err := daoStories.ToEntity()
	if err != nil {
		return nil, err
	}
	return &entityStory, nil
}

func (s *Story) CreateStory(story model.Story) (*model.Story, error) {
	var daoStory dao.Story
	daoStory = daoStory.ToDAO(story)
	if err := s.tx.Create(&daoStory).Error; err != nil {
		return nil, errors.New("[gateway.CreateStory] failed create story")
	}
	entityStory, err := daoStory.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityStory, nil
}
