package taskcollaboratorrepository

import (
	"context"
	"project-management-demo-backend/ent/taskcollaborator"
	"project-management-demo-backend/pkg/adapter/repository/respositoryutil"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *taskCollaboratorRepository) Create(ctx context.Context, input model.CreateTaskCollaboratorInput) (*model.TaskCollaborator, error) {
	client := respositoryutil.WithTransactionalMutation(ctx)

	res, err := client.
		TaskCollaborator.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	q := client.TaskCollaborator.Query().Where(taskcollaborator.ID(res.ID))

	respositoryutil.WithTaskCollaborator(q)

	taskCollaborator, err := q.Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return taskCollaborator, nil
}
