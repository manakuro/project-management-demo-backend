package tasklistcompletedstatusrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *taskListCompletedStatusRepository) Update(ctx context.Context, input model.UpdateTaskListCompletedStatusInput) (*model.TaskListCompletedStatus, error) {
	res, err := r.client.
		TaskListCompletedStatus.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return res, nil
}
