package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectTeammateController() controller.ProjectTeammate {
	repo := repository.NewProjectTeammateRepository(r.client)
	u := usecase.NewProjectTeammateUsecase(repo)

	return controller.NewProjectTeammateController(u)
}
