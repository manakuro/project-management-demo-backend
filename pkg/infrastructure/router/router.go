package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	rm "project-management-demo-backend/pkg/infrastructure/router/middleware"
)

const (
	apiPath          = "api"
	graphQLPath      = "/graphql"
	subscriptionPath = "/subscription"
	playgroundPath   = "/playground"
)

// GraphQLPath is an endpoint of graphql server.
const (
	GraphQLPath = "/" + apiPath + graphQLPath
)

// Options of router
type Options struct {
	Auth bool
}

// New creates route endpoint
func New(srv *handler.Server, options Options) *echo.Echo {
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	g := e.Group(apiPath, rm.Auth(rm.Options{
		Skip: !options.Auth,
	}))
	{
		g.POST(graphQLPath, echo.WrapHandler(srv))

		g.GET(subscriptionPath, echo.WrapHandler(srv))

		e.GET(playgroundPath, func(c echo.Context) error {
			playground.Handler("GraphQL", GraphQLPath).ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	return e
}
