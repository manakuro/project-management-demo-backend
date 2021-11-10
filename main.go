package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/migrate"
	"project-management-demo-backend/graph/resolver"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	loc, err := time.LoadLocation("Local")
	if !errors.Is(err, nil) {
		log.Fatalf("failed loading time location: %v\n", err)
	}

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3308",
		DBName:               "project_management_demo",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  loc,
	}
	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if !errors.Is(err, nil) {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer func(client *ent.Client) {
		err = client.Close()
		if err != nil {
			log.Fatalf("failed closing mysql connection: %v", err)
		}
	}(client)

	if err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); !errors.Is(err, nil) {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
