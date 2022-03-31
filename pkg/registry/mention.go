package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewMentionController() controller.Mention {
	repo := repository.NewMentionRepository(r.client)
	u := usecase.NewMentionUsecase(repo)

	return controller.NewMentionController(u)
}
