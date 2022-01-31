package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TaskFile generates tasks data
func TaskFile(ctx context.Context, client *ent.Client) {
	_, err := client.TaskFile.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskFile failed to delete data: %v", err)
	}
	appDevelopmentProject := feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name)
	task1 := feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name)
	feeds := feedutil.GetTaskFeeds(ctx, client, task1.ID)

	imageFileType := feedutil.GetFileType(ctx, client, fileTypeFeed.image.Name)
	pdfFileType := feedutil.GetFileType(ctx, client, fileTypeFeed.pdf.Name)
	textFileType := feedutil.GetFileType(ctx, client, fileTypeFeed.text.Name)

	ts := []ent.CreateTaskFileInput{
		{
			ProjectID:  appDevelopmentProject.ID,
			TaskID:     task1.ID,
			TaskFeedID: feeds[0].ID,
			FileTypeID: imageFileType.ID,
			Name:       "/images/cat_img.png",
			Src:        "/images/cat_img.png",
		},
		{
			ProjectID:  appDevelopmentProject.ID,
			TaskID:     task1.ID,
			TaskFeedID: feeds[0].ID,
			FileTypeID: pdfFileType.ID,
			Name:       "/files/pdf-test.pdf",
			Src:        "/files/pdf-test.pdf",
		},
		{
			ProjectID:  appDevelopmentProject.ID,
			TaskID:     task1.ID,
			TaskFeedID: feeds[0].ID,
			FileTypeID: pdfFileType.ID,
			Name:       "/files/pdf-test-2.pdf",
			Src:        "/files/pdf-test-2.pdf",
		},
		{
			ProjectID:  appDevelopmentProject.ID,
			TaskID:     task1.ID,
			TaskFeedID: feeds[0].ID,
			FileTypeID: textFileType.ID,
			Name:       "/files/test.js",
			Src:        "/files/test.js",
		},
	}
	bulk := make([]*ent.TaskFileCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskFile.Create().SetInput(t)
	}
	if _, err = client.TaskFile.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskFile failed to feed data: %v", err)
	}
}
