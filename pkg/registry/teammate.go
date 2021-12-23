package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateController() controller.Teammate {
	repo := repository.NewTeammateRepository(r.client)
	u := usecase.NewTeammateUsecase(repo)

	return controller.NewTeammateController(u)
}
