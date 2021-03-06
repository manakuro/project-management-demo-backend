// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceactivity"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// WorkspaceActivity is the model entity for the WorkspaceActivity schema.
type WorkspaceActivity struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// ActivityTypeID holds the value of the "activity_type_id" field.
	ActivityTypeID ulid.ID `json:"activity_type_id,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID ulid.ID `json:"project_id,omitempty"`
	// TeammateID holds the value of the "teammate_id" field.
	TeammateID ulid.ID `json:"teammate_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkspaceActivityQuery when eager-loading is set.
	Edges WorkspaceActivityEdges `json:"edges"`
}

// WorkspaceActivityEdges holds the relations/edges for other nodes in the graph.
type WorkspaceActivityEdges struct {
	// ActivityType holds the value of the activityType edge.
	ActivityType *ActivityType `json:"activityType,omitempty"`
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace `json:"workspace,omitempty"`
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Teammate holds the value of the teammate edge.
	Teammate *Teammate `json:"teammate,omitempty"`
	// WorkspaceActivityTasks holds the value of the workspaceActivityTasks edge.
	WorkspaceActivityTasks []*WorkspaceActivityTask `json:"workspaceActivityTasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ActivityTypeOrErr returns the ActivityType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceActivityEdges) ActivityTypeOrErr() (*ActivityType, error) {
	if e.loadedTypes[0] {
		if e.ActivityType == nil {
			// The edge activityType was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: activitytype.Label}
		}
		return e.ActivityType, nil
	}
	return nil, &NotLoadedError{edge: "activityType"}
}

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceActivityEdges) WorkspaceOrErr() (*Workspace, error) {
	if e.loadedTypes[1] {
		if e.Workspace == nil {
			// The edge workspace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspace, nil
	}
	return nil, &NotLoadedError{edge: "workspace"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceActivityEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[2] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// TeammateOrErr returns the Teammate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceActivityEdges) TeammateOrErr() (*Teammate, error) {
	if e.loadedTypes[3] {
		if e.Teammate == nil {
			// The edge teammate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: teammate.Label}
		}
		return e.Teammate, nil
	}
	return nil, &NotLoadedError{edge: "teammate"}
}

// WorkspaceActivityTasksOrErr returns the WorkspaceActivityTasks value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceActivityEdges) WorkspaceActivityTasksOrErr() ([]*WorkspaceActivityTask, error) {
	if e.loadedTypes[4] {
		return e.WorkspaceActivityTasks, nil
	}
	return nil, &NotLoadedError{edge: "workspaceActivityTasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkspaceActivity) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workspaceactivity.FieldCreatedAt, workspaceactivity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case workspaceactivity.FieldID, workspaceactivity.FieldActivityTypeID, workspaceactivity.FieldWorkspaceID, workspaceactivity.FieldProjectID, workspaceactivity.FieldTeammateID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkspaceActivity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkspaceActivity fields.
func (wa *WorkspaceActivity) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workspaceactivity.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wa.ID = *value
			}
		case workspaceactivity.FieldActivityTypeID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field activity_type_id", values[i])
			} else if value != nil {
				wa.ActivityTypeID = *value
			}
		case workspaceactivity.FieldWorkspaceID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workspace_id", values[i])
			} else if value != nil {
				wa.WorkspaceID = *value
			}
		case workspaceactivity.FieldProjectID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				wa.ProjectID = *value
			}
		case workspaceactivity.FieldTeammateID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field teammate_id", values[i])
			} else if value != nil {
				wa.TeammateID = *value
			}
		case workspaceactivity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wa.CreatedAt = value.Time
			}
		case workspaceactivity.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wa.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryActivityType queries the "activityType" edge of the WorkspaceActivity entity.
func (wa *WorkspaceActivity) QueryActivityType() *ActivityTypeQuery {
	return (&WorkspaceActivityClient{config: wa.config}).QueryActivityType(wa)
}

// QueryWorkspace queries the "workspace" edge of the WorkspaceActivity entity.
func (wa *WorkspaceActivity) QueryWorkspace() *WorkspaceQuery {
	return (&WorkspaceActivityClient{config: wa.config}).QueryWorkspace(wa)
}

// QueryProject queries the "project" edge of the WorkspaceActivity entity.
func (wa *WorkspaceActivity) QueryProject() *ProjectQuery {
	return (&WorkspaceActivityClient{config: wa.config}).QueryProject(wa)
}

// QueryTeammate queries the "teammate" edge of the WorkspaceActivity entity.
func (wa *WorkspaceActivity) QueryTeammate() *TeammateQuery {
	return (&WorkspaceActivityClient{config: wa.config}).QueryTeammate(wa)
}

// QueryWorkspaceActivityTasks queries the "workspaceActivityTasks" edge of the WorkspaceActivity entity.
func (wa *WorkspaceActivity) QueryWorkspaceActivityTasks() *WorkspaceActivityTaskQuery {
	return (&WorkspaceActivityClient{config: wa.config}).QueryWorkspaceActivityTasks(wa)
}

// Update returns a builder for updating this WorkspaceActivity.
// Note that you need to call WorkspaceActivity.Unwrap() before calling this method if this WorkspaceActivity
// was returned from a transaction, and the transaction was committed or rolled back.
func (wa *WorkspaceActivity) Update() *WorkspaceActivityUpdateOne {
	return (&WorkspaceActivityClient{config: wa.config}).UpdateOne(wa)
}

// Unwrap unwraps the WorkspaceActivity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wa *WorkspaceActivity) Unwrap() *WorkspaceActivity {
	tx, ok := wa.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkspaceActivity is not a transactional entity")
	}
	wa.config.driver = tx.drv
	return wa
}

// String implements the fmt.Stringer.
func (wa *WorkspaceActivity) String() string {
	var builder strings.Builder
	builder.WriteString("WorkspaceActivity(")
	builder.WriteString(fmt.Sprintf("id=%v", wa.ID))
	builder.WriteString(", activity_type_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.ActivityTypeID))
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.WorkspaceID))
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.ProjectID))
	builder.WriteString(", teammate_id=")
	builder.WriteString(fmt.Sprintf("%v", wa.TeammateID))
	builder.WriteString(", created_at=")
	builder.WriteString(wa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(wa.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WorkspaceActivities is a parsable slice of WorkspaceActivity.
type WorkspaceActivities []*WorkspaceActivity

func (wa WorkspaceActivities) config(cfg config) {
	for _i := range wa {
		wa[_i].config = cfg
	}
}
