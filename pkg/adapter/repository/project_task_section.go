package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type projectTaskSectionRepository struct {
	client *ent.Client
}

// NewProjectTaskSectionRepository generates projectTaskSection repository
func NewProjectTaskSectionRepository(client *ent.Client) ur.ProjectTaskSection {
	return &projectTaskSectionRepository{client: client}
}

func (r *projectTaskSectionRepository) Get(ctx context.Context, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSection, error) {
	q := r.client.ProjectTaskSection.Query()

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

func (r *projectTaskSectionRepository) List(ctx context.Context) ([]*model.ProjectTaskSection, error) {
	res, err := r.client.ProjectTaskSection.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTaskSectionRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskSectionWhereInput) (*model.ProjectTaskSectionConnection, error) {
	q := r.client.ProjectTaskSection.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTaskSectionFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectTaskSectionRepository) ListByTaskID(ctx context.Context, taskID model.ID, where *model.ProjectTaskSectionWhereInput) ([]*model.ProjectTaskSection, error) {
	projectTask, err := r.client.ProjectTask.Query().Where(projecttask.TaskID(taskID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	q := r.client.ProjectTaskSection.
		Query().Where(projecttasksection.ProjectID(projectTask.ProjectID))

	q, err = where.Filter(q)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	q.WithProject(func(pq *ent.ProjectQuery) {
		WithProject(pq)
	})

	projectTaskSections, err := q.All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return projectTaskSections, nil
}

func (r *projectTaskSectionRepository) Create(ctx context.Context, input model.CreateProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	res, err := r.client.
		ProjectTaskSection.
		Create().
		SetInput(input).
		SetName("").
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

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

func (r *projectTaskSectionRepository) Delete(ctx context.Context, input model.DeleteProjectTaskSectionInput) (*model.ProjectTaskSection, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.ProjectTaskSection.Query().Where(projecttasksection.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.ProjectTaskSection.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

func (r *projectTaskSectionRepository) DeleteAndKeepTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndKeepTasksInput) (*model.DeleteProjectTaskSectionAndKeepTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

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

	projectTasks, err := client.ProjectTask.Query().Where(projecttask.ProjectTaskSectionIDEQ(deleted.ID)).All(ctx)
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

func (r *projectTaskSectionRepository) DeleteAndDeleteTasks(ctx context.Context, input model.DeleteProjectTaskSectionAndDeleteTasksInput) (*model.DeleteProjectTaskSectionAndDeleteTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

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

	err = client.ProjectTaskSection.DeleteOneID(deleted.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	taskIDs := make([]model.ID, len(projectTasks))
	for i, t := range projectTasks {
		taskIDs[i] = t.TaskID
	}

	// TODO: Task repository can be called in usecase/usecase package
	taskRepo := taskRepository{
		client: r.client,
	}
	_, err = taskRepo.DeleteAll(ctx, model.DeleteAllTaskInput{
		TaskIDs:     taskIDs,
		WorkspaceID: input.WorkspaceID,
		RequestID:   "",
	})
	if err != nil {
		return nil, err
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

func (r *projectTaskSectionRepository) UndeleteAndKeepTasks(ctx context.Context, input model.UndeleteProjectTaskSectionAndKeepTasksInput) (*model.UndeleteProjectTaskSectionAndKeepTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

	createdProjectTaskSection, err := client.ProjectTaskSection.
		Create().
		SetProjectID(input.ProjectID).
		SetName(input.Name).
		SetCreatedAt(*input.CreatedAt).
		SetUpdatedAt(*input.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	projectTasks, err := client.ProjectTask.
		Query().
		Where(projecttask.IDIn(input.KeptProjectTaskIDs...)).
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.KeptProjectTaskIDs)
		}
		return nil, model.NewDBError(err)
	}
	bulk := make([]*ent.ProjectTaskCreate, len(projectTasks))
	for i, t := range projectTasks {
		bulk[i] = client.ProjectTask.Create().
			SetID(t.ID).
			SetTaskID(t.TaskID).
			SetProjectID(t.ProjectID).
			SetProjectTaskSectionID(createdProjectTaskSection.ID).
			SetCreatedAt(t.CreatedAt)
	}
	err = client.ProjectTask.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	projectTaskIDs := make([]model.ID, len(projectTasks))
	for i, t := range projectTasks {
		projectTaskIDs[i] = t.ID
	}

	return &model.UndeleteProjectTaskSectionAndKeepTasksPayload{
		ProjectTaskSection: createdProjectTaskSection,
		ProjectTaskIDs:     projectTaskIDs,
	}, nil
}

func (r *projectTaskSectionRepository) UndeleteAndDeleteTasks(ctx context.Context, input model.UndeleteProjectTaskSectionAndDeleteTasksInput) (*model.UndeleteProjectTaskSectionAndDeleteTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

	createdProjectTaskSection, err := client.ProjectTaskSection.
		Create().
		SetProjectID(input.ProjectID).
		SetName(input.Name).
		SetCreatedAt(*input.CreatedAt).
		SetUpdatedAt(*input.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	deletedTasks, err := client.DeletedTask.
		Query().
		Where(
			deletedtask.TaskIDIn(input.DeletedTaskIDs...),
			deletedtask.TaskTypeEQ(deletedtask.TaskTypeProject),
		).
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	if deletedTasks == nil {
		return nil, model.NewNotFoundError(err, input)
	}

	taskIDs := make([]model.ID, len(deletedTasks))
	for i, t := range deletedTasks {
		taskIDs[i] = t.TaskID
	}

	// TODO: Task repository can be called in usecase/usecase package
	taskRepo := taskRepository{
		client: r.client,
	}

	p, rerr := taskRepo.UndeleteAll(ctx, model.UndeleteAllTaskInput{
		TaskIDs:              taskIDs,
		WorkspaceID:          input.WorkspaceID,
		ProjectTaskSectionID: &createdProjectTaskSection.ID,
		RequestID:            "",
	})
	if rerr != nil {
		return nil, model.NewDBError(err)
	}

	bulk := make([]*ent.ProjectTaskCreate, len(p.ProjectTasks))
	for i, t := range p.ProjectTasks {
		bulk[i] = client.ProjectTask.Create().
			SetID(t.ID).
			SetCreatedAt(t.CreatedAt).
			SetUpdatedAt(t.UpdatedAt).
			SetTaskID(t.TaskID).
			SetProjectID(t.ProjectID).
			SetProjectTaskSectionID(createdProjectTaskSection.ID)
	}
	err = client.ProjectTask.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	projectTaskIDs := make([]model.ID, len(p.ProjectTasks))
	for i, t := range p.TeammateTasks {
		projectTaskIDs[i] = t.ID
	}
	ts, err := client.ProjectTask.Query().Where(projecttask.IDIn(projectTaskIDs...)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"IDs": projectTaskIDs,
			})
		}
		return nil, model.NewDBError(err)
	}

	createdProjectTasks := make([]*model.ProjectTask, len(ts))
	for i, t := range ts {
		createdProjectTasks[i] = t.Unwrap()
	}

	// Eager-loads associations with task.
	projectTaskSection, err := client.ProjectTaskSection.Query().
		WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
			WithProjectTask(ptq)
		}).
		Where(projecttasksection.ID(createdProjectTaskSection.ID)).
		Only(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.UndeleteProjectTaskSectionAndDeleteTasksPayload{
		ProjectTaskSection: projectTaskSection,
		ProjectTasks:       createdProjectTasks,
	}, nil
}
