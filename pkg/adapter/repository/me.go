package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type meRepository struct {
	client *ent.Client
}

// NewMeRepository generates teammate repository
func NewMeRepository(client *ent.Client) ur.Me {
	return &meRepository{client: client}
}

func (r *meRepository) Get(ctx context.Context, id model.ID) (*model.Me, error) {
	q := r.client.Teammate.Query()

	if id == "" {
		return nil, model.NewInvalidParamError(map[string]interface{}{
			"id": id,
		})
	}
	q.Where(teammate.IDEQ(id))

	me, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return &model.Me{
		ID:        me.ID,
		Name:      me.Name,
		Image:     me.Image,
		Email:     me.Email,
		CreatedAt: me.CreatedAt,
		UpdatedAt: me.UpdatedAt,
	}, nil
}

func (r *meRepository) Update(ctx context.Context, input model.UpdateMeInput) (*model.Me, error) {
	teammateInput := model.UpdateTeammateInput{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}

	t, err := r.client.
		Teammate.UpdateOneID(input.ID).
		SetInput(teammateInput).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return &model.Me{
		ID:        t.ID,
		Name:      t.Name,
		Image:     t.Image,
		Email:     t.Email,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}
