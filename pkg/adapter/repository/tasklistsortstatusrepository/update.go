package tasklistsortstatusrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *taskListSortStatusRepository) Update(ctx context.Context, input model.UpdateTaskListSortStatusInput) (*model.TaskListSortStatus, error) {
	res, err := r.client.
		TaskListSortStatus.UpdateOneID(input.ID).
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
