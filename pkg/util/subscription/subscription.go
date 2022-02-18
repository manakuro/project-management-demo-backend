package subscription

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/entity/model"

	"github.com/thanhpk/randstr"
)

// TestUserUpdated is a channel for subscription.
type TestUserUpdated struct {
	ID ulid.ID
	Ch chan *ent.TestUser
}

// TeammateUpdated is a channel for subscription.
type TeammateUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.Teammate
}

// WorkspaceUpdated is a channel for subscription.
type WorkspaceUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.Workspace
}

// ColorUpdated is a channel for subscription.
type ColorUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.Color
}

// IconUpdated is a channel for subscription.
type IconUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.Icon
}

// ProjectUpdated is a channel for subscription.
type ProjectUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.Project
}

// ProjectTeammateUpdated is a channel for subscription.
type ProjectTeammateUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.ProjectTeammate
}

// ProjectBaseColorUpdated is a channel for subscription.
type ProjectBaseColorUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.ProjectBaseColor
}

// ProjectLightColorUpdated is a channel for subscription.
type ProjectLightColorUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.ProjectLightColor
}

// ProjectIconUpdated is a channel for subscription.
type ProjectIconUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.ProjectIcon
}

// MeUpdated is a channel for subscription.
type MeUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *model.Me
}

// WorkspaceTeammateUpdated is a channel for subscription.
type WorkspaceTeammateUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *model.WorkspaceTeammate
}

// FavoriteProjectCreated is a channel for subscription.
type FavoriteProjectCreated struct {
	TeammateID ulid.ID
	RequestID  string
	Ch         chan *model.FavoriteProject
}

// FavoriteProjectIDsUpdated is a channel for subscription.
type FavoriteProjectIDsUpdated struct {
	TeammateID model.ID
	RequestID  string
	Ch         chan []model.ID
}

// FavoriteWorkspaceIDsUpdated is a channel for subscription.
type FavoriteWorkspaceIDsUpdated struct {
	TeammateID model.ID
	RequestID  string
	Ch         chan []model.ID
}

// TeammateTaskTabStatusUpdated is a channel for subscription.
type TeammateTaskTabStatusUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TeammateTaskTabStatus
}

// TaskColumnUpdated is a channel for subscription.
type TaskColumnUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TaskColumn
}

// TeammateTaskColumnUpdated is a channel for subscription.
type TeammateTaskColumnUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TeammateTaskColumn
}

// ProjectTaskColumnUpdated is a channel for subscription.
type ProjectTaskColumnUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.ProjectTaskColumn
}

// TaskSectionUpdated is a channel for subscription.
type TaskSectionUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TaskSection
}

// TeammateTaskListStatusUpdated is a channel for subscription.
type TeammateTaskListStatusUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TeammateTaskListStatus
}

// ProjectTaskListStatusUpdated is a channel for subscription.
type ProjectTaskListStatusUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.ProjectTaskListStatus
}

// TeammateTaskSectionUpdated is a channel for subscription.
type TeammateTaskSectionUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TeammateTaskSection
}

// TeammateTaskSectionCreated is a channel for subscription.
type TeammateTaskSectionCreated struct {
	TeammateID  model.ID
	WorkspaceID model.ID
	RequestID   string
	Ch          chan *model.TeammateTaskSection
}

// ProjectTaskSectionUpdated is a channel for subscription.
type ProjectTaskSectionUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.ProjectTaskSection
}

// ProjectTaskSectionCreated is a channel for subscription.
type ProjectTaskSectionCreated struct {
	ProjectID model.ID
	RequestID string
	Ch        chan *model.ProjectTaskSection
}

// TaskUpdated is a channel for subscription.
type TaskUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.Task
}

// TeammateTaskUpdated is a channel for subscription.
type TeammateTaskUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TeammateTask
}

// TeammateTaskCreated is a channel for subscription.
type TeammateTaskCreated struct {
	TeammateID  ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
	Ch          chan *model.TeammateTask
}

// ProjectTaskUpdated is a channel for subscription.
type ProjectTaskUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.ProjectTask
}

// ProjectTaskCreated is a channel for subscription.
type ProjectTaskCreated struct {
	ProjectID model.ID
	RequestID string
	Ch        chan *model.ProjectTask
}

// TaskLikesUpdated is a channel for subscription.
type TaskLikesUpdated struct {
	WorkspaceID model.ID
	RequestID   string
	Ch          chan *model.TaskLike
}

// TagUpdated is a channel for subscription.
type TagUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.Tag
}

