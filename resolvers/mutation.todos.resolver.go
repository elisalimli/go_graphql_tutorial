package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/elisalimli/meetmeup/common"
	"github.com/elisalimli/meetmeup/graph/generated"
	"github.com/elisalimli/meetmeup/models"
	"github.com/elisalimli/meetmeup/utils"
	validator "github.com/go-playground/validator/v10"
)

// TodoCreate is the resolver for the todoCreate field.
func (r *mutationResolver) TodoCreate(ctx context.Context, todo models.TodoInput) (*models.ResponseTodoCreate, error) {
	context := common.GetContext(ctx)
	// returns nil or ValidationErrors ( []FieldError )
	validate = validator.New()
	err := validate.Struct(todo)
	if err != nil {
		errors := utils.FormatErrors(err)
		return &models.ResponseTodoCreate{Errors: errors}, nil
	}
	cTodo := &models.Todo{
		Name: todo.Name,
	}
	context.Database.Create(cTodo)
	return &models.ResponseTodoCreate{Todo: cTodo}, nil
}

// TodoComplete is the resolver for the todoComplete field.
func (r *mutationResolver) TodoComplete(ctx context.Context, id int, updatedBy int) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: TodoComplete - todoComplete"))
}

// TodoDelete is the resolver for the todoDelete field.
func (r *mutationResolver) TodoDelete(ctx context.Context, id int, updatedBy int) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: TodoDelete - todoDelete"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.

var validate *validator.Validate
