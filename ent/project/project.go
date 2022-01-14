// Code generated by entc, DO NOT EDIT.

package project

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldProjectBaseColorID holds the string denoting the project_base_color_id field in the database.
	FieldProjectBaseColorID = "project_base_color_id"
	// FieldProjectLightColorID holds the string denoting the project_light_color_id field in the database.
	FieldProjectLightColorID = "project_light_color_id"
	// FieldProjectIconID holds the string denoting the project_icon_id field in the database.
	FieldProjectIconID = "project_icon_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDescriptionTitle holds the string denoting the description_title field in the database.
	FieldDescriptionTitle = "description_title"
	// FieldDueDate holds the string denoting the due_date field in the database.
	FieldDueDate = "due_date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWorkspace holds the string denoting the workspace edge name in mutations.
	EdgeWorkspace = "workspace"
	// EdgeProjectBaseColor holds the string denoting the project_base_color edge name in mutations.
	EdgeProjectBaseColor = "project_base_color"
	// EdgeProjectLightColor holds the string denoting the project_light_color edge name in mutations.
	EdgeProjectLightColor = "project_light_color"
	// EdgeProjectIcon holds the string denoting the project_icon edge name in mutations.
	EdgeProjectIcon = "project_icon"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeProjectTeammates holds the string denoting the project_teammates edge name in mutations.
	EdgeProjectTeammates = "project_teammates"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "projects"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
	// ProjectBaseColorTable is the table that holds the project_base_color relation/edge.
	ProjectBaseColorTable = "projects"
	// ProjectBaseColorInverseTable is the table name for the ProjectBaseColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectbasecolor" package.
	ProjectBaseColorInverseTable = "project_base_colors"
	// ProjectBaseColorColumn is the table column denoting the project_base_color relation/edge.
	ProjectBaseColorColumn = "project_base_color_id"
	// ProjectLightColorTable is the table that holds the project_light_color relation/edge.
	ProjectLightColorTable = "projects"
	// ProjectLightColorInverseTable is the table name for the ProjectLightColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectlightcolor" package.
	ProjectLightColorInverseTable = "project_light_colors"
	// ProjectLightColorColumn is the table column denoting the project_light_color relation/edge.
	ProjectLightColorColumn = "project_light_color_id"
	// ProjectIconTable is the table that holds the project_icon relation/edge.
	ProjectIconTable = "projects"
	// ProjectIconInverseTable is the table name for the ProjectIcon entity.
	// It exists in this package in order to avoid circular dependency with the "projecticon" package.
	ProjectIconInverseTable = "project_icons"
	// ProjectIconColumn is the table column denoting the project_icon relation/edge.
	ProjectIconColumn = "project_icon_id"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "projects"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "created_by"
	// ProjectTeammatesTable is the table that holds the project_teammates relation/edge.
	ProjectTeammatesTable = "project_teammates"
	// ProjectTeammatesInverseTable is the table name for the ProjectTeammate entity.
	// It exists in this package in order to avoid circular dependency with the "projectteammate" package.
	ProjectTeammatesInverseTable = "project_teammates"
	// ProjectTeammatesColumn is the table column denoting the project_teammates relation/edge.
	ProjectTeammatesColumn = "project_id"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldWorkspaceID,
	FieldProjectBaseColorID,
	FieldProjectLightColorID,
	FieldProjectIconID,
	FieldCreatedBy,
	FieldName,
	FieldDescription,
	FieldDescriptionTitle,
	FieldDueDate,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionTitleValidator is a validator for the "description_title" field. It is called by the builders before save.
	DescriptionTitleValidator func(string) error
	// DefaultDueDate holds the default value on creation for the "due_date" field.
	DefaultDueDate func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
