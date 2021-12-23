package main

import (
	"context"
	"log"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/infrastructure/datastore"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	defer client.Close()

	ctx := context.Background()

	FeedTeammate(ctx, client)
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}

// FeedTeammate creates teammate data
func FeedTeammate(ctx context.Context, client *ent.Client) {
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
	if _, err := client.Teammate.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed teammate")
	}
}
