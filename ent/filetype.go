// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/schema/ulid"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// FileType is the model entity for the FileType schema.
type FileType struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TypeCode holds the value of the "type_code" field.
	TypeCode filetype.TypeCode `json:"type_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case filetype.FieldName, filetype.FieldTypeCode:
			values[i] = new(sql.NullString)
		case filetype.FieldCreatedAt, filetype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case filetype.FieldID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FileType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileType fields.
func (ft *FileType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filetype.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ft.ID = *value
			}
		case filetype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ft.Name = value.String
			}
		case filetype.FieldTypeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type_code", values[i])
			} else if value.Valid {
				ft.TypeCode = filetype.TypeCode(value.String)
			}
		case filetype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ft.CreatedAt = value.Time
			}
		case filetype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ft.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FileType.
// Note that you need to call FileType.Unwrap() before calling this method if this FileType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FileType) Update() *FileTypeUpdateOne {
	return (&FileTypeClient{config: ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the FileType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FileType) Unwrap() *FileType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FileType) String() string {
	var builder strings.Builder
	builder.WriteString("FileType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", name=")
	builder.WriteString(ft.Name)
	builder.WriteString(", type_code=")
	builder.WriteString(fmt.Sprintf("%v", ft.TypeCode))
	builder.WriteString(", created_at=")
	builder.WriteString(ft.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ft.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FileTypes is a parsable slice of FileType.
type FileTypes []*FileType

func (ft FileTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
