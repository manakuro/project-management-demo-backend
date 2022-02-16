package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// TeammateTaskColumn generates project_teammates data
func TeammateTaskColumn(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTaskColumn.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTaskColumn failed to delete data: %v", err)
	}

	workspace := seedutil.GetWorkspace(ctx, client)
	manato := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)

	ts := []ent.CreateTeammateTaskColumnInput{
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.taskName.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "600px",
			Disabled:     false,
			Customizable: false,
			Order:        1,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.dueDate.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        2,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.project.Name).ID,
			TeammateID:   manato.ID,
			WorkspaceID:  workspace.ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        3,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.tags.Name).ID,
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
