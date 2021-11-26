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

func (r *testTodoRepository) Get(id *model.ID) (*model.TestTodo, error) {
	ctx := context.Background()

	q := r.client.TestTodo.Query()
	if id != nil {
		q.Where(testtodo.IDEQ(*id))
	}

	u, err := q.Only(ctx)

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

	return u, nil
}

func (r *testTodoRepository) List() ([]*model.TestTodo, error) {
	ctx := context.Background()

	ts, err := r.client.
		TestTodo.
		Query().
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ts, nil
}

func (r *testTodoRepository) Create(input model.CreateTestTodoInput) (*model.TestTodo, error) {
	ctx := context.Background()
	u, err := r.client.
		TestTodo.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *testTodoRepository) Update(input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	ctx := context.Background()

	u, err := r.client.
		TestTodo.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return u, nil
}
