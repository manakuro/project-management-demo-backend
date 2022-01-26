package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTaskTabStatusController() controller.TeammateTaskTabStatus {
	repo := repository.NewTeammateTaskTabStatusRepository(r.client)
	u := usecase.NewTeammateTaskTabStatusUsecase(repo)

	return controller.NewTeammateTaskTabStatusController(u)
}
