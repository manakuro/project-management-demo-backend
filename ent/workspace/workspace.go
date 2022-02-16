// Code generated by entc, DO NOT EDIT.

package workspace

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the workspace type in the database.
	Label = "workspace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeWorkspaceTeammates holds the string denoting the workspaceteammates edge name in mutations.
	EdgeWorkspaceTeammates = "workspaceTeammates"
	// EdgeFavoriteWorkspaces holds the string denoting the favoriteworkspaces edge name in mutations.
	EdgeFavoriteWorkspaces = "favoriteWorkspaces"
	// EdgeTeammateTaskTabStatuses holds the string denoting the teammatetasktabstatuses edge name in mutations.
	EdgeTeammateTaskTabStatuses = "teammateTaskTabStatuses"
	// EdgeTeammateTaskListStatuses holds the string denoting the teammatetaskliststatuses edge name in mutations.
	EdgeTeammateTaskListStatuses = "teammateTaskListStatuses"
	// EdgeTeammateTaskSections holds the string denoting the teammatetasksections edge name in mutations.
	EdgeTeammateTaskSections = "teammateTaskSections"
	// EdgeTaskLikes holds the string denoting the tasklikes edge name in mutations.
	EdgeTaskLikes = "taskLikes"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeTeammateTaskColumns holds the string denoting the teammatetaskcolumns edge name in mutations.
	EdgeTeammateTaskColumns = "teammateTaskColumns"
	// EdgeTeammateTasks holds the string denoting the teammatetasks edge name in mutations.
	EdgeTeammateTasks = "teammateTasks"
	// Table holds the table name of the workspace in the database.
	Table = "workspaces"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "workspaces"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "created_by"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "workspace_id"
	// WorkspaceTeammatesTable is the table that holds the workspaceTeammates relation/edge.
	WorkspaceTeammatesTable = "workspace_teammates"
	// WorkspaceTeammatesInverseTable is the table name for the WorkspaceTeammate entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceteammate" package.
	WorkspaceTeammatesInverseTable = "workspace_teammates"
	// WorkspaceTeammatesColumn is the table column denoting the workspaceTeammates relation/edge.
	WorkspaceTeammatesColumn = "workspace_id"
	// FavoriteWorkspacesTable is the table that holds the favoriteWorkspaces relation/edge.
	FavoriteWorkspacesTable = "favorite_workspaces"
	// FavoriteWorkspacesInverseTable is the table name for the FavoriteWorkspace entity.
	// It exists in this package in order to avoid circular dependency with the "favoriteworkspace" package.
	FavoriteWorkspacesInverseTable = "favorite_workspaces"
	// FavoriteWorkspacesColumn is the table column denoting the favoriteWorkspaces relation/edge.
	FavoriteWorkspacesColumn = "workspace_id"
	// TeammateTaskTabStatusesTable is the table that holds the teammateTaskTabStatuses relation/edge.
	TeammateTaskTabStatusesTable = "teammate_task_tab_status"
	// TeammateTaskTabStatusesInverseTable is the table name for the TeammateTaskTabStatus entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetasktabstatus" package.
	TeammateTaskTabStatusesInverseTable = "teammate_task_tab_status"
	// TeammateTaskTabStatusesColumn is the table column denoting the teammateTaskTabStatuses relation/edge.
	TeammateTaskTabStatusesColumn = "workspace_id"
	// TeammateTaskListStatusesTable is the table that holds the teammateTaskListStatuses relation/edge.
	TeammateTaskListStatusesTable = "teammate_task_list_status"
	// TeammateTaskListStatusesInverseTable is the table name for the TeammateTaskListStatus entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskliststatus" package.
	TeammateTaskListStatusesInverseTable = "teammate_task_list_status"
	// TeammateTaskListStatusesColumn is the table column denoting the teammateTaskListStatuses relation/edge.
	TeammateTaskListStatusesColumn = "workspace_id"
	// TeammateTaskSectionsTable is the table that holds the teammateTaskSections relation/edge.
	TeammateTaskSectionsTable = "teammate_task_sections"
	// TeammateTaskSectionsInverseTable is the table name for the TeammateTaskSection entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetasksection" package.
	TeammateTaskSectionsInverseTable = "teammate_task_sections"
	// TeammateTaskSectionsColumn is the table column denoting the teammateTaskSections relation/edge.
	TeammateTaskSectionsColumn = "workspace_id"
	// TaskLikesTable is the table that holds the taskLikes relation/edge.
	TaskLikesTable = "task_likes"
	// TaskLikesInverseTable is the table name for the TaskLike entity.
	// It exists in this package in order to avoid circular dependency with the "tasklike" package.
	TaskLikesInverseTable = "task_likes"
	// TaskLikesColumn is the table column denoting the taskLikes relation/edge.
	TaskLikesColumn = "workspace_id"
	// TagsTable is the table that holds the tags relation/edge.
	TagsTable = "tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TagsColumn is the table column denoting the tags relation/edge.
	TagsColumn = "workspace_id"
	// TeammateTaskColumnsTable is the table that holds the teammateTaskColumns relation/edge.
	TeammateTaskColumnsTable = "teammate_task_columns"
	// TeammateTaskColumnsInverseTable is the table name for the TeammateTaskColumn entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskcolumn" package.
	TeammateTaskColumnsInverseTable = "teammate_task_columns"
	// TeammateTaskColumnsColumn is the table column denoting the teammateTaskColumns relation/edge.
	TeammateTaskColumnsColumn = "workspace_id"
	// TeammateTasksTable is the table that holds the teammateTasks relation/edge.
	TeammateTasksTable = "teammate_tasks"
	// TeammateTasksInverseTable is the table name for the TeammateTask entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetask" package.
	TeammateTasksInverseTable = "teammate_tasks"
	// TeammateTasksColumn is the table column denoting the teammateTasks relation/edge.
	TeammateTasksColumn = "workspace_id"
)

// Columns holds all SQL columns for workspace fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldName,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
