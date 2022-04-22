package projecttasksectionrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/pkg/adapter/repository/respositoryutil"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *projectTaskSectionRepository) DeleteAndKeepTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndKeepTasksInput) (*model.DeleteProjectTaskSectionAndKeepTasksPayload, error) {
	client := respositoryutil.WithTransactionalMutation(ctx)

	deleted, err := client.ProjectTaskSection.
		Query().
		Where(projecttasksection.IDEQ(input.ID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	projectTasks, err := client.ProjectTask.
		Query().
		Where(projecttask.ProjectTaskSectionIDEQ(deleted.ID)).
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, deleted.ID)
		}
		return nil, model.NewDBError(err)
	}

	projectTaskSections, err := client.
		ProjectTaskSection.
		Query().
		Where(
			projecttasksection.ProjectID(deleted.ProjectID),
		).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"ProjectID": deleted.ProjectID,
			})
		}
		return nil, model.NewDBError(err)
	}
	projectTaskSection := projectTaskSections[0]

	bulk := make([]*ent.ProjectTaskCreate, len(projectTasks))
	for i, t := range projectTasks {
		bulk[i] = client.ProjectTask.Create().
			SetID(t.ID).
			SetTaskID(t.TaskID).
			SetProjectID(t.ProjectID).
			SetProjectTaskSectionID(projectTaskSection.ID).
			SetCreatedAt(t.CreatedAt)
	}
	err = client.ProjectTask.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	err = client.ProjectTaskSection.DeleteOneID(deleted.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	projectTaskIDs := make([]model.ID, len(projectTasks))
	for i, task := range projectTasks {
		projectTaskIDs[i] = task.ID
	}

	return &model.DeleteProjectTaskSectionAndKeepTasksPayload{
		ProjectTaskSection:     deleted,
		KeptProjectTaskSection: projectTaskSection,
		ProjectTaskIDs:         projectTaskIDs,
	}, nil
}
