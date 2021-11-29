package mixin

import (
	"project-management-demo-backend/ent/schema/ulid"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NewUlidMixin creates a Mixin that encodes the provided prefix.
func NewUlidMixin(prefix string) *UlidMixin {
	return &UlidMixin{prefix: prefix}
}

// UlidMixin defines an ent Mixin that captures the PULID prefix for a type.
type UlidMixin struct {
	mixin.Schema
	prefix string
}

// Fields provides the id field.
func (m UlidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID { return ulid.MustNew(m.prefix) }),
	}
}

// Annotation captures the id prefix for a type.
type Annotation struct {
	Prefix string
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return "ULID"
}

// Annotations returns the annotations for a Mixin instance.
func (m UlidMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: m.prefix},
	}
}
