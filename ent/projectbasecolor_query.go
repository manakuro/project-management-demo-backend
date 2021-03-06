// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectBaseColorQuery is the builder for querying ProjectBaseColor entities.
type ProjectBaseColorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProjectBaseColor
	// eager-loading edges.
	withProjects *ProjectQuery
	withColor    *ColorQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectBaseColorQuery builder.
func (pbcq *ProjectBaseColorQuery) Where(ps ...predicate.ProjectBaseColor) *ProjectBaseColorQuery {
	pbcq.predicates = append(pbcq.predicates, ps...)
	return pbcq
}

// Limit adds a limit step to the query.
func (pbcq *ProjectBaseColorQuery) Limit(limit int) *ProjectBaseColorQuery {
	pbcq.limit = &limit
	return pbcq
}

// Offset adds an offset step to the query.
func (pbcq *ProjectBaseColorQuery) Offset(offset int) *ProjectBaseColorQuery {
	pbcq.offset = &offset
	return pbcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pbcq *ProjectBaseColorQuery) Unique(unique bool) *ProjectBaseColorQuery {
	pbcq.unique = &unique
	return pbcq
}

// Order adds an order step to the query.
func (pbcq *ProjectBaseColorQuery) Order(o ...OrderFunc) *ProjectBaseColorQuery {
	pbcq.order = append(pbcq.order, o...)
	return pbcq
}

