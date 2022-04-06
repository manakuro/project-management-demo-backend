package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfeedlike"
	"project-management-demo-backend/ent/taskfile"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskFeedRepository struct {
	client *ent.Client
}

// NewTaskFeedRepository generates taskFeed repository
func NewTaskFeedRepository(client *ent.Client) ur.TaskFeed {
	return &taskFeedRepository{client: client}
}

func (r *taskFeedRepository) Get(ctx context.Context, where *model.TaskFeedWhereInput) (*model.TaskFeed, error) {
	q := r.client.TaskFeed.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	res, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) List(ctx context.Context, where *model.TaskFeedWhereInput) ([]*model.TaskFeed, error) {
	q := r.client.TaskFeed.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	res, err := q.All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFeedWhereInput) (*model.TaskFeedConnection, error) {
	q := r.client.TaskFeed.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFeedFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskFeedRepository) Create(ctx context.Context, input model.CreateTaskFeedInput) (*model.TaskFeed, error) {
	res, err := r.client.
		TaskFeed.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) Update(ctx context.Context, input model.UpdateTaskFeedInput) (*model.TaskFeed, error) {
	res, err := r.client.
		TaskFeed.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) Delete(ctx context.Context, input model.DeleteTaskFeedInput) (*model.DeleteTaskFeedInputPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.DeleteTaskFeedInputPayload{
		TaskFeed:      nil,
		TaskFeedLikes: []*model.TaskFeedLike{},
		TaskFiles:     []*model.TaskFile{},
	}

	deleted, err := client.TaskFeed.Query().Where(taskfeed.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	taskFeedLikes, err := client.TaskFeedLike.Query().Where(taskfeedlike.TaskFeedID(input.ID)).All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if len(taskFeedLikes) > 0 {
		taskFeedLikeIds := make([]model.ID, len(taskFeedLikes))
		for i, t := range taskFeedLikes {
			taskFeedLikeIds[i] = t.ID
		}

		_, err = client.TaskFeedLike.Delete().Where(taskfeedlike.IDIn(taskFeedLikeIds...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
	}

	taskFiles, err := client.TaskFile.Query().Where(taskfile.TaskFeedID(input.ID)).All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, model.NewDBError(err)
	}

	if len(taskFiles) > 0 {
		taskFileIds := make([]model.ID, len(taskFiles))
		for i, t := range taskFiles {
			taskFileIds[i] = t.ID
		}

		_, err = client.TaskFile.Delete().Where(taskfile.IDIn(taskFileIds...)).Exec(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
	}

	err = client.TaskFeed.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	if taskFeedLikes != nil {
		payload.TaskFeedLikes = taskFeedLikes
	}
	if taskFiles != nil {
		payload.TaskFiles = taskFiles
	}

	payload.TaskFeed = deleted

	return payload, nil
}

func (r *taskFeedRepository) Undelete(ctx context.Context, input model.UndeleteTaskFeedInput) (*model.UndeleteTaskFeedInputPayload, error) {
	client := WithTransactionalMutation(ctx)

	payload := &model.UndeleteTaskFeedInputPayload{
		TaskFeed:      nil,
		TaskFeedLikes: []*model.TaskFeedLike{},
		TaskFiles:     []*model.TaskFile{},
	}

	createdTaskFeed, err := client.
		TaskFeed.
		Create().
		SetID(input.TaskFeed.ID).
		SetTaskID(input.TaskFeed.TaskID).
		SetTeammateID(input.TaskFeed.TeammateID).
		SetIsFirst(input.TaskFeed.IsFirst).
		SetIsPinned(input.TaskFeed.IsPinned).
		SetDescription(input.TaskFeed.Description).
		SetCreatedAt(input.TaskFeed.CreatedAt).
		SetUpdatedAt(input.TaskFeed.UpdatedAt).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	var createdTaskFeedLikes []*model.TaskFeedLike
	if len(input.TaskFeedLikes) > 0 {
		bulk := make([]*ent.TaskFeedLikeCreate, len(input.TaskFeedLikes))
		for i, t := range input.TaskFeedLikes {
			bulk[i] = client.TaskFeedLike.Create().
				SetID(t.ID).
				SetTaskID(t.TaskID).
				SetTaskFeedID(t.TaskFeedID).
				SetTeammateID(t.TeammateID).
				SetCreatedAt(t.CreatedAt).
				SetUpdatedAt(t.UpdatedAt)
		}
		createdTaskFeedLikes, err = client.TaskFeedLike.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
	}

	var createdTaskFiles []*model.TaskFile
	if len(input.TaskFiles) > 0 {
		bulk := make([]*ent.TaskFileCreate, len(input.TaskFiles))
		for i, t := range input.TaskFiles {
			bulk[i] = client.TaskFile.Create().
				SetID(t.ID).
				SetTaskID(t.TaskID).
				SetTaskFeedID(t.TaskFeedID).
				SetName(t.Name).
				SetSrc(t.Src).
				SetAttached(t.Attached).
				SetFileTypeID(t.FileTypeID).
				SetProjectID(t.ProjectID).
				SetCreatedAt(t.CreatedAt).
				SetUpdatedAt(t.UpdatedAt)
		}
		createdTaskFiles, err = client.TaskFile.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
	}

	if createdTaskFeedLikes != nil {
		payload.TaskFeedLikes = createdTaskFeedLikes
	}
	if createdTaskFiles != nil {
		payload.TaskFiles = createdTaskFiles
	}
	payload.TaskFeed = createdTaskFeed

	return payload, nil
}
