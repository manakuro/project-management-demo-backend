// Code generated by entc, DO NOT EDIT.

package color

import (
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
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
func IDNotIn(ids ...ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
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
func IDGT(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// Hex applies equality check predicate on the "hex" field. It's identical to HexEQ.
func Hex(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHex), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldColor), v))
	})
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldColor), v...))
	})
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldColor), v...))
	})
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldColor), v))
	})
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldColor), v))
	})
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldColor), v))
	})
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldColor), v))
	})
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldColor), v))
	})
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldColor), v))
	})
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldColor), v))
	})
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldColor), v))
	})
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldColor), v))
	})
}

// HexEQ applies the EQ predicate on the "hex" field.
func HexEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHex), v))
	})
}

// HexNEQ applies the NEQ predicate on the "hex" field.
func HexNEQ(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHex), v))
	})
}

// HexIn applies the In predicate on the "hex" field.
func HexIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHex), v...))
	})
}

// HexNotIn applies the NotIn predicate on the "hex" field.
func HexNotIn(vs ...string) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHex), v...))
	})
}

// HexGT applies the GT predicate on the "hex" field.
func HexGT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHex), v))
	})
}

// HexGTE applies the GTE predicate on the "hex" field.
func HexGTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHex), v))
	})
}

// HexLT applies the LT predicate on the "hex" field.
func HexLT(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHex), v))
	})
}

// HexLTE applies the LTE predicate on the "hex" field.
func HexLTE(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHex), v))
	})
}

// HexContains applies the Contains predicate on the "hex" field.
func HexContains(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHex), v))
	})
}

// HexHasPrefix applies the HasPrefix predicate on the "hex" field.
func HexHasPrefix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHex), v))
	})
}

// HexHasSuffix applies the HasSuffix predicate on the "hex" field.
func HexHasSuffix(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHex), v))
	})
}

// HexEqualFold applies the EqualFold predicate on the "hex" field.
func HexEqualFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHex), v))
	})
}

// HexContainsFold applies the ContainsFold predicate on the "hex" field.
func HexContainsFold(v string) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHex), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Color {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Color(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasProjectBaseColors applies the HasEdge predicate on the "project_base_colors" edge.
func HasProjectBaseColors() predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectBaseColorsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectBaseColorsTable, ProjectBaseColorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectBaseColorsWith applies the HasEdge predicate on the "project_base_colors" edge with a given conditions (other predicates).
func HasProjectBaseColorsWith(preds ...predicate.ProjectBaseColor) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectBaseColorsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectBaseColorsTable, ProjectBaseColorsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectLightColors applies the HasEdge predicate on the "project_light_colors" edge.
func HasProjectLightColors() predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLightColorsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLightColorsTable, ProjectLightColorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectLightColorsWith applies the HasEdge predicate on the "project_light_colors" edge with a given conditions (other predicates).
func HasProjectLightColorsWith(preds ...predicate.ProjectLightColor) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLightColorsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLightColorsTable, ProjectLightColorsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaskPriorities applies the HasEdge predicate on the "task_priorities" edge.
func HasTaskPriorities() predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskPrioritiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskPrioritiesTable, TaskPrioritiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskPrioritiesWith applies the HasEdge predicate on the "task_priorities" edge with a given conditions (other predicates).
func HasTaskPrioritiesWith(preds ...predicate.TaskPriority) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaskPrioritiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskPrioritiesTable, TaskPrioritiesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Color) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Color) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
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
func Not(p predicate.Color) predicate.Color {
	return predicate.Color(func(s *sql.Selector) {
		p(s.Not())
	})
}
