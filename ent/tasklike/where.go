// Code generated by entc, DO NOT EDIT.

package tasklike

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TeammateID applies equality check predicate on the "teammate_id" field. It's identical to TeammateIDEQ.
func TeammateID(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeammateID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskID), v))
	})
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...ulid.ID) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskID), v...))
	})
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...ulid.ID) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskID), v...))
	})
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskID), v))
	})
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskID), v))
	})
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskID), v))
	})
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskID), v))
	})
}

// TaskIDContains applies the Contains predicate on the "task_id" field.
func TaskIDContains(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasPrefix applies the HasPrefix predicate on the "task_id" field.
func TaskIDHasPrefix(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasSuffix applies the HasSuffix predicate on the "task_id" field.
func TaskIDHasSuffix(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskID), vc))
	})
}

// TaskIDEqualFold applies the EqualFold predicate on the "task_id" field.
func TaskIDEqualFold(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskID), vc))
	})
}

// TaskIDContainsFold applies the ContainsFold predicate on the "task_id" field.
func TaskIDContainsFold(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskID), vc))
	})
}

// TeammateIDEQ applies the EQ predicate on the "teammate_id" field.
func TeammateIDEQ(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeammateID), v))
	})
}

// TeammateIDNEQ applies the NEQ predicate on the "teammate_id" field.
func TeammateIDNEQ(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeammateID), v))
	})
}

// TeammateIDIn applies the In predicate on the "teammate_id" field.
func TeammateIDIn(vs ...ulid.ID) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func TeammateIDNotIn(vs ...ulid.ID) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func TeammateIDGT(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeammateID), v))
	})
}

// TeammateIDGTE applies the GTE predicate on the "teammate_id" field.
func TeammateIDGTE(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeammateID), v))
	})
}

// TeammateIDLT applies the LT predicate on the "teammate_id" field.
func TeammateIDLT(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeammateID), v))
	})
}

// TeammateIDLTE applies the LTE predicate on the "teammate_id" field.
func TeammateIDLTE(v ulid.ID) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeammateID), v))
	})
}

// TeammateIDContains applies the Contains predicate on the "teammate_id" field.
func TeammateIDContains(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDHasPrefix applies the HasPrefix predicate on the "teammate_id" field.
func TeammateIDHasPrefix(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDHasSuffix applies the HasSuffix predicate on the "teammate_id" field.
func TeammateIDHasSuffix(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDEqualFold applies the EqualFold predicate on the "teammate_id" field.
func TeammateIDEqualFold(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeammateID), vc))
	})
}

// TeammateIDContainsFold applies the ContainsFold predicate on the "teammate_id" field.
func TeammateIDContainsFold(v ulid.ID) predicate.TaskLike {
	vc := string(v)
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeammateID), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.TaskLike {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskLike(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeammate applies the HasEdge predicate on the "teammate" edge.
func HasTeammate() predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeammateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeammateTable, TeammateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeammateWith applies the HasEdge predicate on the "teammate" edge with a given conditions (other predicates).
func HasTeammateWith(preds ...predicate.Teammate) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskLike) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskLike) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
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
func Not(p predicate.TaskLike) predicate.TaskLike {
	return predicate.TaskLike(func(s *sql.Selector) {
		p(s.Not())
	})
}
