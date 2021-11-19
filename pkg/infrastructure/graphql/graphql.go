package graphql

import (
	"context"
	"errors"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/resolver"
	"project-management-demo-backend/pkg/entity/model"

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
		addError(ctx, err)
		return err
	})

	return srv
}

func addError(ctx context.Context, err error) {
	var extendedError interface{ Extensions() map[string]interface{} }

	for err != nil {
		u, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			break
		}
		e := u.Unwrap()
		err = e

		if e != nil {
			// Skip when its stack strace
			var stackTraceErr model.StackTrace
			if !errors.As(e, &stackTraceErr) {
				continue
			}

			// Skip when it's not the standard error type
			var modelErr model.Error
			if !errors.As(e, &modelErr) {
				continue
			}

			gqlerr := &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: e.Error(),
			}
			if errors.As(e, &extendedError) {
				gqlerr.Extensions = extendedError.Extensions()
			}

			graphql.AddError(ctx, gqlerr)
		}
	}
}
