package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedprojecttask"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/deletedtaskactivitytask"
	"project-management-demo-backend/ent/deletedteammatetask"
	"project-management-demo-backend/ent/deletedworkspaceactivitytask"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskactivitytask"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspaceactivitytask"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskRepository struct {
	client *ent.Client
}

// NewTaskRepository generates task repository.
func NewTaskRepository(client *ent.Client) ur.Task {
	return &taskRepository{client: client}
}

func (r *taskRepository) Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error) {
	q := r.client.Task.Query()

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

func (r *taskRepository) List(ctx context.Context) ([]*model.Task, error) {
	res, err := r.client.Task.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error) {
	q := r.client.Task.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskRepository) Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.
		Create().
		SetInput(input).
		SetName("").
		SetIsNew(true).
		SetDescription(model.DefaultEditorDescription()).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.UpdateOneID(input.ID).
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

func (r *taskRepository) Assign(ctx context.Context, input model.AssignTaskInput) (*model.AssignTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	updatedTask, err := client.Task.UpdateOneID(input.ID).SetAssigneeID(input.AssigneeID).Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	t, err := client.TeammateTask.Query().Where(teammatetask.TaskID(updatedTask.ID)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if t != nil {
		_, derr := client.TeammateTask.Delete().Where(teammatetask.ID(t.ID)).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}
	}

	assignedTeammateTaskSection, err := client.
		TeammateTaskSection.
		Query().
		Where(
			teammatetasksection.TeammateID(input.AssigneeID),
			teammatetasksection.Assigned(true),
		).
		Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	t, err = client.TeammateTask.Create().
		SetTaskID(input.ID).
		SetTeammateID(input.AssigneeID).
		SetTeammateTaskSectionID(assignedTeammateTaskSection.ID).
		SetWorkspaceID(input.WorkspaceID).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	teammateTask, err := client.TeammateTask.
		Query().
		WithTask(func(tq *ent.TaskQuery) {
			WithTask(tq)
		}).
		Where(teammatetask.ID(t.ID)).
		Only(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.AssignTaskPayload{
		Task:         updatedTask,
		TeammateTask: teammateTask.Unwrap(),
	}, nil
}

func (r *taskRepository) Unassign(ctx context.Context, input model.UnassignTaskInput) (*model.UnassignTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	updatedTask, err := client.Task.UpdateOneID(input.ID).ClearAssigneeID().Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	teammateTask, err := client.TeammateTask.Query().Where(teammatetask.TaskID(updatedTask.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	_, err = client.TeammateTask.Delete().Where(teammatetask.ID(teammateTask.ID)).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.UnassignTaskPayload{
		Task:           updatedTask,
		TeammateTaskID: teammateTask.ID,
	}, nil
}

func (r *taskRepository) Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteTaskPayload{
		ProjectTasks: []*model.ProjectTask{},
		TeammateTask: &model.TeammateTask{},
		DeletedTask:  &model.DeletedTask{},
	}

	taskRes, err := client.Task.Query().Where(task.ID(input.TaskID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	teammateTask, err := client.
		TeammateTask.Query().
		WithTask().
		Where(teammatetask.TaskID(input.TaskID)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if teammateTask != nil {
		err = client.TeammateTask.DeleteOneID(teammateTask.ID).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !taskRes.IsNew {
			_, derr := client.DeletedTeammateTask.Create().
				SetTeammateTaskID(teammateTask.ID).
				SetTaskID(teammateTask.TaskID).
				SetTeammateID(teammateTask.TeammateID).
				SetWorkspaceID(teammateTask.WorkspaceID).
				SetTeammateTaskSectionID(teammateTask.TeammateTaskSectionID).
				SetTeammateTaskCreatedAt(teammateTask.CreatedAt).
				SetTeammateTaskUpdatedAt(teammateTask.UpdatedAt).
				Save(ctx)

			if derr != nil {
				return nil, model.NewDBError(derr)
			}
		}
	}

	projectTasks, err := client.ProjectTask.
		Query().
		WithTask().
		Where(projecttask.TaskID(input.TaskID)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if len(projectTasks) > 0 {
		deletedProjectTaskIDs := make([]model.ID, len(projectTasks))
		for i, p := range projectTasks {
			deletedProjectTaskIDs[i] = p.ID
		}

		_, err = client.ProjectTask.Delete().Where(projecttask.IDIn(deletedProjectTaskIDs...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !taskRes.IsNew {
			bulk := make([]*ent.DeletedProjectTaskCreate, len(projectTasks))
			for i, p := range projectTasks {
				bulk[i] = client.DeletedProjectTask.Create().
					SetProjectTaskID(p.ID).
					SetProjectID(p.ProjectID).
					SetTaskID(p.TaskID).
					SetProjectTaskSectionID(p.ProjectTaskSectionID).
					SetProjectTaskCreatedAt(p.CreatedAt).
					SetProjectTaskUpdatedAt(p.UpdatedAt)
			}
			derr := client.DeletedProjectTask.CreateBulk(bulk...).Exec(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
		}
	}

	taskActivityTasks, err := client.TaskActivityTask.Query().Where(taskactivitytask.TaskID(input.TaskID)).All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(taskActivityTasks) > 0 {
		ids := make([]model.ID, len(taskActivityTasks))
		for i, p := range taskActivityTasks {
			ids[i] = p.ID
		}
		_, err = client.TaskActivityTask.Delete().Where(taskactivitytask.IDIn(ids...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !taskRes.IsNew {
			bulk := make([]*ent.DeletedTaskActivityTaskCreate, len(taskActivityTasks))
			for i, t := range taskActivityTasks {
				bulk[i] = client.DeletedTaskActivityTask.Create().
					SetTaskID(t.TaskID).
					SetTaskActivityTaskCreatedAt(t.CreatedAt).
					SetTaskActivityTaskUpdatedAt(t.UpdatedAt).
					SetTaskActivityID(t.TaskActivityID).
					SetTaskActivityTaskID(t.ID)
			}
			derr := client.DeletedTaskActivityTask.CreateBulk(bulk...).Exec(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
		}
	}

	workspaceActivityTasks, err := client.WorkspaceActivityTask.Query().Where(workspaceactivitytask.TaskID(input.TaskID)).All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(workspaceActivityTasks) > 0 {
		ids := make([]model.ID, len(workspaceActivityTasks))
		for i, p := range workspaceActivityTasks {
			ids[i] = p.ID
		}
		_, err = client.WorkspaceActivityTask.Delete().Where(workspaceactivitytask.IDIn(ids...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !taskRes.IsNew {
			bulk := make([]*ent.DeletedWorkspaceActivityTaskCreate, len(workspaceActivityTasks))
			for i, t := range workspaceActivityTasks {
				bulk[i] = client.DeletedWorkspaceActivityTask.Create().
					SetTaskID(t.TaskID).
					SetWorkspaceActivityTaskCreatedAt(t.CreatedAt).
					SetWorkspaceActivityTaskUpdatedAt(t.UpdatedAt).
					SetWorkspaceActivityID(t.WorkspaceActivityID).
					SetWorkspaceActivityTaskID(t.ID)
			}
			derr := client.DeletedWorkspaceActivityTask.CreateBulk(bulk...).Exec(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
		}
	}

	d, err := client.DeletedTask.Create().SetTaskID(taskRes.ID).SetWorkspaceID(input.WorkspaceID).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	deletedTask, err := client.DeletedTask.Query().WithTask(func(q *ent.TaskQuery) {
		WithTask(q)
	}).Where(deletedtask.ID(d.ID)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	if teammateTask != nil {
		payload.TeammateTask = teammateTask
	}
	if len(projectTasks) > 0 {
		payload.ProjectTasks = projectTasks
	}

	payload.DeletedTask = deletedTask

	return payload, nil
}

func (r *taskRepository) Undelete(ctx context.Context, input model.UndeleteTaskInput) (*model.UndeleteTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.UndeleteTaskPayload{
		TeammateTask: nil,
		ProjectTasks: []*model.ProjectTask{},
		DeletedTask:  &model.DeletedTask{},
	}

	deletedTask, err := client.DeletedTask.
		Query().
		WithTask().
		Where(deletedtask.TaskID(input.TaskID)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.TaskID)
		}
		return nil, model.NewDBError(err)
	}

	var undeletedTeammateTask *model.TeammateTask
	deletedTeammateTask, err := client.DeletedTeammateTask.
		Query().
		Where(deletedteammatetask.TaskID(input.TaskID)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if deletedTeammateTask != nil {
		t, terr := client.TeammateTask.Create().
			SetID(deletedTeammateTask.TeammateTaskID).
			SetWorkspaceID(deletedTeammateTask.WorkspaceID).
			SetTaskID(deletedTeammateTask.TaskID).
			SetTeammateID(deletedTeammateTask.TeammateID).
			SetTeammateTaskSectionID(deletedTeammateTask.TeammateTaskSectionID).
			SetCreatedAt(deletedTeammateTask.TeammateTaskCreatedAt).
			SetUpdatedAt(deletedTeammateTask.TeammateTaskUpdatedAt).
			Save(ctx)
		if terr != nil {
			return nil, model.NewDBError(terr)
		}

		terr = client.DeletedTeammateTask.DeleteOneID(deletedTeammateTask.ID).Exec(ctx)
		if terr != nil {
			return nil, model.NewDBError(terr)
		}

		undeletedTeammateTask = t
	}

	var undeletedProjectTasks []*model.ProjectTask
	deletedProjectTasks, err := client.DeletedProjectTask.
		Query().
		Where(deletedprojecttask.TaskID(input.TaskID)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(deletedProjectTasks) > 0 {
		bulk := make([]*ent.ProjectTaskCreate, len(deletedProjectTasks))
		for i, d := range deletedProjectTasks {
			bulk[i] = client.ProjectTask.Create().
				SetID(d.ProjectTaskID).
				SetProjectID(d.ProjectID).
				SetTaskID(d.TaskID).
				SetProjectTaskSectionID(d.ProjectTaskSectionID).
				SetCreatedAt(d.ProjectTaskCreatedAt).
				SetUpdatedAt(d.ProjectTaskUpdatedAt)
		}
		ps, derr := client.ProjectTask.CreateBulk(bulk...).Save(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		ids := make([]model.ID, len(deletedProjectTasks))
		for i, d := range deletedProjectTasks {
			ids[i] = d.ID
		}
		_, derr = client.DeletedProjectTask.Delete().Where(deletedprojecttask.IDIn(ids...)).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		undeletedProjectTasks = append(undeletedProjectTasks, ps...)
	}

	deletedTaskActivityTasks, err := client.DeletedTaskActivityTask.
		Query().
		Where(deletedtaskactivitytask.TaskID(input.TaskID)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(deletedTaskActivityTasks) > 0 {
		bulk := make([]*ent.TaskActivityTaskCreate, len(deletedTaskActivityTasks))
		for i, d := range deletedTaskActivityTasks {
			bulk[i] = client.TaskActivityTask.Create().
				SetID(d.TaskActivityTaskID).
				SetTaskID(d.TaskID).
				SetCreatedAt(d.TaskActivityTaskCreatedAt).
				SetUpdatedAt(d.TaskActivityTaskUpdatedAt).
				SetTaskActivityID(d.TaskActivityID)

		}
		_, derr := client.TaskActivityTask.CreateBulk(bulk...).Save(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		ids := make([]model.ID, len(deletedTaskActivityTasks))
		for i, d := range deletedTaskActivityTasks {
			ids[i] = d.ID
		}
		_, derr = client.DeletedTaskActivityTask.Delete().Where(deletedtaskactivitytask.IDIn(ids...)).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}
	}

	deletedWorkspaceActivityTasks, err := client.DeletedWorkspaceActivityTask.
		Query().
		Where(deletedworkspaceactivitytask.TaskID(input.TaskID)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(deletedWorkspaceActivityTasks) > 0 {
		bulk := make([]*ent.WorkspaceActivityTaskCreate, len(deletedWorkspaceActivityTasks))
		for i, d := range deletedWorkspaceActivityTasks {
			bulk[i] = client.WorkspaceActivityTask.Create().
				SetID(d.WorkspaceActivityTaskID).
				SetTaskID(d.TaskID).
				SetCreatedAt(d.WorkspaceActivityTaskCreatedAt).
				SetUpdatedAt(d.WorkspaceActivityTaskUpdatedAt).
				SetWorkspaceActivityID(d.WorkspaceActivityID)

		}
		_, derr := client.WorkspaceActivityTask.CreateBulk(bulk...).Save(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		ids := make([]model.ID, len(deletedWorkspaceActivityTasks))
		for i, d := range deletedWorkspaceActivityTasks {
			ids[i] = d.ID
		}
		_, derr = client.DeletedWorkspaceActivityTask.Delete().Where(deletedworkspaceactivitytask.IDIn(ids...)).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}
	}

	_, err = client.DeletedTask.
		Delete().
		Where(deletedtask.TaskID(input.TaskID)).
		Exec(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	if undeletedTeammateTask != nil {
		// Restore the state in order to query en entity after successful transaction.
		t, terr := client.TeammateTask.
			Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).
			Where(teammatetask.ID(undeletedTeammateTask.ID)).Only(ctx)

		if terr != nil {
			return nil, model.NewDBError(err)
		}
		payload.TeammateTask = t.Unwrap()
	}
	if undeletedProjectTasks != nil {
		ids := make([]model.ID, len(undeletedProjectTasks))
		for i, p := range undeletedProjectTasks {
			ids[i] = p.ID
		}

		// Restore the state in order to query en entity after successful transaction.
		q := client.ProjectTask.
			Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			})

		WithProjectTask(q)

		ps, terr := q.Where(projecttask.IDIn(ids...)).All(ctx)
		if terr != nil {
			return nil, model.NewDBError(err)
		}
		for _, p := range ps {
			payload.ProjectTasks = append(payload.ProjectTasks, p.Unwrap())
		}
	}

	payload.DeletedTask = deletedTask

	return payload, nil
}

func (r *taskRepository) DeleteAll(ctx context.Context, input model.DeleteAllTaskInput) (*model.DeleteAllTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteAllTaskPayload{
		TeammateTasks: []*model.TeammateTask{},
		ProjectTasks:  []*model.ProjectTask{},
		DeletedTasks:  []*model.DeletedTask{},
	}

	tasks, err := client.Task.Query().Where(task.IDIn(input.TaskIDs...)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
	}
	taskIDs := make([]ulid.ID, len(tasks))
	for i, t := range tasks {
		taskIDs[i] = t.ID
	}

	teammateTasks, err := client.
		TeammateTask.Query().
		WithTask().
		Where(teammatetask.TaskIDIn(taskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if len(teammateTasks) > 0 {
		teammateTaskDeletedTaskIDs := make([]ulid.ID, len(teammateTasks))
		for i, t := range teammateTasks {
			teammateTaskDeletedTaskIDs[i] = t.ID
		}
		_, err = client.TeammateTask.Delete().Where(teammatetask.IDIn(teammateTaskDeletedTaskIDs...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}

		bulk := make([]*ent.DeletedTeammateTaskCreate, len(teammateTasks))
		for i, t := range teammateTasks {
			bulk[i] = client.DeletedTeammateTask.Create().
				SetTeammateTaskID(t.ID).
				SetTaskID(t.TaskID).
				SetTeammateID(t.TeammateID).
				SetWorkspaceID(t.WorkspaceID).
				SetTeammateTaskSectionID(t.TeammateTaskSectionID).
				SetTeammateTaskCreatedAt(t.CreatedAt).
				SetTeammateTaskUpdatedAt(t.UpdatedAt)
		}
		derr := client.DeletedTeammateTask.CreateBulk(bulk...).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}
	}

	deletedProjectTasks, err := client.
		ProjectTask.Query().
		WithTask().
		Where(projecttask.TaskIDIn(taskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if deletedProjectTasks != nil {
		deletedProjectTaskDeletedTaskIDs := make([]ulid.ID, len(deletedProjectTasks))
		for i, t := range deletedProjectTasks {
			deletedProjectTaskDeletedTaskIDs[i] = t.ID
		}
		_, err = client.ProjectTask.Delete().Where(projecttask.IDIn(deletedProjectTaskDeletedTaskIDs...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}

		bulk := make([]*ent.DeletedProjectTaskCreate, len(deletedProjectTasks))
		for i, d := range deletedProjectTasks {
			bulk[i] = client.DeletedProjectTask.Create().
				SetProjectTaskID(d.ID).
				SetProjectID(d.ProjectID).
				SetTaskID(d.TaskID).
				SetProjectTaskSectionID(d.ProjectTaskSectionID).
				SetProjectTaskCreatedAt(d.CreatedAt).
				SetProjectTaskUpdatedAt(d.UpdatedAt)
		}
		derr := client.DeletedProjectTask.CreateBulk(bulk...).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}
	}

	bulk := make([]*ent.DeletedTaskCreate, len(taskIDs))
	for i, id := range taskIDs {
		bulk[i] = client.DeletedTask.Create().
			SetTaskID(id).
			SetWorkspaceID(input.WorkspaceID)
	}
	createdDeletedTasks, err := client.DeletedTask.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	if teammateTasks != nil {
		payload.TeammateTasks = teammateTasks
	}
	if deletedProjectTasks != nil {
		payload.ProjectTasks = deletedProjectTasks
	}
	payload.DeletedTasks = createdDeletedTasks

	return payload, nil
}

func (r *taskRepository) UndeleteAll(ctx context.Context, input model.UndeleteAllTaskInput) (*model.UndeleteAllTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.UndeleteAllTaskPayload{
		TeammateTasks: []*model.TeammateTask{},
		ProjectTasks:  []*model.ProjectTask{},
		DeletedTasks:  []*model.DeletedTask{},
	}

	taskIDs := input.TaskIDs

	var createdTeammateTasks []*model.TeammateTask
	deletedTeammateTasks, err := client.DeletedTeammateTask.
		Query().
		Where(deletedteammatetask.TaskIDIn(taskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(deletedTeammateTasks) > 0 {
		bulk := make([]*ent.TeammateTaskCreate, len(deletedTeammateTasks))
		for i, d := range deletedTeammateTasks {
			bulk[i] = client.TeammateTask.Create().
				SetID(d.TeammateTaskID).
				SetTaskID(d.TaskID).
				SetTeammateID(d.TeammateID).
				SetWorkspaceID(d.WorkspaceID).
				SetCreatedAt(d.TeammateTaskCreatedAt).
				SetUpdatedAt(d.TeammateTaskUpdatedAt)

			if input.TeammateTaskSectionID != nil {
				bulk[i].SetTeammateTaskSectionID(*input.TeammateTaskSectionID)
			} else {
				bulk[i].SetTeammateTaskSectionID(d.TeammateTaskSectionID)
			}
		}
		ts, terr := client.TeammateTask.CreateBulk(bulk...).Save(ctx)
		if terr != nil {
			return nil, model.NewDBError(terr)
		}
		for _, t := range ts {
			createdTeammateTasks = append(createdTeammateTasks, t.Unwrap())
		}

		ids := make([]model.ID, len(deletedTeammateTasks))
		for i, d := range deletedTeammateTasks {
			ids[i] = d.ID
		}
		_, terr = client.DeletedTeammateTask.Delete().Where(deletedteammatetask.IDIn(ids...)).Exec(ctx)
		if terr != nil {
			return nil, model.NewDBError(terr)
		}
	}

	var createdProjectTasks []*model.ProjectTask
	deletedProjectTasks, err := client.DeletedProjectTask.
		Query().
		Where(deletedprojecttask.TaskIDIn(taskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}
	if len(deletedProjectTasks) > 0 {
		bulk := make([]*ent.ProjectTaskCreate, len(deletedProjectTasks))
		for i, d := range deletedProjectTasks {
			bulk[i] = client.ProjectTask.Create().
				SetID(d.ProjectTaskID).
				SetTaskID(d.TaskID).
				SetProjectID(d.ProjectID).
				SetCreatedAt(d.ProjectTaskCreatedAt).
				SetUpdatedAt(d.ProjectTaskUpdatedAt)

			if input.ProjectTaskSectionID != nil {
				bulk[i].SetProjectTaskSectionID(*input.ProjectTaskSectionID)
			} else {
				bulk[i].SetProjectTaskSectionID(d.ProjectTaskSectionID)
			}
		}
		ps, perr := client.ProjectTask.CreateBulk(bulk...).Save(ctx)
		if perr != nil {
			return nil, model.NewDBError(perr)
		}
		for _, t := range ps {
			createdProjectTasks = append(createdProjectTasks, t.Unwrap())
		}

		ids := make([]model.ID, len(deletedProjectTasks))
		for i, d := range deletedProjectTasks {
			ids[i] = d.ID
		}
		_, perr = client.DeletedProjectTask.Delete().Where(deletedprojecttask.IDIn(ids...)).Exec(ctx)
		if perr != nil {
			return nil, model.NewDBError(perr)
		}
	}

	deletedTasks, err := client.DeletedTask.
		Query().
		Where(deletedtask.TaskIDIn(taskIDs...)).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, taskIDs)
		}
		return nil, model.NewDBError(err)
	}

	deletedTaskIds := make([]model.ID, len(deletedTasks))
	for i, d := range deletedTasks {
		deletedTaskIds[i] = d.ID
	}

	_, err = client.DeletedTask.Delete().Where(deletedtask.IDIn(deletedTaskIds...)).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	if createdTeammateTasks != nil {
		payload.TeammateTasks = createdTeammateTasks
	}

	if createdProjectTasks != nil {
		payload.ProjectTasks = createdProjectTasks
	}

	payload.DeletedTasks = deletedTasks

	return payload, nil
}

// WithTask eager-loads associations with task entity.
func WithTask(query *ent.TaskQuery) {
	query.WithTaskFeeds()
	query.WithTaskFiles(func(tfq *ent.TaskFileQuery) {
		WithTaskFiles(tfq)
	})
	query.WithTaskFeedLikes()
	query.WithTaskPriority(func(tpq *ent.TaskPriorityQuery) {
		WithTaskPriority(tpq)
	})
	query.WithSubTasks(func(subTaskQuery *ent.TaskQuery) {
		WithSubTask(subTaskQuery)
	})
	query.WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
		ptq.WithProject(func(pq *ent.ProjectQuery) {
			WithProject(pq)
		})
	})
	query.WithTaskTags(func(ttq *ent.TaskTagQuery) {
		WithTaskTag(ttq)
	})
	query.WithTaskLikes()
	query.WithTaskCollaborators(func(tcq *ent.TaskCollaboratorQuery) {
		WithTaskCollaborator(tcq)
	})
	query.WithParentTask(func(parentTaskQuery *ent.TaskQuery) {
		WithParentTask(parentTaskQuery)
	})
}

// WithParentTask eager-loads associations with task parent entity.
func WithParentTask(query *ent.TaskQuery) {
	query.WithTaskFeeds()
	query.WithTaskFiles(func(tfq *ent.TaskFileQuery) {
		WithTaskFiles(tfq)
	})
	query.WithTaskFeedLikes()
	query.WithTaskPriority(func(tpq *ent.TaskPriorityQuery) {
		WithTaskPriority(tpq)
	})
	query.WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
		ptq.WithProject(func(pq *ent.ProjectQuery) {
			WithProject(pq)
		})
	})
	query.WithTaskTags(func(ttq *ent.TaskTagQuery) {
		WithTaskTag(ttq)
	})
	query.WithTaskLikes()
	query.WithTaskCollaborators(func(tcq *ent.TaskCollaboratorQuery) {
		WithTaskCollaborator(tcq)
	})
}

// WithSubTask eager-loads associations with subtask entity.
func WithSubTask(query *ent.TaskQuery) {
	query.WithTaskFeeds()
	query.WithTaskFiles(func(tfq *ent.TaskFileQuery) {
		WithTaskFiles(tfq)
	})
	query.WithTaskFeedLikes()
	query.WithTaskPriority(func(tpq *ent.TaskPriorityQuery) {
		WithTaskPriority(tpq)
	})
	query.WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
		ptq.WithProject(func(pq *ent.ProjectQuery) {
			WithProject(pq)
		})
	})
	query.WithTaskTags(func(ttq *ent.TaskTagQuery) {
		WithTaskTag(ttq)
	})
	query.WithTaskLikes()
	query.WithTaskCollaborators(func(tcq *ent.TaskCollaboratorQuery) {
		WithTaskCollaborator(tcq)
	})
}
