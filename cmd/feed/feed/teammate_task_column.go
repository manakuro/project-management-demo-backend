package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TeammateTaskColumn generates project_teammates data
func TeammateTaskColumn(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTaskColumn.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTaskColumn failed to delete data: %v", err)
	}

	workspace := feedutil.GetWorkspace(ctx, client)
	manato := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)

	ts := []ent.CreateTeammateTaskColumnInput{
		{
			TaskColumnID: feedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.taskName.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "600px",
			Disabled:     false,
			Customizable: false,
			Order:        1,
		},
		{
			TaskColumnID: feedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.dueDate.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        2,
		},
		{
			TaskColumnID: feedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.project.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        3,
		},
		{
			TaskColumnID: feedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.tags.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        4,
		},
	}
	bulk := make([]*ent.TeammateTaskColumnCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTaskColumn.Create().SetInput(t)
	}
	if _, err = client.TeammateTaskColumn.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTaskColumn failed to feed data: %v", err)
	}
}
