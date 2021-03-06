// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/projecticon"
	"project-management-demo-backend/ent/schema/ulid"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ProjectIcon is the model entity for the ProjectIcon schema.
type ProjectIcon struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// IconID holds the value of the "icon_id" field.
	IconID ulid.ID `json:"icon_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectIconQuery when eager-loading is set.
	Edges ProjectIconEdges `json:"edges"`
}

// ProjectIconEdges holds the relations/edges for other nodes in the graph.
type ProjectIconEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Icon holds the value of the icon edge.
	Icon *Icon `json:"icon,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectIconEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// IconOrErr returns the Icon value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectIconEdges) IconOrErr() (*Icon, error) {
	if e.loadedTypes[1] {
		if e.Icon == nil {
			// The edge icon was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: icon.Label}
		}
		return e.Icon, nil
	}
	return nil, &NotLoadedError{edge: "icon"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProjectIcon) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case projecticon.FieldCreatedAt, projecticon.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case projecticon.FieldID, projecticon.FieldIconID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProjectIcon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProjectIcon fields.
func (pi *ProjectIcon) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case projecticon.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pi.ID = *value
			}
		case projecticon.FieldIconID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field icon_id", values[i])
			} else if value != nil {
				pi.IconID = *value
			}
		case projecticon.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case projecticon.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProjects queries the "projects" edge of the ProjectIcon entity.
func (pi *ProjectIcon) QueryProjects() *ProjectQuery {
	return (&ProjectIconClient{config: pi.config}).QueryProjects(pi)
}

// QueryIcon queries the "icon" edge of the ProjectIcon entity.
func (pi *ProjectIcon) QueryIcon() *IconQuery {
	return (&ProjectIconClient{config: pi.config}).QueryIcon(pi)
}

// Update returns a builder for updating this ProjectIcon.
// Note that you need to call ProjectIcon.Unwrap() before calling this method if this ProjectIcon
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *ProjectIcon) Update() *ProjectIconUpdateOne {
	return (&ProjectIconClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the ProjectIcon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *ProjectIcon) Unwrap() *ProjectIcon {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProjectIcon is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *ProjectIcon) String() string {
	var builder strings.Builder
	builder.WriteString("ProjectIcon(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", icon_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.IconID))
	builder.WriteString(", created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProjectIcons is a parsable slice of ProjectIcon.
type ProjectIcons []*ProjectIcon

func (pi ProjectIcons) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
