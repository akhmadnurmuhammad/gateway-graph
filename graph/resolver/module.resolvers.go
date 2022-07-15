package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	module1 "xti-gateway-graph-go/graph/authorization/module"
	module "xti-gateway-graph-go/graph/authorization/module/model"
)

// ModuleStore is the resolver for the moduleStore field.
func (r *mutationResolver) ModuleStore(ctx context.Context, in *module.ModuleRequestInput) (*module.MainResponseModule, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModuleUpdate is the resolver for the moduleUpdate field.
func (r *mutationResolver) ModuleUpdate(ctx context.Context, in *module.ModuleRequestInput) (*module.MainResponseModule, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModuleDelete is the resolver for the moduleDelete field.
func (r *mutationResolver) ModuleDelete(ctx context.Context, in *module.FindByIDModuleRequestInput) (*module.MainResponseModule, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModuleFetch is the resolver for the moduleFetch field.
func (r *mutationResolver) ModuleFetch(ctx context.Context, in *module.FetchRequestModuleInput) (*module.FetchResponseModule, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModuleFindByID is the resolver for the moduleFindByID field.
func (r *mutationResolver) ModuleFindByID(ctx context.Context, in *module.FindByIDModuleRequestInput) (*module.MainResponseModule, error) {
	panic(fmt.Errorf("not implemented"))
}

// Dummy is the resolver for the dummy field.
func (r *queryResolver) Dummy(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns module1.MutationResolver implementation.
func (r *Resolver) Mutation() module1.MutationResolver { return &mutationResolver{r} }

// Query returns module1.QueryResolver implementation.
func (r *Resolver) Query() module1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
