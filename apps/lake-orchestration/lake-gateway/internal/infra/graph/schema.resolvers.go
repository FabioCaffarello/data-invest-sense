package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"apps/lake-orchestration/lake-gateway/internal/infra/graph/model"
	"context"
	"fmt"
)

// CreateInput is the resolver for the createInput field.
func (r *mutationResolver) CreateInput(ctx context.Context, input model.InputInput, source string, service string) (*model.Input, error) {
	panic(fmt.Errorf("not implemented: CreateInput - createInput"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
