package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewProjectTaskSectionController() controller.ProjectTaskSection {
	repo := repository.NewProjectTaskSectionRepository(r.client)
	u := usecase.NewProjectTaskSectionUsecase(repo)

	return controller.NewProjectTaskSectionController(u)
}
