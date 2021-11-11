package main

import (
	"context"
	"errors"
	"log"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/migrate"
	"project-management-demo-backend/infrastructure/datastore"
)

func main() {
	config.ReadConfig()

	client, err := datastore.NewDB()
	if !errors.Is(err, nil) {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
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
