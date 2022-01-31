// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tag"
	"project-management-demo-backend/ent/workspace"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
	// ColorID holds the value of the "color_id" field.
	ColorID ulid.ID `json:"color_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges TagEdges `json:"edges"`
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace `json:"workspace,omitempty"`
	// Color holds the value of the color edge.
	Color *Color `json:"color,omitempty"`
	// TaskTags holds the value of the task_tags edge.
	TaskTags []*TaskTag `json:"task_tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TagEdges) WorkspaceOrErr() (*Workspace, error) {
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
func (e TagEdges) ColorOrErr() (*Color, error) {
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

// TaskTagsOrErr returns the TaskTags value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) TaskTagsOrErr() ([]*TaskTag, error) {
	if e.loadedTypes[2] {
		return e.TaskTags, nil
	}
	return nil, &NotLoadedError{edge: "task_tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldName:
			values[i] = new(sql.NullString)
		case tag.FieldCreatedAt, tag.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case tag.FieldID, tag.FieldWorkspaceID, tag.FieldColorID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldWorkspaceID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workspace_id", values[i])
			} else if value != nil {
				t.WorkspaceID = *value
			}
		case tag.FieldColorID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field color_id", values[i])
			} else if value != nil {
				t.ColorID = *value
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryWorkspace queries the "workspace" edge of the Tag entity.
func (t *Tag) QueryWorkspace() *WorkspaceQuery {
	return (&TagClient{config: t.config}).QueryWorkspace(t)
}

// QueryColor queries the "color" edge of the Tag entity.
func (t *Tag) QueryColor() *ColorQuery {
	return (&TagClient{config: t.config}).QueryColor(t)
}

// QueryTaskTags queries the "task_tags" edge of the Tag entity.
func (t *Tag) QueryTaskTags() *TaskTagQuery {
	return (&TagClient{config: t.config}).QueryTaskTags(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return (&TagClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", t.WorkspaceID))
	builder.WriteString(", color_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ColorID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

func (t Tags) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}