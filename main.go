package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/migrate"
	"project-management-demo-backend/graph/resolver"
	"project-management-demo-backend/infrastructure/datastore"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ReadConfig()

	client := newDBClient()
	defer client.Close()
	createDBSchema(client)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	{
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})
		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewDB()
	if !errors.Is(err, nil) {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); !errors.Is(err, nil) {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
