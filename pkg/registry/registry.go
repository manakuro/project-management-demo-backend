package registry

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
)

type registry struct {
	client *ent.Client
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		TestUser: r.NewTestUserController(),
		TestTodo: r.NewTestTodoController(),
	}
}
