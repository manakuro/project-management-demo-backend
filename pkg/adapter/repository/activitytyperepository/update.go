package activitytyperepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *activityTypeRepository) Update(ctx context.Context, input model.UpdateActivityTypeInput) (*model.ActivityType, error) {
	res, err := r.client.
		ActivityType.UpdateOneID(input.ID).
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
