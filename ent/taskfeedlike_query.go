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
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfeedlike"
	"project-management-demo-backend/ent/teammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFeedLikeQuery is the builder for querying TaskFeedLike entities.
type TaskFeedLikeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TaskFeedLike
	// eager-loading edges.
	withTask     *TaskQuery
	withTeammate *TeammateQuery
	withFeed     *TaskFeedQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskFeedLikeQuery builder.
func (tflq *TaskFeedLikeQuery) Where(ps ...predicate.TaskFeedLike) *TaskFeedLikeQuery {
	tflq.predicates = append(tflq.predicates, ps...)
	return tflq
}

// Limit adds a limit step to the query.
func (tflq *TaskFeedLikeQuery) Limit(limit int) *TaskFeedLikeQuery {
	tflq.limit = &limit
	return tflq
}

// Offset adds an offset step to the query.
func (tflq *TaskFeedLikeQuery) Offset(offset int) *TaskFeedLikeQuery {
	tflq.offset = &offset
	return tflq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tflq *TaskFeedLikeQuery) Unique(unique bool) *TaskFeedLikeQuery {
	tflq.unique = &unique
	return tflq
}

// Order adds an order step to the query.
func (tflq *TaskFeedLikeQuery) Order(o ...OrderFunc) *TaskFeedLikeQuery {
	tflq.order = append(tflq.order, o...)
	return tflq
}

