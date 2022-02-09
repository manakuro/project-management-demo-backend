// Code generated by entc, DO NOT EDIT.

package tasklistsortstatus

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func NameGT(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StatusCodeEQ applies the EQ predicate on the "status_code" field.
func StatusCodeEQ(v StatusCode) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusCode), v))
	})
}

// StatusCodeNEQ applies the NEQ predicate on the "status_code" field.
func StatusCodeNEQ(v StatusCode) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusCode), v))
	})
}

// StatusCodeIn applies the In predicate on the "status_code" field.
func StatusCodeIn(vs ...StatusCode) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusCode), v...))
	})
}

// StatusCodeNotIn applies the NotIn predicate on the "status_code" field.
func StatusCodeNotIn(vs ...StatusCode) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusCode), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.TaskListSortStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasTeammateTaskListStatuses applies the HasEdge predicate on the "teammateTaskListStatuses" edge.
func HasTeammateTaskListStatuses() predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeammateTaskListStatusesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeammateTaskListStatusesTable, TeammateTaskListStatusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeammateTaskListStatusesWith applies the HasEdge predicate on the "teammateTaskListStatuses" edge with a given conditions (other predicates).
func HasTeammateTaskListStatusesWith(preds ...predicate.TeammateTaskListStatus) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeammateTaskListStatusesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeammateTaskListStatusesTable, TeammateTaskListStatusesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectTaskListStatuses applies the HasEdge predicate on the "projectTaskListStatuses" edge.
func HasProjectTaskListStatuses() predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTaskListStatusesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectTaskListStatusesTable, ProjectTaskListStatusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectTaskListStatusesWith applies the HasEdge predicate on the "projectTaskListStatuses" edge with a given conditions (other predicates).
func HasProjectTaskListStatusesWith(preds ...predicate.ProjectTaskListStatus) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTaskListStatusesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectTaskListStatusesTable, ProjectTaskListStatusesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskListSortStatus) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskListSortStatus) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
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
func Not(p predicate.TaskListSortStatus) predicate.TaskListSortStatus {
	return predicate.TaskListSortStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
