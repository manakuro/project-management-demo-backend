package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type testUserRepository struct {
	client *ent.Client
}

// NewTestUserRepository generates test user repository
func NewTestUserRepository(client *ent.Client) ur.TestUser {
	return &testUserRepository{client: client}
}

func (r *testUserRepository) Get(ctx context.Context, id model.ID, age *int) (*model.TestUser, error) {
	q := r.client.TestUser.Query()

	if id == "" {
		return nil, model.NewInvalidParamError(map[string]interface{}{
			"id": id,
		})
	}
	q.Where(testuser.IDEQ(id))

	if age != nil {
		q.Where(testuser.AgeEQ(*age))
	}

	res, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id":  id,
				"age": age,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testUserRepository) List(ctx context.Context) ([]*model.TestUser, error) {
	res, err := r.client.
		TestUser.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testUserRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestUserWhereInput, requestedFields []string) (*model.TestUserConnection, error) {
	q := r.client.TestUser.Query()

	if collection.Contains(requestedFields, "edges.node.testTodos") {
		q.WithTestTodos()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTestUserFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *testUserRepository) Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	res, err := r.client.
		TestUser.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *testUserRepository) CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	client := WithTransactionalMutation(ctx)

	todo, err := client.
		TestTodo.
		Create().
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	u, err := client.TestUser.
		Create().
		SetInput(input).
		AddTestTodos(todo).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *testUserRepository) Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	res, err := r.client.
		TestUser.UpdateOneID(input.ID).
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
