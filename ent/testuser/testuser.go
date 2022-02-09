// Code generated by entc, DO NOT EDIT.

package testuser

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the testuser type in the database.
	Label = "test_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTestTodos holds the string denoting the testtodos edge name in mutations.
	EdgeTestTodos = "testTodos"
	// Table holds the table name of the testuser in the database.
	Table = "test_users"
	// TestTodosTable is the table that holds the testTodos relation/edge.
	TestTodosTable = "test_todos"
	// TestTodosInverseTable is the table name for the TestTodo entity.
	// It exists in this package in order to avoid circular dependency with the "testtodo" package.
	TestTodosInverseTable = "test_todos"
	// TestTodosColumn is the table column denoting the testTodos relation/edge.
	TestTodosColumn = "test_user_id"
)

// Columns holds all SQL columns for testuser fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
	FieldProfile,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
