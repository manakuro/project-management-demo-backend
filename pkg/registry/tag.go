package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTagController() controller.Tag {
	repo := repository.NewTagRepository(r.client)
	u := usecase.NewTagUsecase(repo)

	return controller.NewTagController(u)
}
