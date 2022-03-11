// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestUserQuery is the builder for querying TestUser entities.
type TestUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TestUser
	// eager-loading edges.
	withTestTodos *TestTodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestUserQuery builder.
func (tuq *TestUserQuery) Where(ps ...predicate.TestUser) *TestUserQuery {
	tuq.predicates = append(tuq.predicates, ps...)
	return tuq
}

// Limit adds a limit step to the query.
func (tuq *TestUserQuery) Limit(limit int) *TestUserQuery {
	tuq.limit = &limit
	return tuq
}

// Offset adds an offset step to the query.
func (tuq *TestUserQuery) Offset(offset int) *TestUserQuery {
	tuq.offset = &offset
	return tuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tuq *TestUserQuery) Unique(unique bool) *TestUserQuery {
	tuq.unique = &unique
	return tuq
}

// Order adds an order step to the query.
func (tuq *TestUserQuery) Order(o ...OrderFunc) *TestUserQuery {
	tuq.order = append(tuq.order, o...)
	return tuq
}

// QueryTestTodos chains the current query on the "testTodos" edge.
func (tuq *TestUserQuery) QueryTestTodos() *TestTodoQuery {
	query := &TestTodoQuery{config: tuq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testuser.Table, testuser.FieldID, selector),
			sqlgraph.To(testtodo.Table, testtodo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, testuser.TestTodosTable, testuser.TestTodosColumn),
		)
		fromU = sqlgraph.SetNeighbors(tuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestUser entity from the query.
// Returns a *NotFoundError when no TestUser was found.
func (tuq *TestUserQuery) First(ctx context.Context) (*TestUser, error) {
	nodes, err := tuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tuq *TestUserQuery) FirstX(ctx context.Context) *TestUser {
	node, err := tuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestUser ID from the query.
// Returns a *NotFoundError when no TestUser ID was found.
func (tuq *TestUserQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tuq *TestUserQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := tuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestUser entity is found.
// Returns a *NotFoundError when no TestUser entities are found.
func (tuq *TestUserQuery) Only(ctx context.Context) (*TestUser, error) {
	nodes, err := tuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testuser.Label}
	default:
		return nil, &NotSingularError{testuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tuq *TestUserQuery) OnlyX(ctx context.Context) *TestUser {
	node, err := tuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestUser ID in the query.
// Returns a *NotSingularError when more than one TestUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (tuq *TestUserQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = tuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = &NotSingularError{testuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tuq *TestUserQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := tuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestUsers.
func (tuq *TestUserQuery) All(ctx context.Context) ([]*TestUser, error) {
	if err := tuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tuq *TestUserQuery) AllX(ctx context.Context) []*TestUser {
	nodes, err := tuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestUser IDs.
func (tuq *TestUserQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := tuq.Select(testuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tuq *TestUserQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := tuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tuq *TestUserQuery) Count(ctx context.Context) (int, error) {
	if err := tuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tuq *TestUserQuery) CountX(ctx context.Context) int {
	count, err := tuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tuq *TestUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := tuq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tuq *TestUserQuery) ExistX(ctx context.Context) bool {
	exist, err := tuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tuq *TestUserQuery) Clone() *TestUserQuery {
	if tuq == nil {
		return nil
	}
	return &TestUserQuery{
		config:        tuq.config,
		limit:         tuq.limit,
		offset:        tuq.offset,
		order:         append([]OrderFunc{}, tuq.order...),
		predicates:    append([]predicate.TestUser{}, tuq.predicates...),
		withTestTodos: tuq.withTestTodos.Clone(),
		// clone intermediate query.
		sql:    tuq.sql.Clone(),
		path:   tuq.path,
		unique: tuq.unique,
	}
}

// WithTestTodos tells the query-builder to eager-load the nodes that are connected to
// the "testTodos" edge. The optional arguments are used to configure the query builder of the edge.
func (tuq *TestUserQuery) WithTestTodos(opts ...func(*TestTodoQuery)) *TestUserQuery {
	query := &TestTodoQuery{config: tuq.config}
	for _, opt := range opts {
		opt(query)
	}
	tuq.withTestTodos = query
	return tuq
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
//	client.TestUser.Query().
//		GroupBy(testuser.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tuq *TestUserQuery) GroupBy(field string, fields ...string) *TestUserGroupBy {
	group := &TestUserGroupBy{config: tuq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tuq.sqlQuery(ctx), nil
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
//	client.TestUser.Query().
//		Select(testuser.FieldName).
//		Scan(ctx, &v)
//
func (tuq *TestUserQuery) Select(fields ...string) *TestUserSelect {
	tuq.fields = append(tuq.fields, fields...)
	return &TestUserSelect{TestUserQuery: tuq}
}

func (tuq *TestUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tuq.fields {
		if !testuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tuq.path != nil {
		prev, err := tuq.path(ctx)
		if err != nil {
			return err
		}
		tuq.sql = prev
	}
	return nil
}

func (tuq *TestUserQuery) sqlAll(ctx context.Context) ([]*TestUser, error) {
	var (
		nodes       = []*TestUser{}
		_spec       = tuq.querySpec()
		loadedTypes = [1]bool{
			tuq.withTestTodos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TestUser{config: tuq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tuq.withTestTodos; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*TestUser)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TestTodos = []*TestTodo{}
		}
		query.Where(predicate.TestTodo(func(s *sql.Selector) {
			s.Where(sql.InValues(testuser.TestTodosColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TestUserID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "test_user_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.TestTodos = append(node.Edges.TestTodos, n)
		}
	}

	return nodes, nil
}

func (tuq *TestUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tuq.querySpec()
	_spec.Node.Columns = tuq.fields
	if len(tuq.fields) > 0 {
		_spec.Unique = tuq.unique != nil && *tuq.unique
	}
	return sqlgraph.CountNodes(ctx, tuq.driver, _spec)
}

func (tuq *TestUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tuq *TestUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testuser.Table,
			Columns: testuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testuser.FieldID,
			},
		},
		From:   tuq.sql,
		Unique: true,
	}
	if unique := tuq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tuq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testuser.FieldID)
		for i := range fields {
			if fields[i] != testuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tuq *TestUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tuq.driver.Dialect())
	t1 := builder.Table(testuser.Table)
	columns := tuq.fields
	if len(columns) == 0 {
		columns = testuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tuq.sql != nil {
		selector = tuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tuq.unique != nil && *tuq.unique {
		selector.Distinct()
	}
	for _, p := range tuq.predicates {
		p(selector)
	}
	for _, p := range tuq.order {
		p(selector)
	}
	if offset := tuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestUserGroupBy is the group-by builder for TestUser entities.
type TestUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tugb *TestUserGroupBy) Aggregate(fns ...AggregateFunc) *TestUserGroupBy {
	tugb.fns = append(tugb.fns, fns...)
	return tugb
}

// Scan applies the group-by query and scans the result into the given value.
func (tugb *TestUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tugb.path(ctx)
	if err != nil {
		return err
	}
	tugb.sql = query
	return tugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tugb *TestUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tugb.fields) > 1 {
		return nil, errors.New("ent: TestUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tugb *TestUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := tugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tugb *TestUserGroupBy) StringX(ctx context.Context) string {
	v, err := tugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tugb.fields) > 1 {
		return nil, errors.New("ent: TestUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tugb *TestUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := tugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tugb *TestUserGroupBy) IntX(ctx context.Context) int {
	v, err := tugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tugb.fields) > 1 {
		return nil, errors.New("ent: TestUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tugb *TestUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tugb *TestUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tugb.fields) > 1 {
		return nil, errors.New("ent: TestUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tugb *TestUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tugb *TestUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tugb *TestUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := tugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tugb *TestUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tugb.fields {
		if !testuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tugb *TestUserGroupBy) sqlQuery() *sql.Selector {
	selector := tugb.sql.Select()
	aggregation := make([]string, 0, len(tugb.fns))
	for _, fn := range tugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tugb.fields)+len(tugb.fns))
		for _, f := range tugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tugb.fields...)...)
}

// TestUserSelect is the builder for selecting fields of TestUser entities.
type TestUserSelect struct {
	*TestUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tus *TestUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tus.prepareQuery(ctx); err != nil {
		return err
	}
	tus.sql = tus.TestUserQuery.sqlQuery(ctx)
	return tus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tus *TestUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tus.fields) > 1 {
		return nil, errors.New("ent: TestUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tus *TestUserSelect) StringsX(ctx context.Context) []string {
	v, err := tus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tus *TestUserSelect) StringX(ctx context.Context) string {
	v, err := tus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tus.fields) > 1 {
		return nil, errors.New("ent: TestUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tus *TestUserSelect) IntsX(ctx context.Context) []int {
	v, err := tus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tus *TestUserSelect) IntX(ctx context.Context) int {
	v, err := tus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tus.fields) > 1 {
		return nil, errors.New("ent: TestUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tus *TestUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tus *TestUserSelect) Float64X(ctx context.Context) float64 {
	v, err := tus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tus.fields) > 1 {
		return nil, errors.New("ent: TestUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tus *TestUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := tus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tus *TestUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{testuser.Label}
	default:
		err = fmt.Errorf("ent: TestUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tus *TestUserSelect) BoolX(ctx context.Context) bool {
	v, err := tus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tus *TestUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tus.sql.Query()
	if err := tus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
