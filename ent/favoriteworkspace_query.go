// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteWorkspaceQuery is the builder for querying FavoriteWorkspace entities.
type FavoriteWorkspaceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FavoriteWorkspace
	// eager-loading edges.
	withWorkspace *WorkspaceQuery
	withTeammate  *TeammateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FavoriteWorkspaceQuery builder.
func (fwq *FavoriteWorkspaceQuery) Where(ps ...predicate.FavoriteWorkspace) *FavoriteWorkspaceQuery {
	fwq.predicates = append(fwq.predicates, ps...)
	return fwq
}

// Limit adds a limit step to the query.
func (fwq *FavoriteWorkspaceQuery) Limit(limit int) *FavoriteWorkspaceQuery {
	fwq.limit = &limit
	return fwq
}

// Offset adds an offset step to the query.
func (fwq *FavoriteWorkspaceQuery) Offset(offset int) *FavoriteWorkspaceQuery {
	fwq.offset = &offset
	return fwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fwq *FavoriteWorkspaceQuery) Unique(unique bool) *FavoriteWorkspaceQuery {
	fwq.unique = &unique
	return fwq
}

// Order adds an order step to the query.
func (fwq *FavoriteWorkspaceQuery) Order(o ...OrderFunc) *FavoriteWorkspaceQuery {
	fwq.order = append(fwq.order, o...)
	return fwq
}

