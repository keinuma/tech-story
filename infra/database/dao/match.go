package dao

import (
	"encoding/json"
	"github.com/keinuma/tech-story/domain/model"
	"time"
)

type Matches []*Match

func (m *Matches) ToEntity() ([]*model.Match, error) {
	var entityMatches []*model.Match
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityMatches)
	if err != nil {
		return nil, err
	}
	return entityMatches, nil
}

type Match struct {
	ID        int       `json:"id"`
	StoryID   int       `gorm:"foreignKey:StoryId" json:"storyId"`
	Story     Story     `json:"story"`
	Date      time.Time `json:"date"`
	Attendees []*User   `gorm:"many2many:matches_users" json:"attendees"`
}

func (m *Match) ToDAO(match model.Match) Match {
	var daoUsers []*User
	for _, entityUser := range match.Attendees {
		daoUser := &User{
			ID: entityUser.ID,
		}
		daoUsers = append(daoUsers, daoUser)
	}
	daoMatch := Match{
		StoryID:   match.Story.ID,
		Date:      match.Date,
		Attendees: daoUsers,
	}
	return daoMatch
}

func (m *Match) ToEntity() (*model.Match, error) {
	var entityMatch model.Match
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &entityMatch)
	if err != nil {
		return nil, err
	}
	return &entityMatch, nil
}
