// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/deletedteammatetask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeletedTeammateTaskQuery is the builder for querying DeletedTeammateTask entities.
type DeletedTeammateTaskQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeletedTeammateTask
	// eager-loading edges.
	withTeammate  *TeammateQuery
	withTask      *TaskQuery
	withWorkspace *WorkspaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeletedTeammateTaskQuery builder.
func (dttq *DeletedTeammateTaskQuery) Where(ps ...predicate.DeletedTeammateTask) *DeletedTeammateTaskQuery {
	dttq.predicates = append(dttq.predicates, ps...)
	return dttq
}

// Limit adds a limit step to the query.
func (dttq *DeletedTeammateTaskQuery) Limit(limit int) *DeletedTeammateTaskQuery {
	dttq.limit = &limit
	return dttq
}

// Offset adds an offset step to the query.
func (dttq *DeletedTeammateTaskQuery) Offset(offset int) *DeletedTeammateTaskQuery {
	dttq.offset = &offset
	return dttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dttq *DeletedTeammateTaskQuery) Unique(unique bool) *DeletedTeammateTaskQuery {
	dttq.unique = &unique
	return dttq
}

// Order adds an order step to the query.
func (dttq *DeletedTeammateTaskQuery) Order(o ...OrderFunc) *DeletedTeammateTaskQuery {
	dttq.order = append(dttq.order, o...)
	return dttq
}

