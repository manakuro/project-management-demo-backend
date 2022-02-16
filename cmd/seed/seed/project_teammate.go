package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// ProjectTeammate generates project_teammates data
func ProjectTeammate(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTeammate.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTeammate failed to delete data: %v", err)
	}

	ts := []ent.CreateProjectTeammateInput{
		{
			ProjectID:  seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			TeammateID: seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
			Role:       "",
			IsOwner:    true,
		},
		{
			ProjectID:  seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			TeammateID: seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID,
			Role:       "",
			IsOwner:    false,
		},
		{
			ProjectID:  seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			TeammateID: seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID,
			Role:       "",
			IsOwner:    false,
		},
	}
	bulk := make([]*ent.ProjectTeammateCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTeammate.Create().SetInput(t)
	}
	if _, err = client.ProjectTeammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectTeammate failed to feed data: %v", err)
	}
}
