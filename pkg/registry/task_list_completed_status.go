package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskListCompletedStatusController() controller.TaskListCompletedStatus {
	repo := repository.NewTaskListCompletedStatusRepository(r.client)
	u := usecase.NewTaskListCompletedStatusUsecase(repo)

	return controller.NewTaskListCompletedStatusController(u)
}
