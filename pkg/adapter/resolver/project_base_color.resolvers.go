package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
)

func (r *projectBaseColorResolver) CreatedAt(ctx context.Context, obj *ent.ProjectBaseColor) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectBaseColorResolver) UpdatedAt(ctx context.Context, obj *ent.ProjectBaseColor) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ProjectBaseColor returns generated.ProjectBaseColorResolver implementation.
func (r *Resolver) ProjectBaseColor() generated.ProjectBaseColorResolver {
	return &projectBaseColorResolver{r}
}

type projectBaseColorResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *projectBaseColorResolver) Color(ctx context.Context, obj *ent.ProjectBaseColor) (ulid.ID, error) {
	panic(fmt.Errorf("not implemented"))
}
