package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectIconController() controller.ProjectIcon {
	repo := repository.NewProjectIconRepository(r.client)
	u := usecase.NewProjectIconUsecase(repo)

	return controller.NewProjectIconController(u)
}
