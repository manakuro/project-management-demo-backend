// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/schema/ulid"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ActivityType is the model entity for the ActivityType schema.
type ActivityType struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TypeCode holds the value of the "type_code" field.
	TypeCode activitytype.TypeCode `json:"type_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActivityTypeQuery when eager-loading is set.
	Edges ActivityTypeEdges `json:"edges"`
}

// ActivityTypeEdges holds the relations/edges for other nodes in the graph.
type ActivityTypeEdges struct {
	// TaskActivities holds the value of the taskActivities edge.
	TaskActivities []*TaskActivity `json:"taskActivities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TaskActivitiesOrErr returns the TaskActivities value or an error if the edge
// was not loaded in eager-loading.
func (e ActivityTypeEdges) TaskActivitiesOrErr() ([]*TaskActivity, error) {
	if e.loadedTypes[0] {
		return e.TaskActivities, nil
	}
	return nil, &NotLoadedError{edge: "taskActivities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ActivityType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case activitytype.FieldName, activitytype.FieldTypeCode:
			values[i] = new(sql.NullString)
		case activitytype.FieldCreatedAt, activitytype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case activitytype.FieldID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ActivityType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ActivityType fields.
func (at *ActivityType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activitytype.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case activitytype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				at.Name = value.String
			}
		case activitytype.FieldTypeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type_code", values[i])
			} else if value.Valid {
				at.TypeCode = activitytype.TypeCode(value.String)
			}
		case activitytype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = value.Time
			}
		case activitytype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				at.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTaskActivities queries the "taskActivities" edge of the ActivityType entity.
func (at *ActivityType) QueryTaskActivities() *TaskActivityQuery {
	return (&ActivityTypeClient{config: at.config}).QueryTaskActivities(at)
}

// Update returns a builder for updating this ActivityType.
// Note that you need to call ActivityType.Unwrap() before calling this method if this ActivityType
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *ActivityType) Update() *ActivityTypeUpdateOne {
	return (&ActivityTypeClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the ActivityType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *ActivityType) Unwrap() *ActivityType {
	tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: ActivityType is not a transactional entity")
	}
	at.config.driver = tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *ActivityType) String() string {
	var builder strings.Builder
	builder.WriteString("ActivityType(")
	builder.WriteString(fmt.Sprintf("id=%v", at.ID))
	builder.WriteString(", name=")
	builder.WriteString(at.Name)
	builder.WriteString(", type_code=")
	builder.WriteString(fmt.Sprintf("%v", at.TypeCode))
	builder.WriteString(", created_at=")
	builder.WriteString(at.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(at.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ActivityTypes is a parsable slice of ActivityType.
type ActivityTypes []*ActivityType

func (at ActivityTypes) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}
