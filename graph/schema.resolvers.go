package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ChihHaoChen/graphql-booking-server/graph/generated"
	"github.com/ChihHaoChen/graphql-booking-server/graph/model"
	"github.com/google/uuid"
)

// SignUpUser is the resolver for the signUpUser field.
func (r *mutationResolver) SignUpUser(ctx context.Context, input model.AuthInput) (*model.User, error) {
	user := &model.User{
		Email:    input.Email,
		Password: input.Password,
		ID:       uuid.New().String(),
		Username: input.Email,
	}

	r.users = append(r.users, user)

	return user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Travels(ctx context.Context) ([]*model.Travel, error) {
	return r.travels, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
