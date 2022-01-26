package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTabStatusController() controller.TeammateTabStatus {
	repo := repository.NewTeammateTabStatusRepository(r.client)
	u := usecase.NewTeammateTabStatusUsecase(repo)

	return controller.NewTeammateTabStatusController(u)
}
