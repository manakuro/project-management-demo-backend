package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var tagFeed = struct {
	design      ent.CreateTagInput
	development ent.CreateTagInput
	bug         ent.CreateTagInput
}{
	design:      ent.CreateTagInput{Name: "Design"},
	development: ent.CreateTagInput{Name: "Development"},
	bug:         ent.CreateTagInput{Name: "Bug"},
}

// Tag generates project light color data
func Tag(ctx context.Context, client *ent.Client) {
	_, err := client.Tag.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("Tag failed to delete data: %v", err)
	}

	workspace := seedutil.GetWorkspace(ctx, client)
	data := []ent.CreateTagInput{
		{
			WorkspaceID: workspace.ID,
			ColorID:     seedutil.GetColor(ctx, client, colorFeed.gray400.Color).ID,
			Name:        tagFeed.design.Name,
		},
		{
			WorkspaceID: workspace.ID,
			ColorID:     seedutil.GetColor(ctx, client, colorFeed.orange400.Color).ID,
			Name:        tagFeed.development.Name,
		},
		{
			WorkspaceID: workspace.ID,
			ColorID:     seedutil.GetColor(ctx, client, colorFeed.red400.Color).ID,
			Name:        tagFeed.bug.Name,
		},
	}
	bulk := make([]*ent.TagCreate, len(data))
	for i, t := range data {
		bulk[i] = client.Tag.Create().SetInput(t)
	}
	if _, err = client.Tag.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Tag failed to feed data: %v", err)
	}
}
