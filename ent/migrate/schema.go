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
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeJSON},
		{Name: "description_title", Type: field.TypeString, Size: 255},
		{Name: "due_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "icon_id", Type: field.TypeString, Nullable: true},
		{Name: "project_base_color_id", Type: field.TypeString, Nullable: true},
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
				Symbol:     "projects_icons_projects",
				Columns:    []*schema.Column{ProjectsColumns[7]},
				RefColumns: []*schema.Column{IconsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_project_base_colors_projects",
				Columns:    []*schema.Column{ProjectsColumns[8]},
				RefColumns: []*schema.Column{ProjectBaseColorsColumns[0]},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ColorsTable,
		IconsTable,
		ProjectsTable,
		ProjectBaseColorsTable,
		ProjectLightColorsTable,
		ProjectTeammatesTable,
		TeammatesTable,
		TestTodosTable,
		TestUsersTable,
		WorkspacesTable,
	}
)

func init() {
	ProjectsTable.ForeignKeys[0].RefTable = IconsTable
	ProjectsTable.ForeignKeys[1].RefTable = ProjectBaseColorsTable
	ProjectsTable.ForeignKeys[2].RefTable = ProjectLightColorsTable
	ProjectsTable.ForeignKeys[3].RefTable = TeammatesTable
	ProjectsTable.ForeignKeys[4].RefTable = WorkspacesTable
	ProjectBaseColorsTable.ForeignKeys[0].RefTable = ColorsTable
	ProjectLightColorsTable.ForeignKeys[0].RefTable = ColorsTable
	ProjectTeammatesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTeammatesTable.ForeignKeys[1].RefTable = TeammatesTable
	TestTodosTable.ForeignKeys[0].RefTable = TestUsersTable
	WorkspacesTable.ForeignKeys[0].RefTable = TeammatesTable
}