// QueryProjects chains the current query on the "projects" edge.
func (pbcq *ProjectBaseColorQuery) QueryProjects() *ProjectQuery {
	query := &ProjectQuery{config: pbcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pbcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pbcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectbasecolor.Table, projectbasecolor.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, projectbasecolor.ProjectsTable, projectbasecolor.ProjectsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pbcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryColor chains the current query on the "color" edge.
func (pbcq *ProjectBaseColorQuery) QueryColor() *ColorQuery {
	query := &ColorQuery{config: pbcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pbcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pbcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectbasecolor.Table, projectbasecolor.FieldID, selector),
			sqlgraph.To(color.Table, color.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectbasecolor.ColorTable, projectbasecolor.ColorColumn),
		)
		fromU = sqlgraph.SetNeighbors(pbcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectBaseColor entity from the query.
// Returns a *NotFoundError when no ProjectBaseColor was found.
func (pbcq *ProjectBaseColorQuery) First(ctx context.Context) (*ProjectBaseColor, error) {
	nodes, err := pbcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectbasecolor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) FirstX(ctx context.Context) *ProjectBaseColor {
	node, err := pbcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectBaseColor ID from the query.
// Returns a *NotFoundError when no ProjectBaseColor ID was found.
func (pbcq *ProjectBaseColorQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = pbcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projectbasecolor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := pbcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectBaseColor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectBaseColor entity is found.
// Returns a *NotFoundError when no ProjectBaseColor entities are found.
func (pbcq *ProjectBaseColorQuery) Only(ctx context.Context) (*ProjectBaseColor, error) {
	nodes, err := pbcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectbasecolor.Label}
	default:
		return nil, &NotSingularError{projectbasecolor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) OnlyX(ctx context.Context) *ProjectBaseColor {
	node, err := pbcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectBaseColor ID in the query.
// Returns a *NotSingularError when more than one ProjectBaseColor ID is found.
// Returns a *NotFoundError when no entities are found.
func (pbcq *ProjectBaseColorQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = pbcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = &NotSingularError{projectbasecolor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := pbcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectBaseColors.
func (pbcq *ProjectBaseColorQuery) All(ctx context.Context) ([]*ProjectBaseColor, error) {
	if err := pbcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pbcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) AllX(ctx context.Context) []*ProjectBaseColor {
	nodes, err := pbcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectBaseColor IDs.
func (pbcq *ProjectBaseColorQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := pbcq.Select(projectbasecolor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := pbcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pbcq *ProjectBaseColorQuery) Count(ctx context.Context) (int, error) {
	if err := pbcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pbcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) CountX(ctx context.Context) int {
	count, err := pbcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pbcq *ProjectBaseColorQuery) Exist(ctx context.Context) (bool, error) {
	if err := pbcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pbcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pbcq *ProjectBaseColorQuery) ExistX(ctx context.Context) bool {
	exist, err := pbcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectBaseColorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pbcq *ProjectBaseColorQuery) Clone() *ProjectBaseColorQuery {
	if pbcq == nil {
		return nil
	}
	return &ProjectBaseColorQuery{
		config:       pbcq.config,
		limit:        pbcq.limit,
		offset:       pbcq.offset,
		order:        append([]OrderFunc{}, pbcq.order...),
		predicates:   append([]predicate.ProjectBaseColor{}, pbcq.predicates...),
		withProjects: pbcq.withProjects.Clone(),
		withColor:    pbcq.withColor.Clone(),
		// clone intermediate query.
		sql:    pbcq.sql.Clone(),
		path:   pbcq.path,
		unique: pbcq.unique,
	}
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (pbcq *ProjectBaseColorQuery) WithProjects(opts ...func(*ProjectQuery)) *ProjectBaseColorQuery {
	query := &ProjectQuery{config: pbcq.config}
	for _, opt := range opts {
		opt(query)
	}
	pbcq.withProjects = query
	return pbcq
}

// WithColor tells the query-builder to eager-load the nodes that are connected to
// the "color" edge. The optional arguments are used to configure the query builder of the edge.
func (pbcq *ProjectBaseColorQuery) WithColor(opts ...func(*ColorQuery)) *ProjectBaseColorQuery {
	query := &ColorQuery{config: pbcq.config}
	for _, opt := range opts {
		opt(query)
	}
	pbcq.withColor = query
	return pbcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ColorID ulid.ID `json:"color_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectBaseColor.Query().
//		GroupBy(projectbasecolor.FieldColorID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pbcq *ProjectBaseColorQuery) GroupBy(field string, fields ...string) *ProjectBaseColorGroupBy {
	group := &ProjectBaseColorGroupBy{config: pbcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pbcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pbcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ColorID ulid.ID `json:"color_id,omitempty"`
//	}
//
//	client.ProjectBaseColor.Query().
//		Select(projectbasecolor.FieldColorID).
//		Scan(ctx, &v)
//
func (pbcq *ProjectBaseColorQuery) Select(fields ...string) *ProjectBaseColorSelect {
	pbcq.fields = append(pbcq.fields, fields...)
	return &ProjectBaseColorSelect{ProjectBaseColorQuery: pbcq}
}

func (pbcq *ProjectBaseColorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pbcq.fields {
		if !projectbasecolor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pbcq.path != nil {
		prev, err := pbcq.path(ctx)
		if err != nil {
			return err
		}
		pbcq.sql = prev
	}
	return nil
}

func (pbcq *ProjectBaseColorQuery) sqlAll(ctx context.Context) ([]*ProjectBaseColor, error) {
	var (
		nodes       = []*ProjectBaseColor{}
		_spec       = pbcq.querySpec()
		loadedTypes = [2]bool{
			pbcq.withProjects != nil,
			pbcq.withColor != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProjectBaseColor{config: pbcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pbcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pbcq.withProjects; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*ProjectBaseColor)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Projects = []*Project{}
		}
		query.Where(predicate.Project(func(s *sql.Selector) {
			s.Where(sql.InValues(projectbasecolor.ProjectsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ProjectBaseColorID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "project_base_color_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Projects = append(node.Edges.Projects, n)
		}
	}

	if query := pbcq.withColor; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectBaseColor)
		for i := range nodes {
			fk := nodes[i].ColorID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(color.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "color_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Color = n
			}
		}
	}

	return nodes, nil
}

func (pbcq *ProjectBaseColorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pbcq.querySpec()
	_spec.Node.Columns = pbcq.fields
	if len(pbcq.fields) > 0 {
		_spec.Unique = pbcq.unique != nil && *pbcq.unique
	}
	return sqlgraph.CountNodes(ctx, pbcq.driver, _spec)
}

func (pbcq *ProjectBaseColorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pbcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pbcq *ProjectBaseColorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectbasecolor.Table,
			Columns: projectbasecolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projectbasecolor.FieldID,
			},
		},
		From:   pbcq.sql,
		Unique: true,
	}
	if unique := pbcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pbcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectbasecolor.FieldID)
		for i := range fields {
			if fields[i] != projectbasecolor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pbcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pbcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pbcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pbcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pbcq *ProjectBaseColorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pbcq.driver.Dialect())
	t1 := builder.Table(projectbasecolor.Table)
	columns := pbcq.fields
	if len(columns) == 0 {
		columns = projectbasecolor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pbcq.sql != nil {
		selector = pbcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pbcq.unique != nil && *pbcq.unique {
		selector.Distinct()
	}
	for _, p := range pbcq.predicates {
		p(selector)
	}
	for _, p := range pbcq.order {
		p(selector)
	}
	if offset := pbcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pbcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectBaseColorGroupBy is the group-by builder for ProjectBaseColor entities.
type ProjectBaseColorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pbcgb *ProjectBaseColorGroupBy) Aggregate(fns ...AggregateFunc) *ProjectBaseColorGroupBy {
	pbcgb.fns = append(pbcgb.fns, fns...)
	return pbcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pbcgb *ProjectBaseColorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pbcgb.path(ctx)
	if err != nil {
		return err
	}
	pbcgb.sql = query
	return pbcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pbcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pbcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pbcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) StringsX(ctx context.Context) []string {
	v, err := pbcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pbcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) StringX(ctx context.Context) string {
	v, err := pbcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pbcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pbcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) IntsX(ctx context.Context) []int {
	v, err := pbcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pbcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) IntX(ctx context.Context) int {
	v, err := pbcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pbcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pbcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pbcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pbcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pbcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pbcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pbcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pbcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pbcgb *ProjectBaseColorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pbcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pbcgb *ProjectBaseColorGroupBy) BoolX(ctx context.Context) bool {
	v, err := pbcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pbcgb *ProjectBaseColorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pbcgb.fields {
		if !projectbasecolor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pbcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pbcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pbcgb *ProjectBaseColorGroupBy) sqlQuery() *sql.Selector {
	selector := pbcgb.sql.Select()
	aggregation := make([]string, 0, len(pbcgb.fns))
	for _, fn := range pbcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pbcgb.fields)+len(pbcgb.fns))
		for _, f := range pbcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pbcgb.fields...)...)
}

// ProjectBaseColorSelect is the builder for selecting fields of ProjectBaseColor entities.
type ProjectBaseColorSelect struct {
	*ProjectBaseColorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pbcs *ProjectBaseColorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pbcs.prepareQuery(ctx); err != nil {
		return err
	}
	pbcs.sql = pbcs.ProjectBaseColorQuery.sqlQuery(ctx)
	return pbcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pbcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pbcs.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pbcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) StringsX(ctx context.Context) []string {
	v, err := pbcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pbcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) StringX(ctx context.Context) string {
	v, err := pbcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pbcs.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pbcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) IntsX(ctx context.Context) []int {
	v, err := pbcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pbcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) IntX(ctx context.Context) int {
	v, err := pbcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pbcs.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pbcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pbcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pbcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) Float64X(ctx context.Context) float64 {
	v, err := pbcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pbcs.fields) > 1 {
		return nil, errors.New("ent: ProjectBaseColorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pbcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) BoolsX(ctx context.Context) []bool {
	v, err := pbcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pbcs *ProjectBaseColorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pbcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectbasecolor.Label}
	default:
		err = fmt.Errorf("ent: ProjectBaseColorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pbcs *ProjectBaseColorSelect) BoolX(ctx context.Context) bool {
	v, err := pbcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pbcs *ProjectBaseColorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pbcs.sql.Query()
	if err := pbcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
