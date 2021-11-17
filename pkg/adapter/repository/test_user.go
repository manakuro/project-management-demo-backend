package repository

import (
	"context"
	"errors"
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

func (r *testUserRepository) Find() (*model.TestUser, error) {
	ctx := context.Background()
	u, err := r.client.TestUser.Query().Where(testuser.IDEQ(1)).Only(ctx)

	if !errors.Is(err, nil) {
		return nil, err
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

	if !errors.Is(err, nil) {
		return nil, err
	}

	return &model.TestUser{TestUser: u}, nil
}

func (r *testUserRepository) Update(input model.UpdateTestUserInput) (*model.TestUser, error) {
	ctx := context.Background()
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if !errors.Is(err, nil) {
		return nil, err
	}

	u, err := r.client.
		TestUser.UpdateOneID(int(id)).
		SetInput(input.UpdateTestUserInput).
		Save(ctx)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return &model.TestUser{TestUser: u}, nil
}
