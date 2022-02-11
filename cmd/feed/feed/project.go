package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
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
		log.Fatalf("Project failed to delete data: %v", err)
	}

	createdBy := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)
	ws := feedutil.GetWorkspace(ctx, client)
	ts := []ent.CreateProjectInput{
		{
			Name:                projectFeed.appDevelopment.name,
			WorkspaceID:         ws.ID,
			ProjectBaseColorID:  feedutil.GetProjectBaseColorByColor(ctx, client, colorFeed.pink400.Color).ID,
			ProjectLightColorID: feedutil.GetProjectLightColorByColor(ctx, client, colorFeed.pink200.Color).ID,
			ProjectIconID:       feedutil.GetProjectIconByIcon(ctx, client, iconFeed.sun.Icon).ID,
			Description:         feedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Welcome to the App Development team! We’ll be using this project to track our progress on our Q1 product launch. Final ad designs are in the “Key Resources” section below. Use this form to submit new ideas! 😆","attrs":{"mentionId":"","mentionType":""}}]},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null}]}`)),
			DescriptionTitle:    "How we'll collaborate",
			DueDate:             getDueDate(3),
			CreatedBy:           createdBy.ID,
		},
		{
			Name:                projectFeed.marketing.name,
			WorkspaceID:         ws.ID,
			ProjectBaseColorID:  feedutil.GetProjectBaseColorByColor(ctx, client, colorFeed.teal400.Color).ID,
			ProjectLightColorID: feedutil.GetProjectLightColorByColor(ctx, client, colorFeed.teal200.Color).ID,
			ProjectIconID:       feedutil.GetProjectIconByIcon(ctx, client, iconFeed.moon.Icon).ID,
			Description:         feedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Welcome to the Marketing team! We’ll be using this project to track our progress on our Q1 product launch. Final ad designs are in the “Key Resources” section below. Use this form to submit new ideas! 😆","attrs":{"mentionId":"","mentionType":""}}]},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null}]}`)),
			DescriptionTitle:    "How we'll collaborate",
			DueDate:             getDueDate(10),
			CreatedBy:           createdBy.ID,
		},
		{
			Name:                projectFeed.customerSuccess.name,
			WorkspaceID:         ws.ID,
			ProjectBaseColorID:  feedutil.GetProjectBaseColorByColor(ctx, client, colorFeed.orange400.Color).ID,
			ProjectLightColorID: feedutil.GetProjectLightColorByColor(ctx, client, colorFeed.orange200.Color).ID,
			ProjectIconID:       feedutil.GetProjectIconByIcon(ctx, client, iconFeed.moon.Icon).ID,
			Description:         feedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Welcome to the Customer Success team! We’ll be using this project to track our progress on our Q1 product launch. Final ad designs are in the “Key Resources” section below. Use this form to submit new ideas! 😆","attrs":{"mentionId":"","mentionType":""}}]},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null},{"type":"paragraph","content":null}]}`)),
			DescriptionTitle:    "How we'll collaborate",
			DueDate:             nil,
			CreatedBy:           createdBy.ID,
		},
	}
	bulk := make([]*ent.ProjectCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Project.Create().SetInput(t)
	}
	if _, err = client.Project.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Project failed to feed data: %v", err)
	}
}

func getDueDate(date int) *time.Time {
	t := time.Now()
	t.AddDate(0, 0, date)

	return &t
}
