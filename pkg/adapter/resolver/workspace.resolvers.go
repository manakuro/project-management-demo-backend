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

func (r *mutationResolver) CreateWorkspace(ctx context.Context, input ent.CreateWorkspaceInput) (*ent.Workspace, error) {
	w, err := r.controller.Workspace.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return w, nil
}

func (r *mutationResolver) UpdateWorkspace(ctx context.Context, input ent.UpdateWorkspaceInput) (*ent.Workspace, error) {
	w, err := r.controller.Workspace.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, wu := range r.subscriptions.WorkspaceUpdated {
		if wu.ID == w.ID {
			wu.Ch <- w
		}
	}

	return w, nil
}

func (r *queryResolver) Workspace(ctx context.Context, where *ent.WorkspaceWhereInput) (*ent.Workspace, error) {
	requestFields := graphqlutil.GetRequestedFields(ctx)

	ws, err := r.controller.Workspace.Get(ctx, where, requestFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return ws, nil
}

func (r *queryResolver) Workspaces(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.WorkspaceWhereInput) (*ent.WorkspaceConnection, error) {
	requestFields := graphqlutil.GetRequestedFields(ctx)

	ws, err := r.controller.Workspace.ListWithPagination(ctx, after, first, before, last, where, requestFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return ws, nil
}

func (r *subscriptionResolver) WorkspaceUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.Workspace, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.Workspace, 1)

	r.mutex.Lock()
	r.subscriptions.WorkspaceUpdated[key] = subscription.WorkspaceUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.WorkspaceUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *workspaceResolver) CreatedAt(ctx context.Context, obj *ent.Workspace) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *workspaceResolver) UpdatedAt(ctx context.Context, obj *ent.Workspace) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Workspace returns generated.WorkspaceResolver implementation.
func (r *Resolver) Workspace() generated.WorkspaceResolver { return &workspaceResolver{r} }

type workspaceResolver struct{ *Resolver }
