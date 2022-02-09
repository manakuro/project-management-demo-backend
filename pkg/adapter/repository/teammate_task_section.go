package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
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

func (r *teammateTaskSectionRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskSectionWhereInput, requestedFields []string) (*model.TeammateTaskSectionConnection, error) {
	q := r.client.TeammateTaskSection.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	if collection.Contains(requestedFields, "edges.node.teammateTasks") {
		q.WithTeammateTasks(func(teammateTaskQuery *ent.TeammateTaskQuery) {
			teammateTaskQuery.WithTask(func(taskQuery *ent.TaskQuery) {
				WithTask(taskQuery, WithTaskOptions{
					SubTasks:          true,
					TaskFiles:         true,
					TaskFeeds:         true,
					TaskCollaborators: true,
					TaskTags:          true,
					ProjectTasks:      true,
					TaskPriority:      true,
				})
			})
		})
	}

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
	deleted, err := r.client.TeammateTaskSection.Query().Where(teammatetasksection.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TeammateTaskSection.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}
