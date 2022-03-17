package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type projectTaskRepository struct {
	client *ent.Client
}

// NewProjectTaskRepository generates projectTask repository
func NewProjectTaskRepository(client *ent.Client) ur.ProjectTask {
	return &projectTaskRepository{client: client}
}

func (r *projectTaskRepository) Get(ctx context.Context, where *model.ProjectTaskWhereInput) (*model.ProjectTask, error) {
	q := r.client.ProjectTask.Query()

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

func (r *projectTaskRepository) List(ctx context.Context) ([]*model.ProjectTask, error) {
	res, err := r.client.ProjectTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskWhereInput) (*model.ProjectTaskConnection, error) {
	q := r.client.ProjectTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectTaskRepository) Create(ctx context.Context, input model.CreateProjectTaskInput) (*model.ProjectTask, error) {
	tc := r.client.Task.
		Create().
		SetIsNew(true).
		SetName("").
		SetCreatedBy(input.CreatedBy).
		SetDescription(model.DefaultEditorDescription())

	if input.TaskParentID != nil {
		tc.SetTaskParentID(*input.TaskParentID)
	}

	newTask, err := tc.Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	_, err = r.client.TaskFeed.Create().
		SetTask(newTask).
		SetIsFirst(true).
		SetDescription(model.DefaultEditorDescription()).
		SetTeammateID(input.CreatedBy).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	res, err := r.client.
		ProjectTask.
		Create().
		SetInput(input).
		SetTask(newTask).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTaskRepository) CreateByTaskID(ctx context.Context, input model.CreateProjectTaskByTaskIDInput) (*model.ProjectTask, error) {
	client := WithTransactionalMutation(ctx)

	ps, err := client.ProjectTaskSection.
		Query().
		Where(projecttasksection.ProjectID(input.ProjectID)).
		All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	projectTaskSection := ps[0]

	p, err := client.ProjectTask.
		Create().
		SetTaskID(input.TaskID).
		SetProjectID(input.ProjectID).
		SetProjectTaskSectionID(projectTaskSection.ID).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return p, nil
}

func (r *projectTaskRepository) Update(ctx context.Context, input model.UpdateProjectTaskInput) (*model.ProjectTask, error) {
	res, err := r.client.
		ProjectTask.UpdateOneID(input.ID).
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

func (r *projectTaskRepository) Delete(ctx context.Context, input model.DeleteProjectTaskInput) (*model.ProjectTask, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.ProjectTask.
		Query().
		Where(projecttask.IDEQ(input.ID)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.ProjectTask.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

// WithProjectTask eager-loads association with project task.
func WithProjectTask(query *ent.ProjectTaskQuery) {
	query.WithTask(func(tq *ent.TaskQuery) {
		WithTask(tq)
	})
	query.WithProject(func(pq *ent.ProjectQuery) {
		WithProject(pq)
	})
}
