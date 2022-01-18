package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type workspaceRepository struct {
	client *ent.Client
}

// NewWorkspaceRepository generates workspace repository
func NewWorkspaceRepository(client *ent.Client) ur.Workspace {
	return &workspaceRepository{client: client}
}

func (r *workspaceRepository) Get(ctx context.Context, where *model.WorkspaceWhereInput, requestFields []string) (*model.Workspace, error) {
	q := r.client.Workspace.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	if collection.Contains(requestFields, "projects") {
		q.WithProjects(func(pq *ent.ProjectQuery) {
			pq.WithProjectTeammates(func(ptq *ent.ProjectTeammateQuery) {
				ptq.WithTeammate()
			})
		})
	}
	if collection.Contains(requestFields, "workspaceTeammates") {
		q.WithWorkspaceTeammates(func(wtq *ent.WorkspaceTeammateQuery) {
			wtq.WithTeammate()
		})
	}

	result, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceRepository) List(ctx context.Context) ([]*model.Workspace, error) {
	result, err := r.client.
		Workspace.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceWhereInput, requestFields []string) (*model.WorkspaceConnection, error) {
	q := r.client.Workspace.Query()

	if collection.Contains(requestFields, "edges.node.projects") {
		q.WithProjects(func(qp *ent.ProjectQuery) {
			qp.WithProjectTeammates(func(ptq *ent.ProjectTeammateQuery) {
				ptq.WithProject()
			})
		})
	}

	result, err := q.Paginate(ctx, after, first, before, last, ent.WithWorkspaceFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return result, nil
}

func (r *workspaceRepository) Create(ctx context.Context, input model.CreateWorkspaceInput) (*model.Workspace, error) {
	result, err := r.client.
		Workspace.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceRepository) Update(ctx context.Context, input model.UpdateWorkspaceInput) (*model.Workspace, error) {
	result, err := r.client.
		Workspace.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return result, nil
}
