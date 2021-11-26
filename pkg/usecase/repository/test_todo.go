package repository

import "project-management-demo-backend/pkg/entity/model"

// TestTodo is interface of repository
type TestTodo interface {
	Get(id *model.ID) (*model.TestTodo, error)
	List() ([]*model.TestTodo, error)
	Create(input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(input model.UpdateTestTodoInput) (*model.TestTodo, error)
}
