package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
)

func (r *archivedActivityResolver) UpdatedAt(ctx context.Context, obj *model.ArchivedActivity) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) ArchivedActivities(ctx context.Context, where model.ArchivedActivityWhereInput) ([]*model.ArchivedActivity, error) {
	as, err := r.controller.ArchivedActivity.List(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return as, nil
}

// ArchivedActivity returns generated.ArchivedActivityResolver implementation.
func (r *Resolver) ArchivedActivity() generated.ArchivedActivityResolver {
	return &archivedActivityResolver{r}
}

type archivedActivityResolver struct{ *Resolver }
