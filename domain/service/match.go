package service

import (
	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/repository"
)

type MatchService interface {
	GetMatches(limit, offset int) ([]*model.Match, error)
	CreateMatch(match model.Match) (*model.Match, error)
}

type Match struct {
	matchRepository repository.MatchRepository
}

func NewMatch(matchRepository repository.MatchRepository) *Match {
	return &Match{
		matchRepository: matchRepository,
	}
}

func (s *Match) GetMatches(limit, offset int) ([]*model.Match, error) {
	matches, err := s.matchRepository.GetMatches(limit, offset)
	if err != nil {
		return nil, err
	}
	return matches, err
}

func (s *Match) CreateMatch(input model.Match) (*model.Match, error) {
	match, err := s.matchRepository.CreateMatch(input)
	if err != nil {
		return nil, err
	}
	return match, err
}
