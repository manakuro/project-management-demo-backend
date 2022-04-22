package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository/testuserrepository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTestUserController() controller.TestUser {
	repo := testuserrepository.New(r.client)
	u := usecase.NewTestUserUsecase(repo)

	return controller.NewTestUserController(u)
}
