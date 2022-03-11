package repository

import (
	"context"
	"errors"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
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
		return nil, model.NewValidationError(errors.New("this task has already been assigned"))
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

	teammateTask, err := client.TeammateTask.Query().WithTask(func(tq *ent.TaskQuery) {
		WithTask(tq)
	}).Where(teammatetask.ID(t.ID)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.AssignTaskPayload{
		Task:         updatedTask,
		TeammateTask: teammateTask,
	}, nil
}

func (r *taskRepository) Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteTaskPayload{
		ProjectTask:  &model.ProjectTask{},
		TeammateTask: &model.TeammateTask{},
		DeletedTasks: []*model.DeletedTask{},
	}

	deletedTask, err := client.Task.Query().Where(task.ID(input.TaskID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	deletedTeammateTask, err := client.
		TeammateTask.Query().
		WithTask().
		Where(teammatetask.TaskID(input.TaskID)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedTeammateTaskDeletedTask *model.DeletedTask
	if deletedTeammateTask != nil {
		err = client.TeammateTask.DeleteOneID(deletedTeammateTask.ID).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !deletedTask.IsNew {
			d, derr := client.DeletedTask.
				Create().
				SetTaskID(input.TaskID).
				SetWorkspaceID(input.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeTeammate).
				SetTaskJoinID(deletedTeammateTask.TeammateID).
				SetTaskSectionID(deletedTeammateTask.TeammateTaskSectionID).
				Save(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			// Restore the state in order to get an entity after a successful transaction.
			d, derr = client.DeletedTask.Query().WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).Where(deletedtask.ID(d.ID)).Only(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			deletedTeammateTaskDeletedTask = d.Unwrap()
		}
	}

	deletedProjectTask, err := client.ProjectTask.
		Query().
		WithTask().
		Where(projecttask.TaskID(input.TaskID)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedProjectTaskDeletedTask *model.DeletedTask
	if deletedProjectTask != nil {
		err = client.ProjectTask.DeleteOneID(deletedProjectTask.ID).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		if !deletedTask.IsNew {
			d, derr := client.DeletedTask.
				Create().
				SetTaskID(input.TaskID).
				SetWorkspaceID(input.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeProject).
				SetTaskJoinID(deletedProjectTask.ProjectID).
				SetTaskSectionID(deletedProjectTask.ProjectTaskSectionID).
				Save(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			// Restore the state in order to get an entity after a successful transaction.
			d, derr = client.DeletedTask.Query().WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).Where(deletedtask.ID(d.ID)).Only(ctx)
			if derr != nil {
				return nil, model.NewDBError(derr)
			}
			deletedProjectTaskDeletedTask = d.Unwrap()
		}
	}

	if deletedTeammateTask != nil {
		payload.TeammateTask = deletedTeammateTask
	}
	if deletedProjectTask != nil {
		payload.ProjectTask = deletedProjectTask
	}
	if deletedTeammateTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedTeammateTaskDeletedTask)
	}
	if deletedProjectTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedProjectTaskDeletedTask)
	}

	return payload, nil
}

func (r *taskRepository) Undelete(ctx context.Context, input model.UndeleteTaskInput) (*model.UndeleteTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.UndeleteTaskPayload{
		TeammateTask: nil,
		ProjectTask:  nil,
		DeletedTasks: []*model.DeletedTask{},
	}

	deletedTasks, err := client.DeletedTask.
		Query().
		WithTask().
		Where(deletedtask.TaskID(input.TaskID)).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.TaskID)
		}
		return nil, model.NewDBError(err)
	}

	// The task to be undeleted will be limited up to two records.
	var undeletedTeammateTask *model.TeammateTask
	var undeletedProjectTask *model.ProjectTask
	for _, t := range deletedTasks {
		if t.TaskType == deletedtask.TaskTypeTeammate {
			undeletedTeammateTask, err = client.TeammateTask.Create().
				SetWorkspaceID(t.WorkspaceID).
				SetTaskID(t.TaskID).
				SetTeammateID(t.TaskJoinID).
				SetTeammateTaskSectionID(t.TaskSectionID).
				Save(ctx)
			if err != nil {
				return nil, model.NewDBError(err)
			}
		}
		if t.TaskType == deletedtask.TaskTypeProject {
			undeletedProjectTask, err = client.ProjectTask.Create().
				SetProjectID(t.TaskJoinID).
				SetTaskID(t.TaskID).
				SetProjectTaskSectionID(t.TaskSectionID).
				Save(ctx)
			if err != nil {
				return nil, model.NewDBError(err)
			}
		}
	}

	deletedIds := make([]model.ID, len(deletedTasks))
	for i, t := range deletedTasks {
		deletedIds[i] = t.ID
	}

	_, err = client.DeletedTask.
		Delete().
		Where(deletedtask.IDIn(deletedIds...)).
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
	if undeletedProjectTask != nil {
		// Restore the state in order to query en entity after successful transaction.
		t, terr := client.ProjectTask.
			Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).
			Where(projecttask.ID(undeletedProjectTask.ID)).Only(ctx)
		if terr != nil {
			return nil, model.NewDBError(err)
		}
		payload.ProjectTask = t.Unwrap()
	}

	payload.DeletedTasks = deletedTasks

	return payload, nil
}

