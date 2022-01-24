package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type workspaceUsecase struct {
	workspaceRepository repository.Workspace
}

// Workspace is an interface of test user
type Workspace interface {
	Get(ctx context.Context, where *model.WorkspaceWhereInput, requestFields []string) (*model.Workspace, error)
	List(ctx context.Context) ([]*model.Workspace, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput, requestFields []string) (*model.WorkspaceConnection, error)
	Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error)
	Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error)
}

// NewWorkspaceUsecase generates test user repository
func NewWorkspaceUsecase(r repository.Workspace) Workspace {
	return &workspaceUsecase{workspaceRepository: r}
}

func (u *workspaceUsecase) Get(ctx context.Context, where *model.WorkspaceWhereInput, requestFields []string) (*model.Workspace, error) {
	return u.workspaceRepository.Get(ctx, where, requestFields)
}

func (u *workspaceUsecase) List(ctx context.Context) ([]*model.Workspace, error) {
	return u.workspaceRepository.List(ctx)
}

func (u *workspaceUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput, requestFields []string) (*model.WorkspaceConnection, error) {
	return u.workspaceRepository.ListWithPagination(ctx, after, first, before, last, where, requestFields)
}

func (u *workspaceUsecase) Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error) {
	return u.workspaceRepository.Create(ctx, input)
}

func (u *workspaceUsecase) Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error) {
	return u.workspaceRepository.Update(ctx, input)
}
