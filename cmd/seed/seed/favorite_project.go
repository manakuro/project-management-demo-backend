package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// FavoriteProject generates project_teammates data
func FavoriteProject(ctx context.Context, client *ent.Client) {
	_, err := client.FavoriteProject.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("FavoriteProject failed to delete data: %v", err)
	}

	ts := []ent.CreateFavoriteProjectInput{
		{
			ProjectID:  seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			TeammateID: seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
		{
			ProjectID:  seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			TeammateID: seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
	}
	bulk := make([]*ent.FavoriteProjectCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.FavoriteProject.Create().SetInput(t)
	}
	if _, err = client.FavoriteProject.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("FavoriteProject failed to feed data: %v", err)
	}
}
