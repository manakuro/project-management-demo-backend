// Code generated by entc, DO NOT EDIT.

package deletedtaskactivitytask

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TaskActivityID applies equality check predicate on the "task_activity_id" field. It's identical to TaskActivityIDEQ.
func TaskActivityID(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityID), v))
	})
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskActivityTaskID applies equality check predicate on the "task_activity_task_id" field. It's identical to TaskActivityTaskIDEQ.
func TaskActivityTaskID(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskCreatedAt applies equality check predicate on the "task_activity_task_created_at" field. It's identical to TaskActivityTaskCreatedAtEQ.
func TaskActivityTaskCreatedAt(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskUpdatedAt applies equality check predicate on the "task_activity_task_updated_at" field. It's identical to TaskActivityTaskUpdatedAtEQ.
func TaskActivityTaskUpdatedAt(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TaskActivityIDEQ applies the EQ predicate on the "task_activity_id" field.
func TaskActivityIDEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDNEQ applies the NEQ predicate on the "task_activity_id" field.
func TaskActivityIDNEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDIn applies the In predicate on the "task_activity_id" field.
func TaskActivityIDIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskActivityID), v...))
	})
}

// TaskActivityIDNotIn applies the NotIn predicate on the "task_activity_id" field.
func TaskActivityIDNotIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskActivityID), v...))
	})
}

// TaskActivityIDGT applies the GT predicate on the "task_activity_id" field.
func TaskActivityIDGT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDGTE applies the GTE predicate on the "task_activity_id" field.
func TaskActivityIDGTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDLT applies the LT predicate on the "task_activity_id" field.
func TaskActivityIDLT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDLTE applies the LTE predicate on the "task_activity_id" field.
func TaskActivityIDLTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskActivityID), v))
	})
}

// TaskActivityIDContains applies the Contains predicate on the "task_activity_id" field.
func TaskActivityIDContains(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskActivityID), vc))
	})
}

// TaskActivityIDHasPrefix applies the HasPrefix predicate on the "task_activity_id" field.
func TaskActivityIDHasPrefix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskActivityID), vc))
	})
}

// TaskActivityIDHasSuffix applies the HasSuffix predicate on the "task_activity_id" field.
func TaskActivityIDHasSuffix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskActivityID), vc))
	})
}

// TaskActivityIDEqualFold applies the EqualFold predicate on the "task_activity_id" field.
func TaskActivityIDEqualFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskActivityID), vc))
	})
}

// TaskActivityIDContainsFold applies the ContainsFold predicate on the "task_activity_id" field.
func TaskActivityIDContainsFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskActivityID), vc))
	})
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskID), v))
	})
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func TaskIDNotIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func TaskIDGT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskID), v))
	})
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskID), v))
	})
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskID), v))
	})
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskID), v))
	})
}

// TaskIDContains applies the Contains predicate on the "task_id" field.
func TaskIDContains(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasPrefix applies the HasPrefix predicate on the "task_id" field.
func TaskIDHasPrefix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasSuffix applies the HasSuffix predicate on the "task_id" field.
func TaskIDHasSuffix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskID), vc))
	})
}

// TaskIDEqualFold applies the EqualFold predicate on the "task_id" field.
func TaskIDEqualFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskID), vc))
	})
}

// TaskIDContainsFold applies the ContainsFold predicate on the "task_id" field.
func TaskIDContainsFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskID), vc))
	})
}

// TaskActivityTaskIDEQ applies the EQ predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDNEQ applies the NEQ predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDNEQ(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDIn applies the In predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskActivityTaskID), v...))
	})
}

// TaskActivityTaskIDNotIn applies the NotIn predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDNotIn(vs ...ulid.ID) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskActivityTaskID), v...))
	})
}

// TaskActivityTaskIDGT applies the GT predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDGT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDGTE applies the GTE predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDGTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDLT applies the LT predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDLT(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDLTE applies the LTE predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDLTE(v ulid.ID) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskActivityTaskID), v))
	})
}

// TaskActivityTaskIDContains applies the Contains predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDContains(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskActivityTaskID), vc))
	})
}

// TaskActivityTaskIDHasPrefix applies the HasPrefix predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDHasPrefix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskActivityTaskID), vc))
	})
}

// TaskActivityTaskIDHasSuffix applies the HasSuffix predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDHasSuffix(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskActivityTaskID), vc))
	})
}

// TaskActivityTaskIDEqualFold applies the EqualFold predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDEqualFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskActivityTaskID), vc))
	})
}

// TaskActivityTaskIDContainsFold applies the ContainsFold predicate on the "task_activity_task_id" field.
func TaskActivityTaskIDContainsFold(v ulid.ID) predicate.DeletedTaskActivityTask {
	vc := string(v)
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskActivityTaskID), vc))
	})
}

// TaskActivityTaskCreatedAtEQ applies the EQ predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskCreatedAtNEQ applies the NEQ predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtNEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskCreatedAtIn applies the In predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskActivityTaskCreatedAt), v...))
	})
}

// TaskActivityTaskCreatedAtNotIn applies the NotIn predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtNotIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskActivityTaskCreatedAt), v...))
	})
}

// TaskActivityTaskCreatedAtGT applies the GT predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtGT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskCreatedAtGTE applies the GTE predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtGTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskCreatedAtLT applies the LT predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtLT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskCreatedAtLTE applies the LTE predicate on the "task_activity_task_created_at" field.
func TaskActivityTaskCreatedAtLTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskActivityTaskCreatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtEQ applies the EQ predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtNEQ applies the NEQ predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtNEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtIn applies the In predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskActivityTaskUpdatedAt), v...))
	})
}

// TaskActivityTaskUpdatedAtNotIn applies the NotIn predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtNotIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskActivityTaskUpdatedAt), v...))
	})
}

// TaskActivityTaskUpdatedAtGT applies the GT predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtGT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtGTE applies the GTE predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtGTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtLT applies the LT predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtLT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// TaskActivityTaskUpdatedAtLTE applies the LTE predicate on the "task_activity_task_updated_at" field.
func TaskActivityTaskUpdatedAtLTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskActivityTaskUpdatedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.DeletedTaskActivityTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeletedTaskActivityTask) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeletedTaskActivityTask) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
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
func Not(p predicate.DeletedTaskActivityTask) predicate.DeletedTaskActivityTask {
	return predicate.DeletedTaskActivityTask(func(s *sql.Selector) {
		p(s.Not())
	})
}