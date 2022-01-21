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
		Auth:              r.NewAuthController(),
		Color:             r.NewColorController(),
		FavoriteProject:   r.NewFavoriteProjectController(),
		FavoriteWorkspace: r.NewFavoriteWorkspaceController(),
		Icon:              r.NewIconController(),
		Me:                r.NewMeController(),
		MyTasksTabStatus:  r.NewMyTasksTabStatusController(),
		Project:           r.NewProjectController(),
		ProjectBaseColor:  r.NewProjectBaseColorController(),
		ProjectIcon:       r.NewProjectIconController(),
		ProjectLightColor: r.NewProjectLightColorController(),
		ProjectTeammate:   r.NewProjectTeammateController(),
		TaskColumn:        r.NewTaskColumnController(),
		Teammate:          r.NewTeammateController(),
		TestTodo:          r.NewTestTodoController(),
		TestUser:          r.NewTestUserController(),
		Workspace:         r.NewWorkspaceController(),
		WorkspaceTeammate: r.NewWorkspaceTeammateController(),
	}
}
