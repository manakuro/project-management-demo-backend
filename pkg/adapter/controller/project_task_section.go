package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectTaskSection is an interface of controller.
type ProjectTaskSection interface {
	Get(ctx context.Context, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSection, error)
	List(ctx context.Context) ([]*model.ProjectTaskSection, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSectionConnection, error)
	Create(ctx context.Context, input model.CreateProjectTaskSectionInput) (*model.ProjectTaskSection, error)
	Update(ctx context.Context, input model.UpdateProjectTaskSectionInput) (*model.ProjectTaskSection, error)
	Delete(ctx context.Context, input model.DeleteProjectTaskSectionInput) (*model.ProjectTaskSection, error)
	DeleteProjectTaskSectionAndKeepTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndKeepTasksInput) (*model.DeleteProjectTaskSectionAndKeepTasksPayload, error)
}

type projectTaskSectionController struct {
	projectTaskSectionUsecase usecase.ProjectTaskSection
}

// NewProjectTaskSectionController generates controller.
func NewProjectTaskSectionController(pt usecase.ProjectTaskSection) ProjectTaskSection {
	return &projectTaskSectionController{
		projectTaskSectionUsecase: pt,
	}
}

func (c *projectTaskSectionController) Get(ctx context.Context, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSection, error) {
	return c.projectTaskSectionUsecase.Get(ctx, where)
}

func (c *projectTaskSectionController) List(ctx context.Context) ([]*model.ProjectTaskSection, error) {
	return c.projectTaskSectionUsecase.List(ctx)
}

func (c *projectTaskSectionController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSectionConnection, error) {
	return c.projectTaskSectionUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (c *projectTaskSectionController) Create(ctx context.Context, input model.CreateProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	return c.projectTaskSectionUsecase.Create(ctx, input)
}

func (c *projectTaskSectionController) Update(ctx context.Context, input model.UpdateProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	return c.projectTaskSectionUsecase.Update(ctx, input)
}

func (c *projectTaskSectionController) Delete(ctx context.Context, input model.DeleteProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	return c.projectTaskSectionUsecase.Delete(ctx, input)
}

func (c *projectTaskSectionController) DeleteProjectTaskSectionAndKeepTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndKeepTasksInput) (*model.DeleteProjectTaskSectionAndKeepTasksPayload, error) {
	return c.projectTaskSectionUsecase.DeleteProjectTaskSectionAndKeepTasks(ctx, input)
}
