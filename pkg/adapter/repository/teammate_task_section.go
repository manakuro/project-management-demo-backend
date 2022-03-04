package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type teammateTaskSectionRepository struct {
	client *ent.Client
}

// NewTeammateTaskSectionRepository generates teammateTaskSection repository
func NewTeammateTaskSectionRepository(client *ent.Client) ur.TeammateTaskSection {
	return &teammateTaskSectionRepository{client: client}
}

func (r *teammateTaskSectionRepository) Get(ctx context.Context, where *model.TeammateTaskSectionWhereInput) (*model.TeammateTaskSection, error) {
	q := r.client.TeammateTaskSection.Query()

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

func (r *teammateTaskSectionRepository) List(ctx context.Context) ([]*model.TeammateTaskSection, error) {
	res, err := r.client.TeammateTaskSection.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskSectionRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskSectionWhereInput) (*model.TeammateTaskSectionConnection, error) {
	q := r.client.TeammateTaskSection.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskSectionFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskSectionRepository) Create(ctx context.Context, input model.CreateTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	res, err := r.client.
		TeammateTaskSection.
		Create().
		SetInput(input).
		SetName("").
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskSectionRepository) Update(ctx context.Context, input model.UpdateTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	res, err := r.client.
		TeammateTaskSection.UpdateOneID(input.ID).
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

func (r *teammateTaskSectionRepository) Delete(ctx context.Context, input model.DeleteTeammateTaskSectionInput) (*model.TeammateTaskSection, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.TeammateTaskSection.
		Query().
		Where(teammatetasksection.IDEQ(input.ID)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.TeammateTaskSection.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

func (r *teammateTaskSectionRepository) DeleteAndKeepTasks(ctx context.Context, input model.DeleteTeammateTaskSectionAndKeepTasksInput) (*model.DeleteTeammateTaskSectionAndKeepTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.TeammateTaskSection.
		Query().
		Where(teammatetasksection.IDEQ(input.ID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	teammateTasks, err := client.TeammateTask.Query().Where(teammatetask.TeammateTaskSectionIDEQ(deleted.ID)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, deleted.ID)
		}
		return nil, model.NewDBError(err)
	}

	teammateTaskSection, err := client.
		TeammateTaskSection.
		Query().
		Where(
			teammatetasksection.TeammateID(deleted.TeammateID),
			teammatetasksection.WorkspaceID(deleted.WorkspaceID),
			teammatetasksection.Assigned(true),
		).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"TeammateID":  deleted.TeammateID,
				"WorkspaceID": deleted.WorkspaceID,
			})
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
			SetTeammateTaskSectionID(teammateTaskSection.ID).
			SetCreatedAt(t.CreatedAt)
	}
	err = client.TeammateTask.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	err = client.TeammateTaskSection.DeleteOneID(deleted.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	teammateTaskIDs := make([]ulid.ID, len(teammateTasks))
	for i, task := range teammateTasks {
		teammateTaskIDs[i] = task.ID
	}

	return &model.DeleteTeammateTaskSectionAndKeepTasksPayload{
		TeammateTaskSection:     deleted,
		KeptTeammateTaskSection: teammateTaskSection,
		TeammateTaskIDs:         teammateTaskIDs,
	}, nil
}

func (r *teammateTaskSectionRepository) DeleteAndDeleteTasks(ctx context.Context, input model.DeleteTeammateTaskSectionAndDeleteTasksInput) (*model.DeleteTeammateTaskSectionAndDeleteTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.TeammateTaskSection.
		Query().
		Where(teammatetasksection.IDEQ(input.ID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	teammateTasks, err := client.TeammateTask.
		Query().
		Where(teammatetask.TeammateTaskSectionIDEQ(deleted.ID)).
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, deleted.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.TeammateTaskSection.DeleteOneID(deleted.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	taskIDs := make([]ulid.ID, len(teammateTasks))
	for i, t := range teammateTasks {
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

	teammateTaskIDs := make([]ulid.ID, len(teammateTasks))
	for i, t := range teammateTasks {
		teammateTaskIDs[i] = t.ID
	}

	return &model.DeleteTeammateTaskSectionAndDeleteTasksPayload{
		TeammateTaskSection: deleted,
		TeammateTaskIDs:     teammateTaskIDs,
	}, nil
}

func (r *teammateTaskSectionRepository) UndeleteAndKeepTasks(ctx context.Context, input model.UndeleteTeammateTaskSectionAndKeepTasksInput) (*model.UndeleteTeammateTaskSectionAndKeepTasksPayload, error) {
	client := WithTransactionalMutation(ctx)

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
	teammateTaskIDs := make([]ulid.ID, len(teammateTasks))
	for i, task := range teammateTasks {
		teammateTaskIDs[i] = task.ID
	}

	return &model.UndeleteTeammateTaskSectionAndKeepTasksPayload{
		TeammateTaskSection: createdTeammateTaskSection,
		TeammateTaskIDs:     teammateTaskIDs,
	}, nil
}
