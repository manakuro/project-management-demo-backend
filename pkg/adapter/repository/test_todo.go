package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type testTodoRepository struct {
	client *ent.Client
}

// NewTestTodoRepository generates testTodo repository
func NewTestTodoRepository(client *ent.Client) ur.TestTodo {
	return &testTodoRepository{client: client}
}

func (r *testTodoRepository) Get(ctx context.Context, where *model.TestTodoWhereInput) (*model.TestTodo, error) {
	q := r.client.TestTodo.Query()

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

func (r *testTodoRepository) List(ctx context.Context) ([]*model.TestTodo, error) {
	res, err := r.client.TestTodo.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testTodoRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestTodoWhereInput) (*model.TestTodoConnection, error) {
	q := r.client.TestTodo.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTestTodoFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *testTodoRepository) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	res, err := r.client.
		TestTodo.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testTodoRepository) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	res, err := r.client.
		TestTodo.UpdateOneID(input.ID).
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
