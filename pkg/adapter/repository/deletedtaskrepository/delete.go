package deletedtaskrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *deletedTaskRepository) Delete(ctx context.Context, input model.DeleteDeletedTaskInput) (*model.DeletedTask, error) {
	deleted, err := r.client.DeletedTask.Query().Where(deletedtask.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.DeletedTask.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}
