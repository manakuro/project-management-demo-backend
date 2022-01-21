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

func (r *mutationResolver) CreateMyTasksTabStatus(ctx context.Context, input ent.CreateMyTasksTabStatusInput) (*ent.MyTasksTabStatus, error) {
	result, err := r.controller.MyTasksTabStatus.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return result, nil
}

func (r *mutationResolver) UpdateMyTasksTabStatus(ctx context.Context, input ent.UpdateMyTasksTabStatusInput) (*ent.MyTasksTabStatus, error) {
	result, err := r.controller.MyTasksTabStatus.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.MyTasksTabStatusUpdated {
		if u.ID == result.ID {
			u.Ch <- result
		}
	}

	return result, nil
}

func (r *myTasksTabStatusResolver) CreatedAt(ctx context.Context, obj *ent.MyTasksTabStatus) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *myTasksTabStatusResolver) UpdatedAt(ctx context.Context, obj *ent.MyTasksTabStatus) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) MyTasksTabStatus(ctx context.Context, where *ent.MyTasksTabStatusWhereInput) (*ent.MyTasksTabStatus, error) {
	result, err := r.controller.MyTasksTabStatus.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return result, nil
}

func (r *queryResolver) MyTasksTabStatuses(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.MyTasksTabStatusWhereInput) (*ent.MyTasksTabStatusConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)

	result, err := r.controller.MyTasksTabStatus.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return result, nil
}

func (r *subscriptionResolver) MyTasksTabStatusUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.MyTasksTabStatus, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.MyTasksTabStatus, 1)

	r.mutex.Lock()
	r.subscriptions.MyTasksTabStatusUpdated[key] = subscription.MyTasksTabStatusUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.MyTasksTabStatusUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// MyTasksTabStatus returns generated.MyTasksTabStatusResolver implementation.
func (r *Resolver) MyTasksTabStatus() generated.MyTasksTabStatusResolver {
	return &myTasksTabStatusResolver{r}
}

type myTasksTabStatusResolver struct{ *Resolver }
