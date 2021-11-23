package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"strconv"
)

type testTodoRepository struct {
	client *ent.Client
}

// NewTestTodoRepository generates test user repository
func NewTestTodoRepository(client *ent.Client) ur.TestTodo {
	return &testTodoRepository{client: client}
}

func (r *testTodoRepository) FindBy(id *string) (*model.TestTodo, error) {
	ctx := context.Background()

	q := r.client.TestTodo.Query()
	if id != nil {
		i, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			return nil, model.NewInvalidParamError(err, id)
		}
		q.Where(testtodo.IDEQ(int(i)))
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

	return &model.TestTodo{TestTodo: u}, nil
}

func (r *testTodoRepository) Create(input model.CreateTestTodoInput) (*model.TestTodo, error) {
	ctx := context.Background()
	u, err := r.client.
		TestTodo.
		Create().
		SetInput(input.CreateTestTodoInput).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.TestTodo{TestTodo: u}, nil
}

func (r *testTodoRepository) Update(input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	ctx := context.Background()
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.NewInvalidParamError(err, id)
	}

	u, err := r.client.
		TestTodo.UpdateOneID(int(id)).
		SetInput(input.UpdateTestTodoInput).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, id)
		}

		return nil, model.NewDBError(err)
	}

	return &model.TestTodo{TestTodo: u}, nil
}
