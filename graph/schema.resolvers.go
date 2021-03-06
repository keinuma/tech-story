package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/keinuma/tech-story/domain/model"
	"github.com/keinuma/tech-story/domain/service"
	"github.com/keinuma/tech-story/graph/generated"
	"github.com/keinuma/tech-story/infra/database/gateway"
	"github.com/keinuma/tech-story/presenter"
	"github.com/keinuma/tech-story/presenter/request"
)

func (r *matchResolver) Attendees(ctx context.Context, obj *model.Match) ([]*model.User, error) {
	return obj.Attendees, nil
}

func (r *mutationResolver) CreateStory(ctx context.Context, input request.NewStory) (*model.Story, error) {
	storyPresenter := presenter.NewStory(*service.NewStory(gateway.NewStory(ctx, r.DB)))
	story, err := storyPresenter.CreateStory(input)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input request.NewUser) (*model.User, error) {
	userPresenter := presenter.NewUser(*service.NewUser(gateway.NewUser(ctx, r.DB)))
	user, err := userPresenter.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) CreateMatch(ctx context.Context, input request.NewMatch) (*model.Match, error) {
	matchPresenter := presenter.NewMatch(*service.NewMatch(gateway.NewMatch(ctx, r.DB)))
	match, err := matchPresenter.CreateMatch(input)
	if err != nil {
		return nil, err
	}

	err = r.Subscriber.Publish(ctx, match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (*bool, error) {
	flag := true
	return &flag, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input request.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetStories(ctx context.Context) ([]*model.Story, error) {
	storyPresenter := presenter.NewStory(*service.NewStory(gateway.NewStory(ctx, r.DB)))
	stories, err := storyPresenter.GetStories(100, 0)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *queryResolver) GetMatches(ctx context.Context) ([]*model.Match, error) {
	matchPresenter := presenter.NewMatch(*service.NewMatch(gateway.NewMatch(ctx, r.DB)))
	matches, err := matchPresenter.GetMatches(100, 0)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *subscriptionResolver) CreateMatch(ctx context.Context, userUID string) (<-chan *model.Match, error) {
	return r.Subscriber.Receive(ctx, userUID), nil
}

// Match returns generated.MatchResolver implementation.
func (r *Resolver) Match() generated.MatchResolver { return &matchResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type matchResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
