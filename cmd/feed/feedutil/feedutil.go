package feedutil

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/ent/workspace"
)

// GetTeammateByEmail gets teammate by email.
func GetTeammateByEmail(ctx context.Context, client *ent.Client, email string) *ent.Teammate {
	t, err := client.Teammate.Query().Where(teammate.EmailEQ(email)).Only(ctx)
	if err != nil {
		log.Fatalf("getTeammateByEmail: failed get teammate: %v", err)
	}

	return t
}

// GetProjectByName gets project by name.
func GetProjectByName(ctx context.Context, client *ent.Client, name string) *ent.Project {
	p, err := client.Project.Query().Where(project.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Fatalf("GetProjectByName: failed get project: %v", err)
	}

	return p
}

// GetColor gets color by name.
func GetColor(ctx context.Context, client *ent.Client, val string) *ent.Color {
	c, err := client.Color.Query().Where(color.ColorEQ(val)).Only(ctx)
	if err != nil {
		log.Fatalf("GetColor: failed to get color: %v", err)
	}

	return c
}

// GetIcon gets icon by name.
func GetIcon(ctx context.Context, client *ent.Client, val string) *ent.Icon {
	i, err := client.Icon.Query().Where(icon.IconEQ(val)).Only(ctx)
	if err != nil {
		log.Fatalf("GetIcon: failed to get icon: %v", err)
	}

	return i
}

// GetWorkspace gets workspace.
func GetWorkspace(ctx context.Context, client *ent.Client) *ent.Workspace {
	w, err := client.Workspace.Query().Where(workspace.NameEQ("My Workspace")).Only(ctx)
	if err != nil {
		log.Fatalf("GetWorkspace: failed get workspace: %v", err)
	}

	return w
}

// GetTestUserByName gets test user.
func GetTestUserByName(ctx context.Context, client *ent.Client, name string) *ent.TestUser {
	t, err := client.TestUser.Query().Where(testuser.NameEQ(name)).Only(ctx)
	if err != nil {
		log.Fatalf("GetTestUser: failed get test user: %v", err)
	}

	return t
}

// GetProjectBaseColorByColor gets project base color data.
func GetProjectBaseColorByColor(ctx context.Context, client *ent.Client, val string) *ent.ProjectBaseColor {
	c := GetColor(ctx, client, val)

	p, err := client.ProjectBaseColor.Query().Where(projectbasecolor.ColorID(c.ID)).Only(ctx)
	if err != nil {
		log.Fatalf("GetProjectBaseColorByColor: failed get project base color: %v", err)
	}

	return p
}

// GetProjectLightColorByColor gets project light color data.
func GetProjectLightColorByColor(ctx context.Context, client *ent.Client, val string) *ent.ProjectLightColor {
	c := GetColor(ctx, client, val)

	p, err := client.ProjectLightColor.Query().Where(projectlightcolor.ColorID(c.ID)).Only(ctx)
	if err != nil {
		log.Fatalf("GetProjectLightColorByColor: failed get project light color: %v", err)
	}

	return p
}
