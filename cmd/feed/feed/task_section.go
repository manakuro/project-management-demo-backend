package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

type taskSectionFeedField struct {
	name string
}

var taskSectionFeed = struct {
	backlog    taskSectionFeedField
	ready      taskSectionFeedField
	inProgress taskSectionFeedField
	done       taskSectionFeedField
}{
	backlog:    taskSectionFeedField{name: "Backlog"},
	ready:      taskSectionFeedField{name: "Ready"},
	inProgress: taskSectionFeedField{name: "In Progress"},
	done:       taskSectionFeedField{name: "Done"},
}

// TaskSection generates taskSections data
func TaskSection(ctx context.Context, client *ent.Client) {
	_, err := client.TaskSection.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskSection failed to delete data: %v", err)
	}

	ts := []ent.CreateTaskSectionInput{
		{
			Name: taskSectionFeed.backlog.name,
		},
		{
			Name: taskSectionFeed.ready.name,
		},
		{
			Name: taskSectionFeed.inProgress.name,
		},
		{
			Name: taskSectionFeed.done.name,
		},
	}
	bulk := make([]*ent.TaskSectionCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskSection.Create().SetInput(t)
	}
	if _, err = client.TaskSection.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskSection failed to feed data: %v", err)
	}
}