func (r *taskRepository) DeleteAll(ctx context.Context, input model.DeleteAllTaskInput) (*model.DeleteAllTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteAllTaskPayload{
		TeammateTasks: []*model.TeammateTask{},
		ProjectTasks:  []*model.ProjectTask{},
		DeletedTasks:  []*model.DeletedTask{},
	}

	deletedTasks, err := client.Task.Query().Where(task.IDIn(input.TaskIDs...)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
	}
	deletedTaskIDs := make([]ulid.ID, len(deletedTasks))
	for i, t := range deletedTasks {
		deletedTaskIDs[i] = t.ID
	}

	deletedTeammateTasks, err := client.
		TeammateTask.Query().
		WithTask().
		Where(teammatetask.TaskIDIn(deletedTaskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedTeammateTaskDeletedTask []*model.DeletedTask
	if deletedTeammateTasks != nil {
		deletedTeammateTaskDeletedTaskIDs := make([]ulid.ID, len(deletedTeammateTasks))
		for i, t := range deletedTeammateTasks {
			deletedTeammateTaskDeletedTaskIDs[i] = t.ID
		}
		_, err = client.TeammateTask.Delete().Where(teammatetask.IDIn(deletedTeammateTaskDeletedTaskIDs...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}

		bulk := make([]*ent.DeletedTaskCreate, len(deletedTeammateTasks))
		for i, t := range deletedTeammateTasks {
			bulk[i] = client.DeletedTask.Create().
				SetTaskID(t.TaskID).
				SetWorkspaceID(t.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeTeammate).
				SetTaskJoinID(t.TeammateID).
				SetTaskSectionID(t.TeammateTaskSectionID)
		}
		derr := client.DeletedTask.CreateBulk(bulk...).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		ds, derr := client.DeletedTask.Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).
			Where(
				deletedtask.TaskIDIn(deletedTaskIDs...),
				deletedtask.TaskTypeEQ(deletedtask.TaskTypeTeammate),
			).All(ctx)
		if derr != nil {
			return nil, model.NewDBError(err)
		}

		for _, d := range ds {
			deletedTeammateTaskDeletedTask = append(deletedTeammateTaskDeletedTask, d.Unwrap())
		}
	}

	deletedProjectTasks, err := client.
		ProjectTask.Query().
		WithTask().
		Where(projecttask.TaskIDIn(deletedTaskIDs...)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	var deletedProjectTaskDeletedTask []*model.DeletedTask
	if deletedProjectTasks != nil {
		deletedProjectTaskDeletedTaskIDs := make([]ulid.ID, len(deletedProjectTasks))
		for i, t := range deletedProjectTasks {
			deletedProjectTaskDeletedTaskIDs[i] = t.ID
		}
		_, err = client.ProjectTask.Delete().Where(projecttask.IDIn(deletedProjectTaskDeletedTaskIDs...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}

		bulk := make([]*ent.DeletedTaskCreate, len(deletedProjectTasks))
		for i, t := range deletedProjectTasks {
			bulk[i] = client.DeletedTask.Create().
				SetTaskID(t.TaskID).
				SetWorkspaceID(input.WorkspaceID).
				SetTaskType(deletedtask.TaskTypeProject).
				SetTaskJoinID(t.ProjectID).
				SetTaskSectionID(t.ProjectTaskSectionID)
		}
		derr := client.DeletedTask.CreateBulk(bulk...).Exec(ctx)
		if derr != nil {
			return nil, model.NewDBError(derr)
		}

		ds, derr := client.DeletedTask.Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).
			Where(
				deletedtask.IDIn(deletedTaskIDs...),
				deletedtask.TaskTypeEQ(deletedtask.TaskTypeProject),
			).
			All(ctx)
		if derr != nil {
			return nil, model.NewDBError(err)
		}

		for _, d := range ds {
			deletedProjectTaskDeletedTask = append(deletedProjectTaskDeletedTask, d.Unwrap())
		}
	}

	if deletedTeammateTasks != nil {
		payload.TeammateTasks = deletedTeammateTasks
	}
	if deletedProjectTasks != nil {
		payload.ProjectTasks = deletedProjectTasks
	}
	if deletedTeammateTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedTeammateTaskDeletedTask...)
	}
	if deletedProjectTaskDeletedTask != nil {
		payload.DeletedTasks = append(payload.DeletedTasks, deletedProjectTaskDeletedTask...)
	}

	return payload, nil
}

func (r *taskRepository) UndeleteAll(ctx context.Context, input model.UndeleteAllTaskInput) (*model.UndeleteAllTaskPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.UndeleteAllTaskPayload{
		TeammateTasks: []*model.TeammateTask{},
		ProjectTasks:  []*model.ProjectTask{},
		DeletedTasks:  []*model.DeletedTask{},
	}

	taskIds := input.TaskIDs

	deletedTasks, err := client.DeletedTask.
		Query().
		Where(deletedtask.TaskIDIn(taskIds...)).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, taskIds)
		}
		return nil, model.NewDBError(err)
	}

	var teammateTaskCreates []*ent.TeammateTaskCreate
	var projectTaskCreates []*ent.ProjectTaskCreate
	for _, t := range deletedTasks {
		if t.TaskType == deletedtask.TaskTypeTeammate {
			c := client.TeammateTask.Create().
				SetWorkspaceID(t.WorkspaceID).
				SetTaskID(t.TaskID).
				SetTeammateID(t.TaskJoinID)

			if input.TeammateTaskSectionID != nil {
				c.SetTeammateTaskSectionID(*input.TeammateTaskSectionID)
			} else {
				c.SetTeammateTaskSectionID(t.TaskSectionID)
			}

			teammateTaskCreates = append(teammateTaskCreates, c)
		}
		if t.TaskType == deletedtask.TaskTypeProject {
			c := client.ProjectTask.Create().
				SetProjectID(t.TaskJoinID).
				SetTaskID(t.TaskID)
			if input.ProjectTaskSectionID != nil {
				c.SetProjectTaskSectionID(*input.ProjectTaskSectionID)
			} else {
				c.SetProjectTaskSectionID(t.TaskSectionID)
			}

			projectTaskCreates = append(projectTaskCreates, c)
		}
	}

	var createdTeammateTasks []*model.TeammateTask
	if teammateTaskCreates != nil {
		ts, terr := client.TeammateTask.CreateBulk(teammateTaskCreates...).Save(ctx)
		if terr != nil {
			return nil, model.NewDBError(terr)
		}

		if terr != nil {
			if ent.IsNotFound(terr) {
				return nil, model.NewNotFoundError(terr, map[string]interface{}{
					"TaskIDs": taskIds,
				})
			}
			return nil, model.NewDBError(terr)
		}

		for _, t := range ts {
			createdTeammateTasks = append(createdTeammateTasks, t.Unwrap())
		}
	}

	var createdProjectTasks []*model.ProjectTask
	if projectTaskCreates != nil {
		perr := client.ProjectTask.CreateBulk(projectTaskCreates...).Exec(ctx)
		if perr != nil {
			return nil, model.NewDBError(perr)
		}
		ts, terr := client.ProjectTask.
			Query().
			WithTask(func(tq *ent.TaskQuery) {
				WithTask(tq)
			}).
			Where(projecttask.TaskIDIn(taskIds...)).
			All(ctx)
		if terr != nil {
			if ent.IsNotFound(terr) {
				return nil, model.NewNotFoundError(terr, map[string]interface{}{
					"TaskIDs": taskIds,
				})
			}
			return nil, model.NewDBError(terr)
		}
		for _, t := range ts {
			createdProjectTasks = append(createdProjectTasks, t.Unwrap())
		}
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
	query.WithTaskFiles()
	query.WithTaskFeedLikes()
	query.WithTaskPriority()
	query.WithSubTasks()
	query.WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
		ptq.WithProject()
	})
	query.WithTaskTags()
	query.WithTaskLikes()
	query.WithTaskCollaborators()
}
