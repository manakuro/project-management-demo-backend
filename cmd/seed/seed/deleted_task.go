package seed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

// DeletedTask generates activity data.
func DeletedTask(ctx context.Context, client *ent.Client) {
	_, err := client.DeletedTask.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("DeletedTask failed to delete data: %v", err)
	}
}
