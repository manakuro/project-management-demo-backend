// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskcolumn"
	"project-management-demo-backend/ent/teammatetaskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskColumnQuery is the builder for querying TaskColumn entities.
type TaskColumnQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TaskColumn
	// eager-loading edges.
	withTeammateTaskColumns *TeammateTaskColumnQuery
	withProjectTaskColumns  *ProjectTaskColumnQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskColumnQuery builder.
func (tcq *TaskColumnQuery) Where(ps ...predicate.TaskColumn) *TaskColumnQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit adds a limit step to the query.
func (tcq *TaskColumnQuery) Limit(limit int) *TaskColumnQuery {
	tcq.limit = &limit
	return tcq
}

// Offset adds an offset step to the query.
func (tcq *TaskColumnQuery) Offset(offset int) *TaskColumnQuery {
	tcq.offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TaskColumnQuery) Unique(unique bool) *TaskColumnQuery {
	tcq.unique = &unique
	return tcq
}

// Order adds an order step to the query.
func (tcq *TaskColumnQuery) Order(o ...OrderFunc) *TaskColumnQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QueryTeammateTaskColumns chains the current query on the "teammateTaskColumns" edge.
func (tcq *TaskColumnQuery) QueryTeammateTaskColumns() *TeammateTaskColumnQuery {
	query := &TeammateTaskColumnQuery{config: tcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskcolumn.Table, taskcolumn.FieldID, selector),
			sqlgraph.To(teammatetaskcolumn.Table, teammatetaskcolumn.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, taskcolumn.TeammateTaskColumnsTable, taskcolumn.TeammateTaskColumnsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjectTaskColumns chains the current query on the "projectTaskColumns" edge.
func (tcq *TaskColumnQuery) QueryProjectTaskColumns() *ProjectTaskColumnQuery {
	query := &ProjectTaskColumnQuery{config: tcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskcolumn.Table, taskcolumn.FieldID, selector),
			sqlgraph.To(projecttaskcolumn.Table, projecttaskcolumn.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, taskcolumn.ProjectTaskColumnsTable, taskcolumn.ProjectTaskColumnsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskColumn entity from the query.
// Returns a *NotFoundError when no TaskColumn was found.
func (tcq *TaskColumnQuery) First(ctx context.Context) (*TaskColumn, error) {
	nodes, err := tcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taskcolumn.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TaskColumnQuery) FirstX(ctx context.Context) *TaskColumn {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskColumn ID from the query.
// Returns a *NotFoundError when no TaskColumn ID was found.
func (tcq *TaskColumnQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taskcolumn.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TaskColumnQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskColumn entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskColumn entity is found.
// Returns a *NotFoundError when no TaskColumn entities are found.
func (tcq *TaskColumnQuery) Only(ctx context.Context) (*TaskColumn, error) {
	nodes, err := tcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taskcolumn.Label}
	default:
		return nil, &NotSingularError{taskcolumn.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TaskColumnQuery) OnlyX(ctx context.Context) *TaskColumn {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskColumn ID in the query.
// Returns a *NotSingularError when more than one TaskColumn ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TaskColumnQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = &NotSingularError{taskcolumn.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TaskColumnQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskColumns.
func (tcq *TaskColumnQuery) All(ctx context.Context) ([]*TaskColumn, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TaskColumnQuery) AllX(ctx context.Context) []*TaskColumn {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskColumn IDs.
func (tcq *TaskColumnQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tcq.Select(taskcolumn.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TaskColumnQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TaskColumnQuery) Count(ctx context.Context) (int, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TaskColumnQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TaskColumnQuery) Exist(ctx context.Context) (bool, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TaskColumnQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskColumnQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TaskColumnQuery) Clone() *TaskColumnQuery {
	if tcq == nil {
		return nil
	}
	return &TaskColumnQuery{
		config:                  tcq.config,
		limit:                   tcq.limit,
		offset:                  tcq.offset,
		order:                   append([]OrderFunc{}, tcq.order...),
		predicates:              append([]predicate.TaskColumn{}, tcq.predicates...),
		withTeammateTaskColumns: tcq.withTeammateTaskColumns.Clone(),
		withProjectTaskColumns:  tcq.withProjectTaskColumns.Clone(),
		// clone intermediate query.
		sql:    tcq.sql.Clone(),
		path:   tcq.path,
		unique: tcq.unique,
	}
}

// WithTeammateTaskColumns tells the query-builder to eager-load the nodes that are connected to
// the "teammateTaskColumns" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TaskColumnQuery) WithTeammateTaskColumns(opts ...func(*TeammateTaskColumnQuery)) *TaskColumnQuery {
	query := &TeammateTaskColumnQuery{config: tcq.config}
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTeammateTaskColumns = query
	return tcq
}

// WithProjectTaskColumns tells the query-builder to eager-load the nodes that are connected to
// the "projectTaskColumns" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TaskColumnQuery) WithProjectTaskColumns(opts ...func(*ProjectTaskColumnQuery)) *TaskColumnQuery {
	query := &ProjectTaskColumnQuery{config: tcq.config}
	for _, opt := range opts {
		opt(query)
	}
	tcq.withProjectTaskColumns = query
	return tcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TaskColumn.Query().
//		GroupBy(taskcolumn.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tcq *TaskColumnQuery) GroupBy(field string, fields ...string) *TaskColumnGroupBy {
	group := &TaskColumnGroupBy{config: tcq.config}
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
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TaskColumn.Query().
//		Select(taskcolumn.FieldName).
//		Scan(ctx, &v)
//
func (tcq *TaskColumnQuery) Select(fields ...string) *TaskColumnSelect {
	tcq.fields = append(tcq.fields, fields...)
	return &TaskColumnSelect{TaskColumnQuery: tcq}
}

func (tcq *TaskColumnQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tcq.fields {
		if !taskcolumn.ValidColumn(f) {
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

func (tcq *TaskColumnQuery) sqlAll(ctx context.Context) ([]*TaskColumn, error) {
	var (
		nodes       = []*TaskColumn{}
		_spec       = tcq.querySpec()
		loadedTypes = [2]bool{
			tcq.withTeammateTaskColumns != nil,
			tcq.withProjectTaskColumns != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TaskColumn{config: tcq.config}
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

	if query := tcq.withTeammateTaskColumns; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*TaskColumn)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TeammateTaskColumns = []*TeammateTaskColumn{}
		}
		query.Where(predicate.TeammateTaskColumn(func(s *sql.Selector) {
			s.Where(sql.InValues(taskcolumn.TeammateTaskColumnsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TaskColumnID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_column_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.TeammateTaskColumns = append(node.Edges.TeammateTaskColumns, n)
		}
	}

	if query := tcq.withProjectTaskColumns; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*TaskColumn)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ProjectTaskColumns = []*ProjectTaskColumn{}
		}
		query.Where(predicate.ProjectTaskColumn(func(s *sql.Selector) {
			s.Where(sql.InValues(taskcolumn.ProjectTaskColumnsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TaskColumnID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_column_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ProjectTaskColumns = append(node.Edges.ProjectTaskColumns, n)
		}
	}

	return nodes, nil
}

func (tcq *TaskColumnQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.fields
	if len(tcq.fields) > 0 {
		_spec.Unique = tcq.unique != nil && *tcq.unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TaskColumnQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tcq *TaskColumnQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskcolumn.Table,
			Columns: taskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcolumn.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, taskcolumn.FieldID)
		for i := range fields {
			if fields[i] != taskcolumn.FieldID {
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

func (tcq *TaskColumnQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(taskcolumn.Table)
	columns := tcq.fields
	if len(columns) == 0 {
		columns = taskcolumn.Columns
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

// TaskColumnGroupBy is the group-by builder for TaskColumn entities.
type TaskColumnGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TaskColumnGroupBy) Aggregate(fns ...AggregateFunc) *TaskColumnGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tcgb *TaskColumnGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tcgb.path(ctx)
	if err != nil {
		return err
	}
	tcgb.sql = query
	return tcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskColumnGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) StringsX(ctx context.Context) []string {
	v, err := tcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) StringX(ctx context.Context) string {
	v, err := tcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskColumnGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) IntsX(ctx context.Context) []int {
	v, err := tcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) IntX(ctx context.Context) int {
	v, err := tcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskColumnGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("ent: TaskColumnGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TaskColumnGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcgb *TaskColumnGroupBy) BoolX(ctx context.Context) bool {
	v, err := tcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcgb *TaskColumnGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tcgb.fields {
		if !taskcolumn.ValidColumn(f) {
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

func (tcgb *TaskColumnGroupBy) sqlQuery() *sql.Selector {
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

// TaskColumnSelect is the builder for selecting fields of TaskColumn entities.
type TaskColumnSelect struct {
	*TaskColumnQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TaskColumnSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	tcs.sql = tcs.TaskColumnQuery.sqlQuery(ctx)
	return tcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcs *TaskColumnSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskColumnSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcs *TaskColumnSelect) StringsX(ctx context.Context) []string {
	v, err := tcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcs *TaskColumnSelect) StringX(ctx context.Context) string {
	v, err := tcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskColumnSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcs *TaskColumnSelect) IntsX(ctx context.Context) []int {
	v, err := tcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcs *TaskColumnSelect) IntX(ctx context.Context) int {
	v, err := tcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskColumnSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcs *TaskColumnSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcs *TaskColumnSelect) Float64X(ctx context.Context) float64 {
	v, err := tcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("ent: TaskColumnSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcs *TaskColumnSelect) BoolsX(ctx context.Context) []bool {
	v, err := tcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tcs *TaskColumnSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{taskcolumn.Label}
	default:
		err = fmt.Errorf("ent: TaskColumnSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcs *TaskColumnSelect) BoolX(ctx context.Context) bool {
	v, err := tcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcs *TaskColumnSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tcs.sql.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
