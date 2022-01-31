package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskFeedLikeController() controller.TaskFeedLike {
	repo := repository.NewTaskFeedLikeRepository(r.client)
	u := usecase.NewTaskFeedLikeUsecase(repo)

	return controller.NewTaskFeedLikeController(u)
}
