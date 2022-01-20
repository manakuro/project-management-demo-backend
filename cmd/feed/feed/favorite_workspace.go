package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// FavoriteWorkspace generates project_teammates data
func FavoriteWorkspace(ctx context.Context, client *ent.Client) {
	_, err := client.FavoriteWorkspace.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("FavoriteWorkspace failed to delete data: %v", err)
	}

	ts := []ent.CreateFavoriteWorkspaceInput{
		{
			WorkspaceID: feedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
	}
	bulk := make([]*ent.FavoriteWorkspaceCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.FavoriteWorkspace.Create().SetInput(t)
	}
	if _, err = client.FavoriteWorkspace.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("FavoriteWorkspace failed to feed data: %v", err)
	}
}
