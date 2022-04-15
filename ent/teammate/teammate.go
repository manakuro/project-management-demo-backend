// Code generated by entc, DO NOT EDIT.

package teammate

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the teammate type in the database.
	Label = "teammate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWorkspaces holds the string denoting the workspaces edge name in mutations.
	EdgeWorkspaces = "workspaces"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeProjectTeammates holds the string denoting the projectteammates edge name in mutations.
	EdgeProjectTeammates = "projectTeammates"
	// EdgeWorkspaceTeammates holds the string denoting the workspaceteammates edge name in mutations.
	EdgeWorkspaceTeammates = "workspaceTeammates"
	// EdgeFavoriteProjects holds the string denoting the favoriteprojects edge name in mutations.
	EdgeFavoriteProjects = "favoriteProjects"
	// EdgeFavoriteWorkspaces holds the string denoting the favoriteworkspaces edge name in mutations.
	EdgeFavoriteWorkspaces = "favoriteWorkspaces"
	// EdgeTeammateTaskTabStatuses holds the string denoting the teammatetasktabstatuses edge name in mutations.
	EdgeTeammateTaskTabStatuses = "teammateTaskTabStatuses"
	// EdgeTeammateTaskColumns holds the string denoting the teammatetaskcolumns edge name in mutations.
	EdgeTeammateTaskColumns = "teammateTaskColumns"
	// EdgeTeammateTaskListStatuses holds the string denoting the teammatetaskliststatuses edge name in mutations.
	EdgeTeammateTaskListStatuses = "teammateTaskListStatuses"
	// EdgeTeammateTaskSections holds the string denoting the teammatetasksections edge name in mutations.
	EdgeTeammateTaskSections = "teammateTaskSections"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgeTeammateTasks holds the string denoting the teammatetasks edge name in mutations.
	EdgeTeammateTasks = "teammateTasks"
	// EdgeTaskLikes holds the string denoting the tasklikes edge name in mutations.
	EdgeTaskLikes = "taskLikes"
	// EdgeTaskCollaborators holds the string denoting the taskcollaborators edge name in mutations.
	EdgeTaskCollaborators = "taskCollaborators"
	// EdgeTaskFeeds holds the string denoting the taskfeeds edge name in mutations.
	EdgeTaskFeeds = "taskFeeds"
	// EdgeTaskFeedLikes holds the string denoting the taskfeedlikes edge name in mutations.
	EdgeTaskFeedLikes = "taskFeedLikes"
	// EdgeTaskActivities holds the string denoting the taskactivities edge name in mutations.
	EdgeTaskActivities = "taskActivities"
	// EdgeWorkspaceActivities holds the string denoting the workspaceactivities edge name in mutations.
	EdgeWorkspaceActivities = "workspaceActivities"
	// EdgeArchivedTaskActivities holds the string denoting the archivedtaskactivities edge name in mutations.
	EdgeArchivedTaskActivities = "archivedTaskActivities"
	// Table holds the table name of the teammate in the database.
	Table = "teammates"
	// WorkspacesTable is the table that holds the workspaces relation/edge.
	WorkspacesTable = "workspaces"
	// WorkspacesInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspacesInverseTable = "workspaces"
	// WorkspacesColumn is the table column denoting the workspaces relation/edge.
	WorkspacesColumn = "created_by"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "created_by"
	// ProjectTeammatesTable is the table that holds the projectTeammates relation/edge.
	ProjectTeammatesTable = "project_teammates"
	// ProjectTeammatesInverseTable is the table name for the ProjectTeammate entity.
	// It exists in this package in order to avoid circular dependency with the "projectteammate" package.
	ProjectTeammatesInverseTable = "project_teammates"
	// ProjectTeammatesColumn is the table column denoting the projectTeammates relation/edge.
	ProjectTeammatesColumn = "teammate_id"
	// WorkspaceTeammatesTable is the table that holds the workspaceTeammates relation/edge.
	WorkspaceTeammatesTable = "workspace_teammates"
	// WorkspaceTeammatesInverseTable is the table name for the WorkspaceTeammate entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceteammate" package.
	WorkspaceTeammatesInverseTable = "workspace_teammates"
	// WorkspaceTeammatesColumn is the table column denoting the workspaceTeammates relation/edge.
	WorkspaceTeammatesColumn = "teammate_id"
	// FavoriteProjectsTable is the table that holds the favoriteProjects relation/edge.
	FavoriteProjectsTable = "favorite_projects"
	// FavoriteProjectsInverseTable is the table name for the FavoriteProject entity.
	// It exists in this package in order to avoid circular dependency with the "favoriteproject" package.
	FavoriteProjectsInverseTable = "favorite_projects"
	// FavoriteProjectsColumn is the table column denoting the favoriteProjects relation/edge.
	FavoriteProjectsColumn = "teammate_id"
	// FavoriteWorkspacesTable is the table that holds the favoriteWorkspaces relation/edge.
	FavoriteWorkspacesTable = "favorite_workspaces"
	// FavoriteWorkspacesInverseTable is the table name for the FavoriteWorkspace entity.
	// It exists in this package in order to avoid circular dependency with the "favoriteworkspace" package.
	FavoriteWorkspacesInverseTable = "favorite_workspaces"
	// FavoriteWorkspacesColumn is the table column denoting the favoriteWorkspaces relation/edge.
	FavoriteWorkspacesColumn = "teammate_id"
	// TeammateTaskTabStatusesTable is the table that holds the teammateTaskTabStatuses relation/edge.
	TeammateTaskTabStatusesTable = "teammate_task_tab_status"
	// TeammateTaskTabStatusesInverseTable is the table name for the TeammateTaskTabStatus entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetasktabstatus" package.
	TeammateTaskTabStatusesInverseTable = "teammate_task_tab_status"
	// TeammateTaskTabStatusesColumn is the table column denoting the teammateTaskTabStatuses relation/edge.
	TeammateTaskTabStatusesColumn = "teammate_id"
	// TeammateTaskColumnsTable is the table that holds the teammateTaskColumns relation/edge.
	TeammateTaskColumnsTable = "teammate_task_columns"
	// TeammateTaskColumnsInverseTable is the table name for the TeammateTaskColumn entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskcolumn" package.
	TeammateTaskColumnsInverseTable = "teammate_task_columns"
	// TeammateTaskColumnsColumn is the table column denoting the teammateTaskColumns relation/edge.
	TeammateTaskColumnsColumn = "teammate_id"
	// TeammateTaskListStatusesTable is the table that holds the teammateTaskListStatuses relation/edge.
	TeammateTaskListStatusesTable = "teammate_task_list_status"
	// TeammateTaskListStatusesInverseTable is the table name for the TeammateTaskListStatus entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskliststatus" package.
	TeammateTaskListStatusesInverseTable = "teammate_task_list_status"
	// TeammateTaskListStatusesColumn is the table column denoting the teammateTaskListStatuses relation/edge.
	TeammateTaskListStatusesColumn = "teammate_id"
	// TeammateTaskSectionsTable is the table that holds the teammateTaskSections relation/edge.
	TeammateTaskSectionsTable = "teammate_task_sections"
	// TeammateTaskSectionsInverseTable is the table name for the TeammateTaskSection entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetasksection" package.
	TeammateTaskSectionsInverseTable = "teammate_task_sections"
	// TeammateTaskSectionsColumn is the table column denoting the teammateTaskSections relation/edge.
	TeammateTaskSectionsColumn = "teammate_id"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "assignee_id"
	// TeammateTasksTable is the table that holds the teammateTasks relation/edge.
	TeammateTasksTable = "teammate_tasks"
	// TeammateTasksInverseTable is the table name for the TeammateTask entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetask" package.
	TeammateTasksInverseTable = "teammate_tasks"
	// TeammateTasksColumn is the table column denoting the teammateTasks relation/edge.
	TeammateTasksColumn = "teammate_id"
	// TaskLikesTable is the table that holds the taskLikes relation/edge.
	TaskLikesTable = "task_likes"
	// TaskLikesInverseTable is the table name for the TaskLike entity.
	// It exists in this package in order to avoid circular dependency with the "tasklike" package.
	TaskLikesInverseTable = "task_likes"
	// TaskLikesColumn is the table column denoting the taskLikes relation/edge.
	TaskLikesColumn = "teammate_id"
	// TaskCollaboratorsTable is the table that holds the taskCollaborators relation/edge.
	TaskCollaboratorsTable = "task_collaborators"
	// TaskCollaboratorsInverseTable is the table name for the TaskCollaborator entity.
	// It exists in this package in order to avoid circular dependency with the "taskcollaborator" package.
	TaskCollaboratorsInverseTable = "task_collaborators"
	// TaskCollaboratorsColumn is the table column denoting the taskCollaborators relation/edge.
	TaskCollaboratorsColumn = "teammate_id"
	// TaskFeedsTable is the table that holds the taskFeeds relation/edge.
	TaskFeedsTable = "task_feeds"
	// TaskFeedsInverseTable is the table name for the TaskFeed entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeed" package.
	TaskFeedsInverseTable = "task_feeds"
	// TaskFeedsColumn is the table column denoting the taskFeeds relation/edge.
	TaskFeedsColumn = "teammate_id"
	// TaskFeedLikesTable is the table that holds the taskFeedLikes relation/edge.
	TaskFeedLikesTable = "task_feed_likes"
	// TaskFeedLikesInverseTable is the table name for the TaskFeedLike entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeedlike" package.
	TaskFeedLikesInverseTable = "task_feed_likes"
	// TaskFeedLikesColumn is the table column denoting the taskFeedLikes relation/edge.
	TaskFeedLikesColumn = "teammate_id"
	// TaskActivitiesTable is the table that holds the taskActivities relation/edge.
	TaskActivitiesTable = "task_activities"
	// TaskActivitiesInverseTable is the table name for the TaskActivity entity.
	// It exists in this package in order to avoid circular dependency with the "taskactivity" package.
	TaskActivitiesInverseTable = "task_activities"
	// TaskActivitiesColumn is the table column denoting the taskActivities relation/edge.
	TaskActivitiesColumn = "teammate_id"
	// WorkspaceActivitiesTable is the table that holds the workspaceActivities relation/edge.
	WorkspaceActivitiesTable = "workspace_activities"
	// WorkspaceActivitiesInverseTable is the table name for the WorkspaceActivity entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceactivity" package.
	WorkspaceActivitiesInverseTable = "workspace_activities"
	// WorkspaceActivitiesColumn is the table column denoting the workspaceActivities relation/edge.
	WorkspaceActivitiesColumn = "teammate_id"
	// ArchivedTaskActivitiesTable is the table that holds the archivedTaskActivities relation/edge.
	ArchivedTaskActivitiesTable = "archived_task_activities"
	// ArchivedTaskActivitiesInverseTable is the table name for the ArchivedTaskActivity entity.
	// It exists in this package in order to avoid circular dependency with the "archivedtaskactivity" package.
	ArchivedTaskActivitiesInverseTable = "archived_task_activities"
	// ArchivedTaskActivitiesColumn is the table column denoting the archivedTaskActivities relation/edge.
	ArchivedTaskActivitiesColumn = "teammate_id"
)

// Columns holds all SQL columns for teammate fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldImage,
	FieldEmail,
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
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
