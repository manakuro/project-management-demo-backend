// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskColumnQuery is the builder for querying ProjectTaskColumn entities.
type ProjectTaskColumnQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProjectTaskColumn
	// eager-loading edges.
	withProject    *ProjectQuery
	withTaskColumn *TaskColumnQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectTaskColumnQuery builder.
func (ptcq *ProjectTaskColumnQuery) Where(ps ...predicate.ProjectTaskColumn) *ProjectTaskColumnQuery {
	ptcq.predicates = append(ptcq.predicates, ps...)
	return ptcq
}

// Limit adds a limit step to the query.
func (ptcq *ProjectTaskColumnQuery) Limit(limit int) *ProjectTaskColumnQuery {
	ptcq.limit = &limit
	return ptcq
}

// Offset adds an offset step to the query.
func (ptcq *ProjectTaskColumnQuery) Offset(offset int) *ProjectTaskColumnQuery {
	ptcq.offset = &offset
	return ptcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptcq *ProjectTaskColumnQuery) Unique(unique bool) *ProjectTaskColumnQuery {
	ptcq.unique = &unique
	return ptcq
}

// Order adds an order step to the query.
func (ptcq *ProjectTaskColumnQuery) Order(o ...OrderFunc) *ProjectTaskColumnQuery {
	ptcq.order = append(ptcq.order, o...)
	return ptcq
}

