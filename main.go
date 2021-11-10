package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/migrate"

	"github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3308",
		DBName:               "project_management_demo",
		AllowNativePasswords: true,
		ParseTime:            true,
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

	e.Logger.Fatal(e.Start(":8080"))
}
