package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Me is an interface of controller
type Me interface {
	Get(ctx context.Context, id model.ID) (*model.Me, error)
	Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error)
}

type me struct {
	meUsecase usecase.Me
}

// NewMeController generates me controller
func NewMeController(tu usecase.Me) Me {
	return &me{
		meUsecase: tu,
	}
}

func (t *me) Get(ctx context.Context, id model.ID) (*model.Me, error) {
	return t.meUsecase.Get(ctx, id)
}

func (t *me) Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error) {
	return t.meUsecase.Update(ctx, input)
}
