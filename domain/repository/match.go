package repository

import "github.com/keinuma/tech-story/domain/model"

type MatchRepository interface {
	GetMatches(limit, offset int) ([]*model.Match, error)
	CreateMatch(story model.Match) (*model.Match, error)
}
