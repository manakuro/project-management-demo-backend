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
		ActivityType:            r.NewActivityTypeController(),
		Auth:                    r.NewAuthController(),
		Color:                   r.NewColorController(),
		DeletedTask:             r.NewDeletedTaskController(),
		FavoriteProject:         r.NewFavoriteProjectController(),
		FavoriteWorkspace:       r.NewFavoriteWorkspaceController(),
		FileType:                r.NewFileTypeController(),
		Icon:                    r.NewIconController(),
		Me:                      r.NewMeController(),
		Mention:                 r.NewMentionController(),
		Project:                 r.NewProjectController(),
		ProjectBaseColor:        r.NewProjectBaseColorController(),
		ProjectIcon:             r.NewProjectIconController(),
		ProjectLightColor:       r.NewProjectLightColorController(),
		ProjectTask:             r.NewProjectTaskController(),
		ProjectTaskColumn:       r.NewProjectTaskColumnController(),
		ProjectTaskListStatus:   r.NewProjectTaskListStatusController(),
		ProjectTaskSection:      r.NewProjectTaskSectionController(),
		ProjectTeammate:         r.NewProjectTeammateController(),
		Tag:                     r.NewTagController(),
		Task:                    r.NewTaskController(),
		TaskActivity:            r.NewTaskActivityController(),
		TaskCollaborator:        r.NewTaskCollaboratorController(),
		TaskColumn:              r.NewTaskColumnController(),
		TaskFeed:                r.NewTaskFeedController(),
		TaskFeedLike:            r.NewTaskFeedLikeController(),
		TaskFile:                r.NewTaskFileController(),
		TaskLike:                r.NewTaskLikeController(),
		TaskListCompletedStatus: r.NewTaskListCompletedStatusController(),
		TaskListSortStatus:      r.NewTaskListSortStatusController(),
		TaskPriority:            r.NewTaskPriorityController(),
		TaskSection:             r.NewTaskSectionController(),
		TaskTag:                 r.NewTaskTagController(),
		Teammate:                r.NewTeammateController(),
		TeammateTask:            r.NewTeammateTaskController(),
		TeammateTaskColumn:      r.NewTeammateTaskColumnController(),
		TeammateTaskListStatus:  r.NewTeammateTaskListStatusController(),
		TeammateTaskSection:     r.NewTeammateTaskSectionController(),
		TeammateTaskTabStatus:   r.NewTeammateTaskTabStatusController(),
		TestTodo:                r.NewTestTodoController(),
		TestUser:                r.NewTestUserController(),
		Workspace:               r.NewWorkspaceController(),
		WorkspaceTeammate:       r.NewWorkspaceTeammateController(),
	}
}
