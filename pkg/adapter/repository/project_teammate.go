package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type projectTeammateRepository struct {
	client *ent.Client
}

// NewProjectTeammateRepository generates projectTeammate repository
func NewProjectTeammateRepository(client *ent.Client) ur.ProjectTeammate {
	return &projectTeammateRepository{client: client}
}

func (r *projectTeammateRepository) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	q := r.client.ProjectTeammate.Query()

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

func (r *projectTeammateRepository) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTeammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammateConnection, error) {
	q := r.client.ProjectTeammate.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectTeammateRepository) Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTeammateRepository) Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.UpdateOneID(input.ID).
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

func (r *projectTeammateRepository) UpdateOwner(ctx context.Context, input model.UpdateProjectTeammateOwnerInput) (*model.ProjectTeammate, error) {
	client := WithTransactionalMutation(ctx)

	_, err := client.ProjectTeammate.
		Update().
		SetIsOwner(false).
		Where(projectteammate.ProjectID(input.ProjectID)).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	q := client.ProjectTeammate.
		Query().
		Where(
			projectteammate.TeammateID(input.TeammateID),
			projectteammate.ProjectID(input.ProjectID),
		)

	WithProjectTeammate(q)

	projectTeammate, err := q.Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if projectTeammate == nil {
		res, perr := client.ProjectTeammate.
			Create().
			SetProjectID(input.ProjectID).
			SetTeammateID(input.TeammateID).
			SetRole("").
			SetIsOwner(true).
			Save(ctx)

		if perr != nil {
			return nil, model.NewDBError(perr)
		}
		return res, perr
	}

	res, err := client.ProjectTeammate.UpdateOneID(projectTeammate.ID).SetIsOwner(true).Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

// WithProjectTeammate eager-loads associations with project teammate entity.
func WithProjectTeammate(query *ent.ProjectTeammateQuery) {
	query.WithTeammate()
}
