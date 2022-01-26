// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetasktabstatus"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskTabStatusQuery is the builder for querying TeammateTaskTabStatus entities.
type TeammateTaskTabStatusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TeammateTaskTabStatus
	// eager-loading edges.
	withWorkspace *WorkspaceQuery
	withTeammate  *TeammateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeammateTaskTabStatusQuery builder.
func (tttsq *TeammateTaskTabStatusQuery) Where(ps ...predicate.TeammateTaskTabStatus) *TeammateTaskTabStatusQuery {
	tttsq.predicates = append(tttsq.predicates, ps...)
	return tttsq
}

// Limit adds a limit step to the query.
func (tttsq *TeammateTaskTabStatusQuery) Limit(limit int) *TeammateTaskTabStatusQuery {
	tttsq.limit = &limit
	return tttsq
}

// Offset adds an offset step to the query.
func (tttsq *TeammateTaskTabStatusQuery) Offset(offset int) *TeammateTaskTabStatusQuery {
	tttsq.offset = &offset
	return tttsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tttsq *TeammateTaskTabStatusQuery) Unique(unique bool) *TeammateTaskTabStatusQuery {
	tttsq.unique = &unique
	return tttsq
}

// Order adds an order step to the query.
func (tttsq *TeammateTaskTabStatusQuery) Order(o ...OrderFunc) *TeammateTaskTabStatusQuery {
	tttsq.order = append(tttsq.order, o...)
	return tttsq
}

