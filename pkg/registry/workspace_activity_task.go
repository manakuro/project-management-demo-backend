package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewWorkspaceActivityTaskController() controller.WorkspaceActivityTask {
	repo := repository.NewWorkspaceActivityTaskRepository(r.client)
	u := usecase.NewWorkspaceActivityTaskUsecase(repo)

	return controller.NewWorkspaceActivityTaskController(u)
}
