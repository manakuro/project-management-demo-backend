// Code generated by entc, DO NOT EDIT.

package teammate

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImage), v))
	})
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImage), v...))
	})
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImage), v...))
	})
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImage), v))
	})
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImage), v))
	})
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImage), v))
	})
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImage), v))
	})
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImage), v))
	})
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImage), v))
	})
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImage), v))
	})
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImage), v))
	})
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImage), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Teammate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teammate(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasWorkspaces applies the HasEdge predicate on the "workspaces" edge.
func HasWorkspaces() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspacesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspacesTable, WorkspacesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspacesWith applies the HasEdge predicate on the "workspaces" edge with a given conditions (other predicates).
func HasWorkspacesWith(preds ...predicate.Workspace) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspacesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspacesTable, WorkspacesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjects applies the HasEdge predicate on the "projects" edge.
func HasProjects() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectsTable, ProjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectsWith applies the HasEdge predicate on the "projects" edge with a given conditions (other predicates).
func HasProjectsWith(preds ...predicate.Project) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectsTable, ProjectsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectTeammates applies the HasEdge predicate on the "project_teammates" edge.
func HasProjectTeammates() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTeammatesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectTeammatesTable, ProjectTeammatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectTeammatesWith applies the HasEdge predicate on the "project_teammates" edge with a given conditions (other predicates).
func HasProjectTeammatesWith(preds ...predicate.ProjectTeammate) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTeammatesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectTeammatesTable, ProjectTeammatesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkspaceTeammates applies the HasEdge predicate on the "workspace_teammates" edge.
func HasWorkspaceTeammates() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceTeammatesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspaceTeammatesTable, WorkspaceTeammatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceTeammatesWith applies the HasEdge predicate on the "workspace_teammates" edge with a given conditions (other predicates).
func HasWorkspaceTeammatesWith(preds ...predicate.WorkspaceTeammate) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspaceTeammatesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkspaceTeammatesTable, WorkspaceTeammatesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFavoriteProjects applies the HasEdge predicate on the "favorite_projects" edge.
func HasFavoriteProjects() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoriteProjectsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoriteProjectsTable, FavoriteProjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFavoriteProjectsWith applies the HasEdge predicate on the "favorite_projects" edge with a given conditions (other predicates).
func HasFavoriteProjectsWith(preds ...predicate.FavoriteProject) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoriteProjectsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoriteProjectsTable, FavoriteProjectsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFavoriteWorkspaces applies the HasEdge predicate on the "favorite_workspaces" edge.
func HasFavoriteWorkspaces() predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoriteWorkspacesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoriteWorkspacesTable, FavoriteWorkspacesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFavoriteWorkspacesWith applies the HasEdge predicate on the "favorite_workspaces" edge with a given conditions (other predicates).
func HasFavoriteWorkspacesWith(preds ...predicate.FavoriteWorkspace) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoriteWorkspacesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoriteWorkspacesTable, FavoriteWorkspacesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Teammate) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Teammate) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
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
func Not(p predicate.Teammate) predicate.Teammate {
	return predicate.Teammate(func(s *sql.Selector) {
		p(s.Not())
	})
}
