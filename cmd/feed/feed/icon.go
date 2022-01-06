package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

// Icon generates color data
func Icon(ctx context.Context, client *ent.Client) {
	_, err := client.Icon.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete icon: %v", err)
	}

	data := []ent.CreateIconInput{
		{Name: "play", Icon: "play"},
		{Name: "home", Icon: "home"},
		{Name: "moon", Icon: "moon"},
		{Name: "sun", Icon: "sun"},
		{Name: "menu", Icon: "menu"},
		{Name: "codeAlt", Icon: "codeAlt"},
		{Name: "rocket", Icon: "rocket"},
		{Name: "idCard", Icon: "idCard"},
		{Name: "trashAlt", Icon: "trashAlt"},
		{Name: "task", Icon: "task"},
		{Name: "bell", Icon: "bell"},
		{Name: "notification", Icon: "notification"},
		{Name: "barChart", Icon: "barChart"},
		{Name: "bookOpen", Icon: "bookOpen"},
		{Name: "layerPlus", Icon: "layerPlus"},
		{Name: "mobile", Icon: "mobile"},
		{Name: "movie", Icon: "movie"},
		{Name: "shapePolygon", Icon: "shapePolygon"},
		{Name: "spreadsheet", Icon: "spreadsheet"},
		{Name: "layout", Icon: "layout"},
	}
	bulk := make([]*ent.IconCreate, len(data))
	for i, t := range data {
		bulk[i] = client.Icon.Create().SetInput(t)
	}
	if _, err = client.Icon.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed icon: %v", err)
	}
}
