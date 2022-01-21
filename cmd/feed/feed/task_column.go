package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskcolumn"
)

var taskColumnFeed = struct {
	taskName ent.CreateTaskColumnInput
	assignee ent.CreateTaskColumnInput
	dueDate  ent.CreateTaskColumnInput
	project  ent.CreateTaskColumnInput
	tags     ent.CreateTaskColumnInput
	custom   ent.CreateTaskColumnInput
}{
	taskName: ent.CreateTaskColumnInput{Name: "Task Name", Type: taskcolumn.TypeTaskName},
	assignee: ent.CreateTaskColumnInput{Name: "Assignee", Type: taskcolumn.TypeAssignee},
	dueDate:  ent.CreateTaskColumnInput{Name: "Due Date", Type: taskcolumn.TypeDueDate},
	project:  ent.CreateTaskColumnInput{Name: "Project", Type: taskcolumn.TypeProject},
	tags:     ent.CreateTaskColumnInput{Name: "Tags", Type: taskcolumn.TypeTags},
	custom:   ent.CreateTaskColumnInput{Name: "Custom", Type: taskcolumn.TypeCustom},
}

// TaskColumn generates color data
func TaskColumn(ctx context.Context, client *ent.Client) {
	_, err := client.TaskColumn.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskColumn failed to delete data: %v", err)
	}

	ts := []ent.CreateTaskColumnInput{
		taskColumnFeed.taskName,
		taskColumnFeed.assignee,
		taskColumnFeed.dueDate,
		taskColumnFeed.project,
		taskColumnFeed.tags,
		taskColumnFeed.custom,
	}
	bulk := make([]*ent.TaskColumnCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskColumn.Create().SetInput(t)
	}
	if _, err = client.TaskColumn.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskColumn failed to feed data: %v", err)
	}
}
