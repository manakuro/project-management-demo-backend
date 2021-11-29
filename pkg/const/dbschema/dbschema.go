package dbschema

import (
	"fmt"
	"log"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"reflect"
)

type field struct {
	Prefix string
	Table  string
}

// DBSchema defines database schema
type DBSchema struct {
	TestUser field
	TestTodo field
}

// New maps unique string to tables names
// This is intended to be used as global identification for node interface query
func New() DBSchema {
	return DBSchema{
		TestUser: field{
			Prefix: "0AA",
			Table:  testuser.Table,
		},
		TestTodo: field{
			Prefix: "OAB",
			Table:  testtodo.Table,
		},
	}
}

var s = New()
var maps = structToMap(&s)

// FindTableByID returns table name by passed id
func (d *DBSchema) FindTableByID(id string) (string, error) {
	v, ok := maps[id]
	if !ok {
		return "", fmt.Errorf("could not map '%s' to a table name", id)
	}
	return v, nil
}

func structToMap(data *DBSchema) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]string, size)

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		m, ok := value.(field)
		if !ok {
			log.Fatalf("Cannot convert struct to map")
		}
		result[m.Prefix] = m.Table
	}

	return result
}
