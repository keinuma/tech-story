package gateway

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/infra/database/dao"
)

type Match struct {
	ctx context.Context
	tx  *gorm.DB
}

func NewMatch(ctx context.Context, tx *gorm.DB) *Match {
	return &Match{
		ctx: ctx,
		tx:  tx,
	}
}

func (s *Match) GetMatches(limit, offset int) ([]*model.Match, error) {
	var daoMatches dao.Matches
	if err := s.tx.Limit(limit).Offset(offset).
		Preload("Attendees").
		Preload("Story").
		Preload("Story.User").
		Find(&daoMatches).
		Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("[gateway.GetMatches] failed get matches")
	}
	entityMatches, err := daoMatches.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityMatches, nil
}

func (s *Match) CreateMatch(match model.Match) (*model.Match, error) {
	var daoMatch dao.Match
	daoMatch = daoMatch.ToDAO(match)
	if err := s.tx.Create(&daoMatch).Error; err != nil {
		return nil, errors.New("[gateway.CreateMatch] failed create match")
	}
	entityMatch, err := daoMatch.ToEntity()
	if err != nil {
		return nil, err
	}
	return entityMatch, nil
}
