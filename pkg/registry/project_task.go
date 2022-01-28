package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectTaskController() controller.ProjectTask {
	repo := repository.NewProjectTaskRepository(r.client)
	u := usecase.NewProjectTaskUsecase(repo)

	return controller.NewProjectTaskController(u)
}
