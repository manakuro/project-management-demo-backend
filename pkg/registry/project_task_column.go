package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectTaskColumnController() controller.ProjectTaskColumn {
	repo := repository.NewProjectTaskColumnRepository(r.client)
	u := usecase.NewProjectTaskColumnUsecase(repo)

	return controller.NewProjectTaskColumnController(u)
}
