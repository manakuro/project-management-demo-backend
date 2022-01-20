package subscription

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/entity/model"

	"github.com/thanhpk/randstr"
)

// TestUserUpdated of channel
type TestUserUpdated struct {
	ID ulid.ID
	Ch chan *ent.TestUser
}

// TeammateUpdated of channel
type TeammateUpdated struct {
	ID ulid.ID
	Ch chan *ent.Teammate
}

// WorkspaceUpdated of channel
type WorkspaceUpdated struct {
	ID ulid.ID
	Ch chan *ent.Workspace
}

// ColorUpdated of channel
type ColorUpdated struct {
	ID ulid.ID
	Ch chan *ent.Color
}

// IconUpdated of channel
type IconUpdated struct {
	ID ulid.ID
	Ch chan *ent.Icon
}

// ProjectUpdated of channel
type ProjectUpdated struct {
	ID ulid.ID
	Ch chan *ent.Project
}

// ProjectTeammateUpdated of channel
type ProjectTeammateUpdated struct {
	ID ulid.ID
	Ch chan *ent.ProjectTeammate
}

// ProjectBaseColorUpdated of channel
type ProjectBaseColorUpdated struct {
	ID ulid.ID
	Ch chan *ent.ProjectBaseColor
}

// ProjectLightColorUpdated of channel
type ProjectLightColorUpdated struct {
	ID ulid.ID
	Ch chan *ent.ProjectLightColor
}

// ProjectIconUpdated of channel
type ProjectIconUpdated struct {
	ID ulid.ID
	Ch chan *ent.ProjectIcon
}

// MeUpdated of channel
type MeUpdated struct {
	ID ulid.ID
	Ch chan *model.Me
}

// WorkspaceTeammateUpdated of channel
type WorkspaceTeammateUpdated struct {
	ID ulid.ID
	Ch chan *model.WorkspaceTeammate
}

// FavoriteProjectCreated of channel
type FavoriteProjectCreated struct {
	TeammateID ulid.ID
	Ch         chan *model.FavoriteProject
}

// FavoriteProjectIDsUpdated of channel
type FavoriteProjectIDsUpdated struct {
	TeammateID model.ID
	Ch         chan []model.ID
}

// Subscriptions hold an id and a channel of subscription.
type Subscriptions struct {
	ColorUpdated              map[string]ColorUpdated
	FavoriteProjectCreated    map[string]FavoriteProjectCreated
	FavoriteProjectIDsUpdated map[string]FavoriteProjectIDsUpdated
	IconUpdated               map[string]IconUpdated
	MeUpdated                 map[string]MeUpdated
	ProjectBaseColorUpdated   map[string]ProjectBaseColorUpdated
	ProjectIconUpdated        map[string]ProjectIconUpdated
	ProjectLightColorUpdated  map[string]ProjectLightColorUpdated
	ProjectTeammateUpdated    map[string]ProjectTeammateUpdated
	ProjectUpdated            map[string]ProjectUpdated
	TeammateUpdated           map[string]TeammateUpdated
	TestUserUpdated           map[string]TestUserUpdated
	WorkspaceTeammateUpdated  map[string]WorkspaceTeammateUpdated
	WorkspaceUpdated          map[string]WorkspaceUpdated
}

// NewKey generates a random hex string with length of 16.
func NewKey() string {
	return randstr.Hex(16)
}
