package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskActivityTaskController() controller.TaskActivityTask {
	repo := repository.NewTaskActivityTaskRepository(r.client)
	u := usecase.NewTaskActivityTaskUsecase(repo)

	return controller.NewTaskActivityTaskController(u)
}
