// Code generated by entc, DO NOT EDIT.

package ent

import (
	"project-management-demo-backend/ent/schema"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"time"
)

// CreateTeammateInput represents a mutation input for creating teammates.
type CreateTeammateInput struct {
	Name        string
	Image       string
	Email       string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	WorkspaceID *ulid.ID
}

// Mutate applies the CreateTeammateInput on the TeammateCreate builder.
func (i *CreateTeammateInput) Mutate(m *TeammateCreate) {
	m.SetName(i.Name)
	m.SetImage(i.Image)
	m.SetEmail(i.Email)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.WorkspaceID; v != nil {
		m.SetWorkspacesID(*v)
	}
}

// SetInput applies the change-set in the CreateTeammateInput on the create builder.
func (c *TeammateCreate) SetInput(i CreateTeammateInput) *TeammateCreate {
	i.Mutate(c)
	return c
}

// UpdateTeammateInput represents a mutation input for updating teammates.
type UpdateTeammateInput struct {
	ID              ulid.ID
	Name            *string
	Image           *string
	Email           *string
	WorkspaceID     *ulid.ID
	ClearWorkspaces bool
}

// Mutate applies the UpdateTeammateInput on the TeammateMutation.
func (i *UpdateTeammateInput) Mutate(m *TeammateMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Image; v != nil {
		m.SetImage(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if i.ClearWorkspaces {
		m.ClearWorkspaces()
	}
	if v := i.WorkspaceID; v != nil {
		m.SetWorkspacesID(*v)
	}
}

// SetInput applies the change-set in the UpdateTeammateInput on the update builder.
func (u *TeammateUpdate) SetInput(i UpdateTeammateInput) *TeammateUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateTeammateInput on the update-one builder.
func (u *TeammateUpdateOne) SetInput(i UpdateTeammateInput) *TeammateUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateTestTodoInput represents a mutation input for creating testtodos.
type CreateTestTodoInput struct {
	Name       *string
	Status     *testtodo.Status
	Priority   *int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	TestUserID *ulid.ID
}

// Mutate applies the CreateTestTodoInput on the TestTodoCreate builder.
func (i *CreateTestTodoInput) Mutate(m *TestTodoCreate) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.TestUserID; v != nil {
		m.SetTestUserID(*v)
	}
}

// SetInput applies the change-set in the CreateTestTodoInput on the create builder.
func (c *TestTodoCreate) SetInput(i CreateTestTodoInput) *TestTodoCreate {
	i.Mutate(c)
	return c
}

// UpdateTestTodoInput represents a mutation input for updating testtodos.
type UpdateTestTodoInput struct {
	ID            ulid.ID
	Name          *string
	Status        *testtodo.Status
	Priority      *int
	TestUserID    *ulid.ID
	ClearTestUser bool
}

// Mutate applies the UpdateTestTodoInput on the TestTodoMutation.
func (i *UpdateTestTodoInput) Mutate(m *TestTodoMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if i.ClearTestUser {
		m.ClearTestUser()
	}
	if v := i.TestUserID; v != nil {
		m.SetTestUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateTestTodoInput on the update builder.
func (u *TestTodoUpdate) SetInput(i UpdateTestTodoInput) *TestTodoUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateTestTodoInput on the update-one builder.
func (u *TestTodoUpdateOne) SetInput(i UpdateTestTodoInput) *TestTodoUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateTestUserInput represents a mutation input for creating testusers.
type CreateTestUserInput struct {
	Name        string
	Age         int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	TestTodoIDs []ulid.ID
}

// Mutate applies the CreateTestUserInput on the TestUserCreate builder.
func (i *CreateTestUserInput) Mutate(m *TestUserCreate) {
	m.SetName(i.Name)
	m.SetAge(i.Age)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if ids := i.TestTodoIDs; len(ids) > 0 {
		m.AddTestTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateTestUserInput on the create builder.
func (c *TestUserCreate) SetInput(i CreateTestUserInput) *TestUserCreate {
	i.Mutate(c)
	return c
}

// UpdateTestUserInput represents a mutation input for updating testusers.
type UpdateTestUserInput struct {
	ID                ulid.ID
	Name              *string
	Age               *int
	AddTestTodoIDs    []ulid.ID
	RemoveTestTodoIDs []ulid.ID
}

// Mutate applies the UpdateTestUserInput on the TestUserMutation.
func (i *UpdateTestUserInput) Mutate(m *TestUserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if ids := i.AddTestTodoIDs; len(ids) > 0 {
		m.AddTestTodoIDs(ids...)
	}
	if ids := i.RemoveTestTodoIDs; len(ids) > 0 {
		m.RemoveTestTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateTestUserInput on the update builder.
func (u *TestUserUpdate) SetInput(i UpdateTestUserInput) *TestUserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateTestUserInput on the update-one builder.
func (u *TestUserUpdateOne) SetInput(i UpdateTestUserInput) *TestUserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateWorkspaceInput represents a mutation input for creating workspaces.
type CreateWorkspaceInput struct {
	Name        string
	Description schema.WorkspaceDescription
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   ulid.ID
}

// Mutate applies the CreateWorkspaceInput on the WorkspaceCreate builder.
func (i *CreateWorkspaceInput) Mutate(m *WorkspaceCreate) {
	m.SetName(i.Name)
	m.SetDescription(i.Description)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTeammateID(i.CreatedBy)
}

// SetInput applies the change-set in the CreateWorkspaceInput on the create builder.
func (c *WorkspaceCreate) SetInput(i CreateWorkspaceInput) *WorkspaceCreate {
	i.Mutate(c)
	return c
}

// UpdateWorkspaceInput represents a mutation input for updating workspaces.
type UpdateWorkspaceInput struct {
	ID            ulid.ID
	Name          *string
	Description   *schema.WorkspaceDescription
	CreatedBy     *ulid.ID
	ClearTeammate bool
}

// Mutate applies the UpdateWorkspaceInput on the WorkspaceMutation.
func (i *UpdateWorkspaceInput) Mutate(m *WorkspaceMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearTeammate {
		m.ClearTeammate()
	}
	if v := i.CreatedBy; v != nil {
		m.SetTeammateID(*v)
	}
}

// SetInput applies the change-set in the UpdateWorkspaceInput on the update builder.
func (u *WorkspaceUpdate) SetInput(i UpdateWorkspaceInput) *WorkspaceUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateWorkspaceInput on the update-one builder.
func (u *WorkspaceUpdateOne) SetInput(i UpdateWorkspaceInput) *WorkspaceUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
