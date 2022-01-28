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

func (r *mutationResolver) CreateTask(ctx context.Context, input ent.CreateTaskInput) (*ent.Task, error) {
	t, err := r.controller.Task.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input ent.UpdateTaskInput) (*ent.Task, error) {
	t, err := r.controller.Task.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.TaskUpdated {
		if u.ID == t.ID {
			u.Ch <- t
		}
	}

	return t, nil
}

func (r *queryResolver) Task(ctx context.Context, where *ent.TaskWhereInput) (*ent.Task, error) {
	t, err := r.controller.Task.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) Tasks(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TaskWhereInput) (*ent.TaskConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)

	ts, err := r.controller.Task.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TaskUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.Task, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.Task, 1)

	r.mutex.Lock()
	r.subscriptions.TaskUpdated[key] = subscription.TaskUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TaskUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *taskResolver) CompletedAt(ctx context.Context, obj *ent.Task) (string, error) {
	if obj.CompletedAt == nil {
		return "", nil
	}

	return datetime.FormatDate(*obj.CompletedAt), nil
}

func (r *taskResolver) DueDate(ctx context.Context, obj *ent.Task) (string, error) {
	if obj.DueDate == nil {
		return "", nil
	}

	return datetime.FormatDate(*obj.DueDate), nil
}

func (r *taskResolver) DueTime(ctx context.Context, obj *ent.Task) (string, error) {
	if obj.DueTime == nil {
		return "", nil
	}

	return datetime.FormatDate(*obj.DueTime), nil
}

func (r *taskResolver) CreatedAt(ctx context.Context, obj *ent.Task) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *taskResolver) UpdatedAt(ctx context.Context, obj *ent.Task) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }