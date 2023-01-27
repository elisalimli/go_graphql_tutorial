package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	generated1 "github.com/elisalimli/meetmeup/graphql/generated"
	"github.com/elisalimli/meetmeup/graphql/models"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }