// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskactivity"
	"project-management-demo-backend/ent/taskactivitytask"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskActivityTaskQuery is the builder for querying TaskActivityTask entities.
type TaskActivityTaskQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TaskActivityTask
	// eager-loading edges.
	withTask         *TaskQuery
	withTaskActivity *TaskActivityQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskActivityTaskQuery builder.
func (tatq *TaskActivityTaskQuery) Where(ps ...predicate.TaskActivityTask) *TaskActivityTaskQuery {
	tatq.predicates = append(tatq.predicates, ps...)
	return tatq
}

// Limit adds a limit step to the query.
func (tatq *TaskActivityTaskQuery) Limit(limit int) *TaskActivityTaskQuery {
	tatq.limit = &limit
	return tatq
}

// Offset adds an offset step to the query.
func (tatq *TaskActivityTaskQuery) Offset(offset int) *TaskActivityTaskQuery {
	tatq.offset = &offset
	return tatq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tatq *TaskActivityTaskQuery) Unique(unique bool) *TaskActivityTaskQuery {
	tatq.unique = &unique
	return tatq
}

// Order adds an order step to the query.
func (tatq *TaskActivityTaskQuery) Order(o ...OrderFunc) *TaskActivityTaskQuery {
	tatq.order = append(tatq.order, o...)
	return tatq
}

