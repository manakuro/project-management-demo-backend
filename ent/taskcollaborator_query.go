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
	"project-management-demo-backend/ent/taskcollaborator"
	"project-management-demo-backend/ent/teammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCollaboratorQuery is the builder for querying TaskCollaborator entities.
type TaskCollaboratorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TaskCollaborator
	// eager-loading edges.
	withTask     *TaskQuery
	withTeammate *TeammateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskCollaboratorQuery builder.
func (tcq *TaskCollaboratorQuery) Where(ps ...predicate.TaskCollaborator) *TaskCollaboratorQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit adds a limit step to the query.
func (tcq *TaskCollaboratorQuery) Limit(limit int) *TaskCollaboratorQuery {
	tcq.limit = &limit
	return tcq
}

// Offset adds an offset step to the query.
func (tcq *TaskCollaboratorQuery) Offset(offset int) *TaskCollaboratorQuery {
	tcq.offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TaskCollaboratorQuery) Unique(unique bool) *TaskCollaboratorQuery {
	tcq.unique = &unique
	return tcq
}

// Order adds an order step to the query.
func (tcq *TaskCollaboratorQuery) Order(o ...OrderFunc) *TaskCollaboratorQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QueryTask chains the current query on the "task" edge.
func (tcq *TaskCollaboratorQuery) QueryTask() *TaskQuery {
	query := &TaskQuery{config: tcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskcollaborator.Table, taskcollaborator.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskcollaborator.TaskTable, taskcollaborator.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeammate chains the current query on the "teammate" edge.
func (tcq *TaskCollaboratorQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: tcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskcollaborator.Table, taskcollaborator.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taskcollaborator.TeammateTable, taskcollaborator.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskCollaborator entity from the query.
// Returns a *NotFoundError when no TaskCollaborator was found.
func (tcq *TaskCollaboratorQuery) First(ctx context.Context) (*TaskCollaborator, error) {
	nodes, err := tcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taskcollaborator.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) FirstX(ctx context.Context) *TaskCollaborator {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskCollaborator ID from the query.
// Returns a *NotFoundError when no TaskCollaborator ID was found.
func (tcq *TaskCollaboratorQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taskcollaborator.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskCollaborator entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskCollaborator entity is found.
// Returns a *NotFoundError when no TaskCollaborator entities are found.
func (tcq *TaskCollaboratorQuery) Only(ctx context.Context) (*TaskCollaborator, error) {
	nodes, err := tcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taskcollaborator.Label}
	default:
		return nil, &NotSingularError{taskcollaborator.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) OnlyX(ctx context.Context) *TaskCollaborator {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskCollaborator ID in the query.
// Returns a *NotSingularError when more than one TaskCollaborator ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TaskCollaboratorQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = &NotSingularError{taskcollaborator.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskCollaborators.
func (tcq *TaskCollaboratorQuery) All(ctx context.Context) ([]*TaskCollaborator, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) AllX(ctx context.Context) []*TaskCollaborator {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskCollaborator IDs.
func (tcq *TaskCollaboratorQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tcq.Select(taskcollaborator.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TaskCollaboratorQuery) Count(ctx context.Context) (int, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TaskCollaboratorQuery) Exist(ctx context.Context) (bool, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TaskCollaboratorQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskCollaboratorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TaskCollaboratorQuery) Clone() *TaskCollaboratorQuery {
	if tcq == nil {
		return nil
	}
	return &TaskCollaboratorQuery{
		config:       tcq.config,
		limit:        tcq.limit,
		offset:       tcq.offset,
		order:        append([]OrderFunc{}, tcq.order...),
		predicates:   append([]predicate.TaskCollaborator{}, tcq.predicates...),
		withTask:     tcq.withTask.Clone(),
		withTeammate: tcq.withTeammate.Clone(),
		// clone intermediate query.
		sql:    tcq.sql.Clone(),
		path:   tcq.path,
		unique: tcq.unique,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TaskCollaboratorQuery) WithTask(opts ...func(*TaskQuery)) *TaskCollaboratorQuery {
	query := &TaskQuery{config: tcq.config}
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTask = query
	return tcq
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TaskCollaboratorQuery) WithTeammate(opts ...func(*TeammateQuery)) *TaskCollaboratorQuery {
	query := &TeammateQuery{config: tcq.config}
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTeammate = query
	return tcq
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
//	client.TaskCollaborator.Query().
//		GroupBy(taskcollaborator.FieldTaskID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tcq *TaskCollaboratorQuery) GroupBy(field string, fields ...string) *TaskCollaboratorGroupBy {
	group := &TaskCollaboratorGroupBy{config: tcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tcq.sqlQuery(ctx), nil
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
//	client.TaskCollaborator.Query().
//		Select(taskcollaborator.FieldTaskID).
//		Scan(ctx, &v)
//
func (tcq *TaskCollaboratorQuery) Select(fields ...string) *TaskCollaboratorSelect {
	tcq.fields = append(tcq.fields, fields...)
	return &TaskCollaboratorSelect{TaskCollaboratorQuery: tcq}
}

func (tcq *TaskCollaboratorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tcq.fields {
		if !taskcollaborator.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TaskCollaboratorQuery) sqlAll(ctx context.Context) ([]*TaskCollaborator, error) {
	var (
		nodes       = []*TaskCollaborator{}
		_spec       = tcq.querySpec()
		loadedTypes = [2]bool{
			tcq.withTask != nil,
			tcq.withTeammate != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TaskCollaborator{config: tcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tcq.withTask; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskCollaborator)
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

	if query := tcq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*TaskCollaborator)
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

func (tcq *TaskCollaboratorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.fields
	if len(tcq.fields) > 0 {
		_spec.Unique = tcq.unique != nil && *tcq.unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TaskCollaboratorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tcq *TaskCollaboratorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskcollaborator.Table,
			Columns: taskcollaborator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcollaborator.FieldID,
			},
		},
		From:   tcq.sql,
		Unique: true,
	}
	if unique := tcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskcollaborator.FieldID)
		for i := range fields {
			if fields[i] != taskcollaborator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TaskCollaboratorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(taskcollaborator.Table)
	columns := tcq.fields
	if len(columns) == 0 {
		columns = taskcollaborator.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.unique != nil && *tcq.unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskCollaboratorGroupBy is the group-by builder for TaskCollaborator entities.
type TaskCollaboratorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TaskCollaboratorGroupBy) Aggregate(fns ...AggregateFunc) *TaskCollaboratorGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tcgb *TaskCollaboratorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tcgb.path(ctx)
	if err != nil {
		return err
	}
	tcgb.sql = query
	return tcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) StringsX(ctx context.Context) []string {
	v, err := tcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) StringX(ctx context.Context) string {
	v, err := tcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) IntsX(ctx context.Context) []int {
	v, err := tcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) IntX(ctx context.Context) int {
	v, err := tcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskCollaboratorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcgb *TaskCollaboratorGroupBy) BoolX(ctx context.Context) bool {
	v, err := tcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcgb *TaskCollaboratorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tcgb.fields {
		if !taskcollaborator.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tcgb *TaskCollaboratorGroupBy) sqlQuery() *sql.Selector {
	selector := tcgb.sql.Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tcgb.fields)+len(tcgb.fns))
		for _, f := range tcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tcgb.fields...)...)
}

// TaskCollaboratorSelect is the builder for selecting fields of TaskCollaborator entities.
type TaskCollaboratorSelect struct {
	*TaskCollaboratorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TaskCollaboratorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	tcs.sql = tcs.TaskCollaboratorQuery.sqlQuery(ctx)
	return tcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) StringsX(ctx context.Context) []string {
	v, err := tcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) StringX(ctx context.Context) string {
	v, err := tcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) IntsX(ctx context.Context) []int {
	v, err := tcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) IntX(ctx context.Context) int {
	v, err := tcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) Float64X(ctx context.Context) float64 {
	v, err := tcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskCollaboratorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) BoolsX(ctx context.Context) []bool {
	v, err := tcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tcs *TaskCollaboratorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcollaborator.Label}
	default:
		err = fmt.Errorf("ent: TaskCollaboratorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcs *TaskCollaboratorSelect) BoolX(ctx context.Context) bool {
	v, err := tcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcs *TaskCollaboratorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tcs.sql.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