// QueryTask chains the current query on the "task" edge.
func (tflq *TaskFeedLikeQuery) QueryTask() *TaskQuery {
	query := &TaskQuery{config: tflq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tflq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tflq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskfeedlike.Table, taskfeedlike.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskfeedlike.TaskTable, taskfeedlike.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(tflq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (tflq *TaskFeedLikeQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: tflq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tflq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tflq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskfeedlike.Table, taskfeedlike.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskfeedlike.TeammateTable, taskfeedlike.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(tflq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeed chains the current query on the "feed" edge.
func (tflq *TaskFeedLikeQuery) QueryFeed() *TaskFeedQuery {
	query := &TaskFeedQuery{config: tflq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tflq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tflq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskfeedlike.Table, taskfeedlike.FieldID, selector),
			sqlgraph.To(taskfeed.Table, taskfeed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskfeedlike.FeedTable, taskfeedlike.FeedColumn),
		)
		fromU = sqlgraph.SetNeighbors(tflq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskFeedLike entity from the query.
// Returns a *NotFoundError when no TaskFeedLike was found.
func (tflq *TaskFeedLikeQuery) First(ctx context.Context) (*TaskFeedLike, error) {
	nodes, err := tflq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taskfeedlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) FirstX(ctx context.Context) *TaskFeedLike {
	node, err := tflq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskFeedLike ID from the query.
// Returns a *NotFoundError when no TaskFeedLike ID was found.
func (tflq *TaskFeedLikeQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tflq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taskfeedlike.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tflq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskFeedLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskFeedLike entity is found.
// Returns a *NotFoundError when no TaskFeedLike entities are found.
func (tflq *TaskFeedLikeQuery) Only(ctx context.Context) (*TaskFeedLike, error) {
	nodes, err := tflq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taskfeedlike.Label}
	default:
		return nil, &NotSingularError{taskfeedlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) OnlyX(ctx context.Context) *TaskFeedLike {
	node, err := tflq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskFeedLike ID in the query.
// Returns a *NotSingularError when more than one TaskFeedLike ID is found.
// Returns a *NotFoundError when no entities are found.
func (tflq *TaskFeedLikeQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tflq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = &NotSingularError{taskfeedlike.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tflq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskFeedLikes.
func (tflq *TaskFeedLikeQuery) All(ctx context.Context) ([]*TaskFeedLike, error) {
	if err := tflq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tflq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) AllX(ctx context.Context) []*TaskFeedLike {
	nodes, err := tflq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskFeedLike IDs.
func (tflq *TaskFeedLikeQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tflq.Select(taskfeedlike.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tflq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tflq *TaskFeedLikeQuery) Count(ctx context.Context) (int, error) {
	if err := tflq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tflq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) CountX(ctx context.Context) int {
	count, err := tflq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tflq *TaskFeedLikeQuery) Exist(ctx context.Context) (bool, error) {
	if err := tflq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tflq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tflq *TaskFeedLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := tflq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskFeedLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tflq *TaskFeedLikeQuery) Clone() *TaskFeedLikeQuery {
	if tflq == nil {
		return nil
	}
	return &TaskFeedLikeQuery{
		config:       tflq.config,
		limit:        tflq.limit,
		offset:       tflq.offset,
		order:        append([]OrderFunc{}, tflq.order...),
		predicates:   append([]predicate.TaskFeedLike{}, tflq.predicates...),
		withTask:     tflq.withTask.Clone(),
		withTeammate: tflq.withTeammate.Clone(),
		withFeed:     tflq.withFeed.Clone(),
		// clone intermediate query.
		sql:    tflq.sql.Clone(),
		path:   tflq.path,
		unique: tflq.unique,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (tflq *TaskFeedLikeQuery) WithTask(opts ...func(*TaskQuery)) *TaskFeedLikeQuery {
	query := &TaskQuery{config: tflq.config}
	for _, opt := range opts {
		opt(query)
	}
	tflq.withTask = query
	return tflq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (tflq *TaskFeedLikeQuery) WithTeammate(opts ...func(*TeammateQuery)) *TaskFeedLikeQuery {
	query := &TeammateQuery{config: tflq.config}
	for _, opt := range opts {
		opt(query)
	}
	tflq.withTeammate = query
	return tflq
}

// WithFeed tells the query-builder to eager-load the nodes that are connected to
// the "feed" edge. The optional arguments are used to configure the query builder of the edge.
func (tflq *TaskFeedLikeQuery) WithFeed(opts ...func(*TaskFeedQuery)) *TaskFeedLikeQuery {
	query := &TaskFeedQuery{config: tflq.config}
	for _, opt := range opts {
		opt(query)
	}
	tflq.withFeed = query
	return tflq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TaskID ulid.ID `json:"task_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TaskFeedLike.Query().
//		GroupBy(taskfeedlike.FieldTaskID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tflq *TaskFeedLikeQuery) GroupBy(field string, fields ...string) *TaskFeedLikeGroupBy {
	group := &TaskFeedLikeGroupBy{config: tflq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tflq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tflq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TaskID ulid.ID `json:"task_id,omitempty"`
//	}
//
//	client.TaskFeedLike.Query().
//		Select(taskfeedlike.FieldTaskID).
//		Scan(ctx, &v)
//
func (tflq *TaskFeedLikeQuery) Select(fields ...string) *TaskFeedLikeSelect {
	tflq.fields = append(tflq.fields, fields...)
	return &TaskFeedLikeSelect{TaskFeedLikeQuery: tflq}
}

func (tflq *TaskFeedLikeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tflq.fields {
		if !taskfeedlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tflq.path != nil {
		prev, err := tflq.path(ctx)
		if err != nil {
			return err
		}
		tflq.sql = prev
	}
	return nil
}

func (tflq *TaskFeedLikeQuery) sqlAll(ctx context.Context) ([]*TaskFeedLike, error) {
	var (
		nodes       = []*TaskFeedLike{}
		_spec       = tflq.querySpec()
		loadedTypes = [3]bool{
			tflq.withTask != nil,
			tflq.withTeammate != nil,
			tflq.withFeed != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TaskFeedLike{config: tflq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tflq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tflq.withTask; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskFeedLike)
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

	if query := tflq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskFeedLike)
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

	if query := tflq.withFeed; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskFeedLike)
		for i := range nodes {
			fk := nodes[i].TaskFeedID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(taskfeed.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_feed_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Feed = n
			}
		}
	}

	return nodes, nil
}

func (tflq *TaskFeedLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tflq.querySpec()
	_spec.Node.Columns = tflq.fields
	if len(tflq.fields) > 0 {
		_spec.Unique = tflq.unique != nil && *tflq.unique
	}
	return sqlgraph.CountNodes(ctx, tflq.driver, _spec)
}

func (tflq *TaskFeedLikeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tflq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tflq *TaskFeedLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskfeedlike.Table,
			Columns: taskfeedlike.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfeedlike.FieldID,
			},
		},
		From:   tflq.sql,
		Unique: true,
	}
	if unique := tflq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tflq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskfeedlike.FieldID)
		for i := range fields {
			if fields[i] != taskfeedlike.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tflq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tflq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tflq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tflq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tflq *TaskFeedLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tflq.driver.Dialect())
	t1 := builder.Table(taskfeedlike.Table)
	columns := tflq.fields
	if len(columns) == 0 {
		columns = taskfeedlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tflq.sql != nil {
		selector = tflq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tflq.unique != nil && *tflq.unique {
		selector.Distinct()
	}
	for _, p := range tflq.predicates {
		p(selector)
	}
	for _, p := range tflq.order {
		p(selector)
	}
	if offset := tflq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tflq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskFeedLikeGroupBy is the group-by builder for TaskFeedLike entities.
type TaskFeedLikeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tflgb *TaskFeedLikeGroupBy) Aggregate(fns ...AggregateFunc) *TaskFeedLikeGroupBy {
	tflgb.fns = append(tflgb.fns, fns...)
	return tflgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tflgb *TaskFeedLikeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tflgb.path(ctx)
	if err != nil {
		return err
	}
	tflgb.sql = query
	return tflgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tflgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tflgb.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tflgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) StringsX(ctx context.Context) []string {
	v, err := tflgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tflgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) StringX(ctx context.Context) string {
	v, err := tflgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tflgb.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tflgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) IntsX(ctx context.Context) []int {
	v, err := tflgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tflgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) IntX(ctx context.Context) int {
	v, err := tflgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tflgb.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tflgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tflgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tflgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tflgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tflgb.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tflgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tflgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tflgb *TaskFeedLikeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tflgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tflgb *TaskFeedLikeGroupBy) BoolX(ctx context.Context) bool {
	v, err := tflgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tflgb *TaskFeedLikeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tflgb.fields {
		if !taskfeedlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tflgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tflgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tflgb *TaskFeedLikeGroupBy) sqlQuery() *sql.Selector {
	selector := tflgb.sql.Select()
	aggregation := make([]string, 0, len(tflgb.fns))
	for _, fn := range tflgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tflgb.fields)+len(tflgb.fns))
		for _, f := range tflgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tflgb.fields...)...)
}

// TaskFeedLikeSelect is the builder for selecting fields of TaskFeedLike entities.
type TaskFeedLikeSelect struct {
	*TaskFeedLikeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tfls *TaskFeedLikeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tfls.prepareQuery(ctx); err != nil {
		return err
	}
	tfls.sql = tfls.TaskFeedLikeQuery.sqlQuery(ctx)
	return tfls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tfls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tfls.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tfls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) StringsX(ctx context.Context) []string {
	v, err := tfls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tfls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) StringX(ctx context.Context) string {
	v, err := tfls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tfls.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tfls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) IntsX(ctx context.Context) []int {
	v, err := tfls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tfls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) IntX(ctx context.Context) int {
	v, err := tfls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tfls.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tfls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tfls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tfls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) Float64X(ctx context.Context) float64 {
	v, err := tfls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tfls.fields) > 1 {
		return nil, errors.New("ent: TaskFeedLikeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tfls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) BoolsX(ctx context.Context) []bool {
	v, err := tfls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tfls *TaskFeedLikeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tfls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskfeedlike.Label}
	default:
		err = fmt.Errorf("ent: TaskFeedLikeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tfls *TaskFeedLikeSelect) BoolX(ctx context.Context) bool {
	v, err := tfls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tfls *TaskFeedLikeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tfls.sql.Query()
	if err := tfls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
