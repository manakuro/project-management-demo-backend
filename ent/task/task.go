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
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTaskPriority holds the string denoting the taskpriority edge name in mutations.
	EdgeTaskPriority = "taskPriority"
	// EdgeSubTasks holds the string denoting the subtasks edge name in mutations.
	EdgeSubTasks = "subTasks"
	// EdgeParentTask holds the string denoting the parenttask edge name in mutations.
	EdgeParentTask = "parentTask"
	// EdgeTeammateTasks holds the string denoting the teammatetasks edge name in mutations.
	EdgeTeammateTasks = "teammateTasks"
	// EdgeProjectTasks holds the string denoting the projecttasks edge name in mutations.
	EdgeProjectTasks = "projectTasks"
	// EdgeTaskLikes holds the string denoting the tasklikes edge name in mutations.
	EdgeTaskLikes = "taskLikes"
	// EdgeTaskTags holds the string denoting the tasktags edge name in mutations.
	EdgeTaskTags = "taskTags"
	// EdgeTaskCollaborators holds the string denoting the taskcollaborators edge name in mutations.
	EdgeTaskCollaborators = "taskCollaborators"
	// EdgeTaskFeeds holds the string denoting the taskfeeds edge name in mutations.
	EdgeTaskFeeds = "taskFeeds"
	// EdgeTaskFeedLikes holds the string denoting the taskfeedlikes edge name in mutations.
	EdgeTaskFeedLikes = "taskFeedLikes"
	// EdgeTaskFiles holds the string denoting the taskfiles edge name in mutations.
	EdgeTaskFiles = "taskFiles"
	// EdgeDeletedTasksRef holds the string denoting the deletedtasksref edge name in mutations.
	EdgeDeletedTasksRef = "deletedTasksRef"
	// EdgeTaskActivityTasks holds the string denoting the taskactivitytasks edge name in mutations.
	EdgeTaskActivityTasks = "taskActivityTasks"
	// EdgeWorkspaceActivityTasks holds the string denoting the workspaceactivitytasks edge name in mutations.
	EdgeWorkspaceActivityTasks = "workspaceActivityTasks"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "tasks"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "assignee_id"
	// TaskPriorityTable is the table that holds the taskPriority relation/edge.
	TaskPriorityTable = "tasks"
	// TaskPriorityInverseTable is the table name for the TaskPriority entity.
	// It exists in this package in order to avoid circular dependency with the "taskpriority" package.
	TaskPriorityInverseTable = "task_priorities"
	// TaskPriorityColumn is the table column denoting the taskPriority relation/edge.
	TaskPriorityColumn = "task_priority_id"
	// SubTasksTable is the table that holds the subTasks relation/edge.
	SubTasksTable = "tasks"
	// SubTasksColumn is the table column denoting the subTasks relation/edge.
	SubTasksColumn = "task_parent_id"
	// ParentTaskTable is the table that holds the parentTask relation/edge.
	ParentTaskTable = "tasks"
	// ParentTaskColumn is the table column denoting the parentTask relation/edge.
	ParentTaskColumn = "task_parent_id"
	// TeammateTasksTable is the table that holds the teammateTasks relation/edge.
	TeammateTasksTable = "teammate_tasks"
	// TeammateTasksInverseTable is the table name for the TeammateTask entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetask" package.
	TeammateTasksInverseTable = "teammate_tasks"
	// TeammateTasksColumn is the table column denoting the teammateTasks relation/edge.
	TeammateTasksColumn = "task_id"
	// ProjectTasksTable is the table that holds the projectTasks relation/edge.
	ProjectTasksTable = "project_tasks"
	// ProjectTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	ProjectTasksInverseTable = "project_tasks"
	// ProjectTasksColumn is the table column denoting the projectTasks relation/edge.
	ProjectTasksColumn = "task_id"
	// TaskLikesTable is the table that holds the taskLikes relation/edge.
	TaskLikesTable = "task_likes"
	// TaskLikesInverseTable is the table name for the TaskLike entity.
	// It exists in this package in order to avoid circular dependency with the "tasklike" package.
	TaskLikesInverseTable = "task_likes"
	// TaskLikesColumn is the table column denoting the taskLikes relation/edge.
	TaskLikesColumn = "task_id"
	// TaskTagsTable is the table that holds the taskTags relation/edge.
	TaskTagsTable = "task_tags"
	// TaskTagsInverseTable is the table name for the TaskTag entity.
	// It exists in this package in order to avoid circular dependency with the "tasktag" package.
	TaskTagsInverseTable = "task_tags"
	// TaskTagsColumn is the table column denoting the taskTags relation/edge.
	TaskTagsColumn = "task_id"
	// TaskCollaboratorsTable is the table that holds the taskCollaborators relation/edge.
	TaskCollaboratorsTable = "task_collaborators"
	// TaskCollaboratorsInverseTable is the table name for the TaskCollaborator entity.
	// It exists in this package in order to avoid circular dependency with the "taskcollaborator" package.
	TaskCollaboratorsInverseTable = "task_collaborators"
	// TaskCollaboratorsColumn is the table column denoting the taskCollaborators relation/edge.
	TaskCollaboratorsColumn = "task_id"
	// TaskFeedsTable is the table that holds the taskFeeds relation/edge.
	TaskFeedsTable = "task_feeds"
	// TaskFeedsInverseTable is the table name for the TaskFeed entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeed" package.
	TaskFeedsInverseTable = "task_feeds"
	// TaskFeedsColumn is the table column denoting the taskFeeds relation/edge.
	TaskFeedsColumn = "task_id"
	// TaskFeedLikesTable is the table that holds the taskFeedLikes relation/edge.
	TaskFeedLikesTable = "task_feed_likes"
	// TaskFeedLikesInverseTable is the table name for the TaskFeedLike entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeedlike" package.
	TaskFeedLikesInverseTable = "task_feed_likes"
	// TaskFeedLikesColumn is the table column denoting the taskFeedLikes relation/edge.
	TaskFeedLikesColumn = "task_id"
	// TaskFilesTable is the table that holds the taskFiles relation/edge.
	TaskFilesTable = "task_files"
	// TaskFilesInverseTable is the table name for the TaskFile entity.
	// It exists in this package in order to avoid circular dependency with the "taskfile" package.
	TaskFilesInverseTable = "task_files"
	// TaskFilesColumn is the table column denoting the taskFiles relation/edge.
	TaskFilesColumn = "task_id"
	// DeletedTasksRefTable is the table that holds the deletedTasksRef relation/edge.
	DeletedTasksRefTable = "deleted_tasks"
	// DeletedTasksRefInverseTable is the table name for the DeletedTask entity.
	// It exists in this package in order to avoid circular dependency with the "deletedtask" package.
	DeletedTasksRefInverseTable = "deleted_tasks"
	// DeletedTasksRefColumn is the table column denoting the deletedTasksRef relation/edge.
	DeletedTasksRefColumn = "task_id"
	// TaskActivityTasksTable is the table that holds the taskActivityTasks relation/edge.
	TaskActivityTasksTable = "task_activity_tasks"
	// TaskActivityTasksInverseTable is the table name for the TaskActivityTask entity.
	// It exists in this package in order to avoid circular dependency with the "taskactivitytask" package.
	TaskActivityTasksInverseTable = "task_activity_tasks"
	// TaskActivityTasksColumn is the table column denoting the taskActivityTasks relation/edge.
	TaskActivityTasksColumn = "task_id"
	// WorkspaceActivityTasksTable is the table that holds the workspaceActivityTasks relation/edge.
	WorkspaceActivityTasksTable = "workspace_activity_tasks"
	// WorkspaceActivityTasksInverseTable is the table name for the WorkspaceActivityTask entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceactivitytask" package.
	WorkspaceActivityTasksInverseTable = "workspace_activity_tasks"
	// WorkspaceActivityTasksColumn is the table column denoting the workspaceActivityTasks relation/edge.
	WorkspaceActivityTasksColumn = "task_id"
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
	FieldDescription,
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
