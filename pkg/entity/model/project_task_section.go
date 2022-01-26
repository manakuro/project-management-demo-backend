package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// ProjectTaskSection is the model entity for the ProjectTaskSection schema.
type ProjectTaskSection = ent.ProjectTaskSection

// ProjectTaskSectionConnection is the connection containing edges to ProjectTaskSection.
type ProjectTaskSectionConnection = ent.ProjectTaskSectionConnection

// CreateProjectTaskSectionInput represents a mutation input.
type CreateProjectTaskSectionInput = ent.CreateProjectTaskSectionInput

// ProjectTaskSectionWhereInput represents a where input.
type ProjectTaskSectionWhereInput = ent.ProjectTaskSectionWhereInput

// UpdateProjectTaskSectionInput represents a mutation input.
type UpdateProjectTaskSectionInput = ent.UpdateProjectTaskSectionInput

// DeleteProjectTaskSectionInput represents a mutation input.
type DeleteProjectTaskSectionInput struct {
	ID ulid.ID
}
