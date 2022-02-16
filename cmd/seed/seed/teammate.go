package seed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

var teammateFeed = struct {
	manato ent.CreateTeammateInput
	dan    ent.CreateTeammateInput
	kent   ent.CreateTeammateInput
}{
	manato: ent.CreateTeammateInput{
		Name:  "Manato Kuroda",
		Image: "/images/cat_img.png",
		Email: "manato.kuroda@example.com",
	},
	dan: ent.CreateTeammateInput{
		Name:  "Dan Abrahmov",
		Image: "https://bit.ly/dan-abramov",
		Email: "dan.abrahmov@example.com",
	},
	kent: ent.CreateTeammateInput{
		Name:  "Kent Dodds",
		Image: "https://bit.ly/kent-c-dodds",
		Email: "kent.dodds@example.com",
	},
}

// Teammate generates teammate data
func Teammate(ctx context.Context, client *ent.Client) {
	_, err := client.Teammate.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("Teammate failed to delete data: %v", err)
	}

	ts := []ent.CreateTeammateInput{
		teammateFeed.manato,
		teammateFeed.dan,
		teammateFeed.kent,
	}
	bulk := make([]*ent.TeammateCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Teammate.Create().SetInput(t)
	}
	if _, err = client.Teammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Teammate failed to feed data: %v", err)
	}
}
