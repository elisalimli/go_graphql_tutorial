package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/elisalimli/meetmeup/common"
	"github.com/elisalimli/meetmeup/graphql/models"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id int) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, limit *int, offset *int) ([]models.Todo, error) {
	context := common.GetContext(ctx)
	var todos []models.Todo
	context.Database.Find(&todos)

	return todos, nil
}
