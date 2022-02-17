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

func (r *mutationResolver) CreateTaskCollaborator(ctx context.Context, input ent.CreateTaskCollaboratorInput) (*ent.TaskCollaborator, error) {
	t, err := r.controller.TaskCollaborator.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	var ts []*ent.TaskCollaborator
	for _, u := range r.subscriptions.TaskCollaboratorUpdated {
		if u.TaskID == t.TaskID && u.RequestID != input.RequestID {
			if ts == nil {
				ts, err = r.controller.TaskCollaborator.List(ctx, &ent.TaskCollaboratorWhereInput{TaskID: &t.TaskID})
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

func (r *mutationResolver) UpdateTaskCollaborator(ctx context.Context, input ent.UpdateTaskCollaboratorInput) (*ent.TaskCollaborator, error) {
	t, err := r.controller.TaskCollaborator.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *mutationResolver) DeleteTaskCollaborator(ctx context.Context, input model.DeleteTaskCollaboratorInput) (*ent.TaskCollaborator, error) {
	t, err := r.controller.TaskCollaborator.Delete(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	var ts []*ent.TaskCollaborator
	for _, u := range r.subscriptions.TaskCollaboratorUpdated {
		if u.TaskID == t.TaskID && u.RequestID != input.RequestID {
			if ts == nil {
				ts, err = r.controller.TaskCollaborator.List(ctx, &ent.TaskCollaboratorWhereInput{TaskID: &t.TaskID})
				if err != nil {
					fmt.Println(err)
				}
			}
			u.Ch <- ts
		}
	}

	return t, nil
}

func (r *queryResolver) TaskCollaborator(ctx context.Context, where *ent.TaskCollaboratorWhereInput) (*ent.TaskCollaborator, error) {
	t, err := r.controller.TaskCollaborator.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return t, nil
}

func (r *queryResolver) TaskCollaborators(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TaskCollaboratorWhereInput) (*ent.TaskCollaboratorConnection, error) {
	ts, err := r.controller.TaskCollaborator.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TaskCollaboratorsUpdated(ctx context.Context, taskID ulid.ID, requestID string) (<-chan []*ent.TaskCollaborator, error) {
	key := subscription.NewKey()
	ch := make(chan []*ent.TaskCollaborator, 1)

	r.mutex.Lock()
	r.subscriptions.TaskCollaboratorUpdated[key] = subscription.TaskCollaboratorUpdated{
		TaskID:    taskID,
		RequestID: requestID,
		Ch:        ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TaskCollaboratorUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *taskCollaboratorResolver) CreatedAt(ctx context.Context, obj *ent.TaskCollaborator) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *taskCollaboratorResolver) UpdatedAt(ctx context.Context, obj *ent.TaskCollaborator) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TaskCollaborator returns generated.TaskCollaboratorResolver implementation.
func (r *Resolver) TaskCollaborator() generated.TaskCollaboratorResolver {
	return &taskCollaboratorResolver{r}
}

type taskCollaboratorResolver struct{ *Resolver }
