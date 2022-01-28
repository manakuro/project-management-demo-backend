// Code generated by entc, DO NOT EDIT.

package task

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTaskParentID holds the string denoting the task_parent_id field in the database.
	FieldTaskParentID = "task_parent_id"
	// FieldTaskPriorityID holds the string denoting the task_priority_id field in the database.
	FieldTaskPriorityID = "task_priority_id"
	// FieldAssigneeID holds the string denoting the assignee_id field in the database.
	FieldAssigneeID = "assignee_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// FieldCompletedAt holds the string denoting the completed_at field in the database.
	FieldCompletedAt = "completed_at"
	// FieldIsNew holds the string denoting the is_new field in the database.
	FieldIsNew = "is_new"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDueDate holds the string denoting the due_date field in the database.
	FieldDueDate = "due_date"
	// FieldDueTime holds the string denoting the due_time field in the database.
	FieldDueTime = "due_time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTaskPriority holds the string denoting the task_priority edge name in mutations.
	EdgeTaskPriority = "task_priority"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeSubTasks holds the string denoting the sub_tasks edge name in mutations.
	EdgeSubTasks = "sub_tasks"
	// EdgeTeammateTasks holds the string denoting the teammate_tasks edge name in mutations.
	EdgeTeammateTasks = "teammate_tasks"
	// EdgeProjectTasks holds the string denoting the project_tasks edge name in mutations.
	EdgeProjectTasks = "project_tasks"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "tasks"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "assignee_id"
	// TaskPriorityTable is the table that holds the task_priority relation/edge.
	TaskPriorityTable = "tasks"
	// TaskPriorityInverseTable is the table name for the TaskPriority entity.
	// It exists in this package in order to avoid circular dependency with the "taskpriority" package.
	TaskPriorityInverseTable = "task_priorities"
	// TaskPriorityColumn is the table column denoting the task_priority relation/edge.
	TaskPriorityColumn = "task_priority_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "tasks"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "task_parent_id"
	// SubTasksTable is the table that holds the sub_tasks relation/edge.
	SubTasksTable = "tasks"
	// SubTasksColumn is the table column denoting the sub_tasks relation/edge.
	SubTasksColumn = "task_parent_id"
	// TeammateTasksTable is the table that holds the teammate_tasks relation/edge.
	TeammateTasksTable = "teammate_tasks"
	// TeammateTasksInverseTable is the table name for the TeammateTask entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetask" package.
	TeammateTasksInverseTable = "teammate_tasks"
	// TeammateTasksColumn is the table column denoting the teammate_tasks relation/edge.
	TeammateTasksColumn = "task_id"
	// ProjectTasksTable is the table that holds the project_tasks relation/edge.
	ProjectTasksTable = "project_tasks"
	// ProjectTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	ProjectTasksInverseTable = "project_tasks"
	// ProjectTasksColumn is the table column denoting the project_tasks relation/edge.
	ProjectTasksColumn = "task_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTaskParentID,
	FieldTaskPriorityID,
	FieldAssigneeID,
	FieldCreatedBy,
	FieldCompleted,
	FieldCompletedAt,
	FieldIsNew,
	FieldName,
	FieldDueDate,
	FieldDueTime,
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
	// DefaultCompleted holds the default value on creation for the "completed" field.
	DefaultCompleted bool
	// DefaultIsNew holds the default value on creation for the "is_new" field.
	DefaultIsNew bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)