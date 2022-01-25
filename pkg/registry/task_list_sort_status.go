package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskListSortStatusController() controller.TaskListSortStatus {
	repo := repository.NewTaskListSortStatusRepository(r.client)
	u := usecase.NewTaskListSortStatusUsecase(repo)

	return controller.NewTaskListSortStatusController(u)
}
