package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateTeammate(ctx context.Context, input ent.CreateTeammateInput) (*ent.Teammate, error) {
	t, err := r.controller.Teammate.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *mutationResolver) UpdateTeammate(ctx context.Context, input ent.UpdateTeammateInput) (*ent.Teammate, error) {
	t, err := r.controller.Teammate.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, tu := range r.subscriptions.TeammateUpdated {
		if tu.ID == t.ID {
			tu.Ch <- t
		}
	}

	return t, nil
}

func (r *queryResolver) Teammate(ctx context.Context, id ulid.ID) (*ent.Teammate, error) {
	t, err := r.controller.Teammate.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) Teammates(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TeammateWhereInput) (*ent.TeammateConnection, error) {
	ts, err := r.controller.Teammate.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return ts, nil
}

func (r *subscriptionResolver) TeammateUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.Teammate, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.Teammate, 1)

	r.mutex.Lock()
	r.subscriptions.TeammateUpdated[key] = subscription.TeammateUpdated{ID: id, Ch: ch}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TeammateUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *teammateResolver) CreatedAt(ctx context.Context, obj *ent.Teammate) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *teammateResolver) UpdatedAt(ctx context.Context, obj *ent.Teammate) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Teammate returns generated.TeammateResolver implementation.
func (r *Resolver) Teammate() generated.TeammateResolver { return &teammateResolver{r} }

type teammateResolver struct{ *Resolver }
