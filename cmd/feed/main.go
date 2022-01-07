package main

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feed"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/infrastructure/datastore"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	defer client.Close()

	ctx := context.Background()

	//client.DisableSQLSafeUpdates()

	feed.Color(ctx, client)
	feed.Icon(ctx, client)
	feed.Teammate(ctx, client)
	feed.Workspace(ctx, client)
	feed.Project(ctx, client)
	feed.ProjectTeammate(ctx, client)

	//client.EnableSQLSafeUpdates()
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}
