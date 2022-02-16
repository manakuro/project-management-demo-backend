package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var taskAssignedFeed = struct {
	task1         ent.CreateTaskInput
	task2         ent.CreateTaskInput
	task3         ent.CreateTaskInput
	task4         ent.CreateTaskInput
	task5         ent.CreateTaskInput
	task6         ent.CreateTaskInput
	task7         ent.CreateTaskInput
	task8         ent.CreateTaskInput
	task9         ent.CreateTaskInput
	task10        ent.CreateTaskInput
	task2Subtask1 ent.CreateTaskInput
	task2Subtask2 ent.CreateTaskInput
	task2Subtask3 ent.CreateTaskInput
}{
	task1:         ent.CreateTaskInput{Name: "Implement new card design", DueDate: seedutil.AddDate(10)},
	task2:         ent.CreateTaskInput{Name: "User bug report", DueDate: seedutil.AddDate(5)},
	task2Subtask1: ent.CreateTaskInput{Name: "User getting sent duplicate notifications", DueDate: seedutil.AddDate(2)},
	task2Subtask2: ent.CreateTaskInput{Name: "User can't invite teammate via modal page", DueDate: seedutil.AddDate(2)},
	task2Subtask3: ent.CreateTaskInput{Name: "Broken links on my page", DueDate: seedutil.AddDate(2)},
	task3:         ent.CreateTaskInput{Name: "Design iOS prototype"},
	task4:         ent.CreateTaskInput{Name: "Scope performance improvements"},
	task5:         ent.CreateTaskInput{Name: "Implement mobile menu"},
	task6:         ent.CreateTaskInput{Name: "Support for offline mode"},
	task7:         ent.CreateTaskInput{Name: "Introduce CI"},
	task8:         ent.CreateTaskInput{Name: "Login with Google"},
	task9:         ent.CreateTaskInput{Name: "Implement undo function"},
	task10:        ent.CreateTaskInput{Name: "Export to PDF file"},
}

var taskNoAssignedFeed = struct {
	task1         ent.CreateTaskInput
	task2         ent.CreateTaskInput
	task2Subtask1 ent.CreateTaskInput
	task2Subtask2 ent.CreateTaskInput
	task2Subtask3 ent.CreateTaskInput
	mTask1        ent.CreateTaskInput
	mTask2        ent.CreateTaskInput
	mTask3        ent.CreateTaskInput
	mTask4        ent.CreateTaskInput
	mTask5        ent.CreateTaskInput
	mTask6        ent.CreateTaskInput
	mTask7        ent.CreateTaskInput
	mTask8        ent.CreateTaskInput
	mTask9        ent.CreateTaskInput
	mTask10       ent.CreateTaskInput
	mTask11       ent.CreateTaskInput
	mTask12       ent.CreateTaskInput
	mTask13       ent.CreateTaskInput
	mTask14       ent.CreateTaskInput
}{
	task1:         ent.CreateTaskInput{Name: "Launch updated task list", DueDate: seedutil.AddDate(10)},
	task2:         ent.CreateTaskInput{Name: "Finalize workspace design", DueDate: seedutil.AddDate(5)},
	task2Subtask1: ent.CreateTaskInput{Name: "Prep for review", DueDate: seedutil.AddDate(2)},
	task2Subtask2: ent.CreateTaskInput{Name: "Tweak project card design", DueDate: seedutil.AddDate(2)},
	task2Subtask3: ent.CreateTaskInput{Name: "Usability testing for top bar", DueDate: seedutil.AddDate(2)},

	// Marketing - Planning
	mTask1: ent.CreateTaskInput{Name: "Weekly Analytics report", DueDate: seedutil.AddDate(5)},
	mTask2: ent.CreateTaskInput{Name: "Finalize campaign brief", DueDate: seedutil.AddDate(5)},
	mTask3: ent.CreateTaskInput{Name: "Kickoff new project", DueDate: seedutil.AddDate(2)},
	mTask4: ent.CreateTaskInput{Name: "Brainstorming", DueDate: seedutil.AddDate(3)},

	// Marketing - Upcoming
	mTask5: ent.CreateTaskInput{Name: "Brief for Black Friday campaign"},
	mTask6: ent.CreateTaskInput{Name: "Q4 brand campaign"},
	mTask7: ent.CreateTaskInput{Name: "Past campaign performance reports"},
	mTask8: ent.CreateTaskInput{Name: "YouTube promotion"},

	// Marketing - Content Development
	mTask9:  ent.CreateTaskInput{Name: "Create new assets design"},
	mTask10: ent.CreateTaskInput{Name: "Develop promotion message framework"},
	mTask11: ent.CreateTaskInput{Name: "Draft content for social media promotion"},

	// Marketing - Campaign Promotion
	mTask12: ent.CreateTaskInput{Name: "Create sales outreach plan"},
	mTask13: ent.CreateTaskInput{Name: "Finalize press release"},
	mTask14: ent.CreateTaskInput{Name: "Interview Ticktocker"},
}

