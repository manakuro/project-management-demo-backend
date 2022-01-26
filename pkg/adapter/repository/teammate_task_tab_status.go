package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type teammateTaskTabStatusRepository struct {
	client *ent.Client
}

// NewTeammateTaskTabStatusRepository generates teammateTaskTabStatus repository
func NewTeammateTaskTabStatusRepository(client *ent.Client) ur.TeammateTaskTabStatus {
	return &teammateTaskTabStatusRepository{client: client}
}

func (r *teammateTaskTabStatusRepository) Get(ctx context.Context, where *model.TeammateTaskTabStatusWhereInput) (*model.TeammateTaskTabStatus, error) {
	q := r.client.TeammateTaskTabStatus.Query()

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

func (r *teammateTaskTabStatusRepository) List(ctx context.Context) ([]*model.TeammateTaskTabStatus, error) {
	res, err := r.client.TeammateTaskTabStatus.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskTabStatusRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskTabStatusWhereInput, requestedFields []string) (*model.TeammateTaskTabStatusConnection, error) {
	q := r.client.TeammateTaskTabStatus.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskTabStatusFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskTabStatusRepository) Create(ctx context.Context, input model.CreateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error) {
	res, err := r.client.
		TeammateTaskTabStatus.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskTabStatusRepository) Update(ctx context.Context, input model.UpdateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error) {
	res, err := r.client.
		TeammateTaskTabStatus.UpdateOneID(input.ID).
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
