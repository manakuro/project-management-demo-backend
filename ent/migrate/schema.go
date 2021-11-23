// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TestTodosColumns holds the columns for the "test_todos" table.
	TestTodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "test_user_id", Type: field.TypeInt, Nullable: true},
	}
	// TestTodosTable holds the schema information for the "test_todos" table.
	TestTodosTable = &schema.Table{
		Name:       "test_todos",
		Columns:    TestTodosColumns,
		PrimaryKey: []*schema.Column{TestTodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_todos_test_users_test_todos",
				Columns:    []*schema.Column{TestTodosColumns[6]},
				RefColumns: []*schema.Column{TestUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestUsersColumns holds the columns for the "test_users" table.
	TestUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "age", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// TestUsersTable holds the schema information for the "test_users" table.
	TestUsersTable = &schema.Table{
		Name:       "test_users",
		Columns:    TestUsersColumns,
		PrimaryKey: []*schema.Column{TestUsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TestTodosTable,
		TestUsersTable,
	}
)

func init() {
	TestTodosTable.ForeignKeys[0].RefTable = TestUsersTable
}
