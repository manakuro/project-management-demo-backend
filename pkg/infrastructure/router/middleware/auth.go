package middleware

import (
	"context"
	"project-management-demo-backend/config"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/infrastructure/router/handler"
	"project-management-demo-backend/pkg/util/auth"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// Options of auth
type Options struct {
	Skip bool
}

// Auth is a middleware of authenticating users
func Auth(opts Options) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if opts.Skip {
				return next(c)
			}

			ctx := c.Request().Context()

			opt := option.WithCredentialsFile(config.C.Firebase.ServiceKey)
			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				return handler.HandleError(c, model.NewInternalServerError(err))
			}

			a, err := app.Auth(ctx)
			if err != nil {
				return handler.HandleError(c, model.NewAuthError(err))
			}

			header := c.Request().Header.Get(echo.HeaderAuthorization)
			idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
			token, err := a.VerifyIDToken(ctx, idToken)
			if err != nil {
				return handler.HandleError(c, model.NewAuthError(err))
			}

			ctx = context.WithValue(ctx, auth.TokenKey, token)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
