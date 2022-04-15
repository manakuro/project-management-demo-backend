// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/archivedtaskactivitytask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskactivity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedTaskActivityTaskQuery is the builder for querying ArchivedTaskActivityTask entities.
type ArchivedTaskActivityTaskQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ArchivedTaskActivityTask
	// eager-loading edges.
	withTask         *TaskQuery
	withTaskActivity *TaskActivityQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArchivedTaskActivityTaskQuery builder.
func (atatq *ArchivedTaskActivityTaskQuery) Where(ps ...predicate.ArchivedTaskActivityTask) *ArchivedTaskActivityTaskQuery {
	atatq.predicates = append(atatq.predicates, ps...)
	return atatq
}

// Limit adds a limit step to the query.
func (atatq *ArchivedTaskActivityTaskQuery) Limit(limit int) *ArchivedTaskActivityTaskQuery {
	atatq.limit = &limit
	return atatq
}

// Offset adds an offset step to the query.
func (atatq *ArchivedTaskActivityTaskQuery) Offset(offset int) *ArchivedTaskActivityTaskQuery {
	atatq.offset = &offset
	return atatq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atatq *ArchivedTaskActivityTaskQuery) Unique(unique bool) *ArchivedTaskActivityTaskQuery {
	atatq.unique = &unique
	return atatq
}

// Order adds an order step to the query.
func (atatq *ArchivedTaskActivityTaskQuery) Order(o ...OrderFunc) *ArchivedTaskActivityTaskQuery {
	atatq.order = append(atatq.order, o...)
	return atatq
}

