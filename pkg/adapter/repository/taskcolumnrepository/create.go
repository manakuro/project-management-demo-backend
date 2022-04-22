package taskcolumnrepository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *taskColumnRepository) Create(ctx context.Context, input model.CreateTaskColumnInput) (*model.TaskColumn, error) {
	res, err := r.client.
		TaskColumn.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}
