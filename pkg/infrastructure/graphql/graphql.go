package graphql

import (
	"context"
	"errors"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/resolver"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	// Add extensions to error
	// @see https://github.com/99designs/gqlgen/issues/1354
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var extendedError interface{ Extensions() map[string]interface{} }
		if errors.As(err, &extendedError) {
			err.Extensions = extendedError.Extensions()
		}
		return err
	})

	return srv
}