// QueryTask chains the current query on the "task" edge.
func (atatq *ArchivedTaskActivityTaskQuery) QueryTask() *TaskQuery {
	query := &TaskQuery{config: atatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivitytask.Table, archivedtaskactivitytask.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, archivedtaskactivitytask.TaskTable, archivedtaskactivitytask.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(atatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskActivity chains the current query on the "taskActivity" edge.
func (atatq *ArchivedTaskActivityTaskQuery) QueryTaskActivity() *TaskActivityQuery {
	query := &TaskActivityQuery{config: atatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivitytask.Table, archivedtaskactivitytask.FieldID, selector),
			sqlgraph.To(taskactivity.Table, taskactivity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, archivedtaskactivitytask.TaskActivityTable, archivedtaskactivitytask.TaskActivityColumn),
		)
		fromU = sqlgraph.SetNeighbors(atatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArchivedTaskActivityTask entity from the query.
// Returns a *NotFoundError when no ArchivedTaskActivityTask was found.
func (atatq *ArchivedTaskActivityTaskQuery) First(ctx context.Context) (*ArchivedTaskActivityTask, error) {
	nodes, err := atatq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{archivedtaskactivitytask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) FirstX(ctx context.Context) *ArchivedTaskActivityTask {
	node, err := atatq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArchivedTaskActivityTask ID from the query.
// Returns a *NotFoundError when no ArchivedTaskActivityTask ID was found.
func (atatq *ArchivedTaskActivityTaskQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = atatq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{archivedtaskactivitytask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := atatq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArchivedTaskActivityTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArchivedTaskActivityTask entity is found.
// Returns a *NotFoundError when no ArchivedTaskActivityTask entities are found.
func (atatq *ArchivedTaskActivityTaskQuery) Only(ctx context.Context) (*ArchivedTaskActivityTask, error) {
	nodes, err := atatq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{archivedtaskactivitytask.Label}
	default:
		return nil, &NotSingularError{archivedtaskactivitytask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) OnlyX(ctx context.Context) *ArchivedTaskActivityTask {
	node, err := atatq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArchivedTaskActivityTask ID in the query.
// Returns a *NotSingularError when more than one ArchivedTaskActivityTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (atatq *ArchivedTaskActivityTaskQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = atatq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = &NotSingularError{archivedtaskactivitytask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := atatq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArchivedTaskActivityTasks.
func (atatq *ArchivedTaskActivityTaskQuery) All(ctx context.Context) ([]*ArchivedTaskActivityTask, error) {
	if err := atatq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return atatq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) AllX(ctx context.Context) []*ArchivedTaskActivityTask {
	nodes, err := atatq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArchivedTaskActivityTask IDs.
func (atatq *ArchivedTaskActivityTaskQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := atatq.Select(archivedtaskactivitytask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := atatq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atatq *ArchivedTaskActivityTaskQuery) Count(ctx context.Context) (int, error) {
	if err := atatq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return atatq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) CountX(ctx context.Context) int {
	count, err := atatq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atatq *ArchivedTaskActivityTaskQuery) Exist(ctx context.Context) (bool, error) {
	if err := atatq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return atatq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (atatq *ArchivedTaskActivityTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := atatq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArchivedTaskActivityTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atatq *ArchivedTaskActivityTaskQuery) Clone() *ArchivedTaskActivityTaskQuery {
	if atatq == nil {
		return nil
	}
	return &ArchivedTaskActivityTaskQuery{
		config:           atatq.config,
		limit:            atatq.limit,
		offset:           atatq.offset,
		order:            append([]OrderFunc{}, atatq.order...),
		predicates:       append([]predicate.ArchivedTaskActivityTask{}, atatq.predicates...),
		withTask:         atatq.withTask.Clone(),
		withTaskActivity: atatq.withTaskActivity.Clone(),
		// clone intermediate query.
		sql:    atatq.sql.Clone(),
		path:   atatq.path,
		unique: atatq.unique,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (atatq *ArchivedTaskActivityTaskQuery) WithTask(opts ...func(*TaskQuery)) *ArchivedTaskActivityTaskQuery {
	query := &TaskQuery{config: atatq.config}
	for _, opt := range opts {
		opt(query)
	}
	atatq.withTask = query
	return atatq
}

// WithTaskActivity tells the query-builder to eager-load the nodes that are connected to
// the "taskActivity" edge. The optional arguments are used to configure the query builder of the edge.
func (atatq *ArchivedTaskActivityTaskQuery) WithTaskActivity(opts ...func(*TaskActivityQuery)) *ArchivedTaskActivityTaskQuery {
	query := &TaskActivityQuery{config: atatq.config}
	for _, opt := range opts {
		opt(query)
	}
	atatq.withTaskActivity = query
	return atatq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TaskActivityID ulid.ID `json:"task_activity_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArchivedTaskActivityTask.Query().
//		GroupBy(archivedtaskactivitytask.FieldTaskActivityID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (atatq *ArchivedTaskActivityTaskQuery) GroupBy(field string, fields ...string) *ArchivedTaskActivityTaskGroupBy {
	group := &ArchivedTaskActivityTaskGroupBy{config: atatq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := atatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return atatq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TaskActivityID ulid.ID `json:"task_activity_id,omitempty"`
//	}
//
//	client.ArchivedTaskActivityTask.Query().
//		Select(archivedtaskactivitytask.FieldTaskActivityID).
//		Scan(ctx, &v)
//
func (atatq *ArchivedTaskActivityTaskQuery) Select(fields ...string) *ArchivedTaskActivityTaskSelect {
	atatq.fields = append(atatq.fields, fields...)
	return &ArchivedTaskActivityTaskSelect{ArchivedTaskActivityTaskQuery: atatq}
}

func (atatq *ArchivedTaskActivityTaskQuery) prepareQuery(ctx context.Context) error {
	for _, f := range atatq.fields {
		if !archivedtaskactivitytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atatq.path != nil {
		prev, err := atatq.path(ctx)
		if err != nil {
			return err
		}
		atatq.sql = prev
	}
	return nil
}

func (atatq *ArchivedTaskActivityTaskQuery) sqlAll(ctx context.Context) ([]*ArchivedTaskActivityTask, error) {
	var (
		nodes       = []*ArchivedTaskActivityTask{}
		_spec       = atatq.querySpec()
		loadedTypes = [2]bool{
			atatq.withTask != nil,
			atatq.withTaskActivity != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ArchivedTaskActivityTask{config: atatq.config}
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
	if err := sqlgraph.QueryNodes(ctx, atatq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := atatq.withTask; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ArchivedTaskActivityTask)
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

	if query := atatq.withTaskActivity; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ArchivedTaskActivityTask)
		for i := range nodes {
			fk := nodes[i].TaskActivityID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(taskactivity.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_activity_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TaskActivity = n
			}
		}
	}

	return nodes, nil
}

func (atatq *ArchivedTaskActivityTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atatq.querySpec()
	_spec.Node.Columns = atatq.fields
	if len(atatq.fields) > 0 {
		_spec.Unique = atatq.unique != nil && *atatq.unique
	}
	return sqlgraph.CountNodes(ctx, atatq.driver, _spec)
}

func (atatq *ArchivedTaskActivityTaskQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := atatq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (atatq *ArchivedTaskActivityTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedtaskactivitytask.Table,
			Columns: archivedtaskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedtaskactivitytask.FieldID,
			},
		},
		From:   atatq.sql,
		Unique: true,
	}
	if unique := atatq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := atatq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, archivedtaskactivitytask.FieldID)
		for i := range fields {
			if fields[i] != archivedtaskactivitytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atatq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atatq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atatq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atatq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atatq *ArchivedTaskActivityTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atatq.driver.Dialect())
	t1 := builder.Table(archivedtaskactivitytask.Table)
	columns := atatq.fields
	if len(columns) == 0 {
		columns = archivedtaskactivitytask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atatq.sql != nil {
		selector = atatq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atatq.unique != nil && *atatq.unique {
		selector.Distinct()
	}
	for _, p := range atatq.predicates {
		p(selector)
	}
	for _, p := range atatq.order {
		p(selector)
	}
	if offset := atatq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atatq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArchivedTaskActivityTaskGroupBy is the group-by builder for ArchivedTaskActivityTask entities.
type ArchivedTaskActivityTaskGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Aggregate(fns ...AggregateFunc) *ArchivedTaskActivityTaskGroupBy {
	atatgb.fns = append(atatgb.fns, fns...)
	return atatgb
}

// Scan applies the group-by query and scans the result into the given value.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := atatgb.path(ctx)
	if err != nil {
		return err
	}
	atatgb.sql = query
	return atatgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := atatgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(atatgb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := atatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) StringsX(ctx context.Context) []string {
	v, err := atatgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atatgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) StringX(ctx context.Context) string {
	v, err := atatgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(atatgb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := atatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) IntsX(ctx context.Context) []int {
	v, err := atatgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atatgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) IntX(ctx context.Context) int {
	v, err := atatgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(atatgb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := atatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := atatgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atatgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Float64X(ctx context.Context) float64 {
	v, err := atatgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(atatgb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := atatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := atatgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atatgb *ArchivedTaskActivityTaskGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atatgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atatgb *ArchivedTaskActivityTaskGroupBy) BoolX(ctx context.Context) bool {
	v, err := atatgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atatgb *ArchivedTaskActivityTaskGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range atatgb.fields {
		if !archivedtaskactivitytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atatgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atatgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atatgb *ArchivedTaskActivityTaskGroupBy) sqlQuery() *sql.Selector {
	selector := atatgb.sql.Select()
	aggregation := make([]string, 0, len(atatgb.fns))
	for _, fn := range atatgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(atatgb.fields)+len(atatgb.fns))
		for _, f := range atatgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(atatgb.fields...)...)
}

// ArchivedTaskActivityTaskSelect is the builder for selecting fields of ArchivedTaskActivityTask entities.
type ArchivedTaskActivityTaskSelect struct {
	*ArchivedTaskActivityTaskQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (atats *ArchivedTaskActivityTaskSelect) Scan(ctx context.Context, v interface{}) error {
	if err := atats.prepareQuery(ctx); err != nil {
		return err
	}
	atats.sql = atats.ArchivedTaskActivityTaskQuery.sqlQuery(ctx)
	return atats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) ScanX(ctx context.Context, v interface{}) {
	if err := atats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Strings(ctx context.Context) ([]string, error) {
	if len(atats.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := atats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) StringsX(ctx context.Context) []string {
	v, err := atats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) StringX(ctx context.Context) string {
	v, err := atats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Ints(ctx context.Context) ([]int, error) {
	if len(atats.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := atats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) IntsX(ctx context.Context) []int {
	v, err := atats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) IntX(ctx context.Context) int {
	v, err := atats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(atats.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := atats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) Float64sX(ctx context.Context) []float64 {
	v, err := atats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) Float64X(ctx context.Context) float64 {
	v, err := atats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(atats.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityTaskSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := atats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) BoolsX(ctx context.Context) []bool {
	v, err := atats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (atats *ArchivedTaskActivityTaskSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityTaskSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atats *ArchivedTaskActivityTaskSelect) BoolX(ctx context.Context) bool {
	v, err := atats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atats *ArchivedTaskActivityTaskSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := atats.sql.Query()
	if err := atats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
