package presenter

import (
	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/service"
	"github.com/keinuma/tech-story/presenter/request"
)

type MatchHandler interface {
	GetMatches(limit, offset int) ([]*model.Match, error)
	CreateMatch() (*model.Match, error)
}

type Match struct {
	matchService service.Match
}

func NewMatch(matchService service.Match) *Match {
	return &Match{
		matchService: matchService,
	}
}

func (m *Match) GetMatches(limit, offset int) ([]*model.Match, error) {
	matches, err := m.matchService.GetMatches(limit, offset)
	if err != nil {
		return nil, err
	}
	return matches, err
}

func (m *Match) CreateMatch(input request.NewMatch) (*model.Match, error) {
	var usersEntity model.Users
	for _, userId := range input.Attendees {
		usersEntity = append(usersEntity, &model.User{
			ID: userId,
		})
	}
	matchEntity := model.Match{
		Story: model.Story{
			ID: input.StoryID,
		},
		Date:      input.Date,
		Attendees: usersEntity,
	}
	match, err := m.matchService.CreateMatch(matchEntity)
	if err != nil {
		return nil, err
	}
	return match, err
}
