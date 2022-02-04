package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Me is an interface of controller.
type Me interface {
	Get(ctx context.Context, id model.ID) (*model.Me, error)
	Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error)
}

type meController struct {
	meUsecase usecase.Me
}

// NewMeController generates me controller.
func NewMeController(tu usecase.Me) Me {
	return &meController{
		meUsecase: tu,
	}
}

func (c *meController) Get(ctx context.Context, id model.ID) (*model.Me, error) {
	return c.meUsecase.Get(ctx, id)
}

func (c *meController) Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error) {
	return c.meUsecase.Update(ctx, input)
}
