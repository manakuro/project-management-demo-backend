package registry

import (
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/usecase/usecase"
)

func (r *registry) NewArchivedTaskActivityController() controller.ArchivedTaskActivity {
	repo := repository.NewArchivedTaskActivityRepository(r.client)
	u := usecase.NewArchivedTaskActivityUsecase(repo)

	return controller.NewArchivedTaskActivityController(u)
}
