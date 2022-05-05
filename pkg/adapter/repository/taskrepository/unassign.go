package taskrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/pkg/adapter/repository/repositoryutil"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *taskRepository) Unassign(ctx context.Context, input model.UnassignTaskInput) (*model.UnassignTaskPayload, error) {
	client := repositoryutil.WithTransactionalMutation(ctx)

	updatedTask, err := client.Task.UpdateOneID(input.ID).ClearAssigneeID().Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	teammateTask, err := client.TeammateTask.Query().Where(teammatetask.TaskID(updatedTask.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	_, err = client.TeammateTask.Delete().Where(teammatetask.ID(teammateTask.ID)).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.UnassignTaskPayload{
		Task:           updatedTask,
		TeammateTaskID: teammateTask.ID,
	}, nil
}