// QueryProject chains the current query on the "project" edge.
func (ptcq *ProjectTaskColumnQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: ptcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttaskcolumn.Table, projecttaskcolumn.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttaskcolumn.ProjectTable, projecttaskcolumn.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskColumn chains the current query on the "task_column" edge.
func (ptcq *ProjectTaskColumnQuery) QueryTaskColumn() *TaskColumnQuery {
	query := &TaskColumnQuery{config: ptcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttaskcolumn.Table, projecttaskcolumn.FieldID, selector),
			sqlgraph.To(taskcolumn.Table, taskcolumn.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttaskcolumn.TaskColumnTable, projecttaskcolumn.TaskColumnColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectTaskColumn entity from the query.
// Returns a *NotFoundError when no ProjectTaskColumn was found.
func (ptcq *ProjectTaskColumnQuery) First(ctx context.Context) (*ProjectTaskColumn, error) {
	nodes, err := ptcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projecttaskcolumn.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) FirstX(ctx context.Context) *ProjectTaskColumn {
	node, err := ptcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectTaskColumn ID from the query.
// Returns a *NotFoundError when no ProjectTaskColumn ID was found.
func (ptcq *ProjectTaskColumnQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projecttaskcolumn.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := ptcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectTaskColumn entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProjectTaskColumn entity is not found.
// Returns a *NotFoundError when no ProjectTaskColumn entities are found.
func (ptcq *ProjectTaskColumnQuery) Only(ctx context.Context) (*ProjectTaskColumn, error) {
	nodes, err := ptcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projecttaskcolumn.Label}
	default:
		return nil, &NotSingularError{projecttaskcolumn.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) OnlyX(ctx context.Context) *ProjectTaskColumn {
	node, err := ptcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectTaskColumn ID in the query.
// Returns a *NotSingularError when exactly one ProjectTaskColumn ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ptcq *ProjectTaskColumnQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = &NotSingularError{projecttaskcolumn.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := ptcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectTaskColumns.
func (ptcq *ProjectTaskColumnQuery) All(ctx context.Context) ([]*ProjectTaskColumn, error) {
	if err := ptcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) AllX(ctx context.Context) []*ProjectTaskColumn {
	nodes, err := ptcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectTaskColumn IDs.
func (ptcq *ProjectTaskColumnQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := ptcq.Select(projecttaskcolumn.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := ptcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptcq *ProjectTaskColumnQuery) Count(ctx context.Context) (int, error) {
	if err := ptcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) CountX(ctx context.Context) int {
	count, err := ptcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptcq *ProjectTaskColumnQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptcq *ProjectTaskColumnQuery) ExistX(ctx context.Context) bool {
	exist, err := ptcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectTaskColumnQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptcq *ProjectTaskColumnQuery) Clone() *ProjectTaskColumnQuery {
	if ptcq == nil {
		return nil
	}
	return &ProjectTaskColumnQuery{
		config:         ptcq.config,
		limit:          ptcq.limit,
		offset:         ptcq.offset,
		order:          append([]OrderFunc{}, ptcq.order...),
		predicates:     append([]predicate.ProjectTaskColumn{}, ptcq.predicates...),
		withProject:    ptcq.withProject.Clone(),
		withTaskColumn: ptcq.withTaskColumn.Clone(),
		// clone intermediate query.
		sql:  ptcq.sql.Clone(),
		path: ptcq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (ptcq *ProjectTaskColumnQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectTaskColumnQuery {
	query := &ProjectQuery{config: ptcq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptcq.withProject = query
	return ptcq
}

// WithTaskColumn tells the query-builder to eager-load the nodes that are connected to
// the "task_column" edge. The optional arguments are used to configure the query builder of the edge.
func (ptcq *ProjectTaskColumnQuery) WithTaskColumn(opts ...func(*TaskColumnQuery)) *ProjectTaskColumnQuery {
	query := &TaskColumnQuery{config: ptcq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptcq.withTaskColumn = query
	return ptcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProjectID ulid.ID `json:"project_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectTaskColumn.Query().
//		GroupBy(projecttaskcolumn.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptcq *ProjectTaskColumnQuery) GroupBy(field string, fields ...string) *ProjectTaskColumnGroupBy {
	group := &ProjectTaskColumnGroupBy{config: ptcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProjectID ulid.ID `json:"project_id,omitempty"`
//	}
//
//	client.ProjectTaskColumn.Query().
//		Select(projecttaskcolumn.FieldProjectID).
//		Scan(ctx, &v)
//
func (ptcq *ProjectTaskColumnQuery) Select(fields ...string) *ProjectTaskColumnSelect {
	ptcq.fields = append(ptcq.fields, fields...)
	return &ProjectTaskColumnSelect{ProjectTaskColumnQuery: ptcq}
}

func (ptcq *ProjectTaskColumnQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptcq.fields {
		if !projecttaskcolumn.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptcq.path != nil {
		prev, err := ptcq.path(ctx)
		if err != nil {
			return err
		}
		ptcq.sql = prev
	}
	return nil
}

func (ptcq *ProjectTaskColumnQuery) sqlAll(ctx context.Context) ([]*ProjectTaskColumn, error) {
	var (
		nodes       = []*ProjectTaskColumn{}
		_spec       = ptcq.querySpec()
		loadedTypes = [2]bool{
			ptcq.withProject != nil,
			ptcq.withTaskColumn != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProjectTaskColumn{config: ptcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ptcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptcq.withProject; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTaskColumn)
		for i := range nodes {
			fk := nodes[i].ProjectID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(project.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Project = n
			}
		}
	}

	if query := ptcq.withTaskColumn; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTaskColumn)
		for i := range nodes {
			fk := nodes[i].TaskColumnID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(taskcolumn.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_column_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TaskColumn = n
			}
		}
	}

	return nodes, nil
}

func (ptcq *ProjectTaskColumnQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptcq.querySpec()
	return sqlgraph.CountNodes(ctx, ptcq.driver, _spec)
}

func (ptcq *ProjectTaskColumnQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ptcq *ProjectTaskColumnQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskcolumn.Table,
			Columns: projecttaskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskcolumn.FieldID,
			},
		},
		From:   ptcq.sql,
		Unique: true,
	}
	if unique := ptcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ptcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttaskcolumn.FieldID)
		for i := range fields {
			if fields[i] != projecttaskcolumn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptcq *ProjectTaskColumnQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptcq.driver.Dialect())
	t1 := builder.Table(projecttaskcolumn.Table)
	columns := ptcq.fields
	if len(columns) == 0 {
		columns = projecttaskcolumn.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptcq.sql != nil {
		selector = ptcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ptcq.predicates {
		p(selector)
	}
	for _, p := range ptcq.order {
		p(selector)
	}
	if offset := ptcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectTaskColumnGroupBy is the group-by builder for ProjectTaskColumn entities.
type ProjectTaskColumnGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptcgb *ProjectTaskColumnGroupBy) Aggregate(fns ...AggregateFunc) *ProjectTaskColumnGroupBy {
	ptcgb.fns = append(ptcgb.fns, fns...)
	return ptcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptcgb *ProjectTaskColumnGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptcgb.path(ctx)
	if err != nil {
		return err
	}
	ptcgb.sql = query
	return ptcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) StringX(ctx context.Context) string {
	v, err := ptcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) IntX(ctx context.Context) int {
	v, err := ptcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptcgb *ProjectTaskColumnGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptcgb *ProjectTaskColumnGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptcgb *ProjectTaskColumnGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptcgb.fields {
		if !projecttaskcolumn.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptcgb *ProjectTaskColumnGroupBy) sqlQuery() *sql.Selector {
	selector := ptcgb.sql.Select()
	aggregation := make([]string, 0, len(ptcgb.fns))
	for _, fn := range ptcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ptcgb.fields)+len(ptcgb.fns))
		for _, f := range ptcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ptcgb.fields...)...)
}

// ProjectTaskColumnSelect is the builder for selecting fields of ProjectTaskColumn entities.
type ProjectTaskColumnSelect struct {
	*ProjectTaskColumnQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ptcs *ProjectTaskColumnSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ptcs.prepareQuery(ctx); err != nil {
		return err
	}
	ptcs.sql = ptcs.ProjectTaskColumnQuery.sqlQuery(ctx)
	return ptcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ptcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ptcs.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ptcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) StringsX(ctx context.Context) []string {
	v, err := ptcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) StringX(ctx context.Context) string {
	v, err := ptcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ptcs.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ptcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) IntsX(ctx context.Context) []int {
	v, err := ptcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) IntX(ctx context.Context) int {
	v, err := ptcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptcs.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ptcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ptcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) Float64X(ctx context.Context) float64 {
	v, err := ptcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ptcs.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskColumnSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ptcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) BoolsX(ctx context.Context) []bool {
	v, err := ptcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ptcs *ProjectTaskColumnSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskcolumn.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskColumnSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptcs *ProjectTaskColumnSelect) BoolX(ctx context.Context) bool {
	v, err := ptcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptcs *ProjectTaskColumnSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ptcs.sql.Query()
	if err := ptcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
