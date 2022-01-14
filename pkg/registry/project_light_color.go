package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectLightColorController() controller.ProjectLightColor {
	repo := repository.NewProjectLightColorRepository(r.client)
	u := usecase.NewProjectLightColorUsecase(repo)

	return controller.NewProjectLightColorController(u)
}
