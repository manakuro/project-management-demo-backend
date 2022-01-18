package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// WorkspaceTeammate is an interface of controller
type WorkspaceTeammate interface {
	Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error)
	List(ctx context.Context) ([]*model.WorkspaceTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput, requestedFields []string) (*model.WorkspaceTeammateConnection, error)
	Create(ctx context.Context, input model.CreateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error)
	Update(ctx context.Context, input model.UpdateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error)
}

type workspaceTeammate struct {
	workspaceTeammateUsecase usecase.WorkspaceTeammate
}

// NewWorkspaceTeammateController generates workspaceTeammate controller
func NewWorkspaceTeammateController(wt usecase.WorkspaceTeammate) WorkspaceTeammate {
	return &workspaceTeammate{
		workspaceTeammateUsecase: wt,
	}
}

func (p *workspaceTeammate) Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error) {
	return p.workspaceTeammateUsecase.Get(ctx, where)
}

func (p *workspaceTeammate) List(ctx context.Context) ([]*model.WorkspaceTeammate, error) {
	return p.workspaceTeammateUsecase.List(ctx)
}

func (p *workspaceTeammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput, requestedFields []string) (*model.WorkspaceTeammateConnection, error) {
	return p.workspaceTeammateUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *workspaceTeammate) Create(ctx context.Context, input model.CreateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	return p.workspaceTeammateUsecase.Create(ctx, input)
}

func (p *workspaceTeammate) Update(ctx context.Context, input model.UpdateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	return p.workspaceTeammateUsecase.Update(ctx, input)
}
