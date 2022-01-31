package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type fileTypeRepository struct {
	client *ent.Client
}

// NewFileTypeRepository generates fileType repository
func NewFileTypeRepository(client *ent.Client) ur.FileType {
	return &fileTypeRepository{client: client}
}

func (r *fileTypeRepository) Get(ctx context.Context, where *model.FileTypeWhereInput) (*model.FileType, error) {
	q := r.client.FileType.Query()

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

func (r *fileTypeRepository) List(ctx context.Context) ([]*model.FileType, error) {
	res, err := r.client.FileType.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *fileTypeRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FileTypeWhereInput, requestedFields []string) (*model.FileTypeConnection, error) {
	q := r.client.FileType.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithFileTypeFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *fileTypeRepository) Create(ctx context.Context, input model.CreateFileTypeInput) (*model.FileType, error) {
	res, err := r.client.
		FileType.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *fileTypeRepository) Update(ctx context.Context, input model.UpdateFileTypeInput) (*model.FileType, error) {
	res, err := r.client.
		FileType.UpdateOneID(input.ID).
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
