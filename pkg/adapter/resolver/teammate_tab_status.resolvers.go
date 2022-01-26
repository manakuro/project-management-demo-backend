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
	"project-management-demo-backend/pkg/util/graphqlutil"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateTeammateTabStatus(ctx context.Context, input ent.CreateTeammateTabStatusInput) (*ent.TeammateTabStatus, error) {
	t, err := r.controller.TeammateTabStatus.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *mutationResolver) UpdateTeammateTabStatus(ctx context.Context, input ent.UpdateTeammateTabStatusInput) (*ent.TeammateTabStatus, error) {
	t, err := r.controller.TeammateTabStatus.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.TeammateTabStatusUpdated {
		if u.ID == t.ID {
			u.Ch <- t
		}
	}

	return t, nil
}

func (r *queryResolver) TeammateTabStatus(ctx context.Context, where *ent.TeammateTabStatusWhereInput) (*ent.TeammateTabStatus, error) {
	t, err := r.controller.TeammateTabStatus.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) TeammateTabStatuses(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TeammateTabStatusWhereInput) (*ent.TeammateTabStatusConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)

	ts, err := r.controller.TeammateTabStatus.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TeammateTabStatusUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.TeammateTabStatus, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.TeammateTabStatus, 1)

	r.mutex.Lock()
	r.subscriptions.TeammateTabStatusUpdated[key] = subscription.TeammateTabStatusUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TeammateTabStatusUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *teammateTabStatusResolver) CreatedAt(ctx context.Context, obj *ent.TeammateTabStatus) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *teammateTabStatusResolver) UpdatedAt(ctx context.Context, obj *ent.TeammateTabStatus) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TeammateTabStatus returns generated.TeammateTabStatusResolver implementation.
func (r *Resolver) TeammateTabStatus() generated.TeammateTabStatusResolver {
	return &teammateTabStatusResolver{r}
}

type teammateTabStatusResolver struct{ *Resolver }