// QueryTask chains the current query on the "task" edge.
func (tatq *TaskActivityTaskQuery) QueryTask() *TaskQuery {
	query := &TaskQuery{config: tatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskactivitytask.Table, taskactivitytask.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskactivitytask.TaskTable, taskactivitytask.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(tatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskActivity chains the current query on the "taskActivity" edge.
func (tatq *TaskActivityTaskQuery) QueryTaskActivity() *TaskActivityQuery {
	query := &TaskActivityQuery{config: tatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskactivitytask.Table, taskactivitytask.FieldID, selector),
			sqlgraph.To(taskactivity.Table, taskactivity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskactivitytask.TaskActivityTable, taskactivitytask.TaskActivityColumn),
		)
		fromU = sqlgraph.SetNeighbors(tatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskActivityTask entity from the query.
// Returns a *NotFoundError when no TaskActivityTask was found.
func (tatq *TaskActivityTaskQuery) First(ctx context.Context) (*TaskActivityTask, error) {
	nodes, err := tatq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taskactivitytask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) FirstX(ctx context.Context) *TaskActivityTask {
	node, err := tatq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskActivityTask ID from the query.
// Returns a *NotFoundError when no TaskActivityTask ID was found.
func (tatq *TaskActivityTaskQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tatq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taskactivitytask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tatq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskActivityTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskActivityTask entity is found.
// Returns a *NotFoundError when no TaskActivityTask entities are found.
func (tatq *TaskActivityTaskQuery) Only(ctx context.Context) (*TaskActivityTask, error) {
	nodes, err := tatq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taskactivitytask.Label}
	default:
		return nil, &NotSingularError{taskactivitytask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) OnlyX(ctx context.Context) *TaskActivityTask {
	node, err := tatq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskActivityTask ID in the query.
// Returns a *NotSingularError when more than one TaskActivityTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (tatq *TaskActivityTaskQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tatq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = &NotSingularError{taskactivitytask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tatq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskActivityTasks.
func (tatq *TaskActivityTaskQuery) All(ctx context.Context) ([]*TaskActivityTask, error) {
	if err := tatq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tatq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) AllX(ctx context.Context) []*TaskActivityTask {
	nodes, err := tatq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskActivityTask IDs.
func (tatq *TaskActivityTaskQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tatq.Select(taskactivitytask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tatq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tatq *TaskActivityTaskQuery) Count(ctx context.Context) (int, error) {
	if err := tatq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tatq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) CountX(ctx context.Context) int {
	count, err := tatq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tatq *TaskActivityTaskQuery) Exist(ctx context.Context) (bool, error) {
	if err := tatq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tatq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tatq *TaskActivityTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := tatq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskActivityTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tatq *TaskActivityTaskQuery) Clone() *TaskActivityTaskQuery {
	if tatq == nil {
		return nil
	}
	return &TaskActivityTaskQuery{
		config:           tatq.config,
		limit:            tatq.limit,
		offset:           tatq.offset,
		order:            append([]OrderFunc{}, tatq.order...),
		predicates:       append([]predicate.TaskActivityTask{}, tatq.predicates...),
		withTask:         tatq.withTask.Clone(),
		withTaskActivity: tatq.withTaskActivity.Clone(),
		// clone intermediate query.
		sql:    tatq.sql.Clone(),
		path:   tatq.path,
		unique: tatq.unique,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (tatq *TaskActivityTaskQuery) WithTask(opts ...func(*TaskQuery)) *TaskActivityTaskQuery {
	query := &TaskQuery{config: tatq.config}
	for _, opt := range opts {
		opt(query)
	}
	tatq.withTask = query
	return tatq
}

// WithTaskActivity tells the query-builder to eager-load the nodes that are connected to
// the "taskActivity" edge. The optional arguments are used to configure the query builder of the edge.
func (tatq *TaskActivityTaskQuery) WithTaskActivity(opts ...func(*TaskActivityQuery)) *TaskActivityTaskQuery {
	query := &TaskActivityQuery{config: tatq.config}
	for _, opt := range opts {
		opt(query)
	}
	tatq.withTaskActivity = query
	return tatq
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
//	client.TaskActivityTask.Query().
//		GroupBy(taskactivitytask.FieldTaskActivityID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tatq *TaskActivityTaskQuery) GroupBy(field string, fields ...string) *TaskActivityTaskGroupBy {
	group := &TaskActivityTaskGroupBy{config: tatq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tatq.sqlQuery(ctx), nil
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
//	client.TaskActivityTask.Query().
//		Select(taskactivitytask.FieldTaskActivityID).
//		Scan(ctx, &v)
//
func (tatq *TaskActivityTaskQuery) Select(fields ...string) *TaskActivityTaskSelect {
	tatq.fields = append(tatq.fields, fields...)
	return &TaskActivityTaskSelect{TaskActivityTaskQuery: tatq}
}

func (tatq *TaskActivityTaskQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tatq.fields {
		if !taskactivitytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tatq.path != nil {
		prev, err := tatq.path(ctx)
		if err != nil {
			return err
		}
		tatq.sql = prev
	}
	return nil
}

func (tatq *TaskActivityTaskQuery) sqlAll(ctx context.Context) ([]*TaskActivityTask, error) {
	var (
		nodes       = []*TaskActivityTask{}
		_spec       = tatq.querySpec()
		loadedTypes = [2]bool{
			tatq.withTask != nil,
			tatq.withTaskActivity != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TaskActivityTask{config: tatq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tatq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tatq.withTask; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskActivityTask)
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

	if query := tatq.withTaskActivity; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskActivityTask)
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

func (tatq *TaskActivityTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tatq.querySpec()
	_spec.Node.Columns = tatq.fields
	if len(tatq.fields) > 0 {
		_spec.Unique = tatq.unique != nil && *tatq.unique
	}
	return sqlgraph.CountNodes(ctx, tatq.driver, _spec)
}

func (tatq *TaskActivityTaskQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tatq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tatq *TaskActivityTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskactivitytask.Table,
			Columns: taskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskactivitytask.FieldID,
			},
		},
		From:   tatq.sql,
		Unique: true,
	}
	if unique := tatq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tatq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskactivitytask.FieldID)
		for i := range fields {
			if fields[i] != taskactivitytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tatq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tatq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tatq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tatq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tatq *TaskActivityTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tatq.driver.Dialect())
	t1 := builder.Table(taskactivitytask.Table)
	columns := tatq.fields
	if len(columns) == 0 {
		columns = taskactivitytask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tatq.sql != nil {
		selector = tatq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tatq.unique != nil && *tatq.unique {
		selector.Distinct()
	}
	for _, p := range tatq.predicates {
		p(selector)
	}
	for _, p := range tatq.order {
		p(selector)
	}
	if offset := tatq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tatq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskActivityTaskGroupBy is the group-by builder for TaskActivityTask entities.
type TaskActivityTaskGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tatgb *TaskActivityTaskGroupBy) Aggregate(fns ...AggregateFunc) *TaskActivityTaskGroupBy {
	tatgb.fns = append(tatgb.fns, fns...)
	return tatgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tatgb *TaskActivityTaskGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tatgb.path(ctx)
	if err != nil {
		return err
	}
	tatgb.sql = query
	return tatgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tatgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tatgb.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) StringsX(ctx context.Context) []string {
	v, err := tatgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tatgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) StringX(ctx context.Context) string {
	v, err := tatgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tatgb.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) IntsX(ctx context.Context) []int {
	v, err := tatgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tatgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) IntX(ctx context.Context) int {
	v, err := tatgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tatgb.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tatgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tatgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tatgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tatgb.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tatgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tatgb *TaskActivityTaskGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tatgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tatgb *TaskActivityTaskGroupBy) BoolX(ctx context.Context) bool {
	v, err := tatgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tatgb *TaskActivityTaskGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tatgb.fields {
		if !taskactivitytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tatgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tatgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tatgb *TaskActivityTaskGroupBy) sqlQuery() *sql.Selector {
	selector := tatgb.sql.Select()
	aggregation := make([]string, 0, len(tatgb.fns))
	for _, fn := range tatgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tatgb.fields)+len(tatgb.fns))
		for _, f := range tatgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tatgb.fields...)...)
}

// TaskActivityTaskSelect is the builder for selecting fields of TaskActivityTask entities.
type TaskActivityTaskSelect struct {
	*TaskActivityTaskQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tats *TaskActivityTaskSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tats.prepareQuery(ctx); err != nil {
		return err
	}
	tats.sql = tats.TaskActivityTaskQuery.sqlQuery(ctx)
	return tats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tats.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) StringsX(ctx context.Context) []string {
	v, err := tats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) StringX(ctx context.Context) string {
	v, err := tats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tats.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) IntsX(ctx context.Context) []int {
	v, err := tats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) IntX(ctx context.Context) int {
	v, err := tats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tats.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) Float64X(ctx context.Context) float64 {
	v, err := tats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tats.fields) > 1 {
		return nil, errors.New("ent: TaskActivityTaskSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) BoolsX(ctx context.Context) []bool {
	v, err := tats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tats *TaskActivityTaskSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskactivitytask.Label}
	default:
		err = fmt.Errorf("ent: TaskActivityTaskSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tats *TaskActivityTaskSelect) BoolX(ctx context.Context) bool {
	v, err := tats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tats *TaskActivityTaskSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tats.sql.Query()
	if err := tats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
