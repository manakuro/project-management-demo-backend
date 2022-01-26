package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTaskListStatusController() controller.TeammateTaskListStatus {
	repo := repository.NewTeammateTaskListStatusRepository(r.client)
	u := usecase.NewTeammateTaskListStatusUsecase(repo)

	return controller.NewTeammateTaskListStatusController(u)
}
