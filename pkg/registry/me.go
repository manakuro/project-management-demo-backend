package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewMeController() controller.Me {
	repo := repository.NewMeRepository(r.client)
	u := usecase.NewMeUsecase(repo)

	return controller.NewMeController(u)
}
