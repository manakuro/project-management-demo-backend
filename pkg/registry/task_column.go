package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskColumnController() controller.TaskColumn {
	repo := repository.NewTaskColumnRepository(r.client)
	u := usecase.NewTaskColumnUsecase(repo)

	return controller.NewTaskColumnController(u)
}
