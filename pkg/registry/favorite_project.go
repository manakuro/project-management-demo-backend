package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewFavoriteProjectController() controller.FavoriteProject {
	repo := repository.NewFavoriteProjectRepository(r.client)
	u := usecase.NewFavoriteProjectUsecase(repo)

	return controller.NewFavoriteProjectController(u)
}
