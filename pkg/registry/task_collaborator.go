package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskCollaboratorController() controller.TaskCollaborator {
	repo := repository.NewTaskCollaboratorRepository(r.client)
	u := usecase.NewTaskCollaboratorUsecase(repo)

	return controller.NewTaskCollaboratorController(u)
}
