package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/tasklistsortstatus"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type teammateTaskListStatusRepository struct {
	client *ent.Client
}

// NewTeammateTaskListStatusRepository generates teammateTaskListStatus repository
func NewTeammateTaskListStatusRepository(client *ent.Client) ur.TeammateTaskListStatus {
	return &teammateTaskListStatusRepository{client: client}
}

func (r *teammateTaskListStatusRepository) Get(ctx context.Context, where *model.TeammateTaskListStatusWhereInput) (*model.TeammateTaskListStatus, error) {
	q := r.client.TeammateTaskListStatus.Query()

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

func (r *teammateTaskListStatusRepository) List(ctx context.Context) ([]*model.TeammateTaskListStatus, error) {
	res, err := r.client.TeammateTaskListStatus.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskListStatusRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskListStatusWhereInput) (*model.TeammateTaskListStatusConnection, error) {
	q := r.client.TeammateTaskListStatus.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskListStatusFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskListStatusRepository) Create(ctx context.Context, input model.CreateTeammateTaskListStatusInput) (*model.TeammateTaskListStatus, error) {
	res, err := r.client.
		TeammateTaskListStatus.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskListStatusRepository) Update(ctx context.Context, input model.UpdateTeammateTaskListStatusInput) (*model.TeammateTaskListStatus, error) {
	client := WithTransactionalMutation(ctx)

	q := client.
		TeammateTaskListStatus.UpdateOneID(input.ID).
		SetInput(input)

	if input.TaskListSortStatusCode != nil {
		taskListSortStatus, terr := client.
			TaskListSortStatus.
			Query().
			Where(tasklistsortstatus.StatusCodeEQ(*input.TaskListSortStatusCode)).
			Only(ctx)

		if terr != nil {
			return nil, model.NewDBError(terr)
		}

		q.SetTaskListSortStatusID(taskListSortStatus.ID)
	}

	if input.TaskListCompletedStatusCode != nil {
		taskListCompletedStatus, terr := client.
			TaskListCompletedStatus.
			Query().
			Where(tasklistcompletedstatus.StatusCodeEQ(*input.TaskListCompletedStatusCode)).
			Only(ctx)

		if terr != nil {
			return nil, model.NewDBError(terr)
		}

		q.SetTaskListCompletedStatusID(taskListCompletedStatus.ID)
	}

	res, err := q.Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return res, nil
}
