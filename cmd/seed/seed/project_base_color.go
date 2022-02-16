package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// ProjectBaseColor generates project base color data
func ProjectBaseColor(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectBaseColor.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectBaseColor failed to delete data: %v", err)
	}

	data := []ent.CreateProjectBaseColorInput{
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.gray400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.red400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.orange400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.yellow400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.green400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.teal400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.blue400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.cyan400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.purple400.Color).ID,
		},
		{
			ColorID: seedutil.GetColor(ctx, client, colorFeed.pink400.Color).ID,
		},
	}
	bulk := make([]*ent.ProjectBaseColorCreate, len(data))
	for i, t := range data {
		bulk[i] = client.ProjectBaseColor.Create().SetInput(t)
	}
	if _, err = client.ProjectBaseColor.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectBaseColor failed to feed data: %v", err)
	}
}
