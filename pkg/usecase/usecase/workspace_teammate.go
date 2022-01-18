package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type workspaceTeammate struct {
	workspaceTeammateRepository repository.WorkspaceTeammate
}

// WorkspaceTeammate is an interface of test user
type WorkspaceTeammate interface {
	Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error)
	List(ctx context.Context) ([]*model.WorkspaceTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput, requestedFields []string) (*model.WorkspaceTeammateConnection, error)
	Create(ctx context.Context, input model.CreateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error)
	Update(ctx context.Context, input model.UpdateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error)
}

// NewWorkspaceTeammateUsecase generates test user repository
func NewWorkspaceTeammateUsecase(r repository.WorkspaceTeammate) WorkspaceTeammate {
	return &workspaceTeammate{workspaceTeammateRepository: r}
}

func (w *workspaceTeammate) Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error) {
	return w.workspaceTeammateRepository.Get(ctx, where)
}

func (w *workspaceTeammate) List(ctx context.Context) ([]*model.WorkspaceTeammate, error) {
	return w.workspaceTeammateRepository.List(ctx)
}

func (w *workspaceTeammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput, requestedFields []string) (*model.WorkspaceTeammateConnection, error) {
	return w.workspaceTeammateRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (w *workspaceTeammate) Create(ctx context.Context, input model.CreateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	return w.workspaceTeammateRepository.Create(ctx, input)
}

func (w *workspaceTeammate) Update(ctx context.Context, input model.UpdateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	return w.workspaceTeammateRepository.Update(ctx, input)
}
