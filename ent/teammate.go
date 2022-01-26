// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Teammate is the model entity for the Teammate schema.
type Teammate struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeammateQuery when eager-loading is set.
	Edges TeammateEdges `json:"edges"`
}

// TeammateEdges holds the relations/edges for other nodes in the graph.
type TeammateEdges struct {
	// Workspaces holds the value of the workspaces edge.
	Workspaces []*Workspace `json:"workspaces,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// ProjectTeammates holds the value of the project_teammates edge.
	ProjectTeammates []*ProjectTeammate `json:"project_teammates,omitempty"`
	// WorkspaceTeammates holds the value of the workspace_teammates edge.
	WorkspaceTeammates []*WorkspaceTeammate `json:"workspace_teammates,omitempty"`
	// FavoriteProjects holds the value of the favorite_projects edge.
	FavoriteProjects []*FavoriteProject `json:"favorite_projects,omitempty"`
	// FavoriteWorkspaces holds the value of the favorite_workspaces edge.
	FavoriteWorkspaces []*FavoriteWorkspace `json:"favorite_workspaces,omitempty"`
	// TeammateTaskTabStatuses holds the value of the teammate_task_tab_statuses edge.
	TeammateTaskTabStatuses []*TeammateTaskTabStatus `json:"teammate_task_tab_statuses,omitempty"`
	// TeammateTaskColumns holds the value of the teammate_task_columns edge.
	TeammateTaskColumns []*TeammateTaskColumn `json:"teammate_task_columns,omitempty"`
	// TeammateTaskListStatuses holds the value of the teammate_task_list_statuses edge.
	TeammateTaskListStatuses []*TeammateTaskListStatus `json:"teammate_task_list_statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// WorkspacesOrErr returns the Workspaces value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) WorkspacesOrErr() ([]*Workspace, error) {
	if e.loadedTypes[0] {
		return e.Workspaces, nil
	}
	return nil, &NotLoadedError{edge: "workspaces"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// ProjectTeammatesOrErr returns the ProjectTeammates value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) ProjectTeammatesOrErr() ([]*ProjectTeammate, error) {
	if e.loadedTypes[2] {
		return e.ProjectTeammates, nil
	}
	return nil, &NotLoadedError{edge: "project_teammates"}
}

// WorkspaceTeammatesOrErr returns the WorkspaceTeammates value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) WorkspaceTeammatesOrErr() ([]*WorkspaceTeammate, error) {
	if e.loadedTypes[3] {
		return e.WorkspaceTeammates, nil
	}
	return nil, &NotLoadedError{edge: "workspace_teammates"}
}

// FavoriteProjectsOrErr returns the FavoriteProjects value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) FavoriteProjectsOrErr() ([]*FavoriteProject, error) {
	if e.loadedTypes[4] {
		return e.FavoriteProjects, nil
	}
	return nil, &NotLoadedError{edge: "favorite_projects"}
}

// FavoriteWorkspacesOrErr returns the FavoriteWorkspaces value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) FavoriteWorkspacesOrErr() ([]*FavoriteWorkspace, error) {
	if e.loadedTypes[5] {
		return e.FavoriteWorkspaces, nil
	}
	return nil, &NotLoadedError{edge: "favorite_workspaces"}
}

// TeammateTaskTabStatusesOrErr returns the TeammateTaskTabStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) TeammateTaskTabStatusesOrErr() ([]*TeammateTaskTabStatus, error) {
	if e.loadedTypes[6] {
		return e.TeammateTaskTabStatuses, nil
	}
	return nil, &NotLoadedError{edge: "teammate_task_tab_statuses"}
}

// TeammateTaskColumnsOrErr returns the TeammateTaskColumns value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) TeammateTaskColumnsOrErr() ([]*TeammateTaskColumn, error) {
	if e.loadedTypes[7] {
		return e.TeammateTaskColumns, nil
	}
	return nil, &NotLoadedError{edge: "teammate_task_columns"}
}

