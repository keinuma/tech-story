package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/keinuma/go-graphql/api/domain/service"
	"github.com/keinuma/go-graphql/api/graph/generated"
	"github.com/keinuma/go-graphql/api/graph/model"
	"github.com/keinuma/go-graphql/api/infra/database/gateway"
	"github.com/keinuma/go-graphql/api/presenter"
)

func (r *mutationResolver) CreateStory(ctx context.Context, input model.NewStory) (*model.Story, error) {
	storyPresenter := presenter.NewStory(*service.NewStory(gateway.NewStory(ctx, r.DB)))
	story, err := storyPresenter.CreateStory(input)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userPresenter := presenter.NewUser(*service.NewUser(gateway.NewUser(ctx, r.DB)))
	user, err := userPresenter.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) GetStories(ctx context.Context) ([]*model.Story, error) {
	storyPresenter := presenter.NewStory(*service.NewStory(gateway.NewStory(ctx, r.DB)))
	stories, err := storyPresenter.GetStories(100, 0)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
