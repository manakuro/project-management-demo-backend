package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewColorController() controller.Color {
	repo := repository.NewColorRepository(r.client)
	u := usecase.NewColorUsecase(repo)

	return controller.NewColorController(u)
}
