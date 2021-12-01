package globalid

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

// GlobalIDs maps unique string to tables names.
type GlobalIDs struct {
	TestUser field
	TestTodo field
}

// New generates a map object that is intended to be used as global identification for node interface query.
// Prefix should maintain constrained to 3 characters for encoding the entity type.
func New() GlobalIDs {
	return GlobalIDs{
		TestUser: field{
			Prefix: "0AA",
			Table:  testuser.Table,
		},
		TestTodo: field{
			Prefix: "0AB",
			Table:  testtodo.Table,
		},
	}
}

var globalIDS = New()
var maps = structToMap(&globalIDS)

// FindTableByID returns table name by passed id.
func (GlobalIDs) FindTableByID(id string) (string, error) {
	v, ok := maps[id]
	if !ok {
		return "", fmt.Errorf("could not map '%s' to a table name", id)
	}
	return v, nil
}

func structToMap(data *GlobalIDs) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]string, size)

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		f, ok := value.(field)
		if !ok {
			log.Fatalf("Cannot convert struct to map")
		}
		result[f.Prefix] = f.Table
	}

	return result
}