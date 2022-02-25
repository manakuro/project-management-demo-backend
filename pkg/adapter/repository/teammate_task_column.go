package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammatetaskcolumn"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
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

func (r *teammateTaskColumnRepository) List(ctx context.Context) ([]*model.TeammateTaskColumn, error) {
	res, err := r.client.TeammateTaskColumn.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskColumnRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumnConnection, error) {
	q := r.client.TeammateTaskColumn.Query().Order(ent.Asc(teammatetaskcolumn.FieldOrder))

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskColumnFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskColumnRepository) Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	res, err := r.client.
		TeammateTaskColumn.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskColumnRepository) Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	res, err := r.client.
		TeammateTaskColumn.UpdateOneID(input.ID).
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

func (r *teammateTaskColumnRepository) UpdateOrder(ctx context.Context, input model.UpdateTeammateTaskColumnOrderInput) ([]*model.TeammateTaskColumn, error) {
	client := WithTransactionalMutation(ctx)
	if len(input.IDs) == 0 {
		return nil, model.NewInvalidParamError(input.IDs)
	}

	ts, err := client.TeammateTaskColumn.
		Query().
		WithTaskColumn().
		Where(teammatetaskcolumn.IDIn(input.IDs...)).
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	groupedByID := make(map[ulid.ID]*model.TeammateTaskColumn, len(ts))
	for _, t := range ts {
		groupedByID[t.ID] = t
	}

	bulk := make([]*ent.TeammateTaskColumnCreate, len(input.IDs))
	for i, id := range input.IDs {
		g, ok := groupedByID[id]
		if ok {
			bulk[i] = client.TeammateTaskColumn.Create().
				SetID(id).
				SetOrder(i).
				SetTeammateID(g.TeammateID).
				SetTaskColumnID(g.TaskColumnID).
				SetWorkspaceID(g.WorkspaceID).
				SetWidth(g.Width).
				SetCustomizable(g.Customizable).
				SetDisabled(g.Disabled).
				SetCreatedAt(g.CreatedAt)
		}
	}

	err = client.TeammateTaskColumn.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return ts, nil
}
