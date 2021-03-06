// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TestTodo is the model entity for the TestTodo schema.
type TestTodo struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// TestUserID holds the value of the "test_user_id" field.
	TestUserID ulid.ID `json:"test_user_id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy ulid.ID `json:"created_by,omitempty"`
	// ParentTodoID holds the value of the "parent_todo_id" field.
	ParentTodoID ulid.ID `json:"parent_todo_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Status holds the value of the "status" field.
	Status testtodo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// DueDate holds the value of the "due_date" field.
	DueDate *time.Time `json:"due_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TestTodoQuery when eager-loading is set.
	Edges TestTodoEdges `json:"edges"`
}

// TestTodoEdges holds the relations/edges for other nodes in the graph.
type TestTodoEdges struct {
	// TestUser holds the value of the testUser edge.
	TestUser *TestUser `json:"testUser,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *TestTodo `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*TestTodo `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TestUserOrErr returns the TestUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TestTodoEdges) TestUserOrErr() (*TestUser, error) {
	if e.loadedTypes[0] {
		if e.TestUser == nil {
			// The edge testUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: testuser.Label}
		}
		return e.TestUser, nil
	}
	return nil, &NotLoadedError{edge: "testUser"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TestTodoEdges) ParentOrErr() (*TestTodo, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: testtodo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TestTodoEdges) ChildrenOrErr() ([]*TestTodo, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestTodo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case testtodo.FieldPriority:
			values[i] = new(sql.NullInt64)
		case testtodo.FieldName, testtodo.FieldStatus:
			values[i] = new(sql.NullString)
		case testtodo.FieldDueDate, testtodo.FieldCreatedAt, testtodo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case testtodo.FieldID, testtodo.FieldTestUserID, testtodo.FieldCreatedBy, testtodo.FieldParentTodoID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestTodo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestTodo fields.
func (tt *TestTodo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testtodo.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tt.ID = *value
			}
		case testtodo.FieldTestUserID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field test_user_id", values[i])
			} else if value != nil {
				tt.TestUserID = *value
			}
		case testtodo.FieldCreatedBy:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				tt.CreatedBy = *value
			}
		case testtodo.FieldParentTodoID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_todo_id", values[i])
			} else if value != nil {
				tt.ParentTodoID = *value
			}
		case testtodo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tt.Name = value.String
			}
		case testtodo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tt.Status = testtodo.Status(value.String)
			}
		case testtodo.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				tt.Priority = int(value.Int64)
			}
		case testtodo.FieldDueDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_date", values[i])
			} else if value.Valid {
				tt.DueDate = new(time.Time)
				*tt.DueDate = value.Time
			}
		case testtodo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tt.CreatedAt = value.Time
			}
		case testtodo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tt.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTestUser queries the "testUser" edge of the TestTodo entity.
func (tt *TestTodo) QueryTestUser() *TestUserQuery {
	return (&TestTodoClient{config: tt.config}).QueryTestUser(tt)
}

// QueryParent queries the "parent" edge of the TestTodo entity.
func (tt *TestTodo) QueryParent() *TestTodoQuery {
	return (&TestTodoClient{config: tt.config}).QueryParent(tt)
}

// QueryChildren queries the "children" edge of the TestTodo entity.
func (tt *TestTodo) QueryChildren() *TestTodoQuery {
	return (&TestTodoClient{config: tt.config}).QueryChildren(tt)
}

// Update returns a builder for updating this TestTodo.
// Note that you need to call TestTodo.Unwrap() before calling this method if this TestTodo
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TestTodo) Update() *TestTodoUpdateOne {
	return (&TestTodoClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TestTodo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TestTodo) Unwrap() *TestTodo {
	tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestTodo is not a transactional entity")
	}
	tt.config.driver = tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TestTodo) String() string {
	var builder strings.Builder
	builder.WriteString("TestTodo(")
	builder.WriteString(fmt.Sprintf("id=%v", tt.ID))
	builder.WriteString(", test_user_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.TestUserID))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", tt.CreatedBy))
	builder.WriteString(", parent_todo_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.ParentTodoID))
	builder.WriteString(", name=")
	builder.WriteString(tt.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", tt.Status))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", tt.Priority))
	if v := tt.DueDate; v != nil {
		builder.WriteString(", due_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", created_at=")
	builder.WriteString(tt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tt.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TestTodos is a parsable slice of TestTodo.
type TestTodos []*TestTodo

func (tt TestTodos) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}
