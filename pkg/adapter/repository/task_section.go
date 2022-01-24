package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskSectionRepository struct {
	client *ent.Client
}

// NewTaskSectionRepository generates taskSection repository
func NewTaskSectionRepository(client *ent.Client) ur.TaskSection {
	return &taskSectionRepository{client: client}
}

func (r *taskSectionRepository) Get(ctx context.Context, where *model.TaskSectionWhereInput) (*model.TaskSection, error) {
	q := r.client.TaskSection.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	result, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *taskSectionRepository) List(ctx context.Context) ([]*model.TaskSection, error) {
	result, err := r.client.TaskSection.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *taskSectionRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskSectionWhereInput, requestedFields []string) (*model.TaskSectionConnection, error) {
	q := r.client.TaskSection.Query()

	result, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskSectionFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return result, nil
}

func (r *taskSectionRepository) Create(ctx context.Context, input model.CreateTaskSectionInput) (*model.TaskSection, error) {
	result, err := r.client.
		TaskSection.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *taskSectionRepository) Update(ctx context.Context, input model.UpdateTaskSectionInput) (*model.TaskSection, error) {
	result, err := r.client.
		TaskSection.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return result, nil
}
