package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewArchivedActivityController() controller.ArchivedActivity {
	repo := repository.NewArchivedActivityRepository(r.client)
	u := usecase.NewArchivedActivityUsecase(repo)

	return controller.NewArchivedActivityController(u)
}