// TaskTagUpdated is a channel for subscription.
type TaskTagUpdated struct {
	TaskID    model.ID
	RequestID string
	Ch        chan []*model.TaskTag
}

// TaskCollaboratorUpdated is a channel for subscription.
type TaskCollaboratorUpdated struct {
	TaskID    model.ID
	RequestID string
	Ch        chan []*model.TaskCollaborator
}

// TaskFeedUpdated is a channel for subscription.
type TaskFeedUpdated struct {
	ID        model.ID
	RequestID string
	Ch        chan *model.TaskFeed
}

// TaskFeedCreated is a channel for subscription.
type TaskFeedCreated struct {
	TaskID    model.ID
	RequestID string
	Ch        chan *model.TaskFeed
}

// TaskFeedDeleted is a channel for subscription.
type TaskFeedDeleted struct {
	TaskID    model.ID
	RequestID string
	Ch        chan *model.TaskFeed
}

// TaskFeedLikeUpdated is a channel for subscription.
type TaskFeedLikeUpdated struct {
	TaskID    model.ID
	RequestID string
	Ch        chan []*model.TaskFeedLike
}

// TaskFileUpdated is a channel for subscription.
type TaskFileUpdated struct {
	ID        ulid.ID
	RequestID string
	Ch        chan *ent.TaskFile
}

// Subscriptions hold an id and a channel of subscription.
type Subscriptions struct {
	ColorUpdated                  map[string]ColorUpdated
	FavoriteProjectCreated        map[string]FavoriteProjectCreated
	FavoriteProjectIDsUpdated     map[string]FavoriteProjectIDsUpdated
	FavoriteWorkspaceIDsUpdated   map[string]FavoriteWorkspaceIDsUpdated
	IconUpdated                   map[string]IconUpdated
	MeUpdated                     map[string]MeUpdated
	ProjectBaseColorUpdated       map[string]ProjectBaseColorUpdated
	ProjectIconUpdated            map[string]ProjectIconUpdated
	ProjectLightColorUpdated      map[string]ProjectLightColorUpdated
	ProjectTaskColumnUpdated      map[string]ProjectTaskColumnUpdated
	ProjectTaskCreated            map[string]ProjectTaskCreated
	ProjectTaskListStatusUpdated  map[string]ProjectTaskListStatusUpdated
	ProjectTaskSectionUpdated     map[string]ProjectTaskSectionUpdated
	ProjectTaskSectionCreated     map[string]ProjectTaskSectionCreated
	ProjectTaskUpdated            map[string]ProjectTaskUpdated
	ProjectTeammateUpdated        map[string]ProjectTeammateUpdated
	ProjectUpdated                map[string]ProjectUpdated
	TagUpdated                    map[string]TagUpdated
	TaskCollaboratorUpdated       map[string]TaskCollaboratorUpdated
	TaskColumnUpdated             map[string]TaskColumnUpdated
	TaskFeedLikeUpdated           map[string]TaskFeedLikeUpdated
	TaskFeedUpdated               map[string]TaskFeedUpdated
	TaskFeedCreated               map[string]TaskFeedCreated
	TaskFeedDeleted               map[string]TaskFeedDeleted
	TaskFileUpdated               map[string]TaskFileUpdated
	TaskLikesUpdated              map[string]TaskLikesUpdated
	TaskSectionUpdated            map[string]TaskSectionUpdated
	TaskTagUpdated                map[string]TaskTagUpdated
	TaskUpdated                   map[string]TaskUpdated
	TeammateTaskColumnUpdated     map[string]TeammateTaskColumnUpdated
	TeammateTaskCreated           map[string]TeammateTaskCreated
	TeammateTaskListStatusUpdated map[string]TeammateTaskListStatusUpdated
	TeammateTaskSectionCreated    map[string]TeammateTaskSectionCreated
	TeammateTaskSectionUpdated    map[string]TeammateTaskSectionUpdated
	TeammateTaskTabStatusUpdated  map[string]TeammateTaskTabStatusUpdated
	TeammateTaskUpdated           map[string]TeammateTaskUpdated
	TeammateUpdated               map[string]TeammateUpdated
	TestUserUpdated               map[string]TestUserUpdated
	WorkspaceTeammateUpdated      map[string]WorkspaceTeammateUpdated
	WorkspaceUpdated              map[string]WorkspaceUpdated
}

// NewKey generates a random hex string with length of 16.
func NewKey() string {
	return randstr.Hex(16)
}
