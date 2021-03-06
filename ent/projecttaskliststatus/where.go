// Code generated by entc, DO NOT EDIT.

package projecttaskliststatus

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// TaskListCompletedStatusID applies equality check predicate on the "task_list_completed_status_id" field. It's identical to TaskListCompletedStatusIDEQ.
func TaskListCompletedStatusID(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListSortStatusID applies equality check predicate on the "task_list_sort_status_id" field. It's identical to TaskListSortStatusIDEQ.
func TaskListSortStatusID(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskListSortStatusID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProjectID), v...))
	})
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProjectID), v...))
	})
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProjectID), v))
	})
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProjectID), v))
	})
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProjectID), v))
	})
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProjectID), v))
	})
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProjectID), vc))
	})
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProjectID), vc))
	})
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProjectID), vc))
	})
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProjectID), vc))
	})
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProjectID), vc))
	})
}

// TaskListCompletedStatusIDEQ applies the EQ predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDNEQ applies the NEQ predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDNEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDIn applies the In predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskListCompletedStatusID), v...))
	})
}

// TaskListCompletedStatusIDNotIn applies the NotIn predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDNotIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskListCompletedStatusID), v...))
	})
}

// TaskListCompletedStatusIDGT applies the GT predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDGT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDGTE applies the GTE predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDGTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDLT applies the LT predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDLT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDLTE applies the LTE predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDLTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskListCompletedStatusID), v))
	})
}

// TaskListCompletedStatusIDContains applies the Contains predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDContains(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskListCompletedStatusID), vc))
	})
}

// TaskListCompletedStatusIDHasPrefix applies the HasPrefix predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDHasPrefix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskListCompletedStatusID), vc))
	})
}

// TaskListCompletedStatusIDHasSuffix applies the HasSuffix predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDHasSuffix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskListCompletedStatusID), vc))
	})
}

// TaskListCompletedStatusIDEqualFold applies the EqualFold predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDEqualFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskListCompletedStatusID), vc))
	})
}

// TaskListCompletedStatusIDContainsFold applies the ContainsFold predicate on the "task_list_completed_status_id" field.
func TaskListCompletedStatusIDContainsFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskListCompletedStatusID), vc))
	})
}

// TaskListSortStatusIDEQ applies the EQ predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDNEQ applies the NEQ predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDNEQ(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDIn applies the In predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskListSortStatusID), v...))
	})
}

// TaskListSortStatusIDNotIn applies the NotIn predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDNotIn(vs ...ulid.ID) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskListSortStatusID), v...))
	})
}

// TaskListSortStatusIDGT applies the GT predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDGT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDGTE applies the GTE predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDGTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDLT applies the LT predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDLT(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDLTE applies the LTE predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDLTE(v ulid.ID) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskListSortStatusID), v))
	})
}

// TaskListSortStatusIDContains applies the Contains predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDContains(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskListSortStatusID), vc))
	})
}

// TaskListSortStatusIDHasPrefix applies the HasPrefix predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDHasPrefix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskListSortStatusID), vc))
	})
}

// TaskListSortStatusIDHasSuffix applies the HasSuffix predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDHasSuffix(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskListSortStatusID), vc))
	})
}

// TaskListSortStatusIDEqualFold applies the EqualFold predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDEqualFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskListSortStatusID), vc))
	})
}

// TaskListSortStatusIDContainsFold applies the ContainsFold predicate on the "task_list_sort_status_id" field.
func TaskListSortStatusIDContainsFold(v ulid.ID) predicate.ProjectTaskListStatus {
	vc := string(v)
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskListSortStatusID), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.ProjectTaskListStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskListCompletedStatus applies the HasEdge predicate on the "taskListCompletedStatus" edge.
func HasTaskListCompletedStatus() predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskListCompletedStatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskListCompletedStatusTable, TaskListCompletedStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskListCompletedStatusWith applies the HasEdge predicate on the "taskListCompletedStatus" edge with a given conditions (other predicates).
func HasTaskListCompletedStatusWith(preds ...predicate.TaskListCompletedStatus) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskListCompletedStatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskListCompletedStatusTable, TaskListCompletedStatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskListSortStatus applies the HasEdge predicate on the "taskListSortStatus" edge.
func HasTaskListSortStatus() predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskListSortStatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskListSortStatusTable, TaskListSortStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskListSortStatusWith applies the HasEdge predicate on the "taskListSortStatus" edge with a given conditions (other predicates).
func HasTaskListSortStatusWith(preds ...predicate.TaskListSortStatus) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskListSortStatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskListSortStatusTable, TaskListSortStatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProjectTaskListStatus) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProjectTaskListStatus) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
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
func Not(p predicate.ProjectTaskListStatus) predicate.ProjectTaskListStatus {
	return predicate.ProjectTaskListStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
