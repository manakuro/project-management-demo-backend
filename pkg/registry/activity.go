package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewActivityController() controller.Activity {
	repo := repository.NewActivityRepository(r.client)
	u := usecase.NewActivityUsecase(repo)

	return controller.NewActivityController(u)
}
