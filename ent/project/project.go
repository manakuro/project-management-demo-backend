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
	// EdgeProjectBaseColor holds the string denoting the projectbasecolor edge name in mutations.
	EdgeProjectBaseColor = "projectBaseColor"
	// EdgeProjectLightColor holds the string denoting the projectlightcolor edge name in mutations.
	EdgeProjectLightColor = "projectLightColor"
	// EdgeProjectIcon holds the string denoting the projecticon edge name in mutations.
	EdgeProjectIcon = "projectIcon"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeProjectTeammates holds the string denoting the projectteammates edge name in mutations.
	EdgeProjectTeammates = "projectTeammates"
	// EdgeFavoriteProjects holds the string denoting the favoriteprojects edge name in mutations.
	EdgeFavoriteProjects = "favoriteProjects"
	// EdgeProjectTaskColumns holds the string denoting the projecttaskcolumns edge name in mutations.
	EdgeProjectTaskColumns = "projectTaskColumns"
	// EdgeProjectTaskListStatuses holds the string denoting the projecttaskliststatuses edge name in mutations.
	EdgeProjectTaskListStatuses = "projectTaskListStatuses"
	// EdgeProjectTaskSections holds the string denoting the projecttasksections edge name in mutations.
	EdgeProjectTaskSections = "projectTaskSections"
	// EdgeProjectTasks holds the string denoting the projecttasks edge name in mutations.
	EdgeProjectTasks = "projectTasks"
	// EdgeTaskFiles holds the string denoting the taskfiles edge name in mutations.
	EdgeTaskFiles = "taskFiles"
	// EdgeWorkspaceActivities holds the string denoting the workspaceactivities edge name in mutations.
	EdgeWorkspaceActivities = "workspaceActivities"
	// EdgeArchivedWorkspaceActivities holds the string denoting the archivedworkspaceactivities edge name in mutations.
	EdgeArchivedWorkspaceActivities = "archivedWorkspaceActivities"
	// EdgeDeletedProjectTasks holds the string denoting the deletedprojecttasks edge name in mutations.
	EdgeDeletedProjectTasks = "deletedProjectTasks"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "projects"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
	// ProjectBaseColorTable is the table that holds the projectBaseColor relation/edge.
	ProjectBaseColorTable = "projects"
	// ProjectBaseColorInverseTable is the table name for the ProjectBaseColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectbasecolor" package.
	ProjectBaseColorInverseTable = "project_base_colors"
	// ProjectBaseColorColumn is the table column denoting the projectBaseColor relation/edge.
	ProjectBaseColorColumn = "project_base_color_id"
	// ProjectLightColorTable is the table that holds the projectLightColor relation/edge.
	ProjectLightColorTable = "projects"
	// ProjectLightColorInverseTable is the table name for the ProjectLightColor entity.
	// It exists in this package in order to avoid circular dependency with the "projectlightcolor" package.
	ProjectLightColorInverseTable = "project_light_colors"
	// ProjectLightColorColumn is the table column denoting the projectLightColor relation/edge.
	ProjectLightColorColumn = "project_light_color_id"
	// ProjectIconTable is the table that holds the projectIcon relation/edge.
	ProjectIconTable = "projects"
	// ProjectIconInverseTable is the table name for the ProjectIcon entity.
	// It exists in this package in order to avoid circular dependency with the "projecticon" package.
	ProjectIconInverseTable = "project_icons"
	// ProjectIconColumn is the table column denoting the projectIcon relation/edge.
	ProjectIconColumn = "project_icon_id"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "projects"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "created_by"
	// ProjectTeammatesTable is the table that holds the projectTeammates relation/edge.
	ProjectTeammatesTable = "project_teammates"
	// ProjectTeammatesInverseTable is the table name for the ProjectTeammate entity.
	// It exists in this package in order to avoid circular dependency with the "projectteammate" package.
	ProjectTeammatesInverseTable = "project_teammates"
	// ProjectTeammatesColumn is the table column denoting the projectTeammates relation/edge.
	ProjectTeammatesColumn = "project_id"
	// FavoriteProjectsTable is the table that holds the favoriteProjects relation/edge.
	FavoriteProjectsTable = "favorite_projects"
	// FavoriteProjectsInverseTable is the table name for the FavoriteProject entity.
	// It exists in this package in order to avoid circular dependency with the "favoriteproject" package.
	FavoriteProjectsInverseTable = "favorite_projects"
	// FavoriteProjectsColumn is the table column denoting the favoriteProjects relation/edge.
	FavoriteProjectsColumn = "project_id"
	// ProjectTaskColumnsTable is the table that holds the projectTaskColumns relation/edge.
	ProjectTaskColumnsTable = "project_task_columns"
	// ProjectTaskColumnsInverseTable is the table name for the ProjectTaskColumn entity.
	// It exists in this package in order to avoid circular dependency with the "projecttaskcolumn" package.
	ProjectTaskColumnsInverseTable = "project_task_columns"
	// ProjectTaskColumnsColumn is the table column denoting the projectTaskColumns relation/edge.
	ProjectTaskColumnsColumn = "project_id"
	// ProjectTaskListStatusesTable is the table that holds the projectTaskListStatuses relation/edge.
	ProjectTaskListStatusesTable = "project_task_list_status"
	// ProjectTaskListStatusesInverseTable is the table name for the ProjectTaskListStatus entity.
	// It exists in this package in order to avoid circular dependency with the "projecttaskliststatus" package.
	ProjectTaskListStatusesInverseTable = "project_task_list_status"
	// ProjectTaskListStatusesColumn is the table column denoting the projectTaskListStatuses relation/edge.
	ProjectTaskListStatusesColumn = "project_id"
	// ProjectTaskSectionsTable is the table that holds the projectTaskSections relation/edge.
	ProjectTaskSectionsTable = "project_task_sections"
	// ProjectTaskSectionsInverseTable is the table name for the ProjectTaskSection entity.
	// It exists in this package in order to avoid circular dependency with the "projecttasksection" package.
	ProjectTaskSectionsInverseTable = "project_task_sections"
	// ProjectTaskSectionsColumn is the table column denoting the projectTaskSections relation/edge.
	ProjectTaskSectionsColumn = "project_id"
	// ProjectTasksTable is the table that holds the projectTasks relation/edge.
	ProjectTasksTable = "project_tasks"
	// ProjectTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	ProjectTasksInverseTable = "project_tasks"
	// ProjectTasksColumn is the table column denoting the projectTasks relation/edge.
	ProjectTasksColumn = "project_id"
	// TaskFilesTable is the table that holds the taskFiles relation/edge.
	TaskFilesTable = "task_files"
	// TaskFilesInverseTable is the table name for the TaskFile entity.
	// It exists in this package in order to avoid circular dependency with the "taskfile" package.
	TaskFilesInverseTable = "task_files"
	// TaskFilesColumn is the table column denoting the taskFiles relation/edge.
	TaskFilesColumn = "project_id"
	// WorkspaceActivitiesTable is the table that holds the workspaceActivities relation/edge.
	WorkspaceActivitiesTable = "workspace_activities"
	// WorkspaceActivitiesInverseTable is the table name for the WorkspaceActivity entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceactivity" package.
	WorkspaceActivitiesInverseTable = "workspace_activities"
	// WorkspaceActivitiesColumn is the table column denoting the workspaceActivities relation/edge.
	WorkspaceActivitiesColumn = "project_id"
	// ArchivedWorkspaceActivitiesTable is the table that holds the archivedWorkspaceActivities relation/edge.
	ArchivedWorkspaceActivitiesTable = "archived_workspace_activities"
	// ArchivedWorkspaceActivitiesInverseTable is the table name for the ArchivedWorkspaceActivity entity.
	// It exists in this package in order to avoid circular dependency with the "archivedworkspaceactivity" package.
	ArchivedWorkspaceActivitiesInverseTable = "archived_workspace_activities"
	// ArchivedWorkspaceActivitiesColumn is the table column denoting the archivedWorkspaceActivities relation/edge.
	ArchivedWorkspaceActivitiesColumn = "project_id"
	// DeletedProjectTasksTable is the table that holds the deletedProjectTasks relation/edge.
	DeletedProjectTasksTable = "deleted_project_tasks"
	// DeletedProjectTasksInverseTable is the table name for the DeletedProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "deletedprojecttask" package.
	DeletedProjectTasksInverseTable = "deleted_project_tasks"
	// DeletedProjectTasksColumn is the table column denoting the deletedProjectTasks relation/edge.
	DeletedProjectTasksColumn = "project_id"
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
