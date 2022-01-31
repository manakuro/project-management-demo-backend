package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewFileTypeController() controller.FileType {
	repo := repository.NewFileTypeRepository(r.client)
	u := usecase.NewFileTypeUsecase(repo)

	return controller.NewFileTypeController(u)
}