// QueryWorkspace chains the current query on the "workspace" edge.
func (fwq *FavoriteWorkspaceQuery) QueryWorkspace() *WorkspaceQuery {
	query := &WorkspaceQuery{config: fwq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fwq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(favoriteworkspace.Table, favoriteworkspace.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favoriteworkspace.WorkspaceTable, favoriteworkspace.WorkspaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fwq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (fwq *FavoriteWorkspaceQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: fwq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fwq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(favoriteworkspace.Table, favoriteworkspace.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favoriteworkspace.TeammateTable, favoriteworkspace.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(fwq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FavoriteWorkspace entity from the query.
// Returns a *NotFoundError when no FavoriteWorkspace was found.
func (fwq *FavoriteWorkspaceQuery) First(ctx context.Context) (*FavoriteWorkspace, error) {
	nodes, err := fwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{favoriteworkspace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) FirstX(ctx context.Context) *FavoriteWorkspace {
	node, err := fwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FavoriteWorkspace ID from the query.
// Returns a *NotFoundError when no FavoriteWorkspace ID was found.
func (fwq *FavoriteWorkspaceQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = fwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{favoriteworkspace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := fwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FavoriteWorkspace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FavoriteWorkspace entity is found.
// Returns a *NotFoundError when no FavoriteWorkspace entities are found.
func (fwq *FavoriteWorkspaceQuery) Only(ctx context.Context) (*FavoriteWorkspace, error) {
	nodes, err := fwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{favoriteworkspace.Label}
	default:
		return nil, &NotSingularError{favoriteworkspace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) OnlyX(ctx context.Context) *FavoriteWorkspace {
	node, err := fwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FavoriteWorkspace ID in the query.
// Returns a *NotSingularError when more than one FavoriteWorkspace ID is found.
// Returns a *NotFoundError when no entities are found.
func (fwq *FavoriteWorkspaceQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = fwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = &NotSingularError{favoriteworkspace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := fwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FavoriteWorkspaces.
func (fwq *FavoriteWorkspaceQuery) All(ctx context.Context) ([]*FavoriteWorkspace, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) AllX(ctx context.Context) []*FavoriteWorkspace {
	nodes, err := fwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FavoriteWorkspace IDs.
func (fwq *FavoriteWorkspaceQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := fwq.Select(favoriteworkspace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := fwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fwq *FavoriteWorkspaceQuery) Count(ctx context.Context) (int, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) CountX(ctx context.Context) int {
	count, err := fwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fwq *FavoriteWorkspaceQuery) Exist(ctx context.Context) (bool, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fwq *FavoriteWorkspaceQuery) ExistX(ctx context.Context) bool {
	exist, err := fwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FavoriteWorkspaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fwq *FavoriteWorkspaceQuery) Clone() *FavoriteWorkspaceQuery {
	if fwq == nil {
		return nil
	}
	return &FavoriteWorkspaceQuery{
		config:        fwq.config,
		limit:         fwq.limit,
		offset:        fwq.offset,
		order:         append([]OrderFunc{}, fwq.order...),
		predicates:    append([]predicate.FavoriteWorkspace{}, fwq.predicates...),
		withWorkspace: fwq.withWorkspace.Clone(),
		withTeammate:  fwq.withTeammate.Clone(),
		// clone intermediate query.
		sql:    fwq.sql.Clone(),
		path:   fwq.path,
		unique: fwq.unique,
	}
}

// WithWorkspace tells the query-builder to eager-load the nodes that are connected to
// the "workspace" edge. The optional arguments are used to configure the query builder of the edge.
func (fwq *FavoriteWorkspaceQuery) WithWorkspace(opts ...func(*WorkspaceQuery)) *FavoriteWorkspaceQuery {
	query := &WorkspaceQuery{config: fwq.config}
	for _, opt := range opts {
		opt(query)
	}
	fwq.withWorkspace = query
	return fwq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (fwq *FavoriteWorkspaceQuery) WithTeammate(opts ...func(*TeammateQuery)) *FavoriteWorkspaceQuery {
	query := &TeammateQuery{config: fwq.config}
	for _, opt := range opts {
		opt(query)
	}
	fwq.withTeammate = query
	return fwq
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
//	client.FavoriteWorkspace.Query().
//		GroupBy(favoriteworkspace.FieldWorkspaceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fwq *FavoriteWorkspaceQuery) GroupBy(field string, fields ...string) *FavoriteWorkspaceGroupBy {
	group := &FavoriteWorkspaceGroupBy{config: fwq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fwq.sqlQuery(ctx), nil
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
//	client.FavoriteWorkspace.Query().
//		Select(favoriteworkspace.FieldWorkspaceID).
//		Scan(ctx, &v)
//
func (fwq *FavoriteWorkspaceQuery) Select(fields ...string) *FavoriteWorkspaceSelect {
	fwq.fields = append(fwq.fields, fields...)
	return &FavoriteWorkspaceSelect{FavoriteWorkspaceQuery: fwq}
}

func (fwq *FavoriteWorkspaceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fwq.fields {
		if !favoriteworkspace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fwq.path != nil {
		prev, err := fwq.path(ctx)
		if err != nil {
			return err
		}
		fwq.sql = prev
	}
	return nil
}

func (fwq *FavoriteWorkspaceQuery) sqlAll(ctx context.Context) ([]*FavoriteWorkspace, error) {
	var (
		nodes       = []*FavoriteWorkspace{}
		_spec       = fwq.querySpec()
		loadedTypes = [2]bool{
			fwq.withWorkspace != nil,
			fwq.withTeammate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FavoriteWorkspace{config: fwq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fwq.withWorkspace; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*FavoriteWorkspace)
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

	if query := fwq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*FavoriteWorkspace)
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

func (fwq *FavoriteWorkspaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fwq.querySpec()
	_spec.Node.Columns = fwq.fields
	if len(fwq.fields) > 0 {
		_spec.Unique = fwq.unique != nil && *fwq.unique
	}
	return sqlgraph.CountNodes(ctx, fwq.driver, _spec)
}

func (fwq *FavoriteWorkspaceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fwq *FavoriteWorkspaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favoriteworkspace.Table,
			Columns: favoriteworkspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteworkspace.FieldID,
			},
		},
		From:   fwq.sql,
		Unique: true,
	}
	if unique := fwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favoriteworkspace.FieldID)
		for i := range fields {
			if fields[i] != favoriteworkspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fwq *FavoriteWorkspaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fwq.driver.Dialect())
	t1 := builder.Table(favoriteworkspace.Table)
	columns := fwq.fields
	if len(columns) == 0 {
		columns = favoriteworkspace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fwq.sql != nil {
		selector = fwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fwq.unique != nil && *fwq.unique {
		selector.Distinct()
	}
	for _, p := range fwq.predicates {
		p(selector)
	}
	for _, p := range fwq.order {
		p(selector)
	}
	if offset := fwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FavoriteWorkspaceGroupBy is the group-by builder for FavoriteWorkspace entities.
type FavoriteWorkspaceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fwgb *FavoriteWorkspaceGroupBy) Aggregate(fns ...AggregateFunc) *FavoriteWorkspaceGroupBy {
	fwgb.fns = append(fwgb.fns, fns...)
	return fwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fwgb *FavoriteWorkspaceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fwgb.path(ctx)
	if err != nil {
		return err
	}
	fwgb.sql = query
	return fwgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fwgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fwgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) StringsX(ctx context.Context) []string {
	v, err := fwgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fwgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) StringX(ctx context.Context) string {
	v, err := fwgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fwgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) IntsX(ctx context.Context) []int {
	v, err := fwgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fwgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) IntX(ctx context.Context) int {
	v, err := fwgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fwgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fwgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fwgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fwgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fwgb.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fwgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fwgb *FavoriteWorkspaceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fwgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fwgb *FavoriteWorkspaceGroupBy) BoolX(ctx context.Context) bool {
	v, err := fwgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fwgb *FavoriteWorkspaceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fwgb.fields {
		if !favoriteworkspace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fwgb *FavoriteWorkspaceGroupBy) sqlQuery() *sql.Selector {
	selector := fwgb.sql.Select()
	aggregation := make([]string, 0, len(fwgb.fns))
	for _, fn := range fwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fwgb.fields)+len(fwgb.fns))
		for _, f := range fwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fwgb.fields...)...)
}

// FavoriteWorkspaceSelect is the builder for selecting fields of FavoriteWorkspace entities.
type FavoriteWorkspaceSelect struct {
	*FavoriteWorkspaceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fws *FavoriteWorkspaceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fws.prepareQuery(ctx); err != nil {
		return err
	}
	fws.sql = fws.FavoriteWorkspaceQuery.sqlQuery(ctx)
	return fws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fws.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) StringsX(ctx context.Context) []string {
	v, err := fws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) StringX(ctx context.Context) string {
	v, err := fws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fws.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) IntsX(ctx context.Context) []int {
	v, err := fws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) IntX(ctx context.Context) int {
	v, err := fws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fws.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) Float64X(ctx context.Context) float64 {
	v, err := fws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fws.fields) > 1 {
		return nil, errors.New("ent: FavoriteWorkspaceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) BoolsX(ctx context.Context) []bool {
	v, err := fws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fws *FavoriteWorkspaceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{favoriteworkspace.Label}
	default:
		err = fmt.Errorf("ent: FavoriteWorkspaceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fws *FavoriteWorkspaceSelect) BoolX(ctx context.Context) bool {
	v, err := fws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fws *FavoriteWorkspaceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fws.sql.Query()
	if err := fws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
