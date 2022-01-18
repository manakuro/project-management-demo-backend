package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewWorkspaceTeammateController() controller.WorkspaceTeammate {
	repo := repository.NewWorkspaceTeammateRepository(r.client)
	u := usecase.NewWorkspaceTeammateUsecase(repo)

	return controller.NewWorkspaceTeammateController(u)
}
