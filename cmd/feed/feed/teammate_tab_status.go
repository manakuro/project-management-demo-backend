package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TeammateTabStatus generates project_teammates data
func TeammateTabStatus(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTabStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTabStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateTeammateTabStatusInput{
		{
			WorkspaceID: feedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
	}
	bulk := make([]*ent.TeammateTabStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTabStatus.Create().SetInput(t)
	}
	if _, err = client.TeammateTabStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTabStatus failed to feed data: %v", err)
	}
}
