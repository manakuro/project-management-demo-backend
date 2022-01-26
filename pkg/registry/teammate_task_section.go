package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTaskSectionController() controller.TeammateTaskSection {
	repo := repository.NewTeammateTaskSectionRepository(r.client)
	u := usecase.NewTeammateTaskSectionUsecase(repo)

	return controller.NewTeammateTaskSectionController(u)
}
