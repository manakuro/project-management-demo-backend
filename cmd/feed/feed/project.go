package feed

import (
	"context"
	"encoding/json"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/editor"
	"time"
)

type projectFeedField struct {
	name string
}

var projectFeed = struct {
	appDevelopment  projectFeedField
	marketing       projectFeedField
	customerSuccess projectFeedField
}{
	appDevelopment:  projectFeedField{name: "App Development"},
	marketing:       projectFeedField{name: "Marketing"},
	customerSuccess: projectFeedField{name: "Customer Success"},
}

// Project generates projects data
func Project(ctx context.Context, client *ent.Client) {
	_, err := client.Project.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("project: failed to delete color: %v", err)
	}

	createdBy := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)
	ws := feedutil.GetWorkspace(ctx, client)
	desc := getDescription()
	ts := []ent.CreateProjectInput{
		{
			Name:             projectFeed.appDevelopment.name,
			WorkspaceID:      ws.ID,
			ColorID:          feedutil.GetColor(ctx, client, colorFeed.pink400.Color).ID,
			IconID:           feedutil.GetIcon(ctx, client, iconFeed.sun.Icon).ID,
			Description:      desc,
			DescriptionTitle: "How we'll collaborate",
			DueDate:          getDueDate(3),
			CreatedBy:        createdBy.ID,
		},
		{
			Name:             projectFeed.marketing.name,
			WorkspaceID:      ws.ID,
			ColorID:          feedutil.GetColor(ctx, client, colorFeed.teal400.Color).ID,
			IconID:           feedutil.GetIcon(ctx, client, iconFeed.moon.Icon).ID,
			Description:      desc,
			DescriptionTitle: "How we'll collaborate",
			DueDate:          getDueDate(10),
			CreatedBy:        createdBy.ID,
		},
		{
			Name:             projectFeed.customerSuccess.name,
			WorkspaceID:      ws.ID,
			ColorID:          feedutil.GetColor(ctx, client, colorFeed.orange400.Color).ID,
			IconID:           feedutil.GetIcon(ctx, client, iconFeed.moon.Icon).ID,
			Description:      desc,
			DescriptionTitle: "How we'll collaborate",
			DueDate:          getDueDate(10),
			CreatedBy:        createdBy.ID,
		},
	}
	bulk := make([]*ent.ProjectCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Project.Create().SetInput(t)
	}
	if _, err = client.Project.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("project: failed to feed color: %v", err)
	}
}

func getDescription() editor.Description {
	b := []byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Welcome to the Marketing team! We’ll be using this project to track our progress on our Q1 product launch. Final ad designs are in the “Key Resources” section below. Use this form to submit new ideas!"}]},{"type":"paragraph"},{"type":"paragraph","content":[{"type":"text","text":"Project Owner: "},{"type":"mention","attrs":{"mentionId":"1","mentionType":"1"}}]},{"type":"paragraph","content":[{"type":"text","text":"Tech Lead:"}]},{"type":"paragraph","content":[{"type":"mention","attrs":{"mentionId":"2","mentionType":"1"}}]},{"type":"paragraph"},{"type":"paragraph"}]}`)

	var description editor.Description
	if err := json.Unmarshal(b, &description); err != nil {
		log.Fatalf("project: failed to encode json")
	}

	return description
}

func getDueDate(date int) *time.Time {
	t := time.Now()
	t.AddDate(0, 0, date)

	return &t
}
