package main

import (
	"errors"
	"log"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/adapter/resolver"
	"project-management-demo-backend/pkg/infrastructure/datastore"
	"project-management-demo-backend/pkg/infrastructure/router"
	"project-management-demo-backend/pkg/registry"

	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	config.ReadConfig()

	client := newDBClient()
	defer client.Close()

	c := newController(client)
	srv := newGraphQLServer(client, c)

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

func newGraphQLServer(client *ent.Client, controller controller.Controller) *handler.Server {
	return handler.NewDefaultServer(resolver.NewSchema(client, controller))
}

func newController(client *ent.Client) controller.Controller {
	r := registry.NewRegistry(client)
	return r.NewController()
}
