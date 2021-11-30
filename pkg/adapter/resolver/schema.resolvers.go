package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *queryResolver) Node(ctx context.Context, id ulid.ID) (ent.Noder, error) {
	n, err := r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, handler.HandleError(ctx, model.NewNotFoundError(err, id))
		}
		return nil, handler.HandleError(ctx, model.NewGraphQLError(err))
	}
	return n, nil
}

func (r *queryResolver) Nodes(ctx context.Context, ids []ulid.ID) ([]ent.Noder, error) {
	ns, err := r.client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, handler.HandleError(ctx, model.NewNotFoundError(err, ids))
		}
		return nil, handler.HandleError(ctx, model.NewGraphQLError(err))
	}
	return ns, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
