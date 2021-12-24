package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type workspace struct {
	workspaceRepository repository.Workspace
}

// Workspace is an interface of test user
type Workspace interface {
	Get(ctx context.Context, id model.ID) (*model.Workspace, error)
	List(ctx context.Context) ([]*model.Workspace, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput) (*model.WorkspaceConnection, error)
	Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error)
	Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error)
}

// NewWorkspaceUsecase generates test user repository
func NewWorkspaceUsecase(r repository.Workspace) Workspace {
	return &workspace{workspaceRepository: r}
}

func (t *workspace) Get(ctx context.Context, id model.ID) (*model.Workspace, error) {
	return t.workspaceRepository.Get(ctx, id)
}

func (t *workspace) List(ctx context.Context) ([]*model.Workspace, error) {
	return t.workspaceRepository.List(ctx)
}

func (t *workspace) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput) (*model.WorkspaceConnection, error) {
	return t.workspaceRepository.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *workspace) Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error) {
	return t.workspaceRepository.Create(ctx, input)
}

func (t *workspace) Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error) {
	return t.workspaceRepository.Update(ctx, input)
}
