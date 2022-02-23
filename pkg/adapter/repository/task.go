package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskRepository struct {
	client *ent.Client
}

// NewTaskRepository generates task repository.
func NewTaskRepository(client *ent.Client) ur.Task {
	return &taskRepository{client: client}
}

func (r *taskRepository) Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error) {
	q := r.client.Task.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	res, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) List(ctx context.Context) ([]*model.Task, error) {
	res, err := r.client.Task.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error) {
	q := r.client.Task.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskRepository) Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.UpdateOneID(input.ID).
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

func (r *taskRepository) Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteTaskPayload{
		ProjectTask:  &model.ProjectTask{},
		TeammateTask: &model.TeammateTask{},
		DeletedTasks: []*model.DeletedTask{},
	}

	deletedTask, err := client.Task.Query().Where(task.ID(input.TaskID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	deletedTeammateTask, err := client.
		TeammateTask.Query().
		WithTask().
		Where(teammatetask.TaskID(input.TaskID)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedTeammateTaskDeletedTask *model.DeletedTask
	if deletedTeammateTask != nil {
		err = client.TeammateTask.DeleteOneID(deletedTeammateTask.ID).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !deletedTask.IsNew {
			d, derr := client.DeletedTask.
				Create().
				SetTaskID(input.TaskID).
				SetWorkspaceID(input.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeTeammate).
				SetTaskJoinID(deletedTeammateTask.TeammateID).
				SetTaskSectionID(deletedTeammateTask.TeammateTaskSectionID).
				Save(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			// Restore the state in order to get an entity after a successful transaction.
			d, derr = client.DeletedTask.Query().WithTask().Where(deletedtask.ID(d.ID)).Only(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			deletedTeammateTaskDeletedTask = d.Unwrap()
		}
	}

	deletedProjectTask, err := client.ProjectTask.
		Query().
		WithTask().
		Where(projecttask.TaskID(input.TaskID)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedProjectTaskDeletedTask *model.DeletedTask
	if deletedProjectTask != nil {
		err = client.ProjectTask.DeleteOneID(deletedProjectTask.ID).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !deletedTask.IsNew {
			d, derr := r.client.DeletedTask.
				Create().
				SetTaskID(input.TaskID).
				SetWorkspaceID(input.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeProject).
				SetTaskJoinID(deletedProjectTask.ProjectID).
				SetTaskSectionID(deletedProjectTask.ProjectTaskSectionID).
				Save(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			// Restore the state in order to get an entity after a successful transaction.
			d, derr = client.DeletedTask.Query().WithTask().Where(deletedtask.ID(d.ID)).Only(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			deletedProjectTaskDeletedTask = d.Unwrap()
		}
	}

	if deletedTeammateTask != nil {
		payload.TeammateTask = deletedTeammateTask
	}
	if deletedProjectTask != nil {
		payload.ProjectTask = deletedProjectTask
	}
	if deletedTeammateTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedTeammateTaskDeletedTask)
	}
	if deletedProjectTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedProjectTaskDeletedTask)
	}

	return payload, nil
}