// Task generates tasks data
func Task(ctx context.Context, client *ent.Client) {
	_, err := client.Task.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("Task failed to delete data: %v", err)
	}
	teammate := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)
	completed := true

	assignedTask2, err := client.Task.Create().SetInput(ent.CreateTaskInput{
		Name:           taskAssignedFeed.task2.Name,
		DueDate:        taskAssignedFeed.task2.DueDate,
		AssigneeID:     &teammate.ID,
		CreatedBy:      teammate.ID,
		TaskPriorityID: &seedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.high.Name).ID,
		Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
	}).Save(ctx)
	if err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}

	noAssignedTask2, err := client.Task.Create().SetInput(ent.CreateTaskInput{
		Name:           taskNoAssignedFeed.task2.Name,
		DueDate:        taskNoAssignedFeed.task2.DueDate,
		CreatedBy:      teammate.ID,
		TaskPriorityID: &seedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.high.Name).ID,
		Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
	}).Save(ctx)
	if err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}

	taskPriorityHighID := &seedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.high.Name).ID
	taskPriorityMediumID := &seedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID
	taskPriorityLowID := &seedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID

	ts := []ent.CreateTaskInput{
		{
			Name:           taskAssignedFeed.task1.Name,
			DueDate:        taskAssignedFeed.task1.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"text":"Implement a new card design for the Home page by the end of this month.","type":"text","attrs":{"mentionId":"","mentionType":""}}]},{"type":"paragraph","content":null},{"type":"paragraph","content":[{"text":"Please look at the attachment below.","type":"text","attrs":{"mentionId":"","mentionType":""}}]},{"type":"paragraph","content":null}]}`)),
		},

		{
			Name:           taskAssignedFeed.task2Subtask1.Name,
			DueDate:        taskAssignedFeed.task2Subtask1.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task2Subtask2.Name,
			DueDate:        taskAssignedFeed.task2Subtask2.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task2Subtask3.Name,
			DueDate:        taskAssignedFeed.task2Subtask3.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},

		{
			Name:           taskAssignedFeed.task3.Name,
			DueDate:        taskAssignedFeed.task3.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task4.Name,
			DueDate:        taskAssignedFeed.task4.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task5.Name,
			DueDate:        taskAssignedFeed.task5.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task6.Name,
			DueDate:        taskAssignedFeed.task6.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task7.Name,
			DueDate:        taskAssignedFeed.task7.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task8.Name,
			DueDate:        taskAssignedFeed.task8.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    seedutil.AddDate(-2),
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task9.Name,
			DueDate:        taskAssignedFeed.task9.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    seedutil.AddDate(-10),
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskAssignedFeed.task10.Name,
			DueDate:        taskAssignedFeed.task10.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    seedutil.AddDate(-7),
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityLowID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},

		// No assigned feed
		{
			Name:           taskNoAssignedFeed.task1.Name,
			DueDate:        taskAssignedFeed.task2Subtask1.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskNoAssignedFeed.task2Subtask1.Name,
			DueDate:        taskAssignedFeed.task2Subtask1.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskNoAssignedFeed.task2Subtask2.Name,
			DueDate:        taskAssignedFeed.task2Subtask2.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskNoAssignedFeed.task2Subtask3.Name,
			DueDate:        taskAssignedFeed.task2Subtask3.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},

		// Marketing - Planning
		{
			Name:           taskNoAssignedFeed.mTask1.Name,
			DueDate:        taskNoAssignedFeed.mTask1.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityMediumID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskNoAssignedFeed.mTask2.Name,
			DueDate:        taskNoAssignedFeed.mTask2.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityHighID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:           taskNoAssignedFeed.mTask3.Name,
			DueDate:        taskNoAssignedFeed.mTask3.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityHighID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask4.Name,
			DueDate:     taskNoAssignedFeed.mTask4.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// Marketing - Upcoming
		{
			Name:        taskNoAssignedFeed.mTask5.Name,
			DueDate:     taskNoAssignedFeed.mTask5.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask6.Name,
			DueDate:     taskNoAssignedFeed.mTask6.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask7.Name,
			DueDate:     taskNoAssignedFeed.mTask7.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask8.Name,
			DueDate:     taskNoAssignedFeed.mTask8.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// Marketing - Content Development
		{
			Name:           taskNoAssignedFeed.mTask9.Name,
			DueDate:        taskNoAssignedFeed.mTask9.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityHighID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask10.Name,
			DueDate:     taskNoAssignedFeed.mTask10.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask11.Name,
			DueDate:     taskNoAssignedFeed.mTask11.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// Marketing - Campaign Promotion
		{
			Name:           taskNoAssignedFeed.mTask12.Name,
			DueDate:        taskNoAssignedFeed.mTask12.DueDate,
			CreatedBy:      teammate.ID,
			TaskPriorityID: taskPriorityHighID,
			Description:    seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask13.Name,
			DueDate:     taskNoAssignedFeed.mTask13.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			Name:        taskNoAssignedFeed.mTask14.Name,
			DueDate:     taskNoAssignedFeed.mTask14.DueDate,
			CreatedBy:   teammate.ID,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
	}
	bulk := make([]*ent.TaskCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Task.Create().SetInput(t)
	}
	if _, err = client.Task.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}
}
