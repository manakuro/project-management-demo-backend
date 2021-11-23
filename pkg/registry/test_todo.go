package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewTestTodoController() controller.TestTodo {
	repo := repository.NewTestTodoRepository(r.client)
	u := usecase.NewTestTodoUsecase(repo)

	return controller.NewTestTodoController(u)
}
