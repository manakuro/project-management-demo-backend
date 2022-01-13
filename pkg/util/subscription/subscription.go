package subscription

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"

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

// Subscriptions hold an id and a channel of subscription.
type Subscriptions struct {
	ColorUpdated           map[string]ColorUpdated
	IconUpdated            map[string]IconUpdated
	ProjectUpdated         map[string]ProjectUpdated
	ProjectTeammateUpdated map[string]ProjectTeammateUpdated
	TeammateUpdated        map[string]TeammateUpdated
	TestUserUpdated        map[string]TestUserUpdated
	WorkspaceUpdated       map[string]WorkspaceUpdated
}

// NewKey generates a random hex string with length of 16.
func NewKey() string {
	return randstr.Hex(16)
}
