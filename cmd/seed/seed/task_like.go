package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// TaskLike generates tasks data
func TaskLike(ctx context.Context, client *ent.Client) {
	_, err := client.TaskLike.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskLike failed to delete data: %v", err)
	}
	manatoID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID
	danID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID
	kentID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID
	workspace := seedutil.GetWorkspace(ctx, client)

	ts := []ent.CreateTaskLikeInput{
		{
			TeammateID:  manatoID,
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			WorkspaceID: workspace.ID,
		},
		{
			TeammateID:  danID,
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			WorkspaceID: workspace.ID,
		},
		{
			TeammateID:  kentID,
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			WorkspaceID: workspace.ID,
		},
		{
			TeammateID:  manatoID,
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			WorkspaceID: workspace.ID,
		},
		{
			TeammateID:  danID,
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			WorkspaceID: workspace.ID,
		},
	}
	bulk := make([]*ent.TaskLikeCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskLike.Create().SetInput(t)
	}
	if _, err = client.TaskLike.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskLike failed to feed data: %v", err)
	}
}
