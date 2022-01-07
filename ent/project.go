// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
	// ColorID holds the value of the "color_id" field.
	ColorID ulid.ID `json:"color_id,omitempty"`
	// IconID holds the value of the "icon_id" field.
	IconID ulid.ID `json:"icon_id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy ulid.ID `json:"created_by,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description editor.Description `json:"description,omitempty"`
	// DescriptionTitle holds the value of the "description_title" field.
	DescriptionTitle string `json:"description_title,omitempty"`
	// DueDate holds the value of the "due_date" field.
	DueDate time.Time `json:"due_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges ProjectEdges `json:"edges"`
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace `json:"workspace,omitempty"`
	// Color holds the value of the color edge.
	Color *Color `json:"color,omitempty"`
	// Icon holds the value of the icon edge.
	Icon *Icon `json:"icon,omitempty"`
	// Teammate holds the value of the teammate edge.
	Teammate *Teammate `json:"teammate,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) WorkspaceOrErr() (*Workspace, error) {
	if e.loadedTypes[0] {
		if e.Workspace == nil {
			// The edge workspace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspace, nil
	}
	return nil, &NotLoadedError{edge: "workspace"}
}

// ColorOrErr returns the Color value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) ColorOrErr() (*Color, error) {
	if e.loadedTypes[1] {
		if e.Color == nil {
			// The edge color was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: color.Label}
		}
		return e.Color, nil
	}
	return nil, &NotLoadedError{edge: "color"}
}

// IconOrErr returns the Icon value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) IconOrErr() (*Icon, error) {
	if e.loadedTypes[2] {
		if e.Icon == nil {
			// The edge icon was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: icon.Label}
		}
		return e.Icon, nil
	}
	return nil, &NotLoadedError{edge: "icon"}
}

// TeammateOrErr returns the Teammate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) TeammateOrErr() (*Teammate, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldDescription:
			values[i] = new([]byte)
		case project.FieldName, project.FieldDescriptionTitle:
			values[i] = new(sql.NullString)
		case project.FieldDueDate, project.FieldCreatedAt, project.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case project.FieldID, project.FieldWorkspaceID, project.FieldColorID, project.FieldIconID, project.FieldCreatedBy:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Project", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case project.FieldWorkspaceID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workspace_id", values[i])
			} else if value != nil {
				pr.WorkspaceID = *value
			}
		case project.FieldColorID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field color_id", values[i])
			} else if value != nil {
				pr.ColorID = *value
			}
		case project.FieldIconID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field icon_id", values[i])
			} else if value != nil {
				pr.IconID = *value
			}
		case project.FieldCreatedBy:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				pr.CreatedBy = *value
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Description); err != nil {
					return fmt.Errorf("unmarshal field description: %w", err)
				}
			}
		case project.FieldDescriptionTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_title", values[i])
			} else if value.Valid {
				pr.DescriptionTitle = value.String
			}
		case project.FieldDueDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_date", values[i])
			} else if value.Valid {
				pr.DueDate = value.Time
			}
		case project.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case project.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryWorkspace queries the "workspace" edge of the Project entity.
func (pr *Project) QueryWorkspace() *WorkspaceQuery {
	return (&ProjectClient{config: pr.config}).QueryWorkspace(pr)
}

// QueryColor queries the "color" edge of the Project entity.
func (pr *Project) QueryColor() *ColorQuery {
	return (&ProjectClient{config: pr.config}).QueryColor(pr)
}

// QueryIcon queries the "icon" edge of the Project entity.
func (pr *Project) QueryIcon() *IconQuery {
	return (&ProjectClient{config: pr.config}).QueryIcon(pr)
}

// QueryTeammate queries the "teammate" edge of the Project entity.
func (pr *Project) QueryTeammate() *TeammateQuery {
	return (&ProjectClient{config: pr.config}).QueryTeammate(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return (&ProjectClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.WorkspaceID))
	builder.WriteString(", color_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.ColorID))
	builder.WriteString(", icon_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.IconID))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedBy))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", description=")
	builder.WriteString(fmt.Sprintf("%v", pr.Description))
	builder.WriteString(", description_title=")
	builder.WriteString(pr.DescriptionTitle)
	builder.WriteString(", due_date=")
	builder.WriteString(pr.DueDate.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project

func (pr Projects) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
