package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type myTasksTabStatusRepository struct {
	client *ent.Client
}

// NewMyTasksTabStatusRepository generates myTasksTabStatus repository
func NewMyTasksTabStatusRepository(client *ent.Client) ur.MyTasksTabStatus {
	return &myTasksTabStatusRepository{client: client}
}

func (r *myTasksTabStatusRepository) Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error) {
	q := r.client.MyTasksTabStatus.Query()

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

func (r *myTasksTabStatusRepository) List(ctx context.Context) ([]*model.MyTasksTabStatus, error) {
	result, err := r.client.MyTasksTabStatus.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *myTasksTabStatusRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error) {
	q := r.client.MyTasksTabStatus.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	result, err := q.Paginate(ctx, after, first, before, last, ent.WithMyTasksTabStatusFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return result, nil
}

func (r *myTasksTabStatusRepository) Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	result, err := r.client.
		MyTasksTabStatus.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *myTasksTabStatusRepository) Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	result, err := r.client.
		MyTasksTabStatus.UpdateOneID(input.ID).
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
