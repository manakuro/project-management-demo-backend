package projecttasksectionrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/pkg/adapter/repository/repositoryutil"
	"project-management-demo-backend/pkg/adapter/repository/taskrepository"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *projectTaskSectionRepository) DeleteAndDeleteTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndDeleteTasksInput) (*model.DeleteProjectTaskSectionAndDeleteTasksPayload, error) {
	client := repositoryutil.WithTransactionalMutation(ctx)

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

	taskIDs := make([]model.ID, len(projectTasks))
	for i, t := range projectTasks {
		taskIDs[i] = t.TaskID
	}

	// TODO: Task repository can be called in usecase/usecase package
	taskRepo := taskrepository.New(r.client)

	_, err = taskRepo.DeleteAll(ctx, model.DeleteAllTaskInput{
		TaskIDs:     taskIDs,
		WorkspaceID: input.WorkspaceID,
		RequestID:   "",
	})
	if err != nil {
		return nil, err
	}

	err = client.ProjectTaskSection.DeleteOneID(deleted.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	projectTaskIDs := make([]model.ID, len(projectTasks))
	for i, t := range projectTasks {
		projectTaskIDs[i] = t.ID
	}

	return &model.DeleteProjectTaskSectionAndDeleteTasksPayload{
		ProjectTaskSection: deleted,
		ProjectTaskIDs:     projectTaskIDs,
		TaskIDs:            taskIDs,
	}, nil
}
