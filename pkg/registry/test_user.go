package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTestUserController() controller.TestUserController {
	repo := repository.NewTestUserRepository(r.client)
	u := usecase.NewTestUserUsecase(repo)

	return controller.NewTestUserController(u)
}
