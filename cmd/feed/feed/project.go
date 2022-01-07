package feed

import (
	"context"
	"encoding/json"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"
)

// Project generates projects data
func Project(ctx context.Context, client *ent.Client) {
	_, err := client.Project.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("project: failed to delete color: %v", err)
	}

	createdBy := getTeammate(ctx, client)
	ws := getWorkspace(ctx, client)
	ts := []ent.CreateProjectInput{
		{
			Name:             "App Development",
			WorkspaceID:      ws.ID,
			ColorID:          getColor(ctx, client, "pink.400").ID,
			IconID:           getIcon(ctx, client, "sun").ID,
			Description:      getDescription(),
			DescriptionTitle: "How we'll collaborate",
			DueDate:          getDueDate(3),
			CreatedBy:        createdBy.ID,
		},
		{
			Name:             "Marketing",
			WorkspaceID:      ws.ID,
			ColorID:          getColor(ctx, client, "teal.400").ID,
			IconID:           getIcon(ctx, client, "moon").ID,
			Description:      getDescription(),
			DescriptionTitle: "How we'll collaborate",
			DueDate:          getDueDate(10),
			CreatedBy:        createdBy.ID,
		},
		{
			Name:             "Customer Success",
			WorkspaceID:      ws.ID,
			ColorID:          getColor(ctx, client, "orange.400").ID,
			IconID:           getIcon(ctx, client, "moon").ID,
			Description:      getDescription(),
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

func getColor(ctx context.Context, client *ent.Client, val string) *ent.Color {
	c, err := client.Color.Query().Where(color.ColorEQ(val)).Only(ctx)
	if err != nil {
		log.Fatalf("project: failed to get color: %v", err)
	}

	return c
}

func getIcon(ctx context.Context, client *ent.Client, val string) *ent.Icon {
	i, err := client.Icon.Query().Where(icon.IconEQ(val)).Only(ctx)
	if err != nil {
		log.Fatalf("project: failed to get icon from: %v", err)
	}

	return i
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

func getTeammate(ctx context.Context, client *ent.Client) *ent.Teammate {
	tm, err := client.Teammate.Query().Where(teammate.EmailEQ("manato.kuroda@example.com")).Only(ctx)
	if err != nil {
		log.Fatalf("project: failed get teammate: %v", err)
	}

	return tm
}

func getWorkspace(ctx context.Context, client *ent.Client) *ent.Workspace {
	w, err := client.Workspace.Query().Where(workspace.NameEQ("My Workspace")).Only(ctx)
	if err != nil {
		log.Fatalf("project: failed get teammate: %v", err)
	}

	return w
}
