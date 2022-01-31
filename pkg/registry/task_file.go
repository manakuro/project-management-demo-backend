package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskFileController() controller.TaskFile {
	repo := repository.NewTaskFileRepository(r.client)
	u := usecase.NewTaskFileUsecase(repo)

	return controller.NewTaskFileController(u)
}
