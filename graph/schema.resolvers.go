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

func (r *subscriptionResolver) CreateComment(ctx context.Context, userID string) (<-chan *model.Match, error) {
	panic(fmt.Errorf("not implemented"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *matchResolver) Comments(ctx context.Context, obj *model.Match) ([]*model.Comment, error) {
	return obj.Comments, nil
}