// QueryTeammate chains the current query on the "teammate" edge.
func (dttq *DeletedTeammateTaskQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: dttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deletedteammatetask.Table, deletedteammatetask.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deletedteammatetask.TeammateTable, deletedteammatetask.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(dttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTask chains the current query on the "task" edge.
func (dttq *DeletedTeammateTaskQuery) QueryTask() *TaskQuery {
	query := &TaskQuery{config: dttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deletedteammatetask.Table, deletedteammatetask.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deletedteammatetask.TaskTable, deletedteammatetask.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(dttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkspace chains the current query on the "workspace" edge.
func (dttq *DeletedTeammateTaskQuery) QueryWorkspace() *WorkspaceQuery {
	query := &WorkspaceQuery{config: dttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deletedteammatetask.Table, deletedteammatetask.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deletedteammatetask.WorkspaceTable, deletedteammatetask.WorkspaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(dttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeletedTeammateTask entity from the query.
// Returns a *NotFoundError when no DeletedTeammateTask was found.
func (dttq *DeletedTeammateTaskQuery) First(ctx context.Context) (*DeletedTeammateTask, error) {
	nodes, err := dttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deletedteammatetask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) FirstX(ctx context.Context) *DeletedTeammateTask {
	node, err := dttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeletedTeammateTask ID from the query.
// Returns a *NotFoundError when no DeletedTeammateTask ID was found.
func (dttq *DeletedTeammateTaskQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = dttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deletedteammatetask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := dttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeletedTeammateTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeletedTeammateTask entity is found.
// Returns a *NotFoundError when no DeletedTeammateTask entities are found.
func (dttq *DeletedTeammateTaskQuery) Only(ctx context.Context) (*DeletedTeammateTask, error) {
	nodes, err := dttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deletedteammatetask.Label}
	default:
		return nil, &NotSingularError{deletedteammatetask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) OnlyX(ctx context.Context) *DeletedTeammateTask {
	node, err := dttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeletedTeammateTask ID in the query.
// Returns a *NotSingularError when more than one DeletedTeammateTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (dttq *DeletedTeammateTaskQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = dttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = &NotSingularError{deletedteammatetask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := dttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeletedTeammateTasks.
func (dttq *DeletedTeammateTaskQuery) All(ctx context.Context) ([]*DeletedTeammateTask, error) {
	if err := dttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) AllX(ctx context.Context) []*DeletedTeammateTask {
	nodes, err := dttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeletedTeammateTask IDs.
func (dttq *DeletedTeammateTaskQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := dttq.Select(deletedteammatetask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := dttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dttq *DeletedTeammateTaskQuery) Count(ctx context.Context) (int, error) {
	if err := dttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) CountX(ctx context.Context) int {
	count, err := dttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dttq *DeletedTeammateTaskQuery) Exist(ctx context.Context) (bool, error) {
	if err := dttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dttq *DeletedTeammateTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := dttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeletedTeammateTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dttq *DeletedTeammateTaskQuery) Clone() *DeletedTeammateTaskQuery {
	if dttq == nil {
		return nil
	}
	return &DeletedTeammateTaskQuery{
		config:        dttq.config,
		limit:         dttq.limit,
		offset:        dttq.offset,
		order:         append([]OrderFunc{}, dttq.order...),
		predicates:    append([]predicate.DeletedTeammateTask{}, dttq.predicates...),
		withTeammate:  dttq.withTeammate.Clone(),
		withTask:      dttq.withTask.Clone(),
		withWorkspace: dttq.withWorkspace.Clone(),
		// clone intermediate query.
		sql:    dttq.sql.Clone(),
		path:   dttq.path,
		unique: dttq.unique,
	}
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (dttq *DeletedTeammateTaskQuery) WithTeammate(opts ...func(*TeammateQuery)) *DeletedTeammateTaskQuery {
	query := &TeammateQuery{config: dttq.config}
	for _, opt := range opts {
		opt(query)
	}
	dttq.withTeammate = query
	return dttq
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (dttq *DeletedTeammateTaskQuery) WithTask(opts ...func(*TaskQuery)) *DeletedTeammateTaskQuery {
	query := &TaskQuery{config: dttq.config}
	for _, opt := range opts {
		opt(query)
	}
	dttq.withTask = query
	return dttq
}

// WithWorkspace tells the query-builder to eager-load the nodes that are connected to
// the "workspace" edge. The optional arguments are used to configure the query builder of the edge.
func (dttq *DeletedTeammateTaskQuery) WithWorkspace(opts ...func(*WorkspaceQuery)) *DeletedTeammateTaskQuery {
	query := &WorkspaceQuery{config: dttq.config}
	for _, opt := range opts {
		opt(query)
	}
	dttq.withWorkspace = query
	return dttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TeammateID ulid.ID `json:"teammate_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeletedTeammateTask.Query().
//		GroupBy(deletedteammatetask.FieldTeammateID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dttq *DeletedTeammateTaskQuery) GroupBy(field string, fields ...string) *DeletedTeammateTaskGroupBy {
	group := &DeletedTeammateTaskGroupBy{config: dttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dttq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TeammateID ulid.ID `json:"teammate_id,omitempty"`
//	}
//
//	client.DeletedTeammateTask.Query().
//		Select(deletedteammatetask.FieldTeammateID).
//		Scan(ctx, &v)
//
func (dttq *DeletedTeammateTaskQuery) Select(fields ...string) *DeletedTeammateTaskSelect {
	dttq.fields = append(dttq.fields, fields...)
	return &DeletedTeammateTaskSelect{DeletedTeammateTaskQuery: dttq}
}

func (dttq *DeletedTeammateTaskQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dttq.fields {
		if !deletedteammatetask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dttq.path != nil {
		prev, err := dttq.path(ctx)
		if err != nil {
			return err
		}
		dttq.sql = prev
	}
	return nil
}

func (dttq *DeletedTeammateTaskQuery) sqlAll(ctx context.Context) ([]*DeletedTeammateTask, error) {
	var (
		nodes       = []*DeletedTeammateTask{}
		_spec       = dttq.querySpec()
		loadedTypes = [3]bool{
			dttq.withTeammate != nil,
			dttq.withTask != nil,
			dttq.withWorkspace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeletedTeammateTask{config: dttq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dttq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*DeletedTeammateTask)
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

	if query := dttq.withTask; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*DeletedTeammateTask)
		for i := range nodes {
			fk := nodes[i].TaskID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(task.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Task = n
			}
		}
	}

	if query := dttq.withWorkspace; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*DeletedTeammateTask)
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

	return nodes, nil
}

func (dttq *DeletedTeammateTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dttq.querySpec()
	_spec.Node.Columns = dttq.fields
	if len(dttq.fields) > 0 {
		_spec.Unique = dttq.unique != nil && *dttq.unique
	}
	return sqlgraph.CountNodes(ctx, dttq.driver, _spec)
}

func (dttq *DeletedTeammateTaskQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dttq *DeletedTeammateTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedteammatetask.Table,
			Columns: deletedteammatetask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedteammatetask.FieldID,
			},
		},
		From:   dttq.sql,
		Unique: true,
	}
	if unique := dttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deletedteammatetask.FieldID)
		for i := range fields {
			if fields[i] != deletedteammatetask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dttq *DeletedTeammateTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dttq.driver.Dialect())
	t1 := builder.Table(deletedteammatetask.Table)
	columns := dttq.fields
	if len(columns) == 0 {
		columns = deletedteammatetask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dttq.sql != nil {
		selector = dttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dttq.unique != nil && *dttq.unique {
		selector.Distinct()
	}
	for _, p := range dttq.predicates {
		p(selector)
	}
	for _, p := range dttq.order {
		p(selector)
	}
	if offset := dttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeletedTeammateTaskGroupBy is the group-by builder for DeletedTeammateTask entities.
type DeletedTeammateTaskGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dttgb *DeletedTeammateTaskGroupBy) Aggregate(fns ...AggregateFunc) *DeletedTeammateTaskGroupBy {
	dttgb.fns = append(dttgb.fns, fns...)
	return dttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dttgb *DeletedTeammateTaskGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dttgb.path(ctx)
	if err != nil {
		return err
	}
	dttgb.sql = query
	return dttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dttgb.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) StringsX(ctx context.Context) []string {
	v, err := dttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) StringX(ctx context.Context) string {
	v, err := dttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dttgb.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) IntsX(ctx context.Context) []int {
	v, err := dttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) IntX(ctx context.Context) int {
	v, err := dttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dttgb.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dttgb.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dttgb *DeletedTeammateTaskGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dttgb *DeletedTeammateTaskGroupBy) BoolX(ctx context.Context) bool {
	v, err := dttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dttgb *DeletedTeammateTaskGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dttgb.fields {
		if !deletedteammatetask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dttgb *DeletedTeammateTaskGroupBy) sqlQuery() *sql.Selector {
	selector := dttgb.sql.Select()
	aggregation := make([]string, 0, len(dttgb.fns))
	for _, fn := range dttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dttgb.fields)+len(dttgb.fns))
		for _, f := range dttgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dttgb.fields...)...)
}

// DeletedTeammateTaskSelect is the builder for selecting fields of DeletedTeammateTask entities.
type DeletedTeammateTaskSelect struct {
	*DeletedTeammateTaskQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dtts *DeletedTeammateTaskSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dtts.prepareQuery(ctx); err != nil {
		return err
	}
	dtts.sql = dtts.DeletedTeammateTaskQuery.sqlQuery(ctx)
	return dtts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dtts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dtts.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) StringsX(ctx context.Context) []string {
	v, err := dtts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dtts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) StringX(ctx context.Context) string {
	v, err := dtts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dtts.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) IntsX(ctx context.Context) []int {
	v, err := dtts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dtts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) IntX(ctx context.Context) int {
	v, err := dtts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dtts.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dtts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dtts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) Float64X(ctx context.Context) float64 {
	v, err := dtts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dtts.fields) > 1 {
		return nil, errors.New("ent: DeletedTeammateTaskSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dtts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) BoolsX(ctx context.Context) []bool {
	v, err := dtts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dtts *DeletedTeammateTaskSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dtts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deletedteammatetask.Label}
	default:
		err = fmt.Errorf("ent: DeletedTeammateTaskSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dtts *DeletedTeammateTaskSelect) BoolX(ctx context.Context) bool {
	v, err := dtts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dtts *DeletedTeammateTaskSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dtts.sql.Query()
	if err := dtts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}