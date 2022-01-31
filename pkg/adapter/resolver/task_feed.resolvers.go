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
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/graphqlutil"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateTaskFeed(ctx context.Context, input ent.CreateTaskFeedInput) (*ent.TaskFeed, error) {
	t, err := r.controller.TaskFeed.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *mutationResolver) UpdateTaskFeed(ctx context.Context, input ent.UpdateTaskFeedInput) (*ent.TaskFeed, error) {
	t, err := r.controller.TaskFeed.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.TaskFeedUpdated {
		if u.ID == t.ID {
			u.Ch <- t
		}
	}

	return t, nil
}

func (r *mutationResolver) DeleteTaskFeed(ctx context.Context, input model.DeleteTaskFeedInput) (*ent.TaskFeed, error) {
	t, err := r.controller.TaskFeed.Delete(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *queryResolver) TaskFeed(ctx context.Context, where *ent.TaskFeedWhereInput) (*ent.TaskFeed, error) {
	t, err := r.controller.TaskFeed.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) TaskFeeds(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TaskFeedWhereInput) (*ent.TaskFeedConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)

	ts, err := r.controller.TaskFeed.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TaskFeedUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.TaskFeed, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.TaskFeed, 1)

	r.mutex.Lock()
	r.subscriptions.TaskFeedUpdated[key] = subscription.TaskFeedUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TaskFeedUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *taskFeedResolver) CreatedAt(ctx context.Context, obj *ent.TaskFeed) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *taskFeedResolver) UpdatedAt(ctx context.Context, obj *ent.TaskFeed) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TaskFeed returns generated.TaskFeedResolver implementation.
func (r *Resolver) TaskFeed() generated.TaskFeedResolver { return &taskFeedResolver{r} }

type taskFeedResolver struct{ *Resolver }