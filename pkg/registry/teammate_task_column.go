package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTeammateTaskColumnController() controller.TeammateTaskColumn {
	repo := repository.NewTeammateTaskColumnRepository(r.client)
	u := usecase.NewTeammateTaskColumnUsecase(repo)

	return controller.NewTeammateTaskColumnController(u)
}
