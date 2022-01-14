package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// ProjectLightColor generates project light color data
func ProjectLightColor(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectLightColor.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete project light color: %v", err)
	}

	data := []ent.CreateProjectLightColorInput{
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.gray200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.red200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.orange200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.yellow200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.green200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.teal200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.blue200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.cyan200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.purple200.Color).ID,
		},
		{
			ColorID: feedutil.GetColor(ctx, client, colorFeed.pink200.Color).ID,
		},
	}
	bulk := make([]*ent.ProjectLightColorCreate, len(data))
	for i, t := range data {
		bulk[i] = client.ProjectLightColor.Create().SetInput(t)
	}
	if _, err = client.ProjectLightColor.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed project light color: %v", err)
	}
}
