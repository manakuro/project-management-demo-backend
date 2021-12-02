package testutil

import (
	"context"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/enttest"
	"project-management-demo-backend/pkg/infrastructure/datastore"
	"project-management-demo-backend/pkg/util/environment"
	"testing"

	"entgo.io/ent/dialect"
)

// ReadConfig reads config file for test.
func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.Test,
	})
}

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	d := datastore.New()
	return enttest.Open(t, dialect.MySQL, d)
}

// DropTables drops data from database.
func DropTables(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	DropTestUser(t, client)
	DropTestTodo(t, client)
}

// DropTestUser drops data from test_users.
func DropTestUser(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.TestUser.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropTestTodo drops data from test_todos.
func DropTestTodo(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.TestTodo.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
