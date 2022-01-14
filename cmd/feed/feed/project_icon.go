package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// ProjectIcon generates project icon data
func ProjectIcon(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectIcon.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete project icon: %v", err)
	}

	data := []ent.CreateProjectIconInput{
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.play.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.home.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.moon.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.sun.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.menu.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.codeAlt.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.rocket.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.idCard.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.trashAlt.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.task.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.bell.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.notification.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.barChart.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.bookOpen.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.layerPlus.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.mobile.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.movie.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.shapePolygon.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.spreadsheet.Icon).ID,
		},
		{
			IconID: feedutil.GetIcon(ctx, client, iconFeed.layout.Icon).ID,
		},
	}
	bulk := make([]*ent.ProjectIconCreate, len(data))
	for i, t := range data {
		bulk[i] = client.ProjectIcon.Create().SetInput(t)
	}
	if _, err = client.ProjectIcon.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed project icon: %v", err)
	}
}
