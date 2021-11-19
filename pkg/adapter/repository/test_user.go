package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"strconv"
)

type testUserRepository struct {
	client *ent.Client
}

// NewTestUserRepository generates test user repository
func NewTestUserRepository(client *ent.Client) ur.TestUser {
	return &testUserRepository{client: client}
}

func (r *testUserRepository) FindBy(id *string, age *int) (*model.TestUser, error) {
	ctx := context.Background()

	q := r.client.TestUser.Query()
	if id != nil {
		i, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			return nil, model.NewInvalidParamError(err)
		}
		q.Where(testuser.IDEQ(int(i)))
	}
	if age != nil {
		q.Where(testuser.AgeEQ(*age))
	}

	u, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return &model.TestUser{TestUser: u}, nil
}

func (r *testUserRepository) Create(input model.CreateTestUserInput) (*model.TestUser, error) {
	ctx := context.Background()
	u, err := r.client.
		TestUser.
		Create().
		SetInput(input.CreateTestUserInput).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.TestUser{TestUser: u}, nil
}

func (r *testUserRepository) Update(input model.UpdateTestUserInput) (*model.TestUser, error) {
	ctx := context.Background()
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.NewInvalidParamError(err)
	}

	u, err := r.client.
		TestUser.UpdateOneID(int(id)).
		SetInput(input.UpdateTestUserInput).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return &model.TestUser{TestUser: u}, nil
}
