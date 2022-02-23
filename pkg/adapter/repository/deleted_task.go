package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type deletedTaskRepository struct {
	client *ent.Client
}

// NewDeletedTaskRepository generates deletedTask repository
func NewDeletedTaskRepository(client *ent.Client) ur.DeletedTask {
	return &deletedTaskRepository{client: client}
}

func (r *deletedTaskRepository) Get(ctx context.Context, where *model.DeletedTaskWhereInput) (*model.DeletedTask, error) {
	q := r.client.DeletedTask.Query()

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

func (r *deletedTaskRepository) List(ctx context.Context) ([]*model.DeletedTask, error) {
	res, err := r.client.DeletedTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *deletedTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.DeletedTaskWhereInput) (*model.DeletedTaskConnection, error) {
	q := r.client.DeletedTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithDeletedTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *deletedTaskRepository) Create(ctx context.Context, input model.CreateDeletedTaskInput) (*model.DeletedTask, error) {
	newTask, err := r.client.Task.
		Create().
		SetIsNew(true).
		SetName("").
		SetDescription(model.DefaultEditorDescription()).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	_, err = r.client.TaskFeed.Create().
		SetTask(newTask).
		SetIsFirst(true).
		SetDescription(model.DefaultEditorDescription()).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	res, err := r.client.
		DeletedTask.
		Create().
		SetInput(input).
		SetTask(newTask).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *deletedTaskRepository) Update(ctx context.Context, input model.UpdateDeletedTaskInput) (*model.DeletedTask, error) {
	res, err := r.client.
		DeletedTask.UpdateOneID(input.ID).
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

func (r *deletedTaskRepository) Undelete(ctx context.Context, input model.UndeleteDeletedTaskInput) ([]*model.DeletedTask, error) {
	client := WithTransactionalMutation(ctx)

	deletedTasks, err := client.DeletedTask.
		Query().
		WithTask().
		Where(deletedtask.TaskID(input.TaskID)).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.TaskID)
		}
		return nil, model.NewDBError(err)
	}

	// The task to be undeleted will be limited up to two records.
	for _, t := range deletedTasks {
		if t.TaskType == deletedtask.TaskTypeTeammate {
			_, err = client.TeammateTask.Create().
				SetWorkspaceID(t.WorkspaceID).
				SetTaskID(t.TaskID).
				SetTeammateID(t.TaskJoinID).
				SetTeammateTaskSectionID(t.TaskSectionID).
				Save(ctx)
			if err != nil {
				return nil, model.NewDBError(err)
			}
		}
		if t.TaskType == deletedtask.TaskTypeProject {
			_, err = client.ProjectTask.Create().
				SetProjectID(t.TaskJoinID).
				SetTaskID(t.TaskID).
				SetProjectTaskSectionID(t.TaskSectionID).
				Save(ctx)
			if err != nil {
				return nil, model.NewDBError(err)
			}
		}
	}

	deletedIds := make([]model.ID, len(deletedTasks))
	for i, t := range deletedTasks {
		deletedIds[i] = t.ID
	}

	_, err = client.DeletedTask.
		Delete().
		Where(deletedtask.IDIn(deletedIds...)).
		Exec(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deletedTasks, nil
}
