package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewArchivedWorkspaceActivityTaskController() controller.ArchivedWorkspaceActivityTask {
	repo := repository.NewArchivedWorkspaceActivityTaskRepository(r.client)
	u := usecase.NewArchivedWorkspaceActivityTaskUsecase(repo)

	return controller.NewArchivedWorkspaceActivityTaskController(u)
}
