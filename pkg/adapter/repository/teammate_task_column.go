package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type teammateTaskColumnRepository struct {
	client *ent.Client
}

// NewTeammateTaskColumnRepository generates teammateTaskColumn repository
func NewTeammateTaskColumnRepository(client *ent.Client) ur.TeammateTaskColumn {
	return &teammateTaskColumnRepository{client: client}
}

func (r *teammateTaskColumnRepository) Get(ctx context.Context, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumn, error) {
	q := r.client.TeammateTaskColumn.Query()

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

func (r *teammateTaskColumnRepository) List(ctx context.Context) ([]*model.TeammateTaskColumn, error) {
	result, err := r.client.TeammateTaskColumn.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *teammateTaskColumnRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput, requestedFields []string) (*model.TeammateTaskColumnConnection, error) {
	q := r.client.TeammateTaskColumn.Query()

	if collection.Contains(requestedFields, "edges.node.taskColumn") {
		q.WithTaskColumn()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	result, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskColumnFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return result, nil
}

func (r *teammateTaskColumnRepository) Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	result, err := r.client.
		TeammateTaskColumn.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *teammateTaskColumnRepository) Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	result, err := r.client.
		TeammateTaskColumn.UpdateOneID(input.ID).
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
