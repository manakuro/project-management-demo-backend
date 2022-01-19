package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// FavoriteProject is the model entity for the FavoriteProject schema
type FavoriteProject = ent.FavoriteProject

// FavoriteProjectConnection is the connection containing edges to FavoriteProject.
type FavoriteProjectConnection = ent.FavoriteProjectConnection

// CreateFavoriteProjectInput represents a mutation input for creating test users.
type CreateFavoriteProjectInput = ent.CreateFavoriteProjectInput

// FavoriteProjectWhereInput represents a where input for filtering FavoriteProject queries.
type FavoriteProjectWhereInput = ent.FavoriteProjectWhereInput

// UpdateFavoriteProjectInput represents a mutation input for updating test users.
type UpdateFavoriteProjectInput = ent.UpdateFavoriteProjectInput

// DeleteFavoriteProjectInput represents a mutation input for updating test users.
type DeleteFavoriteProjectInput struct {
	ProjectID  ulid.ID
	TeammateID ulid.ID
}
