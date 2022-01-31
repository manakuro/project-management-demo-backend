// Code generated by entc, DO NOT EDIT.

package taskfile

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskFeedID applies equality check predicate on the "task_feed_id" field. It's identical to TaskFeedIDEQ.
func TaskFeedID(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskFeedID), v))
	})
}

// FileTypeID applies equality check predicate on the "file_type_id" field. It's identical to FileTypeIDEQ.
func FileTypeID(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileTypeID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Src applies equality check predicate on the "src" field. It's identical to SrcEQ.
func Src(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSrc), v))
	})
}

// Attached applies equality check predicate on the "attached" field. It's identical to AttachedEQ.
func Attached(v bool) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttached), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func ProjectIDNotIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func ProjectIDGT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProjectID), v))
	})
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProjectID), v))
	})
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProjectID), v))
	})
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProjectID), v))
	})
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProjectID), vc))
	})
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProjectID), vc))
	})
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProjectID), vc))
	})
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProjectID), vc))
	})
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProjectID), vc))
	})
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskID), v))
	})
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskID), v))
	})
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func TaskIDNotIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func TaskIDGT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskID), v))
	})
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskID), v))
	})
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskID), v))
	})
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskID), v))
	})
}

// TaskIDContains applies the Contains predicate on the "task_id" field.
func TaskIDContains(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasPrefix applies the HasPrefix predicate on the "task_id" field.
func TaskIDHasPrefix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskID), vc))
	})
}

// TaskIDHasSuffix applies the HasSuffix predicate on the "task_id" field.
func TaskIDHasSuffix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskID), vc))
	})
}

// TaskIDEqualFold applies the EqualFold predicate on the "task_id" field.
func TaskIDEqualFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskID), vc))
	})
}

// TaskIDContainsFold applies the ContainsFold predicate on the "task_id" field.
func TaskIDContainsFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskID), vc))
	})
}

// TaskFeedIDEQ applies the EQ predicate on the "task_feed_id" field.
func TaskFeedIDEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDNEQ applies the NEQ predicate on the "task_feed_id" field.
func TaskFeedIDNEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDIn applies the In predicate on the "task_feed_id" field.
func TaskFeedIDIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskFeedID), v...))
	})
}

// TaskFeedIDNotIn applies the NotIn predicate on the "task_feed_id" field.
func TaskFeedIDNotIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskFeedID), v...))
	})
}

// TaskFeedIDGT applies the GT predicate on the "task_feed_id" field.
func TaskFeedIDGT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDGTE applies the GTE predicate on the "task_feed_id" field.
func TaskFeedIDGTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDLT applies the LT predicate on the "task_feed_id" field.
func TaskFeedIDLT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDLTE applies the LTE predicate on the "task_feed_id" field.
func TaskFeedIDLTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskFeedID), v))
	})
}

// TaskFeedIDContains applies the Contains predicate on the "task_feed_id" field.
func TaskFeedIDContains(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskFeedID), vc))
	})
}

// TaskFeedIDHasPrefix applies the HasPrefix predicate on the "task_feed_id" field.
func TaskFeedIDHasPrefix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskFeedID), vc))
	})
}

// TaskFeedIDHasSuffix applies the HasSuffix predicate on the "task_feed_id" field.
func TaskFeedIDHasSuffix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskFeedID), vc))
	})
}

// TaskFeedIDEqualFold applies the EqualFold predicate on the "task_feed_id" field.
func TaskFeedIDEqualFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskFeedID), vc))
	})
}

// TaskFeedIDContainsFold applies the ContainsFold predicate on the "task_feed_id" field.
func TaskFeedIDContainsFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskFeedID), vc))
	})
}

// FileTypeIDEQ applies the EQ predicate on the "file_type_id" field.
func FileTypeIDEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDNEQ applies the NEQ predicate on the "file_type_id" field.
func FileTypeIDNEQ(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDIn applies the In predicate on the "file_type_id" field.
func FileTypeIDIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFileTypeID), v...))
	})
}

// FileTypeIDNotIn applies the NotIn predicate on the "file_type_id" field.
func FileTypeIDNotIn(vs ...ulid.ID) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFileTypeID), v...))
	})
}

