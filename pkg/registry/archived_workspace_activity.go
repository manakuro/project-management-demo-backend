package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewArchivedWorkspaceActivityController() controller.ArchivedWorkspaceActivity {
	repo := repository.NewArchivedWorkspaceActivityRepository(r.client)
	u := usecase.NewArchivedWorkspaceActivityUsecase(repo)

	return controller.NewArchivedWorkspaceActivityController(u)
}
