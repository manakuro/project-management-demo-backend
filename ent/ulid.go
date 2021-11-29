package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/dbschema"
)

// IDToType maps an ulid.ID to the underlying table.
func IDToType(ctx context.Context, id ulid.ID) (string, error) {
	if len(id) < 3 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:3]
	schema := dbschema.New()
	typ, err := schema.FindTableByID(string(prefix))
	if err != nil {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
