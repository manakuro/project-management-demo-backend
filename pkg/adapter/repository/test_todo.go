package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type testTodoRepository struct {
	client *ent.Client
}

// NewTestTodoRepository generates test user repository
func NewTestTodoRepository(client *ent.Client) ur.TestTodo {
	return &testTodoRepository{client: client}
}

func (r *testTodoRepository) Get(ctx context.Context, id *model.ID) (*model.TestTodo, error) {
	q := r.client.TestTodo.Query()
	if id != nil {
		q.Where(testtodo.IDEQ(*id))
	}

	res, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testTodoRepository) List(ctx context.Context) ([]*model.TestTodo, error) {
	res, err := r.client.
		TestTodo.
		Query().
		All(ctx)
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
