package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectTaskColumn is an interface of controller
type ProjectTaskColumn interface {
	Get(ctx context.Context, where *model.ProjectTaskColumnWhereInput) (*model.ProjectTaskColumn, error)
	List(ctx context.Context) ([]*model.ProjectTaskColumn, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskColumnWhereInput, requestedFields []string) (*model.ProjectTaskColumnConnection, error)
	Create(ctx context.Context, input model.CreateProjectTaskColumnInput) (*model.ProjectTaskColumn, error)
	Update(ctx context.Context, input model.UpdateProjectTaskColumnInput) (*model.ProjectTaskColumn, error)
}

type projectTaskColumn struct {
	projectTaskColumnUsecase usecase.ProjectTaskColumn
}

// NewProjectTaskColumnController generates projectTaskColumn controller
func NewProjectTaskColumnController(pt usecase.ProjectTaskColumn) ProjectTaskColumn {
	return &projectTaskColumn{
		projectTaskColumnUsecase: pt,
	}
}

func (p *projectTaskColumn) Get(ctx context.Context, where *model.ProjectTaskColumnWhereInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnUsecase.Get(ctx, where)
}

func (p *projectTaskColumn) List(ctx context.Context) ([]*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnUsecase.List(ctx)
}

func (p *projectTaskColumn) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskColumnWhereInput, requestedFields []string) (*model.ProjectTaskColumnConnection, error) {
	return p.projectTaskColumnUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectTaskColumn) Create(ctx context.Context, input model.CreateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnUsecase.Create(ctx, input)
}

func (p *projectTaskColumn) Update(ctx context.Context, input model.UpdateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnUsecase.Update(ctx, input)
}
