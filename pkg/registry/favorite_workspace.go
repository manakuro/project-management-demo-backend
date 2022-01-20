package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewFavoriteWorkspaceController() controller.FavoriteWorkspace {
	repo := repository.NewFavoriteWorkspaceRepository(r.client)
	u := usecase.NewFavoriteWorkspaceUsecase(repo)

	return controller.NewFavoriteWorkspaceController(u)
}
