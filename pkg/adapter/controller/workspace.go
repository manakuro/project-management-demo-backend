package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Workspace is an interface of controller
type Workspace interface {
	Get(ctx context.Context, where *model.WorkspaceWhereInput, requestFields []string) (*model.Workspace, error)
	List(ctx context.Context) ([]*model.Workspace, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput, requestFields []string) (*model.WorkspaceConnection, error)
	Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error)
	Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error)
}

type workspace struct {
	workspaceUsecase usecase.Workspace
}

// NewWorkspaceController generates workspace controller
func NewWorkspaceController(tu usecase.Workspace) Workspace {
	return &workspace{
		workspaceUsecase: tu,
	}
}

func (t *workspace) Get(ctx context.Context, where *model.WorkspaceWhereInput, requestFields []string) (*model.Workspace, error) {
	return t.workspaceUsecase.Get(ctx, where, requestFields)
}

func (t *workspace) List(ctx context.Context) ([]*model.Workspace, error) {
	return t.workspaceUsecase.List(ctx)
}

func (t *workspace) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput, requestFields []string) (*model.WorkspaceConnection, error) {
	return t.workspaceUsecase.ListWithPagination(ctx, after, first, before, last, where, requestFields)
}

func (t *workspace) Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error) {
	return t.workspaceUsecase.Create(ctx, input)
}

func (t *workspace) Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error) {
	return t.workspaceUsecase.Update(ctx, input)
}
