// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTeammateQuery is the builder for querying ProjectTeammate entities.
type ProjectTeammateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProjectTeammate
	// eager-loading edges.
	withProject  *ProjectQuery
	withTeammate *TeammateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectTeammateQuery builder.
func (ptq *ProjectTeammateQuery) Where(ps ...predicate.ProjectTeammate) *ProjectTeammateQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit adds a limit step to the query.
func (ptq *ProjectTeammateQuery) Limit(limit int) *ProjectTeammateQuery {
	ptq.limit = &limit
	return ptq
}

// Offset adds an offset step to the query.
func (ptq *ProjectTeammateQuery) Offset(offset int) *ProjectTeammateQuery {
	ptq.offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *ProjectTeammateQuery) Unique(unique bool) *ProjectTeammateQuery {
	ptq.unique = &unique
	return ptq
}

// Order adds an order step to the query.
func (ptq *ProjectTeammateQuery) Order(o ...OrderFunc) *ProjectTeammateQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryProject chains the current query on the "project" edge.
func (ptq *ProjectTeammateQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: ptq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectteammate.Table, projectteammate.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectteammate.ProjectTable, projectteammate.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (ptq *ProjectTeammateQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: ptq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectteammate.Table, projectteammate.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectteammate.TeammateTable, projectteammate.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectTeammate entity from the query.
// Returns a *NotFoundError when no ProjectTeammate was found.
func (ptq *ProjectTeammateQuery) First(ctx context.Context) (*ProjectTeammate, error) {
	nodes, err := ptq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectteammate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) FirstX(ctx context.Context) *ProjectTeammate {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectTeammate ID from the query.
// Returns a *NotFoundError when no ProjectTeammate ID was found.
func (ptq *ProjectTeammateQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projectteammate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectTeammate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectTeammate entity is found.
// Returns a *NotFoundError when no ProjectTeammate entities are found.
func (ptq *ProjectTeammateQuery) Only(ctx context.Context) (*ProjectTeammate, error) {
	nodes, err := ptq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectteammate.Label}
	default:
		return nil, &NotSingularError{projectteammate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) OnlyX(ctx context.Context) *ProjectTeammate {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectTeammate ID in the query.
// Returns a *NotSingularError when more than one ProjectTeammate ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *ProjectTeammateQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = &NotSingularError{projectteammate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectTeammates.
func (ptq *ProjectTeammateQuery) All(ctx context.Context) ([]*ProjectTeammate, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) AllX(ctx context.Context) []*ProjectTeammate {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectTeammate IDs.
func (ptq *ProjectTeammateQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := ptq.Select(projectteammate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *ProjectTeammateQuery) Count(ctx context.Context) (int, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *ProjectTeammateQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *ProjectTeammateQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectTeammateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *ProjectTeammateQuery) Clone() *ProjectTeammateQuery {
	if ptq == nil {
		return nil
	}
	return &ProjectTeammateQuery{
		config:       ptq.config,
		limit:        ptq.limit,
		offset:       ptq.offset,
		order:        append([]OrderFunc{}, ptq.order...),
		predicates:   append([]predicate.ProjectTeammate{}, ptq.predicates...),
		withProject:  ptq.withProject.Clone(),
		withTeammate: ptq.withTeammate.Clone(),
		// clone intermediate query.
		sql:    ptq.sql.Clone(),
		path:   ptq.path,
		unique: ptq.unique,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTeammateQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectTeammateQuery {
	query := &ProjectQuery{config: ptq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptq.withProject = query
	return ptq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *ProjectTeammateQuery) WithTeammate(opts ...func(*TeammateQuery)) *ProjectTeammateQuery {
	query := &TeammateQuery{config: ptq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptq.withTeammate = query
	return ptq
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
//	client.ProjectTeammate.Query().
//		GroupBy(projectteammate.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptq *ProjectTeammateQuery) GroupBy(field string, fields ...string) *ProjectTeammateGroupBy {
	group := &ProjectTeammateGroupBy{config: ptq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptq.sqlQuery(ctx), nil
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
//	client.ProjectTeammate.Query().
//		Select(projectteammate.FieldProjectID).
//		Scan(ctx, &v)
//
func (ptq *ProjectTeammateQuery) Select(fields ...string) *ProjectTeammateSelect {
	ptq.fields = append(ptq.fields, fields...)
	return &ProjectTeammateSelect{ProjectTeammateQuery: ptq}
}

func (ptq *ProjectTeammateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptq.fields {
		if !projectteammate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *ProjectTeammateQuery) sqlAll(ctx context.Context) ([]*ProjectTeammate, error) {
	var (
		nodes       = []*ProjectTeammate{}
		_spec       = ptq.querySpec()
		loadedTypes = [2]bool{
			ptq.withProject != nil,
			ptq.withTeammate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProjectTeammate{config: ptq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptq.withProject; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTeammate)
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

	if query := ptq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTeammate)
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

func (ptq *ProjectTeammateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.fields
	if len(ptq.fields) > 0 {
		_spec.Unique = ptq.unique != nil && *ptq.unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *ProjectTeammateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ptq *ProjectTeammateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectteammate.Table,
			Columns: projectteammate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projectteammate.FieldID,
			},
		},
		From:   ptq.sql,
		Unique: true,
	}
	if unique := ptq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ptq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectteammate.FieldID)
		for i := range fields {
			if fields[i] != projectteammate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *ProjectTeammateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(projectteammate.Table)
	columns := ptq.fields
	if len(columns) == 0 {
		columns = projectteammate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.unique != nil && *ptq.unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectTeammateGroupBy is the group-by builder for ProjectTeammate entities.
type ProjectTeammateGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *ProjectTeammateGroupBy) Aggregate(fns ...AggregateFunc) *ProjectTeammateGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptgb *ProjectTeammateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptgb.path(ctx)
	if err != nil {
		return err
	}
	ptgb.sql = query
	return ptgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) StringX(ctx context.Context) string {
	v, err := ptgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) IntX(ctx context.Context) int {
	v, err := ptgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *ProjectTeammateGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptgb *ProjectTeammateGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptgb *ProjectTeammateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptgb.fields {
		if !projectteammate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptgb *ProjectTeammateGroupBy) sqlQuery() *sql.Selector {
	selector := ptgb.sql.Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ptgb.fields)+len(ptgb.fns))
		for _, f := range ptgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ptgb.fields...)...)
}

// ProjectTeammateSelect is the builder for selecting fields of ProjectTeammate entities.
type ProjectTeammateSelect struct {
	*ProjectTeammateQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pts *ProjectTeammateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	pts.sql = pts.ProjectTeammateQuery.sqlQuery(ctx)
	return pts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pts *ProjectTeammateSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pts *ProjectTeammateSelect) StringsX(ctx context.Context) []string {
	v, err := pts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pts *ProjectTeammateSelect) StringX(ctx context.Context) string {
	v, err := pts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pts *ProjectTeammateSelect) IntsX(ctx context.Context) []int {
	v, err := pts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pts *ProjectTeammateSelect) IntX(ctx context.Context) int {
	v, err := pts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pts *ProjectTeammateSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pts *ProjectTeammateSelect) Float64X(ctx context.Context) float64 {
	v, err := pts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: ProjectTeammateSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pts *ProjectTeammateSelect) BoolsX(ctx context.Context) []bool {
	v, err := pts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pts *ProjectTeammateSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectteammate.Label}
	default:
		err = fmt.Errorf("ent: ProjectTeammateSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pts *ProjectTeammateSelect) BoolX(ctx context.Context) bool {
	v, err := pts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pts *ProjectTeammateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pts.sql.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