// TeammateTaskListStatusesOrErr returns the TeammateTaskListStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e TeammateEdges) TeammateTaskListStatusesOrErr() ([]*TeammateTaskListStatus, error) {
	if e.loadedTypes[8] {
		return e.TeammateTaskListStatuses, nil
	}
	return nil, &NotLoadedError{edge: "teammate_task_list_statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teammate) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case teammate.FieldName, teammate.FieldImage, teammate.FieldEmail:
			values[i] = new(sql.NullString)
		case teammate.FieldCreatedAt, teammate.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case teammate.FieldID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Teammate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teammate fields.
func (t *Teammate) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teammate.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case teammate.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case teammate.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				t.Image = value.String
			}
		case teammate.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				t.Email = value.String
			}
		case teammate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case teammate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryWorkspaces queries the "workspaces" edge of the Teammate entity.
func (t *Teammate) QueryWorkspaces() *WorkspaceQuery {
	return (&TeammateClient{config: t.config}).QueryWorkspaces(t)
}

// QueryProjects queries the "projects" edge of the Teammate entity.
func (t *Teammate) QueryProjects() *ProjectQuery {
	return (&TeammateClient{config: t.config}).QueryProjects(t)
}

// QueryProjectTeammates queries the "project_teammates" edge of the Teammate entity.
func (t *Teammate) QueryProjectTeammates() *ProjectTeammateQuery {
	return (&TeammateClient{config: t.config}).QueryProjectTeammates(t)
}

// QueryWorkspaceTeammates queries the "workspace_teammates" edge of the Teammate entity.
func (t *Teammate) QueryWorkspaceTeammates() *WorkspaceTeammateQuery {
	return (&TeammateClient{config: t.config}).QueryWorkspaceTeammates(t)
}

// QueryFavoriteProjects queries the "favorite_projects" edge of the Teammate entity.
func (t *Teammate) QueryFavoriteProjects() *FavoriteProjectQuery {
	return (&TeammateClient{config: t.config}).QueryFavoriteProjects(t)
}

// QueryFavoriteWorkspaces queries the "favorite_workspaces" edge of the Teammate entity.
func (t *Teammate) QueryFavoriteWorkspaces() *FavoriteWorkspaceQuery {
	return (&TeammateClient{config: t.config}).QueryFavoriteWorkspaces(t)
}

// QueryTeammateTaskTabStatuses queries the "teammate_task_tab_statuses" edge of the Teammate entity.
func (t *Teammate) QueryTeammateTaskTabStatuses() *TeammateTaskTabStatusQuery {
	return (&TeammateClient{config: t.config}).QueryTeammateTaskTabStatuses(t)
}

// QueryTeammateTaskColumns queries the "teammate_task_columns" edge of the Teammate entity.
func (t *Teammate) QueryTeammateTaskColumns() *TeammateTaskColumnQuery {
	return (&TeammateClient{config: t.config}).QueryTeammateTaskColumns(t)
}

// QueryTeammateTaskListStatuses queries the "teammate_task_list_statuses" edge of the Teammate entity.
func (t *Teammate) QueryTeammateTaskListStatuses() *TeammateTaskListStatusQuery {
	return (&TeammateClient{config: t.config}).QueryTeammateTaskListStatuses(t)
}

// Update returns a builder for updating this Teammate.
// Note that you need to call Teammate.Unwrap() before calling this method if this Teammate
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teammate) Update() *TeammateUpdateOne {
	return (&TeammateClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Teammate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teammate) Unwrap() *Teammate {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teammate is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teammate) String() string {
	var builder strings.Builder
	builder.WriteString("Teammate(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", image=")
	builder.WriteString(t.Image)
	builder.WriteString(", email=")
	builder.WriteString(t.Email)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Teammates is a parsable slice of Teammate.
type Teammates []*Teammate

func (t Teammates) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
