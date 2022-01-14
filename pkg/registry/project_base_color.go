package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectBaseColorController() controller.ProjectBaseColor {
	repo := repository.NewProjectBaseColorRepository(r.client)
	u := usecase.NewProjectBaseColorUsecase(repo)

	return controller.NewProjectBaseColorController(u)
}
