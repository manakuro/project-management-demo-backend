package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TeammateTaskTabStatus generates project_teammates data
func TeammateTaskTabStatus(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTaskTabStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTaskTabStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateTeammateTaskTabStatusInput{
		{
			WorkspaceID: feedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
	}
	bulk := make([]*ent.TeammateTaskTabStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTaskTabStatus.Create().SetInput(t)
	}
	if _, err = client.TeammateTaskTabStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTaskTabStatus failed to feed data: %v", err)
	}
}
