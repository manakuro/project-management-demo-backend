package feedutil

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/teammate"
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
