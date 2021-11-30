package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/dbschema"
)

var databaseSchema = dbschema.New()

// IDToType maps an ulid.ID to the underlying table.
func IDToType(ctx context.Context, id ulid.ID) (string, error) {
	if len(id) < 3 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:3]
	typ, err := databaseSchema.FindTableByID(string(prefix))
	if err != nil {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
