package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewIconController() controller.Icon {
	repo := repository.NewIconRepository(r.client)
	u := usecase.NewIconUsecase(repo)

	return controller.NewIconController(u)
}
