package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewWorkspaceActivityController() controller.WorkspaceActivity {
	repo := repository.NewWorkspaceActivityRepository(r.client)
	u := usecase.NewWorkspaceActivityUsecase(repo)

	return controller.NewWorkspaceActivityController(u)
}
