package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// WorkspaceTeammate generates workspace_teammates data
func WorkspaceTeammate(ctx context.Context, client *ent.Client) {
	_, err := client.WorkspaceTeammate.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("WorkspaceTeammate failed to delete data: %v", err)
	}

	ts := []ent.CreateWorkspaceTeammateInput{
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
			Role:        "",
			IsOwner:     true,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID,
			Role:        "",
			IsOwner:     false,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID,
			Role:        "",
			IsOwner:     false,
		},
	}
	bulk := make([]*ent.WorkspaceTeammateCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.WorkspaceTeammate.Create().SetInput(t)
	}
	if _, err = client.WorkspaceTeammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("WorkspaceTeammate failed to feed data: %v", err)
	}
}
