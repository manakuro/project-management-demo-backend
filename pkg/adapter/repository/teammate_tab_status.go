package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type teammateTabStatusRepository struct {
	client *ent.Client
}

// NewTeammateTabStatusRepository generates teammateTabStatus repository
func NewTeammateTabStatusRepository(client *ent.Client) ur.TeammateTabStatus {
	return &teammateTabStatusRepository{client: client}
}

func (r *teammateTabStatusRepository) Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error) {
	q := r.client.TeammateTabStatus.Query()

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

func (r *teammateTabStatusRepository) List(ctx context.Context) ([]*model.TeammateTabStatus, error) {
	res, err := r.client.TeammateTabStatus.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTabStatusRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error) {
	q := r.client.TeammateTabStatus.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTabStatusFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTabStatusRepository) Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	res, err := r.client.
		TeammateTabStatus.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTabStatusRepository) Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	res, err := r.client.
		TeammateTabStatus.UpdateOneID(input.ID).
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
