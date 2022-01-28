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

func (r *mutationResolver) CreateProjectTask(ctx context.Context, input ent.CreateProjectTaskInput) (*ent.ProjectTask, error) {
	p, err := r.controller.ProjectTask.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return p, nil
}

func (r *mutationResolver) UpdateProjectTask(ctx context.Context, input ent.UpdateProjectTaskInput) (*ent.ProjectTask, error) {
	p, err := r.controller.ProjectTask.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.ProjectTaskUpdated {
		if u.ID == p.ID {
			u.Ch <- p
		}
	}

	return p, nil
}

func (r *projectTaskResolver) CreatedAt(ctx context.Context, obj *ent.ProjectTask) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *projectTaskResolver) UpdatedAt(ctx context.Context, obj *ent.ProjectTask) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) ProjectTask(ctx context.Context, where *ent.ProjectTaskWhereInput) (*ent.ProjectTask, error) {
	p, err := r.controller.ProjectTask.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return p, nil
}

func (r *queryResolver) ProjectTasks(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectTaskWhereInput) (*ent.ProjectTaskConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)
	ps, err := r.controller.ProjectTask.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ps, nil
}

func (r *subscriptionResolver) ProjectTaskUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.ProjectTask, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.ProjectTask, 1)

	r.mutex.Lock()
	r.subscriptions.ProjectTaskUpdated[key] = subscription.ProjectTaskUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.ProjectTaskUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// ProjectTask returns generated.ProjectTaskResolver implementation.
func (r *Resolver) ProjectTask() generated.ProjectTaskResolver { return &projectTaskResolver{r} }

type projectTaskResolver struct{ *Resolver }