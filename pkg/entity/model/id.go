package model

import "project-management-demo-backend/ent/schema/pulid"

// ID implements a PULID - a prefixed ULID.
type ID = pulid.ID

// Value returns pulid.ID.
//func (i *ID) Value() pulid.ID {
//	return *i.ID
//}
