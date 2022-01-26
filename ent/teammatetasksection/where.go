// Code generated by entc, DO NOT EDIT.

package teammatetasksection

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TeammateID applies equality check predicate on the "teammate_id" field. It's identical to TeammateIDEQ.
func TeammateID(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeammateID), v))
	})
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Assigned applies equality check predicate on the "assigned" field. It's identical to AssignedEQ.
func Assigned(v bool) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssigned), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TeammateIDEQ applies the EQ predicate on the "teammate_id" field.
func TeammateIDEQ(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeammateID), v))
	})
}

// TeammateIDNEQ applies the NEQ predicate on the "teammate_id" field.
func TeammateIDNEQ(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeammateID), v))
	})
}

// TeammateIDIn applies the In predicate on the "teammate_id" field.
func TeammateIDIn(vs ...ulid.ID) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeammateID), v...))
	})
}

// TeammateIDNotIn applies the NotIn predicate on the "teammate_id" field.
func TeammateIDNotIn(vs ...ulid.ID) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeammateID), v...))
	})
}

// TeammateIDGT applies the GT predicate on the "teammate_id" field.
func TeammateIDGT(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeammateID), v))
	})
}

// TeammateIDGTE applies the GTE predicate on the "teammate_id" field.
func TeammateIDGTE(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeammateID), v))
	})
}

// TeammateIDLT applies the LT predicate on the "teammate_id" field.
func TeammateIDLT(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeammateID), v))
	})
}

// TeammateIDLTE applies the LTE predicate on the "teammate_id" field.
func TeammateIDLTE(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeammateID), v))
	})
}

// TeammateIDContains applies the Contains predicate on the "teammate_id" field.
func TeammateIDContains(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDHasPrefix applies the HasPrefix predicate on the "teammate_id" field.
func TeammateIDHasPrefix(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDHasSuffix applies the HasSuffix predicate on the "teammate_id" field.
func TeammateIDHasSuffix(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDEqualFold applies the EqualFold predicate on the "teammate_id" field.
func TeammateIDEqualFold(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDContainsFold applies the ContainsFold predicate on the "teammate_id" field.
func TeammateIDContainsFold(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeammateID), vc))
	})
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...ulid.ID) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...ulid.ID) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v ulid.ID) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDContains applies the Contains predicate on the "workspace_id" field.
func WorkspaceIDContains(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWorkspaceID), vc))
	})
}

// WorkspaceIDHasPrefix applies the HasPrefix predicate on the "workspace_id" field.
func WorkspaceIDHasPrefix(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWorkspaceID), vc))
	})
}

// WorkspaceIDHasSuffix applies the HasSuffix predicate on the "workspace_id" field.
func WorkspaceIDHasSuffix(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWorkspaceID), vc))
	})
}

// WorkspaceIDEqualFold applies the EqualFold predicate on the "workspace_id" field.
func WorkspaceIDEqualFold(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWorkspaceID), vc))
	})
}

// WorkspaceIDContainsFold applies the ContainsFold predicate on the "workspace_id" field.
func WorkspaceIDContainsFold(v ulid.ID) predicate.TeammateTaskSection {
	vc := string(v)
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWorkspaceID), vc))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AssignedEQ applies the EQ predicate on the "assigned" field.
func AssignedEQ(v bool) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssigned), v))
	})
}

// AssignedNEQ applies the NEQ predicate on the "assigned" field.
func AssignedNEQ(v bool) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssigned), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TeammateTaskSection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasTeammate applies the HasEdge predicate on the "teammate" edge.
func HasTeammate() predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeammateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeammateTable, TeammateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeammateWith applies the HasEdge predicate on the "teammate" edge with a given conditions (other predicates).
func HasTeammateWith(preds ...predicate.Teammate) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeammateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeammateTable, TeammateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkspace applies the HasEdge predicate on the "workspace" edge.
func HasWorkspace() predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceWith applies the HasEdge predicate on the "workspace" edge with a given conditions (other predicates).
func HasWorkspaceWith(preds ...predicate.Workspace) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TeammateTaskSection) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TeammateTaskSection) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TeammateTaskSection) predicate.TeammateTaskSection {
	return predicate.TeammateTaskSection(func(s *sql.Selector) {
		p(s.Not())
	})
}
