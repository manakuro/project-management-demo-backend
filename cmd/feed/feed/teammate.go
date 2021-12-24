package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

// Teammate generates teammate data
func Teammate(ctx context.Context, client *ent.Client) {
	_, err := client.Teammate.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete teammate: %v", err)
	}

	ts := []ent.CreateTeammateInput{
		{
			Name:  "Manato Kuroda",
			Image: "/images/cat_img.png",
			Email: "manato.kuroda@example.com",
		},
		{
			Name:  "Dan Abrahmov",
			Image: "https://bit.ly/dan-abramov",
			Email: "dan.abrahmov@example.com",
		},
		{
			Name:  "Kent Dodds",
			Image: "https://bit.ly/kent-c-dodds",
			Email: "kent.dodds@example.com",
		},
	}
	bulk := make([]*ent.TeammateCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Teammate.Create().SetInput(t)
	}
	if _, err = client.Teammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed teammate: %v", err)
	}
}
