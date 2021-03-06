package teammatetaskrepository

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/pkg/adapter/repository/repositoryutil"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
	"time"
)

func (r *teammateTaskRepository) Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error) {
	start := time.Now()
	q := r.client.TeammateTask.Query()

	q, err := where.Filter(q)

	// Eager-loading with task explicitly.
	q.WithTask(func(query *ent.TaskQuery) {
		repositoryutil.WithTask(query)
	})

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

	fmt.Println("\n\n========================================================================")
	fmt.Println("duration: ", time.Since(start).String())
	fmt.Print("========================================================================\n\n")

	return res, nil
}

func (r *teammateTaskRepository) List(ctx context.Context) ([]*model.TeammateTask, error) {
	res, err := r.client.TeammateTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskRepository) TasksDueSoon(ctx context.Context, workspaceID model.ID, teammateID model.ID) ([]*model.TeammateTask, error) {
	start := time.Now()

	q := r.client.TeammateTask.Query()

	q.Where(teammatetask.TeammateIDEQ(teammateID))
	q.Where(teammatetask.HasTaskWith(
		task.DueDateGTE(datetime.StartOfDay(datetime.Now())),
		task.DueDateLTE(datetime.EndOfDay(datetime.AddDate(5))),
	))
	q.Where(teammatetask.HasTeammateTaskSectionWith(
		teammatetasksection.WorkspaceIDEQ(workspaceID),
	))
	q.WithTask(func(q *ent.TaskQuery) {
		q.WithProjectTasks(func(ptq *ent.ProjectTaskQuery) {
			repositoryutil.WithProjectTasks(ptq)
		})
	})

	res, err := q.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	fmt.Println("\n\n========================================================================")
	fmt.Println("duration: ", time.Since(start).String())
	fmt.Print("========================================================================\n\n")
	return res, err
}

func (r *teammateTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput) (*model.TeammateTaskConnection, error) {
	q := r.client.TeammateTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}
