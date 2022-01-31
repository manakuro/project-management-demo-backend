package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskTagController() controller.TaskTag {
	repo := repository.NewTaskTagRepository(r.client)
	u := usecase.NewTaskTagUsecase(repo)

	return controller.NewTaskTagController(u)
}
