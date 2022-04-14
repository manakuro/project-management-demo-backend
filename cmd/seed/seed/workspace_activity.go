package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var workspaceActivitySeed = make([]*ent.WorkspaceActivity, 2)

// WorkspaceActivity generates task activity data.
func WorkspaceActivity(ctx context.Context, client *ent.Client) {
	_, err := client.WorkspaceActivity.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("WorkspaceActivity failed to delete data: %v", err)
	}
	activityTypeID := seedutil.GetActivityType(ctx, client, "Workspace").ID
	teammateID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID
	workspaceID := seedutil.GetWorkspace(ctx, client).ID
	projectID := seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID

	inputs := []ent.CreateWorkspaceActivityInput{
		{
			TeammateID:     teammateID,
			ActivityTypeID: activityTypeID,
			ProjectID:      projectID,
			WorkspaceID:    workspaceID,
		},
		{
			TeammateID:     teammateID,
			ActivityTypeID: activityTypeID,
			ProjectID:      projectID,
			WorkspaceID:    workspaceID,
		},
	}
	bulk := make([]*ent.WorkspaceActivityCreate, len(inputs))
	for i, t := range inputs {
		bulk[i] = client.WorkspaceActivity.Create().SetInput(t)
	}
	if _, err = client.WorkspaceActivity.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("WorkspaceActivity failed to feed data: %v", err)
	}

	res, err := client.WorkspaceActivity.Query().All(ctx)
	if err != nil {
		log.Fatalf("WorkspaceActivity failed to feed data: %v", err)
	}

	for i, r := range res {
		workspaceActivitySeed[i] = r
	}
}
