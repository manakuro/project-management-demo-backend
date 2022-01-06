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

// Subscriptions hold an id and a channel of subscription.
type Subscriptions struct {
	TestUserUpdated  map[string]TestUserUpdated
	TeammateUpdated  map[string]TeammateUpdated
	WorkspaceUpdated map[string]WorkspaceUpdated
	ColorUpdated     map[string]ColorUpdated
}

// NewKey generates a random hex string with length of 16.
func NewKey() string {
	return randstr.Hex(16)
}
