package projecttaskrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/pkg/adapter/repository/respositoryutil"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *projectTaskRepository) Delete(ctx context.Context, input model.DeleteProjectTaskInput) (*model.ProjectTask, error) {
	client := respositoryutil.WithTransactionalMutation(ctx)

	deleted, err := client.ProjectTask.
		Query().
		Where(projecttask.IDEQ(input.ID)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.ProjectTask.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}
