package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/datetime"
	"time"
)

type teammateTaskRepository struct {
	client *ent.Client
}

// NewTeammateTaskRepository generates teammateTask repository
func NewTeammateTaskRepository(client *ent.Client) ur.TeammateTask {
	return &teammateTaskRepository{client: client}
}

func (r *teammateTaskRepository) Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error) {
	q := r.client.TeammateTask.Query()

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

func (r *teammateTaskRepository) List(ctx context.Context) ([]*model.TeammateTask, error) {
	res, err := r.client.TeammateTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskRepository) TasksDueSoon(ctx context.Context, workspaceID model.ID, teammateID model.ID) ([]*model.TeammateTask, error) {
	q := r.client.TeammateTask.Query()

	q.Where(teammatetask.TeammateIDEQ(teammateID))
	q.Where(teammatetask.HasTaskWith(
		task.DueDateGTE(datetime.StartOfDay(time.Now())),
		task.DueDateLTE(datetime.EndOfDay(datetime.AddDate(5))),
	))
	q.Where(teammatetask.HasTeammateTaskSectionWith(
		teammatetasksection.WorkspaceIDEQ(workspaceID),
	))

	res, err := q.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, err
}

func (r *teammateTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput) (*model.TeammateTaskConnection, error) {
	q := r.client.TeammateTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskRepository) Create(ctx context.Context, input model.CreateTeammateTaskInput) (*model.TeammateTask, error) {
	client := WithTransactionalMutation(ctx)

	newTask, err := client.Task.
		Create().
		SetIsNew(true).
		SetCreatedBy(input.TeammateID).
		SetAssigneeID(input.TeammateID).
		SetName("").
		SetDescription(model.DefaultEditorDescription()).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	res, err := client.
		TeammateTask.
		Create().
		SetInput(input).
		SetTask(newTask).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskRepository) Update(ctx context.Context, input model.UpdateTeammateTaskInput) (*model.TeammateTask, error) {
	res, err := r.client.
		TeammateTask.UpdateOneID(input.ID).
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

func (r *teammateTaskRepository) Delete(ctx context.Context, input model.DeleteTeammateTaskInput) (*model.TeammateTask, error) {
	deleted, err := r.client.TeammateTask.Query().Where(teammatetask.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TeammateTask.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}
