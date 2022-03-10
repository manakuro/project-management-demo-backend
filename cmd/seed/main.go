package main

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seed"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/infrastructure/datastore"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	defer client.Close()

	ctx := context.Background()

	client.DisableForeignKeyChecks()
	client.DisableSQLSafeUpdates()

	seed.Color(ctx, client)
	seed.Icon(ctx, client)
	seed.FileType(ctx, client)
	seed.Teammate(ctx, client)
	seed.Workspace(ctx, client)
	seed.ProjectBaseColor(ctx, client)
	seed.ProjectLightColor(ctx, client)
	seed.ProjectIcon(ctx, client)
	seed.Project(ctx, client)
	seed.ProjectTeammate(ctx, client)
	seed.WorkspaceTeammate(ctx, client)
	seed.TaskColumn(ctx, client)
	seed.TaskListCompletedStatus(ctx, client)
	seed.TaskListSortStatus(ctx, client)
	seed.TaskSection(ctx, client)
	seed.TaskPriority(ctx, client)
	seed.Tag(ctx, client)

	seed.Task(ctx, client)
	seed.TeammateTaskSection(ctx, client)
	seed.TeammateTask(ctx, client)
	seed.ProjectTaskSection(ctx, client)
	seed.ProjectTask(ctx, client)

	seed.TaskLike(ctx, client)
	seed.TaskTag(ctx, client)
	seed.TaskCollaborator(ctx, client)
	seed.TaskFeed(ctx, client)
	seed.TaskFeedLike(ctx, client)
	seed.TaskFile(ctx, client)

	seed.FavoriteProject(ctx, client)
	seed.FavoriteWorkspace(ctx, client)
	seed.TeammateTaskTabStatus(ctx, client)
	seed.TeammateTaskColumn(ctx, client)
	seed.TeammateTaskListStatus(ctx, client)
	seed.ProjectTaskColumn(ctx, client)
	seed.ProjectTaskListStatus(ctx, client)

	seed.TestUser(ctx, client)
	seed.TestTodo(ctx, client)

	client.EnableForeignKeyChecks()
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient(datastore.NewClientOptions{})
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}