// QueryWorkspace chains the current query on the "workspace" edge.
func (tttsq *TeammateTaskTabStatusQuery) QueryWorkspace() *WorkspaceQuery {
	query := &WorkspaceQuery{config: tttsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tttsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tttsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(teammatetasktabstatus.Table, teammatetasktabstatus.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, teammatetasktabstatus.WorkspaceTable, teammatetasktabstatus.WorkspaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(tttsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (tttsq *TeammateTaskTabStatusQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: tttsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tttsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tttsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(teammatetasktabstatus.Table, teammatetasktabstatus.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, teammatetasktabstatus.TeammateTable, teammatetasktabstatus.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(tttsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TeammateTaskTabStatus entity from the query.
// Returns a *NotFoundError when no TeammateTaskTabStatus was found.
func (tttsq *TeammateTaskTabStatusQuery) First(ctx context.Context) (*TeammateTaskTabStatus, error) {
	nodes, err := tttsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{teammatetasktabstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) FirstX(ctx context.Context) *TeammateTaskTabStatus {
	node, err := tttsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TeammateTaskTabStatus ID from the query.
// Returns a *NotFoundError when no TeammateTaskTabStatus ID was found.
func (tttsq *TeammateTaskTabStatusQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tttsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{teammatetasktabstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tttsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TeammateTaskTabStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TeammateTaskTabStatus entity is not found.
// Returns a *NotFoundError when no TeammateTaskTabStatus entities are found.
func (tttsq *TeammateTaskTabStatusQuery) Only(ctx context.Context) (*TeammateTaskTabStatus, error) {
	nodes, err := tttsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{teammatetasktabstatus.Label}
	default:
		return nil, &NotSingularError{teammatetasktabstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) OnlyX(ctx context.Context) *TeammateTaskTabStatus {
	node, err := tttsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TeammateTaskTabStatus ID in the query.
// Returns a *NotSingularError when exactly one TeammateTaskTabStatus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tttsq *TeammateTaskTabStatusQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tttsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = &NotSingularError{teammatetasktabstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tttsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TeammateTaskTabStatusSlice.
func (tttsq *TeammateTaskTabStatusQuery) All(ctx context.Context) ([]*TeammateTaskTabStatus, error) {
	if err := tttsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tttsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) AllX(ctx context.Context) []*TeammateTaskTabStatus {
	nodes, err := tttsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TeammateTaskTabStatus IDs.
func (tttsq *TeammateTaskTabStatusQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tttsq.Select(teammatetasktabstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tttsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tttsq *TeammateTaskTabStatusQuery) Count(ctx context.Context) (int, error) {
	if err := tttsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tttsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) CountX(ctx context.Context) int {
	count, err := tttsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tttsq *TeammateTaskTabStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := tttsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tttsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tttsq *TeammateTaskTabStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := tttsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeammateTaskTabStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tttsq *TeammateTaskTabStatusQuery) Clone() *TeammateTaskTabStatusQuery {
	if tttsq == nil {
		return nil
	}
	return &TeammateTaskTabStatusQuery{
		config:        tttsq.config,
		limit:         tttsq.limit,
		offset:        tttsq.offset,
		order:         append([]OrderFunc{}, tttsq.order...),
		predicates:    append([]predicate.TeammateTaskTabStatus{}, tttsq.predicates...),
		withWorkspace: tttsq.withWorkspace.Clone(),
		withTeammate:  tttsq.withTeammate.Clone(),
		// clone intermediate query.
		sql:  tttsq.sql.Clone(),
		path: tttsq.path,
	}
}

// WithWorkspace tells the query-builder to eager-load the nodes that are connected to
// the "workspace" edge. The optional arguments are used to configure the query builder of the edge.
func (tttsq *TeammateTaskTabStatusQuery) WithWorkspace(opts ...func(*WorkspaceQuery)) *TeammateTaskTabStatusQuery {
	query := &WorkspaceQuery{config: tttsq.config}
	for _, opt := range opts {
		opt(query)
	}
	tttsq.withWorkspace = query
	return tttsq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (tttsq *TeammateTaskTabStatusQuery) WithTeammate(opts ...func(*TeammateQuery)) *TeammateTaskTabStatusQuery {
	query := &TeammateQuery{config: tttsq.config}
	for _, opt := range opts {
		opt(query)
	}
	tttsq.withTeammate = query
	return tttsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TeammateTaskTabStatus.Query().
//		GroupBy(teammatetasktabstatus.FieldWorkspaceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tttsq *TeammateTaskTabStatusQuery) GroupBy(field string, fields ...string) *TeammateTaskTabStatusGroupBy {
	group := &TeammateTaskTabStatusGroupBy{config: tttsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tttsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tttsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
//	}
//
//	client.TeammateTaskTabStatus.Query().
//		Select(teammatetasktabstatus.FieldWorkspaceID).
//		Scan(ctx, &v)
//
func (tttsq *TeammateTaskTabStatusQuery) Select(fields ...string) *TeammateTaskTabStatusSelect {
	tttsq.fields = append(tttsq.fields, fields...)
	return &TeammateTaskTabStatusSelect{TeammateTaskTabStatusQuery: tttsq}
}

func (tttsq *TeammateTaskTabStatusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tttsq.fields {
		if !teammatetasktabstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tttsq.path != nil {
		prev, err := tttsq.path(ctx)
		if err != nil {
			return err
		}
		tttsq.sql = prev
	}
	return nil
}

func (tttsq *TeammateTaskTabStatusQuery) sqlAll(ctx context.Context) ([]*TeammateTaskTabStatus, error) {
	var (
		nodes       = []*TeammateTaskTabStatus{}
		_spec       = tttsq.querySpec()
		loadedTypes = [2]bool{
			tttsq.withWorkspace != nil,
			tttsq.withTeammate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TeammateTaskTabStatus{config: tttsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tttsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tttsq.withWorkspace; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TeammateTaskTabStatus)
		for i := range nodes {
			fk := nodes[i].WorkspaceID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workspace.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workspace_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Workspace = n
			}
		}
	}

	if query := tttsq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TeammateTaskTabStatus)
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

func (tttsq *TeammateTaskTabStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tttsq.querySpec()
	return sqlgraph.CountNodes(ctx, tttsq.driver, _spec)
}

func (tttsq *TeammateTaskTabStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tttsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tttsq *TeammateTaskTabStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teammatetasktabstatus.Table,
			Columns: teammatetasktabstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasktabstatus.FieldID,
			},
		},
		From:   tttsq.sql,
		Unique: true,
	}
	if unique := tttsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tttsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teammatetasktabstatus.FieldID)
		for i := range fields {
			if fields[i] != teammatetasktabstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tttsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tttsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tttsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tttsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tttsq *TeammateTaskTabStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tttsq.driver.Dialect())
	t1 := builder.Table(teammatetasktabstatus.Table)
	columns := tttsq.fields
	if len(columns) == 0 {
		columns = teammatetasktabstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tttsq.sql != nil {
		selector = tttsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range tttsq.predicates {
		p(selector)
	}
	for _, p := range tttsq.order {
		p(selector)
	}
	if offset := tttsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tttsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TeammateTaskTabStatusGroupBy is the group-by builder for TeammateTaskTabStatus entities.
type TeammateTaskTabStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tttsgb *TeammateTaskTabStatusGroupBy) Aggregate(fns ...AggregateFunc) *TeammateTaskTabStatusGroupBy {
	tttsgb.fns = append(tttsgb.fns, fns...)
	return tttsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tttsgb *TeammateTaskTabStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tttsgb.path(ctx)
	if err != nil {
		return err
	}
	tttsgb.sql = query
	return tttsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tttsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tttsgb.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tttsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := tttsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tttsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) StringX(ctx context.Context) string {
	v, err := tttsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tttsgb.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tttsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := tttsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tttsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) IntX(ctx context.Context) int {
	v, err := tttsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tttsgb.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tttsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tttsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tttsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tttsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tttsgb.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tttsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tttsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tttsgb *TeammateTaskTabStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tttsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tttsgb *TeammateTaskTabStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := tttsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tttsgb *TeammateTaskTabStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tttsgb.fields {
		if !teammatetasktabstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tttsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tttsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tttsgb *TeammateTaskTabStatusGroupBy) sqlQuery() *sql.Selector {
	selector := tttsgb.sql.Select()
	aggregation := make([]string, 0, len(tttsgb.fns))
	for _, fn := range tttsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tttsgb.fields)+len(tttsgb.fns))
		for _, f := range tttsgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tttsgb.fields...)...)
}

// TeammateTaskTabStatusSelect is the builder for selecting fields of TeammateTaskTabStatus entities.
type TeammateTaskTabStatusSelect struct {
	*TeammateTaskTabStatusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tttss *TeammateTaskTabStatusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tttss.prepareQuery(ctx); err != nil {
		return err
	}
	tttss.sql = tttss.TeammateTaskTabStatusQuery.sqlQuery(ctx)
	return tttss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tttss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tttss.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tttss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) StringsX(ctx context.Context) []string {
	v, err := tttss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tttss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) StringX(ctx context.Context) string {
	v, err := tttss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tttss.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tttss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) IntsX(ctx context.Context) []int {
	v, err := tttss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tttss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) IntX(ctx context.Context) int {
	v, err := tttss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tttss.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tttss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tttss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tttss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := tttss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tttss.fields) > 1 {
		return nil, errors.New("ent: TeammateTaskTabStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tttss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := tttss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tttss *TeammateTaskTabStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tttss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{teammatetasktabstatus.Label}
	default:
		err = fmt.Errorf("ent: TeammateTaskTabStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tttss *TeammateTaskTabStatusSelect) BoolX(ctx context.Context) bool {
	v, err := tttss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tttss *TeammateTaskTabStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tttss.sql.Query()
	if err := tttss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
