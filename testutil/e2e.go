package testutil

import (
	"net/http/httptest"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/infrastructure/graphql"
	"project-management-demo-backend/pkg/infrastructure/router"
	"project-management-demo-backend/pkg/registry"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// SetupE2EOption is an option of SetupE2E
type SetupE2EOption struct {
	Teardown func(t *testing.T, client *ent.Client)
}

// SetupE2E set up database and server for E2E test
func SetupE2E(t *testing.T, option SetupE2EOption) (expect *httpexpect.Expect, teardown func()) {
	t.Helper()
	ReadConfig()

	client := NewDBClient(t)
	ctrl := newController(client)
	gqlsrv := graphql.NewServer(client, ctrl)
	e := router.New(gqlsrv)

	srv := httptest.NewServer(e)

	return httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  srv.URL,
			Reporter: httpexpect.NewAssertReporter(t),
			Printers: []httpexpect.Printer{
				httpexpect.NewDebugPrinter(t, true),
			},
		}), func() {
			option.Teardown(t, client)
			defer client.Close()
			defer srv.Close()
		}
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
