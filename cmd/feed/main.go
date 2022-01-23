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
	feed.ProjectBaseColor(ctx, client)
	feed.ProjectLightColor(ctx, client)
	feed.ProjectIcon(ctx, client)
	feed.Project(ctx, client)
	feed.ProjectTeammate(ctx, client)
	feed.WorkspaceTeammate(ctx, client)
	feed.TaskColumn(ctx, client)

	feed.FavoriteProject(ctx, client)
	feed.FavoriteWorkspace(ctx, client)
	feed.MyTasksTabStatus(ctx, client)
	feed.TeammateTaskColumn(ctx, client)
	feed.ProjectTaskColumn(ctx, client)

	feed.TestUser(ctx, client)
	feed.TestTodo(ctx, client)

	//client.EnableSQLSafeUpdates()
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}
