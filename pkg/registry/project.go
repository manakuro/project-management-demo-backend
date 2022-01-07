package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectController() controller.Project {
	repo := repository.NewProjectRepository(r.client)
	u := usecase.NewProjectUsecase(repo)

	return controller.NewProjectController(u)
}
