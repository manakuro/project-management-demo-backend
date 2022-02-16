package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// TaskCollaborator generates tasks data
func TaskCollaborator(ctx context.Context, client *ent.Client) {
	_, err := client.TaskCollaborator.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskCollaborator failed to delete data: %v", err)
	}
	manatoID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID
	danID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID
	kentID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID

	ts := []ent.CreateTaskCollaboratorInput{
		{
			TeammateID: manatoID,
			TaskID:     seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
		},
		{
			TeammateID: danID,
			TaskID:     seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
		},
		{
			TeammateID: kentID,
			TaskID:     seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
		},
		{
			TeammateID: manatoID,
			TaskID:     seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
		},
		{
			TeammateID: danID,
			TaskID:     seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
		},
	}
	bulk := make([]*ent.TaskCollaboratorCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskCollaborator.Create().SetInput(t)
	}
	if _, err = client.TaskCollaborator.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskCollaborator failed to feed data: %v", err)
	}
}
