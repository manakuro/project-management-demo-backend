package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TeammateTaskSection is an interface of controller.
type TeammateTaskSection interface {
	Get(ctx context.Context, where *model.TeammateTaskSectionWhereInput) (*model.TeammateTaskSection, error)
	List(ctx context.Context) ([]*model.TeammateTaskSection, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskSectionWhereInput, requestedFields []string) (*model.TeammateTaskSectionConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskSectionInput) (*model.TeammateTaskSection, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskSectionInput) (*model.TeammateTaskSection, error)
	Delete(ctx context.Context, input model.DeleteTeammateTaskSectionInput) (*model.TeammateTaskSection, error)
}

type teammateTaskSectionController struct {
	teammateTaskSectionUsecase usecase.TeammateTaskSection
}

// NewTeammateTaskSectionController generates controller.
func NewTeammateTaskSectionController(pt usecase.TeammateTaskSection) TeammateTaskSection {
	return &teammateTaskSectionController{
		teammateTaskSectionUsecase: pt,
	}
}

func (c *teammateTaskSectionController) Get(ctx context.Context, where *model.TeammateTaskSectionWhereInput) (*model.TeammateTaskSection, error) {
	return c.teammateTaskSectionUsecase.Get(ctx, where)
}

func (c *teammateTaskSectionController) List(ctx context.Context) ([]*model.TeammateTaskSection, error) {
	return c.teammateTaskSectionUsecase.List(ctx)
}

func (c *teammateTaskSectionController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskSectionWhereInput, requestedFields []string) (*model.TeammateTaskSectionConnection, error) {
	return c.teammateTaskSectionUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *teammateTaskSectionController) Create(ctx context.Context, input model.CreateTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	return c.teammateTaskSectionUsecase.Create(ctx, input)
}

func (c *teammateTaskSectionController) Update(ctx context.Context, input model.UpdateTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	return c.teammateTaskSectionUsecase.Update(ctx, input)
}

func (c *teammateTaskSectionController) Delete(ctx context.Context, input model.DeleteTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	return c.teammateTaskSectionUsecase.Delete(ctx, input)
}
