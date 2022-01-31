package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTaskFeedController() controller.TaskFeed {
	repo := repository.NewTaskFeedRepository(r.client)
	u := usecase.NewTaskFeedUsecase(repo)

	return controller.NewTaskFeedController(u)
}
