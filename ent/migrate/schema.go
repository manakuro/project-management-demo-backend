// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ColorsColumns holds the columns for the "colors" table.
	ColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "color", Type: field.TypeString, Size: 255},
		{Name: "hex", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// ColorsTable holds the schema information for the "colors" table.
	ColorsTable = &schema.Table{
		Name:       "colors",
		Columns:    ColorsColumns,
		PrimaryKey: []*schema.Column{ColorsColumns[0]},
	}
	// FavoriteProjectsColumns holds the columns for the "favorite_projects" table.
	FavoriteProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "project_id", Type: field.TypeString, Nullable: true},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
	}
	// FavoriteProjectsTable holds the schema information for the "favorite_projects" table.
	FavoriteProjectsTable = &schema.Table{
		Name:       "favorite_projects",
		Columns:    FavoriteProjectsColumns,
		PrimaryKey: []*schema.Column{FavoriteProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "favorite_projects_projects_favorite_projects",
				Columns:    []*schema.Column{FavoriteProjectsColumns[3]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "favorite_projects_teammates_favorite_projects",
				Columns:    []*schema.Column{FavoriteProjectsColumns[4]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FavoriteWorkspacesColumns holds the columns for the "favorite_workspaces" table.
	FavoriteWorkspacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
		{Name: "workspace_id", Type: field.TypeString, Nullable: true},
	}
	// FavoriteWorkspacesTable holds the schema information for the "favorite_workspaces" table.
	FavoriteWorkspacesTable = &schema.Table{
		Name:       "favorite_workspaces",
		Columns:    FavoriteWorkspacesColumns,
		PrimaryKey: []*schema.Column{FavoriteWorkspacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "favorite_workspaces_teammates_favorite_workspaces",
				Columns:    []*schema.Column{FavoriteWorkspacesColumns[3]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "favorite_workspaces_workspaces_favorite_workspaces",
				Columns:    []*schema.Column{FavoriteWorkspacesColumns[4]},
				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// IconsColumns holds the columns for the "icons" table.
	IconsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "icon", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// IconsTable holds the schema information for the "icons" table.
	IconsTable = &schema.Table{
		Name:       "icons",
		Columns:    IconsColumns,
		PrimaryKey: []*schema.Column{IconsColumns[0]},
	}
	// MyTasksTabStatusColumns holds the columns for the "my_tasks_tab_status" table.
	MyTasksTabStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"LIST", "BOARD", "CALENDAR", "FILES"}, Default: "LIST"},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
		{Name: "workspace_id", Type: field.TypeString, Nullable: true},
	}
	// MyTasksTabStatusTable holds the schema information for the "my_tasks_tab_status" table.
	MyTasksTabStatusTable = &schema.Table{
		Name:       "my_tasks_tab_status",
		Columns:    MyTasksTabStatusColumns,
		PrimaryKey: []*schema.Column{MyTasksTabStatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "my_tasks_tab_status_teammates_my_tasks_tab_statuses",
				Columns:    []*schema.Column{MyTasksTabStatusColumns[4]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "my_tasks_tab_status_workspaces_my_tasks_tab_statuses",
				Columns:    []*schema.Column{MyTasksTabStatusColumns[5]},
				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeJSON},
		{Name: "description_title", Type: field.TypeString, Size: 255},
		{Name: "due_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "project_base_color_id", Type: field.TypeString, Nullable: true},
		{Name: "project_icon_id", Type: field.TypeString, Nullable: true},
		{Name: "project_light_color_id", Type: field.TypeString, Nullable: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "workspace_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_project_base_colors_projects",
				Columns:    []*schema.Column{ProjectsColumns[7]},
				RefColumns: []*schema.Column{ProjectBaseColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_project_icons_projects",
				Columns:    []*schema.Column{ProjectsColumns[8]},
				RefColumns: []*schema.Column{ProjectIconsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_project_light_colors_projects",
				Columns:    []*schema.Column{ProjectsColumns[9]},
				RefColumns: []*schema.Column{ProjectLightColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_teammates_projects",
				Columns:    []*schema.Column{ProjectsColumns[10]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_workspaces_projects",
				Columns:    []*schema.Column{ProjectsColumns[11]},
				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectBaseColorsColumns holds the columns for the "project_base_colors" table.
	ProjectBaseColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "color_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectBaseColorsTable holds the schema information for the "project_base_colors" table.
	ProjectBaseColorsTable = &schema.Table{
		Name:       "project_base_colors",
		Columns:    ProjectBaseColorsColumns,
		PrimaryKey: []*schema.Column{ProjectBaseColorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_base_colors_colors_project_base_colors",
				Columns:    []*schema.Column{ProjectBaseColorsColumns[3]},
				RefColumns: []*schema.Column{ColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectIconsColumns holds the columns for the "project_icons" table.
	ProjectIconsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "icon_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectIconsTable holds the schema information for the "project_icons" table.
	ProjectIconsTable = &schema.Table{
		Name:       "project_icons",
		Columns:    ProjectIconsColumns,
		PrimaryKey: []*schema.Column{ProjectIconsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_icons_icons_project_icons",
				Columns:    []*schema.Column{ProjectIconsColumns[3]},
				RefColumns: []*schema.Column{IconsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectLightColorsColumns holds the columns for the "project_light_colors" table.
	ProjectLightColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "color_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectLightColorsTable holds the schema information for the "project_light_colors" table.
	ProjectLightColorsTable = &schema.Table{
		Name:       "project_light_colors",
		Columns:    ProjectLightColorsColumns,
		PrimaryKey: []*schema.Column{ProjectLightColorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_light_colors_colors_project_light_colors",
				Columns:    []*schema.Column{ProjectLightColorsColumns[3]},
				RefColumns: []*schema.Column{ColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectTaskColumnsColumns holds the columns for the "project_task_columns" table.
	ProjectTaskColumnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "width", Type: field.TypeString, Size: 255},
		{Name: "disabled", Type: field.TypeBool},
		{Name: "customizable", Type: field.TypeBool},
		{Name: "order", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "project_id", Type: field.TypeString, Nullable: true},
		{Name: "task_column_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectTaskColumnsTable holds the schema information for the "project_task_columns" table.
	ProjectTaskColumnsTable = &schema.Table{
		Name:       "project_task_columns",
		Columns:    ProjectTaskColumnsColumns,
		PrimaryKey: []*schema.Column{ProjectTaskColumnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_task_columns_projects_project_task_columns",
				Columns:    []*schema.Column{ProjectTaskColumnsColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_task_columns_task_columns_project_task_columns",
				Columns:    []*schema.Column{ProjectTaskColumnsColumns[8]},
				RefColumns: []*schema.Column{TaskColumnsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectTeammatesColumns holds the columns for the "project_teammates" table.
	ProjectTeammatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Size: 255},
		{Name: "is_owner", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "project_id", Type: field.TypeString, Nullable: true},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
	}
	// ProjectTeammatesTable holds the schema information for the "project_teammates" table.
	ProjectTeammatesTable = &schema.Table{
		Name:       "project_teammates",
		Columns:    ProjectTeammatesColumns,
		PrimaryKey: []*schema.Column{ProjectTeammatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_teammates_projects_project_teammates",
				Columns:    []*schema.Column{ProjectTeammatesColumns[5]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_teammates_teammates_project_teammates",
				Columns:    []*schema.Column{ProjectTeammatesColumns[6]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TaskColumnsColumns holds the columns for the "task_columns" table.
	TaskColumnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"TASK_NAME", "ASSIGNEE", "DUE_DATE", "PROJECT", "PROJECTS", "PRIORITY", "TAGS", "CUSTOM"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// TaskColumnsTable holds the schema information for the "task_columns" table.
	TaskColumnsTable = &schema.Table{
		Name:       "task_columns",
		Columns:    TaskColumnsColumns,
		PrimaryKey: []*schema.Column{TaskColumnsColumns[0]},
	}
	// TaskSectionsColumns holds the columns for the "task_sections" table.
	TaskSectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// TaskSectionsTable holds the schema information for the "task_sections" table.
	TaskSectionsTable = &schema.Table{
		Name:       "task_sections",
		Columns:    TaskSectionsColumns,
		PrimaryKey: []*schema.Column{TaskSectionsColumns[0]},
	}
	// TeammatesColumns holds the columns for the "teammates" table.
	TeammatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "image", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// TeammatesTable holds the schema information for the "teammates" table.
	TeammatesTable = &schema.Table{
		Name:       "teammates",
		Columns:    TeammatesColumns,
		PrimaryKey: []*schema.Column{TeammatesColumns[0]},
	}
	// TeammateTaskColumnsColumns holds the columns for the "teammate_task_columns" table.
	TeammateTaskColumnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "width", Type: field.TypeString, Size: 255},
		{Name: "disabled", Type: field.TypeBool},
		{Name: "customizable", Type: field.TypeBool},
		{Name: "order", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "task_column_id", Type: field.TypeString, Nullable: true},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
	}
	// TeammateTaskColumnsTable holds the schema information for the "teammate_task_columns" table.
	TeammateTaskColumnsTable = &schema.Table{
		Name:       "teammate_task_columns",
		Columns:    TeammateTaskColumnsColumns,
		PrimaryKey: []*schema.Column{TeammateTaskColumnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teammate_task_columns_task_columns_teammate_task_columns",
				Columns:    []*schema.Column{TeammateTaskColumnsColumns[7]},
				RefColumns: []*schema.Column{TaskColumnsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "teammate_task_columns_teammates_teammate_task_columns",
				Columns:    []*schema.Column{TeammateTaskColumnsColumns[8]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestTodosColumns holds the columns for the "test_todos" table.
	TestTodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "test_user_id", Type: field.TypeString, Nullable: true},
	}
	// TestTodosTable holds the schema information for the "test_todos" table.
	TestTodosTable = &schema.Table{
		Name:       "test_todos",
		Columns:    TestTodosColumns,
		PrimaryKey: []*schema.Column{TestTodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_todos_test_users_test_todos",
				Columns:    []*schema.Column{TestTodosColumns[6]},
				RefColumns: []*schema.Column{TestUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestUsersColumns holds the columns for the "test_users" table.
	TestUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "age", Type: field.TypeInt},
		{Name: "profile", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// TestUsersTable holds the schema information for the "test_users" table.
	TestUsersTable = &schema.Table{
		Name:       "test_users",
		Columns:    TestUsersColumns,
		PrimaryKey: []*schema.Column{TestUsersColumns[0]},
	}
	// WorkspacesColumns holds the columns for the "workspaces" table.
	WorkspacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
	}
	// WorkspacesTable holds the schema information for the "workspaces" table.
	WorkspacesTable = &schema.Table{
		Name:       "workspaces",
		Columns:    WorkspacesColumns,
		PrimaryKey: []*schema.Column{WorkspacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workspaces_teammates_workspaces",
				Columns:    []*schema.Column{WorkspacesColumns[5]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkspaceTeammatesColumns holds the columns for the "workspace_teammates" table.
	WorkspaceTeammatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Size: 255},
		{Name: "is_owner", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "teammate_id", Type: field.TypeString, Nullable: true},
		{Name: "workspace_id", Type: field.TypeString, Nullable: true},
	}
	// WorkspaceTeammatesTable holds the schema information for the "workspace_teammates" table.
	WorkspaceTeammatesTable = &schema.Table{
		Name:       "workspace_teammates",
		Columns:    WorkspaceTeammatesColumns,
		PrimaryKey: []*schema.Column{WorkspaceTeammatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workspace_teammates_teammates_workspace_teammates",
				Columns:    []*schema.Column{WorkspaceTeammatesColumns[5]},
				RefColumns: []*schema.Column{TeammatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workspace_teammates_workspaces_workspace_teammates",
				Columns:    []*schema.Column{WorkspaceTeammatesColumns[6]},
				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ColorsTable,
		FavoriteProjectsTable,
		FavoriteWorkspacesTable,
		IconsTable,
		MyTasksTabStatusTable,
		ProjectsTable,
		ProjectBaseColorsTable,
		ProjectIconsTable,
		ProjectLightColorsTable,
		ProjectTaskColumnsTable,
		ProjectTeammatesTable,
		TaskColumnsTable,
		TaskSectionsTable,
		TeammatesTable,
		TeammateTaskColumnsTable,
		TestTodosTable,
		TestUsersTable,
		WorkspacesTable,
		WorkspaceTeammatesTable,
	}
)

func init() {
	FavoriteProjectsTable.ForeignKeys[0].RefTable = ProjectsTable
	FavoriteProjectsTable.ForeignKeys[1].RefTable = TeammatesTable
	FavoriteWorkspacesTable.ForeignKeys[0].RefTable = TeammatesTable
	FavoriteWorkspacesTable.ForeignKeys[1].RefTable = WorkspacesTable
	MyTasksTabStatusTable.ForeignKeys[0].RefTable = TeammatesTable
	MyTasksTabStatusTable.ForeignKeys[1].RefTable = WorkspacesTable
	ProjectsTable.ForeignKeys[0].RefTable = ProjectBaseColorsTable
	ProjectsTable.ForeignKeys[1].RefTable = ProjectIconsTable
	ProjectsTable.ForeignKeys[2].RefTable = ProjectLightColorsTable
	ProjectsTable.ForeignKeys[3].RefTable = TeammatesTable
	ProjectsTable.ForeignKeys[4].RefTable = WorkspacesTable
	ProjectBaseColorsTable.ForeignKeys[0].RefTable = ColorsTable
	ProjectIconsTable.ForeignKeys[0].RefTable = IconsTable
	ProjectLightColorsTable.ForeignKeys[0].RefTable = ColorsTable
	ProjectTaskColumnsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTaskColumnsTable.ForeignKeys[1].RefTable = TaskColumnsTable
	ProjectTeammatesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTeammatesTable.ForeignKeys[1].RefTable = TeammatesTable
	TeammateTaskColumnsTable.ForeignKeys[0].RefTable = TaskColumnsTable
	TeammateTaskColumnsTable.ForeignKeys[1].RefTable = TeammatesTable
	TestTodosTable.ForeignKeys[0].RefTable = TestUsersTable
	WorkspacesTable.ForeignKeys[0].RefTable = TeammatesTable
	WorkspaceTeammatesTable.ForeignKeys[0].RefTable = TeammatesTable
	WorkspaceTeammatesTable.ForeignKeys[1].RefTable = WorkspacesTable
}
