package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type me struct {
	meRepository repository.Me
}

// Me is an interface of test user
type Me interface {
	Get(ctx context.Context, id model.ID) (*model.Me, error)
	Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error)
}

// NewMeUsecase generates test user repository
func NewMeUsecase(r repository.Me) Me {
	return &me{meRepository: r}
}

func (t *me) Get(ctx context.Context, id model.ID) (*model.Me, error) {
	return t.meRepository.Get(ctx, id)
}

func (t *me) Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error) {
	return t.meRepository.Update(ctx, input)
}
