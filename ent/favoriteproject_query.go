// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteProjectQuery is the builder for querying FavoriteProject entities.
type FavoriteProjectQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FavoriteProject
	// eager-loading edges.
	withProject  *ProjectQuery
	withTeammate *TeammateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FavoriteProjectQuery builder.
func (fpq *FavoriteProjectQuery) Where(ps ...predicate.FavoriteProject) *FavoriteProjectQuery {
	fpq.predicates = append(fpq.predicates, ps...)
	return fpq
}

// Limit adds a limit step to the query.
func (fpq *FavoriteProjectQuery) Limit(limit int) *FavoriteProjectQuery {
	fpq.limit = &limit
	return fpq
}

// Offset adds an offset step to the query.
func (fpq *FavoriteProjectQuery) Offset(offset int) *FavoriteProjectQuery {
	fpq.offset = &offset
	return fpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fpq *FavoriteProjectQuery) Unique(unique bool) *FavoriteProjectQuery {
	fpq.unique = &unique
	return fpq
}

// Order adds an order step to the query.
func (fpq *FavoriteProjectQuery) Order(o ...OrderFunc) *FavoriteProjectQuery {
	fpq.order = append(fpq.order, o...)
	return fpq
}

// QueryProject chains the current query on the "project" edge.
func (fpq *FavoriteProjectQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: fpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(favoriteproject.Table, favoriteproject.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favoriteproject.ProjectTable, favoriteproject.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (fpq *FavoriteProjectQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: fpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(favoriteproject.Table, favoriteproject.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favoriteproject.TeammateTable, favoriteproject.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FavoriteProject entity from the query.
// Returns a *NotFoundError when no FavoriteProject was found.
func (fpq *FavoriteProjectQuery) First(ctx context.Context) (*FavoriteProject, error) {
	nodes, err := fpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{favoriteproject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) FirstX(ctx context.Context) *FavoriteProject {
	node, err := fpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FavoriteProject ID from the query.
// Returns a *NotFoundError when no FavoriteProject ID was found.
func (fpq *FavoriteProjectQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = fpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{favoriteproject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := fpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FavoriteProject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FavoriteProject entity is found.
// Returns a *NotFoundError when no FavoriteProject entities are found.
func (fpq *FavoriteProjectQuery) Only(ctx context.Context) (*FavoriteProject, error) {
	nodes, err := fpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{favoriteproject.Label}
	default:
		return nil, &NotSingularError{favoriteproject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) OnlyX(ctx context.Context) *FavoriteProject {
	node, err := fpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FavoriteProject ID in the query.
// Returns a *NotSingularError when more than one FavoriteProject ID is found.
// Returns a *NotFoundError when no entities are found.
func (fpq *FavoriteProjectQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = fpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = &NotSingularError{favoriteproject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := fpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FavoriteProjects.
func (fpq *FavoriteProjectQuery) All(ctx context.Context) ([]*FavoriteProject, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) AllX(ctx context.Context) []*FavoriteProject {
	nodes, err := fpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FavoriteProject IDs.
func (fpq *FavoriteProjectQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := fpq.Select(favoriteproject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := fpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fpq *FavoriteProjectQuery) Count(ctx context.Context) (int, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) CountX(ctx context.Context) int {
	count, err := fpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fpq *FavoriteProjectQuery) Exist(ctx context.Context) (bool, error) {
	if err := fpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fpq *FavoriteProjectQuery) ExistX(ctx context.Context) bool {
	exist, err := fpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FavoriteProjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fpq *FavoriteProjectQuery) Clone() *FavoriteProjectQuery {
	if fpq == nil {
		return nil
	}
	return &FavoriteProjectQuery{
		config:       fpq.config,
		limit:        fpq.limit,
		offset:       fpq.offset,
		order:        append([]OrderFunc{}, fpq.order...),
		predicates:   append([]predicate.FavoriteProject{}, fpq.predicates...),
		withProject:  fpq.withProject.Clone(),
		withTeammate: fpq.withTeammate.Clone(),
		// clone intermediate query.
		sql:    fpq.sql.Clone(),
		path:   fpq.path,
		unique: fpq.unique,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *FavoriteProjectQuery) WithProject(opts ...func(*ProjectQuery)) *FavoriteProjectQuery {
	query := &ProjectQuery{config: fpq.config}
	for _, opt := range opts {
		opt(query)
	}
	fpq.withProject = query
	return fpq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *FavoriteProjectQuery) WithTeammate(opts ...func(*TeammateQuery)) *FavoriteProjectQuery {
	query := &TeammateQuery{config: fpq.config}
	for _, opt := range opts {
		opt(query)
	}
	fpq.withTeammate = query
	return fpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProjectID ulid.ID `json:"project_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FavoriteProject.Query().
//		GroupBy(favoriteproject.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fpq *FavoriteProjectQuery) GroupBy(field string, fields ...string) *FavoriteProjectGroupBy {
	group := &FavoriteProjectGroupBy{config: fpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProjectID ulid.ID `json:"project_id,omitempty"`
//	}
//
//	client.FavoriteProject.Query().
//		Select(favoriteproject.FieldProjectID).
//		Scan(ctx, &v)
//
func (fpq *FavoriteProjectQuery) Select(fields ...string) *FavoriteProjectSelect {
	fpq.fields = append(fpq.fields, fields...)
	return &FavoriteProjectSelect{FavoriteProjectQuery: fpq}
}

func (fpq *FavoriteProjectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fpq.fields {
		if !favoriteproject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fpq.path != nil {
		prev, err := fpq.path(ctx)
		if err != nil {
			return err
		}
		fpq.sql = prev
	}
	return nil
}

func (fpq *FavoriteProjectQuery) sqlAll(ctx context.Context) ([]*FavoriteProject, error) {
	var (
		nodes       = []*FavoriteProject{}
		_spec       = fpq.querySpec()
		loadedTypes = [2]bool{
			fpq.withProject != nil,
			fpq.withTeammate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FavoriteProject{config: fpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fpq.withProject; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*FavoriteProject)
		for i := range nodes {
			fk := nodes[i].ProjectID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(project.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Project = n
			}
		}
	}

	if query := fpq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*FavoriteProject)
		for i := range nodes {
			fk := nodes[i].TeammateID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(teammate.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "teammate_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Teammate = n
			}
		}
	}

	return nodes, nil
}

func (fpq *FavoriteProjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fpq.querySpec()
	_spec.Node.Columns = fpq.fields
	if len(fpq.fields) > 0 {
		_spec.Unique = fpq.unique != nil && *fpq.unique
	}
	return sqlgraph.CountNodes(ctx, fpq.driver, _spec)
}

func (fpq *FavoriteProjectQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fpq *FavoriteProjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favoriteproject.Table,
			Columns: favoriteproject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteproject.FieldID,
			},
		},
		From:   fpq.sql,
		Unique: true,
	}
	if unique := fpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favoriteproject.FieldID)
		for i := range fields {
			if fields[i] != favoriteproject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fpq *FavoriteProjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fpq.driver.Dialect())
	t1 := builder.Table(favoriteproject.Table)
	columns := fpq.fields
	if len(columns) == 0 {
		columns = favoriteproject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fpq.sql != nil {
		selector = fpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fpq.unique != nil && *fpq.unique {
		selector.Distinct()
	}
	for _, p := range fpq.predicates {
		p(selector)
	}
	for _, p := range fpq.order {
		p(selector)
	}
	if offset := fpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FavoriteProjectGroupBy is the group-by builder for FavoriteProject entities.
type FavoriteProjectGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fpgb *FavoriteProjectGroupBy) Aggregate(fns ...AggregateFunc) *FavoriteProjectGroupBy {
	fpgb.fns = append(fpgb.fns, fns...)
	return fpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fpgb *FavoriteProjectGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fpgb.path(ctx)
	if err != nil {
		return err
	}
	fpgb.sql = query
	return fpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) StringsX(ctx context.Context) []string {
	v, err := fpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) StringX(ctx context.Context) string {
	v, err := fpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) IntsX(ctx context.Context) []int {
	v, err := fpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) IntX(ctx context.Context) int {
	v, err := fpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fpgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fpgb *FavoriteProjectGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fpgb *FavoriteProjectGroupBy) BoolX(ctx context.Context) bool {
	v, err := fpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fpgb *FavoriteProjectGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fpgb.fields {
		if !favoriteproject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fpgb *FavoriteProjectGroupBy) sqlQuery() *sql.Selector {
	selector := fpgb.sql.Select()
	aggregation := make([]string, 0, len(fpgb.fns))
	for _, fn := range fpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fpgb.fields)+len(fpgb.fns))
		for _, f := range fpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fpgb.fields...)...)
}

// FavoriteProjectSelect is the builder for selecting fields of FavoriteProject entities.
type FavoriteProjectSelect struct {
	*FavoriteProjectQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fps *FavoriteProjectSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fps.prepareQuery(ctx); err != nil {
		return err
	}
	fps.sql = fps.FavoriteProjectQuery.sqlQuery(ctx)
	return fps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fps *FavoriteProjectSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fps *FavoriteProjectSelect) StringsX(ctx context.Context) []string {
	v, err := fps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fps *FavoriteProjectSelect) StringX(ctx context.Context) string {
	v, err := fps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fps *FavoriteProjectSelect) IntsX(ctx context.Context) []int {
	v, err := fps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fps *FavoriteProjectSelect) IntX(ctx context.Context) int {
	v, err := fps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fps *FavoriteProjectSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fps *FavoriteProjectSelect) Float64X(ctx context.Context) float64 {
	v, err := fps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fps.fields) > 1 {
		return nil, errors.New("ent: FavoriteProjectSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fps *FavoriteProjectSelect) BoolsX(ctx context.Context) []bool {
	v, err := fps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fps *FavoriteProjectSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteproject.Label}
	default:
		err = fmt.Errorf("ent: FavoriteProjectSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fps *FavoriteProjectSelect) BoolX(ctx context.Context) bool {
	v, err := fps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fps *FavoriteProjectSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fps.sql.Query()
	if err := fps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
