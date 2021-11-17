// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/testuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TestUser is the model entity for the TestUser schema.
type TestUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case testuser.FieldID, testuser.FieldAge:
			values[i] = new(sql.NullInt64)
		case testuser.FieldName:
			values[i] = new(sql.NullString)
		case testuser.FieldCreatedAt, testuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestUser fields.
func (tu *TestUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tu.ID = int(value.Int64)
		case testuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tu.Name = value.String
			}
		case testuser.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				tu.Age = int(value.Int64)
			}
		case testuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tu.CreatedAt = value.Time
			}
		case testuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tu.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TestUser.
// Note that you need to call TestUser.Unwrap() before calling this method if this TestUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (tu *TestUser) Update() *TestUserUpdateOne {
	return (&TestUserClient{config: tu.config}).UpdateOne(tu)
}

// Unwrap unwraps the TestUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tu *TestUser) Unwrap() *TestUser {
	tx, ok := tu.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestUser is not a transactional entity")
	}
	tu.config.driver = tx.drv
	return tu
}

// String implements the fmt.Stringer.
func (tu *TestUser) String() string {
	var builder strings.Builder
	builder.WriteString("TestUser(")
	builder.WriteString(fmt.Sprintf("id=%v", tu.ID))
	builder.WriteString(", name=")
	builder.WriteString(tu.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", tu.Age))
	builder.WriteString(", created_at=")
	builder.WriteString(tu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tu.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TestUsers is a parsable slice of TestUser.
type TestUsers []*TestUser

func (tu TestUsers) config(cfg config) {
	for _i := range tu {
		tu[_i].config = cfg
	}
}
