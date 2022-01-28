package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTaskController() controller.TeammateTask {
	repo := repository.NewTeammateTaskRepository(r.client)
	u := usecase.NewTeammateTaskUsecase(repo)

	return controller.NewTeammateTaskController(u)
}
