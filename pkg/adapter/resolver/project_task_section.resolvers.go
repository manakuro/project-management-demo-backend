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
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateProjectTaskSection(ctx context.Context, input ent.CreateProjectTaskSectionInput) (*ent.ProjectTaskSection, error) {
	p, err := r.controller.ProjectTaskSection.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	go func() {
		for _, c := range r.subscriptions.ProjectTaskSectionCreated {
			if c.WorkspaceID == input.WorkspaceID && c.RequestID != input.RequestID {
				c.Ch <- p
			}
		}
	}()

	return p, nil
}

func (r *mutationResolver) UpdateProjectTaskSection(ctx context.Context, input ent.UpdateProjectTaskSectionInput) (*ent.ProjectTaskSection, error) {
	p, err := r.controller.ProjectTaskSection.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	go func() {
		for _, u := range r.subscriptions.ProjectTaskSectionUpdated {
			if u.WorkspaceID == input.WorkspaceID && u.RequestID != input.RequestID {
				u.Ch <- p
			}
		}
	}()

	return p, nil
}

func (r *mutationResolver) DeleteProjectTaskSection(ctx context.Context, input model.DeleteProjectTaskSectionInput) (*ent.ProjectTaskSection, error) {
	p, err := r.controller.ProjectTaskSection.Delete(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return p, nil
}

func (r *mutationResolver) DeleteProjectTaskSectionAndKeepTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndKeepTasksInput) (*model.DeleteProjectTaskSectionAndKeepTasksPayload, error) {
	p, err := r.controller.ProjectTaskSection.DeleteProjectTaskSectionAndKeepTasks(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return p, nil
}

func (r *mutationResolver) DeleteProjectTaskSectionAndDeleteTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndDeleteTasksInput) (*model.DeleteProjectTaskSectionAndDeleteTasksPayload, error) {
	p, err := r.controller.ProjectTaskSection.DeleteProjectTaskSectionAndDeleteTasks(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return p, nil
}

func (r *projectTaskSectionResolver) CreatedAt(ctx context.Context, obj *ent.ProjectTaskSection) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *projectTaskSectionResolver) UpdatedAt(ctx context.Context, obj *ent.ProjectTaskSection) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) ProjectTaskSection(ctx context.Context, where *ent.ProjectTaskSectionWhereInput) (*ent.ProjectTaskSection, error) {
	p, err := r.controller.ProjectTaskSection.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return p, nil
}

func (r *queryResolver) ProjectTaskSections(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectTaskSectionWhereInput) (*ent.ProjectTaskSectionConnection, error) {
	ps, err := r.controller.ProjectTaskSection.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ps, nil
}

func (r *subscriptionResolver) ProjectTaskSectionUpdated(ctx context.Context, workspaceID ulid.ID, requestID string) (<-chan *ent.ProjectTaskSection, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.ProjectTaskSection, 1)

	r.mutex.Lock()
	r.subscriptions.ProjectTaskSectionUpdated[key] = subscription.ProjectTaskSectionUpdated{
		WorkspaceID: workspaceID,
		RequestID:   requestID,
		Ch:          ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.ProjectTaskSectionUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) ProjectTaskSectionCreated(ctx context.Context, workspaceID ulid.ID, requestID string) (<-chan *ent.ProjectTaskSection, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.ProjectTaskSection, 1)

	r.mutex.Lock()
	r.subscriptions.ProjectTaskSectionCreated[key] = subscription.ProjectTaskSectionCreated{
		WorkspaceID: workspaceID,
		RequestID:   requestID,
		Ch:          ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.ProjectTaskSectionCreated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// ProjectTaskSection returns generated.ProjectTaskSectionResolver implementation.
func (r *Resolver) ProjectTaskSection() generated.ProjectTaskSectionResolver {
	return &projectTaskSectionResolver{r}
}

type projectTaskSectionResolver struct{ *Resolver }
