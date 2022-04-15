package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// WorkspaceActivityTask generates workspace activity task data.
func WorkspaceActivityTask(ctx context.Context, client *ent.Client) {
	_, err := client.WorkspaceActivityTask.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("WorkspaceActivityTask failed to delete data: %v", err)
	}

	workspaceActivity1ID := seedutil.GetWorkspaceActivity(ctx, client, workspaceActivitySeed[0].ID).ID
	task1ID := seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID
	task2ID := seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID

	workspaceActivity2ID := seedutil.GetWorkspaceActivity(ctx, client, workspaceActivitySeed[1].ID).ID
	task3ID := seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task3.Name).ID

	inputs := []ent.CreateWorkspaceActivityTaskInput{
		// workspaceActivity 1
		{
			WorkspaceActivityID: workspaceActivity1ID,
			TaskID:              task1ID,
		},
		{
			WorkspaceActivityID: workspaceActivity1ID,
			TaskID:              task2ID,
		},

		// workspaceActivity 2
		{
			WorkspaceActivityID: workspaceActivity2ID,
			TaskID:              task1ID,
		},
		{
			WorkspaceActivityID: workspaceActivity1ID,
			TaskID:              task3ID,
		},
	}
	bulk := make([]*ent.WorkspaceActivityTaskCreate, len(inputs))
	for i, t := range inputs {
		bulk[i] = client.WorkspaceActivityTask.Create().SetInput(t)
	}
	if _, err = client.WorkspaceActivityTask.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("WorkspaceActivityTask failed to seed data: %v", err)
	}
}
