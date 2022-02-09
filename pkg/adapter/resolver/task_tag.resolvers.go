package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateTaskTag(ctx context.Context, input ent.CreateTaskTagInput) (*ent.TaskTag, error) {
	t, err := r.controller.TaskTag.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	var ts []*ent.TaskTag
	for _, u := range r.subscriptions.TaskTagUpdated {
		if u.TaskID == t.TaskID {
			if ts == nil {
				ts, err = r.controller.TaskTag.List(ctx, &ent.TaskTagWhereInput{TaskID: &t.TaskID})
				if err != nil {
					// TODO: Add error handling
					fmt.Println(err)
				}
			}
			u.Ch <- ts
		}
	}

	return t, nil
}

func (r *mutationResolver) UpdateTaskTag(ctx context.Context, input ent.UpdateTaskTagInput) (*ent.TaskTag, error) {
	t, err := r.controller.TaskTag.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *mutationResolver) DeleteTaskTag(ctx context.Context, input model.DeleteTaskTagInput) (*ent.TaskTag, error) {
	t, err := r.controller.TaskTag.Delete(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	var ts []*ent.TaskTag
	for _, u := range r.subscriptions.TaskTagUpdated {
		if u.TaskID == t.TaskID {
			if ts == nil {
				ts, err = r.controller.TaskTag.List(ctx, &ent.TaskTagWhereInput{TaskID: &t.TaskID})
				if err != nil {
					fmt.Println(err)
				}
			}
			u.Ch <- ts
		}
	}

	return t, nil
}

func (r *queryResolver) TaskTag(ctx context.Context, where *ent.TaskTagWhereInput) (*ent.TaskTag, error) {
	t, err := r.controller.TaskTag.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *queryResolver) TaskTags(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TaskTagWhereInput) (*ent.TaskTagConnection, error) {

	ts, err := r.controller.TaskTag.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TaskTagsUpdated(ctx context.Context, taskID ulid.ID) (<-chan []*ent.TaskTag, error) {
	key := subscription.NewKey()
	ch := make(chan []*ent.TaskTag, 1)

	r.mutex.Lock()
	r.subscriptions.TaskTagUpdated[key] = subscription.TaskTagUpdated{
		TaskID: taskID,
		Ch:     ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TaskTagUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *taskTagResolver) CreatedAt(ctx context.Context, obj *ent.TaskTag) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *taskTagResolver) UpdatedAt(ctx context.Context, obj *ent.TaskTag) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TaskTag returns generated.TaskTagResolver implementation.
func (r *Resolver) TaskTag() generated.TaskTagResolver { return &taskTagResolver{r} }

type taskTagResolver struct{ *Resolver }
