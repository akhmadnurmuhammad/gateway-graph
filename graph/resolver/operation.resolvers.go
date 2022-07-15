package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	operation1 "xti-gateway-graph-go/graph/authorization/operation"
	operation "xti-gateway-graph-go/graph/authorization/operation/model"
)

// OperationStore is the resolver for the operationStore field.
func (r *mutationResolver) OperationStore(ctx context.Context, in *operation.OperationRequestInput) (*operation.MainResponseOperation, error) {
	panic(fmt.Errorf("not implemented"))
}

// OperationUpdate is the resolver for the operationUpdate field.
func (r *mutationResolver) OperationUpdate(ctx context.Context, in *operation.OperationRequestInput) (*operation.MainResponseOperation, error) {
	panic(fmt.Errorf("not implemented"))
}

// OperationDelete is the resolver for the operationDelete field.
func (r *mutationResolver) OperationDelete(ctx context.Context, in *operation.FindByIDOperationRequestInput) (*operation.MainResponseOperation, error) {
	panic(fmt.Errorf("not implemented"))
}

// OperationFetch is the resolver for the operationFetch field.
func (r *queryResolver) OperationFetch(ctx context.Context, in *operation.FetchRequestOperationInput) (*operation.FetchResponseOperation, error) {
	panic(fmt.Errorf("not implemented"))
}

// OperationFindByID is the resolver for the operationFindByID field.
func (r *queryResolver) OperationFindByID(ctx context.Context, in *operation.FindByIDOperationRequestInput) (*operation.MainResponseOperation, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns operation1.MutationResolver implementation.
func (r *Resolver) Mutation() operation1.MutationResolver { return &mutationResolver{r} }

// Query returns operation1.QueryResolver implementation.
func (r *Resolver) Query() operation1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
