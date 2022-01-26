package globalid

import (
	"fmt"
	"log"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projecticon"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/taskcolumn"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/tasklistsortstatus"
	"project-management-demo-backend/ent/tasksection"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetaskcolumn"
	"project-management-demo-backend/ent/teammatetaskliststatus"
	"project-management-demo-backend/ent/teammatetasktabstatus"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"
	"reflect"
)

type field struct {
	Prefix string
	Table  string
}

// GlobalIDs maps unique string to tables names.
type GlobalIDs struct {
	Color                   field
	FavoriteProject         field
	FavoriteWorkspace       field
	Icon                    field
	Project                 field
	ProjectBaseColor        field
	ProjectIcon             field
	ProjectLightColor       field
	ProjectTeammate         field
	Teammate                field
	TestTodo                field
	TestUser                field
	Workspace               field
	WorkspaceTeammate       field
	TeammateTaskTabStatus   field
	TaskColumn              field
	TeammateTaskColumn      field
	ProjectTaskColumn       field
	TaskSection             field
	TaskListCompletedStatus field
	TaskListSortStatus      field
	TeammateTaskListStatus  field
	ProjectTaskListStatus   field
}

// New generates a map object that is intended to be used as global identification for node interface query.
// Prefix should maintain constrained to 3 characters for encoding the entity type.
func New() GlobalIDs {
	return GlobalIDs{
		TestUser: field{
			Prefix: "0AA",
			Table:  testuser.Table,
		},
		TestTodo: field{
			Prefix: "0AB",
			Table:  testtodo.Table,
		},
		Teammate: field{
			Prefix: "0AC",
			Table:  teammate.Table,
		},
		Workspace: field{
			Prefix: "0AD",
			Table:  workspace.Table,
		},
		Color: field{
			Prefix: "0AE",
			Table:  color.Table,
		},
		Icon: field{
			Prefix: "0AF",
			Table:  icon.Table,
		},
		Project: field{
			Prefix: "0AG",
			Table:  project.Table,
		},
		ProjectTeammate: field{
			Prefix: "0AH",
			Table:  projectteammate.Table,
		},
		ProjectBaseColor: field{
			Prefix: "0AI",
			Table:  projectbasecolor.Table,
		},
		ProjectLightColor: field{
			Prefix: "0AJ",
			Table:  projectlightcolor.Table,
		},
		ProjectIcon: field{
			Prefix: "0AK",
			Table:  projecticon.Table,
		},
		WorkspaceTeammate: field{
			Prefix: "0AL",
			Table:  workspaceteammate.Table,
		},
		FavoriteProject: field{
			Prefix: "0AM",
			Table:  favoriteproject.Table,
		},
		FavoriteWorkspace: field{
			Prefix: "0AN",
			Table:  favoriteworkspace.Table,
		},
		TeammateTaskTabStatus: field{
			Prefix: "0AO",
			Table:  teammatetasktabstatus.Table,
		},
		TaskColumn: field{
			Prefix: "0AP",
			Table:  taskcolumn.Table,
		},
		TeammateTaskColumn: field{
			Prefix: "0AQ",
			Table:  teammatetaskcolumn.Table,
		},
		ProjectTaskColumn: field{
			Prefix: "0AR",
			Table:  projecttaskcolumn.Table,
		},
		TaskSection: field{
			Prefix: "0AS",
			Table:  tasksection.Table,
		},
		TaskListCompletedStatus: field{
			Prefix: "0AT",
			Table:  tasklistcompletedstatus.Table,
		},
		TaskListSortStatus: field{
			Prefix: "0AU",
			Table:  tasklistsortstatus.Table,
		},
		TeammateTaskListStatus: field{
			Prefix: "0AV",
			Table:  teammatetaskliststatus.Table,
		},
		ProjectTaskListStatus: field{
			Prefix: "0AW",
			Table:  projecttaskliststatus.Table,
		},
	}
}

var globalIDS = New()
var maps = structToMap(&globalIDS)

// FindTableByID returns table name by passed id.
func (GlobalIDs) FindTableByID(id string) (string, error) {
	v, ok := maps[id]
	if !ok {
		return "", fmt.Errorf("could not map '%s' to a table name", id)
	}
	return v, nil
}

func structToMap(data *GlobalIDs) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]string, size)

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		f, ok := value.(field)
		if !ok {
			log.Fatalf("Cannot convert struct to map")
		}
		result[f.Prefix] = f.Table
	}

	return result
}
