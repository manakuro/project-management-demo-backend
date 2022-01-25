// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistsortstatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskListSortStatusQuery is the builder for querying TaskListSortStatus entities.
type TaskListSortStatusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TaskListSortStatus
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskListSortStatusQuery builder.
func (tlssq *TaskListSortStatusQuery) Where(ps ...predicate.TaskListSortStatus) *TaskListSortStatusQuery {
	tlssq.predicates = append(tlssq.predicates, ps...)
	return tlssq
}

// Limit adds a limit step to the query.
func (tlssq *TaskListSortStatusQuery) Limit(limit int) *TaskListSortStatusQuery {
	tlssq.limit = &limit
	return tlssq
}

// Offset adds an offset step to the query.
func (tlssq *TaskListSortStatusQuery) Offset(offset int) *TaskListSortStatusQuery {
	tlssq.offset = &offset
	return tlssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlssq *TaskListSortStatusQuery) Unique(unique bool) *TaskListSortStatusQuery {
	tlssq.unique = &unique
	return tlssq
}

// Order adds an order step to the query.
func (tlssq *TaskListSortStatusQuery) Order(o ...OrderFunc) *TaskListSortStatusQuery {
	tlssq.order = append(tlssq.order, o...)
	return tlssq
}

// First returns the first TaskListSortStatus entity from the query.
// Returns a *NotFoundError when no TaskListSortStatus was found.
func (tlssq *TaskListSortStatusQuery) First(ctx context.Context) (*TaskListSortStatus, error) {
	nodes, err := tlssq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tasklistsortstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) FirstX(ctx context.Context) *TaskListSortStatus {
	node, err := tlssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskListSortStatus ID from the query.
// Returns a *NotFoundError when no TaskListSortStatus ID was found.
func (tlssq *TaskListSortStatusQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tlssq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tasklistsortstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tlssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskListSortStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TaskListSortStatus entity is not found.
// Returns a *NotFoundError when no TaskListSortStatus entities are found.
func (tlssq *TaskListSortStatusQuery) Only(ctx context.Context) (*TaskListSortStatus, error) {
	nodes, err := tlssq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tasklistsortstatus.Label}
	default:
		return nil, &NotSingularError{tasklistsortstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) OnlyX(ctx context.Context) *TaskListSortStatus {
	node, err := tlssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskListSortStatus ID in the query.
// Returns a *NotSingularError when exactly one TaskListSortStatus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tlssq *TaskListSortStatusQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tlssq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = &NotSingularError{tasklistsortstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tlssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskListSortStatusSlice.
func (tlssq *TaskListSortStatusQuery) All(ctx context.Context) ([]*TaskListSortStatus, error) {
	if err := tlssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tlssq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) AllX(ctx context.Context) []*TaskListSortStatus {
	nodes, err := tlssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskListSortStatus IDs.
func (tlssq *TaskListSortStatusQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tlssq.Select(tasklistsortstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tlssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tlssq *TaskListSortStatusQuery) Count(ctx context.Context) (int, error) {
	if err := tlssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tlssq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) CountX(ctx context.Context) int {
	count, err := tlssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlssq *TaskListSortStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := tlssq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tlssq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tlssq *TaskListSortStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := tlssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskListSortStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlssq *TaskListSortStatusQuery) Clone() *TaskListSortStatusQuery {
	if tlssq == nil {
		return nil
	}
	return &TaskListSortStatusQuery{
		config:     tlssq.config,
		limit:      tlssq.limit,
		offset:     tlssq.offset,
		order:      append([]OrderFunc{}, tlssq.order...),
		predicates: append([]predicate.TaskListSortStatus{}, tlssq.predicates...),
		// clone intermediate query.
		sql:  tlssq.sql.Clone(),
		path: tlssq.path,
	}
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
//	client.TaskListSortStatus.Query().
//		GroupBy(tasklistsortstatus.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tlssq *TaskListSortStatusQuery) GroupBy(field string, fields ...string) *TaskListSortStatusGroupBy {
	group := &TaskListSortStatusGroupBy{config: tlssq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tlssq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tlssq.sqlQuery(ctx), nil
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
//	client.TaskListSortStatus.Query().
//		Select(tasklistsortstatus.FieldName).
//		Scan(ctx, &v)
//
func (tlssq *TaskListSortStatusQuery) Select(fields ...string) *TaskListSortStatusSelect {
	tlssq.fields = append(tlssq.fields, fields...)
	return &TaskListSortStatusSelect{TaskListSortStatusQuery: tlssq}
}

func (tlssq *TaskListSortStatusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tlssq.fields {
		if !tasklistsortstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlssq.path != nil {
		prev, err := tlssq.path(ctx)
		if err != nil {
			return err
		}
		tlssq.sql = prev
	}
	return nil
}

func (tlssq *TaskListSortStatusQuery) sqlAll(ctx context.Context) ([]*TaskListSortStatus, error) {
	var (
		nodes = []*TaskListSortStatus{}
		_spec = tlssq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TaskListSortStatus{config: tlssq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tlssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tlssq *TaskListSortStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlssq.querySpec()
	return sqlgraph.CountNodes(ctx, tlssq.driver, _spec)
}

func (tlssq *TaskListSortStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tlssq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tlssq *TaskListSortStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklistsortstatus.Table,
			Columns: tasklistsortstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistsortstatus.FieldID,
			},
		},
		From:   tlssq.sql,
		Unique: true,
	}
	if unique := tlssq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tlssq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklistsortstatus.FieldID)
		for i := range fields {
			if fields[i] != tasklistsortstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tlssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlssq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlssq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlssq *TaskListSortStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlssq.driver.Dialect())
	t1 := builder.Table(tasklistsortstatus.Table)
	columns := tlssq.fields
	if len(columns) == 0 {
		columns = tasklistsortstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlssq.sql != nil {
		selector = tlssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range tlssq.predicates {
		p(selector)
	}
	for _, p := range tlssq.order {
		p(selector)
	}
	if offset := tlssq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlssq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskListSortStatusGroupBy is the group-by builder for TaskListSortStatus entities.
type TaskListSortStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlssgb *TaskListSortStatusGroupBy) Aggregate(fns ...AggregateFunc) *TaskListSortStatusGroupBy {
	tlssgb.fns = append(tlssgb.fns, fns...)
	return tlssgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tlssgb *TaskListSortStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tlssgb.path(ctx)
	if err != nil {
		return err
	}
	tlssgb.sql = query
	return tlssgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tlssgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tlssgb.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tlssgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := tlssgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tlssgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) StringX(ctx context.Context) string {
	v, err := tlssgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tlssgb.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tlssgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := tlssgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tlssgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) IntX(ctx context.Context) int {
	v, err := tlssgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tlssgb.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tlssgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tlssgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tlssgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tlssgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tlssgb.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tlssgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tlssgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tlssgb *TaskListSortStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tlssgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tlssgb *TaskListSortStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := tlssgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tlssgb *TaskListSortStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tlssgb.fields {
		if !tasklistsortstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tlssgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlssgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tlssgb *TaskListSortStatusGroupBy) sqlQuery() *sql.Selector {
	selector := tlssgb.sql.Select()
	aggregation := make([]string, 0, len(tlssgb.fns))
	for _, fn := range tlssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tlssgb.fields)+len(tlssgb.fns))
		for _, f := range tlssgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tlssgb.fields...)...)
}

// TaskListSortStatusSelect is the builder for selecting fields of TaskListSortStatus entities.
type TaskListSortStatusSelect struct {
	*TaskListSortStatusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tlsss *TaskListSortStatusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tlsss.prepareQuery(ctx); err != nil {
		return err
	}
	tlsss.sql = tlsss.TaskListSortStatusQuery.sqlQuery(ctx)
	return tlsss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tlsss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tlsss.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tlsss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) StringsX(ctx context.Context) []string {
	v, err := tlsss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tlsss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) StringX(ctx context.Context) string {
	v, err := tlsss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tlsss.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tlsss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) IntsX(ctx context.Context) []int {
	v, err := tlsss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tlsss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) IntX(ctx context.Context) int {
	v, err := tlsss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tlsss.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tlsss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tlsss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tlsss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := tlsss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tlsss.fields) > 1 {
		return nil, errors.New("ent: TaskListSortStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tlsss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := tlsss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tlsss *TaskListSortStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tlsss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tasklistsortstatus.Label}
	default:
		err = fmt.Errorf("ent: TaskListSortStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tlsss *TaskListSortStatusSelect) BoolX(ctx context.Context) bool {
	v, err := tlsss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tlsss *TaskListSortStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tlsss.sql.Query()
	if err := tlsss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
