package store

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/infra/store/channel"
)

type Subscriber struct {
	Store         Store
	MatchChannels map[string]chan *model.Match
	Mutex         sync.Mutex
}

func NewSubscriber(ctx context.Context, store Store) *Subscriber {
	subscriber := &Subscriber{
		Store:         store,
		MatchChannels: map[string]chan *model.Match{},
		Mutex:         sync.Mutex{},
	}
	subscriber.StartMatchChannel(ctx)
	return subscriber
}

func (s *Subscriber) StartMatchChannel(ctx context.Context) {
	go func() {
		pubsub := s.Store.Client.Subscribe(ctx, channel.MatchChannel)
		defer pubsub.Close()

		for {
			msgInterface, _ := pubsub.Receive(ctx)
			switch msg := msgInterface.(type) {
			case *redis.Message:
				m := model.Match{}
				if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
					continue
				}
				s.Mutex.Lock()
				for _, ch := range s.MatchChannels {
					ch <- &m
				}
				s.Mutex.Unlock()
			default:
			}
		}
	}()
}

func (s *Subscriber) Publish(ctx context.Context, match *model.Match) error {
	userUID := match.Story.User.UID
	result, err := s.Store.Client.SetXX(ctx, userUID, userUID, 60*time.Minute).Result()
	if result == false {
		return err
	}

	matchJSON, err := json.Marshal(match)
	if err != nil {
		return err
	}
	s.Store.Client.Publish(ctx, channel.MatchChannel, matchJSON)
	return nil
}

func (s *Subscriber) Receive(ctx context.Context, userUID string) <-chan *model.Match {
	match := make(chan *model.Match, 1)
	s.Mutex.Lock()
	s.MatchChannels[userUID] = match
	s.Mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.Mutex.Lock()
		delete(s.MatchChannels, userUID)
		s.Mutex.Unlock()
		s.Store.Client.Del(ctx, userUID)
	}()

	return match
}

func (s *Subscriber) SetUser(ctx context.Context, userUID string) error {
	result, err := s.Store.Client.SetNX(ctx, userUID, userUID, 60*time.Minute).Result()
	if err != nil {
		return err
	}
	if result == false {
		return errors.New("This User name has already used")
	}
	return nil
}
