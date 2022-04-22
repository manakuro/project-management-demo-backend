package teammatetasksectionrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/pkg/adapter/repository/respositoryutil"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *teammateTaskSectionRepository) UndeleteAndKeepTasks(ctx context.Context, input model.UndeleteTeammateTaskSectionAndKeepTasksInput) (*model.UndeleteTeammateTaskSectionAndKeepTasksPayload, error) {
	client := respositoryutil.WithTransactionalMutation(ctx)

	createdTeammateTaskSection, err := client.TeammateTaskSection.
		Create().
		SetTeammateID(input.TeammateID).
		SetWorkspaceID(input.WorkspaceID).
		SetName(input.Name).
		SetAssigned(false).
		SetCreatedAt(*input.CreatedAt).
		SetUpdatedAt(*input.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	teammateTasks, err := client.TeammateTask.Query().Where(teammatetask.IDIn(input.KeptTeammateTaskIDs...)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.KeptTeammateTaskIDs)
		}
		return nil, model.NewDBError(err)
	}
	bulk := make([]*ent.TeammateTaskCreate, len(teammateTasks))
	for i, t := range teammateTasks {
		bulk[i] = client.TeammateTask.Create().
			SetID(t.ID).
			SetTaskID(t.TaskID).
			SetTeammateID(t.TeammateID).
			SetWorkspaceID(t.WorkspaceID).
			SetTeammateTaskSectionID(createdTeammateTaskSection.ID).
			SetCreatedAt(t.CreatedAt)
	}
	err = client.TeammateTask.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	teammateTaskIDs := make([]model.ID, len(teammateTasks))
	for i, t := range teammateTasks {
		teammateTaskIDs[i] = t.ID
	}

	return &model.UndeleteTeammateTaskSectionAndKeepTasksPayload{
		TeammateTaskSection: createdTeammateTaskSection,
		TeammateTaskIDs:     teammateTaskIDs,
	}, nil
}
