package main

import (
	"context"
	"errors"
	"log"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/migrate"
	"project-management-demo-backend/graph/resolver"
	"project-management-demo-backend/infrastructure/datastore"
	"project-management-demo-backend/infrastructure/router"

	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	config.ReadConfig()

	client := newDBClient()
	defer client.Close()
	createDBSchema(client)

	srv := newGraphQLServer(client)

	e := router.NewRouter(srv)

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

func newGraphQLServer(client *ent.Client) *handler.Server {
	return handler.NewDefaultServer(resolver.NewSchema(client))
}