// FileTypeIDGT applies the GT predicate on the "file_type_id" field.
func FileTypeIDGT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDGTE applies the GTE predicate on the "file_type_id" field.
func FileTypeIDGTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDLT applies the LT predicate on the "file_type_id" field.
func FileTypeIDLT(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDLTE applies the LTE predicate on the "file_type_id" field.
func FileTypeIDLTE(v ulid.ID) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFileTypeID), v))
	})
}

// FileTypeIDContains applies the Contains predicate on the "file_type_id" field.
func FileTypeIDContains(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFileTypeID), vc))
	})
}

// FileTypeIDHasPrefix applies the HasPrefix predicate on the "file_type_id" field.
func FileTypeIDHasPrefix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFileTypeID), vc))
	})
}

// FileTypeIDHasSuffix applies the HasSuffix predicate on the "file_type_id" field.
func FileTypeIDHasSuffix(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFileTypeID), vc))
	})
}

// FileTypeIDEqualFold applies the EqualFold predicate on the "file_type_id" field.
func FileTypeIDEqualFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFileTypeID), vc))
	})
}

// FileTypeIDContainsFold applies the ContainsFold predicate on the "file_type_id" field.
func FileTypeIDContainsFold(v ulid.ID) predicate.TaskFile {
	vc := string(v)
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFileTypeID), vc))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func NameGT(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SrcEQ applies the EQ predicate on the "src" field.
func SrcEQ(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSrc), v))
	})
}

// SrcNEQ applies the NEQ predicate on the "src" field.
func SrcNEQ(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSrc), v))
	})
}

// SrcIn applies the In predicate on the "src" field.
func SrcIn(vs ...string) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSrc), v...))
	})
}

// SrcNotIn applies the NotIn predicate on the "src" field.
func SrcNotIn(vs ...string) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSrc), v...))
	})
}

// SrcGT applies the GT predicate on the "src" field.
func SrcGT(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSrc), v))
	})
}

// SrcGTE applies the GTE predicate on the "src" field.
func SrcGTE(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSrc), v))
	})
}

// SrcLT applies the LT predicate on the "src" field.
func SrcLT(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSrc), v))
	})
}

// SrcLTE applies the LTE predicate on the "src" field.
func SrcLTE(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSrc), v))
	})
}

// SrcContains applies the Contains predicate on the "src" field.
func SrcContains(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSrc), v))
	})
}

// SrcHasPrefix applies the HasPrefix predicate on the "src" field.
func SrcHasPrefix(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSrc), v))
	})
}

// SrcHasSuffix applies the HasSuffix predicate on the "src" field.
func SrcHasSuffix(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSrc), v))
	})
}

// SrcEqualFold applies the EqualFold predicate on the "src" field.
func SrcEqualFold(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSrc), v))
	})
}

// SrcContainsFold applies the ContainsFold predicate on the "src" field.
func SrcContainsFold(v string) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSrc), v))
	})
}

// AttachedEQ applies the EQ predicate on the "attached" field.
func AttachedEQ(v bool) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttached), v))
	})
}

// AttachedNEQ applies the NEQ predicate on the "attached" field.
func AttachedNEQ(v bool) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttached), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.TaskFile {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskFile(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
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

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
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

// HasTaskFeed applies the HasEdge predicate on the "task_feed" edge.
func HasTaskFeed() predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskFeedTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskFeedTable, TaskFeedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskFeedWith applies the HasEdge predicate on the "task_feed" edge with a given conditions (other predicates).
func HasTaskFeedWith(preds ...predicate.TaskFeed) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskFeedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskFeedTable, TaskFeedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFileType applies the HasEdge predicate on the "file_type" edge.
func HasFileType() predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTypeTable, FileTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFileTypeWith applies the HasEdge predicate on the "file_type" edge with a given conditions (other predicates).
func HasFileTypeWith(preds ...predicate.FileType) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTypeTable, FileTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskFile) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskFile) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
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
func Not(p predicate.TaskFile) predicate.TaskFile {
	return predicate.TaskFile(func(s *sql.Selector) {
		p(s.Not())
	})
}