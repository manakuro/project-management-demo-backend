package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectTeammate is an interface of controller
type ProjectTeammate interface {
	Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error)
	List(ctx context.Context) ([]*model.ProjectTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error)
	Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error)
	Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error)
}

type projectTeammate struct {
	projectTeammateUsecase usecase.ProjectTeammate
}

// NewProjectTeammateController generates projectTeammate controller
func NewProjectTeammateController(pt usecase.ProjectTeammate) ProjectTeammate {
	return &projectTeammate{
		projectTeammateUsecase: pt,
	}
}

func (p *projectTeammate) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateUsecase.Get(ctx, where)
}

func (p *projectTeammate) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	return p.projectTeammateUsecase.List(ctx)
}

func (p *projectTeammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error) {
	return p.projectTeammateUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectTeammate) Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateUsecase.Create(ctx, input)
}

func (p *projectTeammate) Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateUsecase.Update(ctx, input)
}
