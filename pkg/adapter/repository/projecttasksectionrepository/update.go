package projecttasksectionrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *projectTaskSectionRepository) Update(ctx context.Context, input model.UpdateProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	res, err := r.client.
		ProjectTaskSection.UpdateOneID(input.ID).
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
