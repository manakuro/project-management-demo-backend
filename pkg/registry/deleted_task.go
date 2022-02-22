package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewDeletedTaskController() controller.DeletedTask {
	repo := repository.NewDeletedTaskRepository(r.client)
	u := usecase.NewDeletedTaskUsecase(repo)

	return controller.NewDeletedTaskController(u)
}
