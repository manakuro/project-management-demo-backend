package graphql

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/resolver"

	"entgo.io/contrib/entgql"

	"github.com/99designs/gqlgen/graphql/handler"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	// Add extensions to error
	// @see https://github.com/99designs/gqlgen/issues/1354
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
