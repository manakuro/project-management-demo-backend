// Code generated by entc, DO NOT EDIT.

package color

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the color type in the database.
	Label = "color"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldHex holds the string denoting the hex field in the database.
	FieldHex = "hex"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProjectBaseColors holds the string denoting the project_base_colors edge name in mutations.
	EdgeProjectBaseColors = "project_base_colors"
	// EdgeProjectLightColors holds the string denoting the project_light_colors edge name in mutations.
	EdgeProjectLightColors = "project_light_colors"
	// EdgeTaskPriorities holds the string denoting the task_priorities edge name in mutations.
	EdgeTaskPriorities = "task_priorities"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// Table holds the table name of the color in the database.
	Table = "colors"
	// ProjectBaseColorsTable is the table that holds the project_base_colors relation/edge.
	ProjectBaseColorsTable = "project_base_colors"
	// ProjectBaseColorsInverseTable is the table name for the ProjectBaseColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectbasecolor" package.
	ProjectBaseColorsInverseTable = "project_base_colors"
	// ProjectBaseColorsColumn is the table column denoting the project_base_colors relation/edge.
	ProjectBaseColorsColumn = "color_id"
	// ProjectLightColorsTable is the table that holds the project_light_colors relation/edge.
	ProjectLightColorsTable = "project_light_colors"
	// ProjectLightColorsInverseTable is the table name for the ProjectLightColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectlightcolor" package.
	ProjectLightColorsInverseTable = "project_light_colors"
	// ProjectLightColorsColumn is the table column denoting the project_light_colors relation/edge.
	ProjectLightColorsColumn = "color_id"
	// TaskPrioritiesTable is the table that holds the task_priorities relation/edge.
	TaskPrioritiesTable = "task_priorities"
	// TaskPrioritiesInverseTable is the table name for the TaskPriority entity.
	// It exists in this package in order to avoid circular dependency with the "taskpriority" package.
	TaskPrioritiesInverseTable = "task_priorities"
	// TaskPrioritiesColumn is the table column denoting the task_priorities relation/edge.
	TaskPrioritiesColumn = "color_id"
	// TagsTable is the table that holds the tags relation/edge.
	TagsTable = "tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TagsColumn is the table column denoting the tags relation/edge.
	TagsColumn = "color_id"
)

// Columns holds all SQL columns for color fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldColor,
	FieldHex,
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
	// ColorValidator is a validator for the "color" field. It is called by the builders before save.
	ColorValidator func(string) error
	// HexValidator is a validator for the "hex" field. It is called by the builders before save.
	HexValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
