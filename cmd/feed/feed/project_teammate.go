package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// ProjectTeammate generates project_teammates data
func ProjectTeammate(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTeammate.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("project_teammate: failed to delete color: %v", err)
	}

	ts := []ent.CreateProjectTeammateInput{
		{
			ProjectID:  feedutil.GetProjectByName(ctx, client, "App Development").ID,
			TeammateID: feedutil.GetTeammateByEmail(ctx, client, "manato.kuroda@example.com").ID,
			Role:       "",
			IsOwner:    true,
		},
		{
			ProjectID:  feedutil.GetProjectByName(ctx, client, "App Development").ID,
			TeammateID: feedutil.GetTeammateByEmail(ctx, client, "dan.abrahmov@example.com").ID,
			Role:       "",
			IsOwner:    false,
		},
		{
			ProjectID:  feedutil.GetProjectByName(ctx, client, "App Development").ID,
			TeammateID: feedutil.GetTeammateByEmail(ctx, client, "kent.dodds@example.com").ID,
			Role:       "",
			IsOwner:    false,
		},
	}
	bulk := make([]*ent.ProjectTeammateCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTeammate.Create().SetInput(t)
	}
	if _, err = client.ProjectTeammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("project_teammate: failed to feed color: %v", err)
	}
}
