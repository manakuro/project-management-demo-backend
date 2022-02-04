package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectTeammate is an interface of controller.
type ProjectTeammate interface {
	Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error)
	List(ctx context.Context) ([]*model.ProjectTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error)
	Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error)
	Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error)
}

type projectTeammateController struct {
	projectTeammateUsecase usecase.ProjectTeammate
}

// NewProjectTeammateController generates projectTeammate controller.
func NewProjectTeammateController(pt usecase.ProjectTeammate) ProjectTeammate {
	return &projectTeammateController{
		projectTeammateUsecase: pt,
	}
}

func (c *projectTeammateController) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	return c.projectTeammateUsecase.Get(ctx, where)
}

func (c *projectTeammateController) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	return c.projectTeammateUsecase.List(ctx)
}

func (c *projectTeammateController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error) {
	return c.projectTeammateUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *projectTeammateController) Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return c.projectTeammateUsecase.Create(ctx, input)
}

func (c *projectTeammateController) Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return c.projectTeammateUsecase.Update(ctx, input)
}
