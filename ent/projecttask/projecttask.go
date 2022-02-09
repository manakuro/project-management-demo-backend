// Code generated by entc, DO NOT EDIT.

package projecttask

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the projecttask type in the database.
	Label = "project_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldProjectTaskSectionID holds the string denoting the project_task_section_id field in the database.
	FieldProjectTaskSectionID = "project_task_section_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeProjectTaskSection holds the string denoting the projecttasksection edge name in mutations.
	EdgeProjectTaskSection = "projectTaskSection"
	// Table holds the table name of the projecttask in the database.
	Table = "project_tasks"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "project_tasks"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "project_tasks"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_id"
	// ProjectTaskSectionTable is the table that holds the projectTaskSection relation/edge.
	ProjectTaskSectionTable = "project_tasks"
	// ProjectTaskSectionInverseTable is the table name for the ProjectTaskSection entity.
	// It exists in this package in order to avoid circular dependency with the "projecttasksection" package.
	ProjectTaskSectionInverseTable = "project_task_sections"
	// ProjectTaskSectionColumn is the table column denoting the projectTaskSection relation/edge.
	ProjectTaskSectionColumn = "project_task_section_id"
)

// Columns holds all SQL columns for projecttask fields.
var Columns = []string{
	FieldID,
	FieldProjectID,
	FieldTaskID,
	FieldProjectTaskSectionID,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
