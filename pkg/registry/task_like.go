package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskLikeController() controller.TaskLike {
	repo := repository.NewTaskLikeRepository(r.client)
	u := usecase.NewTaskLikeUsecase(repo)

	return controller.NewTaskLikeController(u)
}
