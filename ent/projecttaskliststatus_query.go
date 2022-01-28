// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/tasklistsortstatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskListStatusQuery is the builder for querying ProjectTaskListStatus entities.
type ProjectTaskListStatusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProjectTaskListStatus
	// eager-loading edges.
	withProject                 *ProjectQuery
	withTaskListCompletedStatus *TaskListCompletedStatusQuery
	withTaskListSortStatus      *TaskListSortStatusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectTaskListStatusQuery builder.
func (ptlsq *ProjectTaskListStatusQuery) Where(ps ...predicate.ProjectTaskListStatus) *ProjectTaskListStatusQuery {
	ptlsq.predicates = append(ptlsq.predicates, ps...)
	return ptlsq
}

// Limit adds a limit step to the query.
func (ptlsq *ProjectTaskListStatusQuery) Limit(limit int) *ProjectTaskListStatusQuery {
	ptlsq.limit = &limit
	return ptlsq
}

// Offset adds an offset step to the query.
func (ptlsq *ProjectTaskListStatusQuery) Offset(offset int) *ProjectTaskListStatusQuery {
	ptlsq.offset = &offset
	return ptlsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptlsq *ProjectTaskListStatusQuery) Unique(unique bool) *ProjectTaskListStatusQuery {
	ptlsq.unique = &unique
	return ptlsq
}

// Order adds an order step to the query.
func (ptlsq *ProjectTaskListStatusQuery) Order(o ...OrderFunc) *ProjectTaskListStatusQuery {
	ptlsq.order = append(ptlsq.order, o...)
	return ptlsq
}

// QueryProject chains the current query on the "project" edge.
func (ptlsq *ProjectTaskListStatusQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: ptlsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptlsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptlsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttaskliststatus.Table, projecttaskliststatus.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttaskliststatus.ProjectTable, projecttaskliststatus.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptlsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskListCompletedStatus chains the current query on the "task_list_completed_status" edge.
func (ptlsq *ProjectTaskListStatusQuery) QueryTaskListCompletedStatus() *TaskListCompletedStatusQuery {
	query := &TaskListCompletedStatusQuery{config: ptlsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptlsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptlsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttaskliststatus.Table, projecttaskliststatus.FieldID, selector),
			sqlgraph.To(tasklistcompletedstatus.Table, tasklistcompletedstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttaskliststatus.TaskListCompletedStatusTable, projecttaskliststatus.TaskListCompletedStatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptlsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskListSortStatus chains the current query on the "task_list_sort_status" edge.
func (ptlsq *ProjectTaskListStatusQuery) QueryTaskListSortStatus() *TaskListSortStatusQuery {
	query := &TaskListSortStatusQuery{config: ptlsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptlsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptlsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projecttaskliststatus.Table, projecttaskliststatus.FieldID, selector),
			sqlgraph.To(tasklistsortstatus.Table, tasklistsortstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projecttaskliststatus.TaskListSortStatusTable, projecttaskliststatus.TaskListSortStatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptlsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectTaskListStatus entity from the query.
// Returns a *NotFoundError when no ProjectTaskListStatus was found.
func (ptlsq *ProjectTaskListStatusQuery) First(ctx context.Context) (*ProjectTaskListStatus, error) {
	nodes, err := ptlsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projecttaskliststatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) FirstX(ctx context.Context) *ProjectTaskListStatus {
	node, err := ptlsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectTaskListStatus ID from the query.
// Returns a *NotFoundError when no ProjectTaskListStatus ID was found.
func (ptlsq *ProjectTaskListStatusQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptlsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projecttaskliststatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := ptlsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectTaskListStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProjectTaskListStatus entity is not found.
// Returns a *NotFoundError when no ProjectTaskListStatus entities are found.
func (ptlsq *ProjectTaskListStatusQuery) Only(ctx context.Context) (*ProjectTaskListStatus, error) {
	nodes, err := ptlsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projecttaskliststatus.Label}
	default:
		return nil, &NotSingularError{projecttaskliststatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) OnlyX(ctx context.Context) *ProjectTaskListStatus {
	node, err := ptlsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectTaskListStatus ID in the query.
// Returns a *NotSingularError when exactly one ProjectTaskListStatus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ptlsq *ProjectTaskListStatusQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ptlsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = &NotSingularError{projecttaskliststatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := ptlsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectTaskListStatusSlice.
func (ptlsq *ProjectTaskListStatusQuery) All(ctx context.Context) ([]*ProjectTaskListStatus, error) {
	if err := ptlsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptlsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) AllX(ctx context.Context) []*ProjectTaskListStatus {
	nodes, err := ptlsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectTaskListStatus IDs.
func (ptlsq *ProjectTaskListStatusQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := ptlsq.Select(projecttaskliststatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := ptlsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptlsq *ProjectTaskListStatusQuery) Count(ctx context.Context) (int, error) {
	if err := ptlsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptlsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) CountX(ctx context.Context) int {
	count, err := ptlsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptlsq *ProjectTaskListStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptlsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptlsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptlsq *ProjectTaskListStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := ptlsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectTaskListStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptlsq *ProjectTaskListStatusQuery) Clone() *ProjectTaskListStatusQuery {
	if ptlsq == nil {
		return nil
	}
	return &ProjectTaskListStatusQuery{
		config:                      ptlsq.config,
		limit:                       ptlsq.limit,
		offset:                      ptlsq.offset,
		order:                       append([]OrderFunc{}, ptlsq.order...),
		predicates:                  append([]predicate.ProjectTaskListStatus{}, ptlsq.predicates...),
		withProject:                 ptlsq.withProject.Clone(),
		withTaskListCompletedStatus: ptlsq.withTaskListCompletedStatus.Clone(),
		withTaskListSortStatus:      ptlsq.withTaskListSortStatus.Clone(),
		// clone intermediate query.
		sql:  ptlsq.sql.Clone(),
		path: ptlsq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (ptlsq *ProjectTaskListStatusQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectTaskListStatusQuery {
	query := &ProjectQuery{config: ptlsq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptlsq.withProject = query
	return ptlsq
}

// WithTaskListCompletedStatus tells the query-builder to eager-load the nodes that are connected to
// the "task_list_completed_status" edge. The optional arguments are used to configure the query builder of the edge.
func (ptlsq *ProjectTaskListStatusQuery) WithTaskListCompletedStatus(opts ...func(*TaskListCompletedStatusQuery)) *ProjectTaskListStatusQuery {
	query := &TaskListCompletedStatusQuery{config: ptlsq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptlsq.withTaskListCompletedStatus = query
	return ptlsq
}

// WithTaskListSortStatus tells the query-builder to eager-load the nodes that are connected to
// the "task_list_sort_status" edge. The optional arguments are used to configure the query builder of the edge.
func (ptlsq *ProjectTaskListStatusQuery) WithTaskListSortStatus(opts ...func(*TaskListSortStatusQuery)) *ProjectTaskListStatusQuery {
	query := &TaskListSortStatusQuery{config: ptlsq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptlsq.withTaskListSortStatus = query
	return ptlsq
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
//	client.ProjectTaskListStatus.Query().
//		GroupBy(projecttaskliststatus.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptlsq *ProjectTaskListStatusQuery) GroupBy(field string, fields ...string) *ProjectTaskListStatusGroupBy {
	group := &ProjectTaskListStatusGroupBy{config: ptlsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptlsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptlsq.sqlQuery(ctx), nil
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
//	client.ProjectTaskListStatus.Query().
//		Select(projecttaskliststatus.FieldProjectID).
//		Scan(ctx, &v)
//
func (ptlsq *ProjectTaskListStatusQuery) Select(fields ...string) *ProjectTaskListStatusSelect {
	ptlsq.fields = append(ptlsq.fields, fields...)
	return &ProjectTaskListStatusSelect{ProjectTaskListStatusQuery: ptlsq}
}

func (ptlsq *ProjectTaskListStatusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptlsq.fields {
		if !projecttaskliststatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptlsq.path != nil {
		prev, err := ptlsq.path(ctx)
		if err != nil {
			return err
		}
		ptlsq.sql = prev
	}
	return nil
}

func (ptlsq *ProjectTaskListStatusQuery) sqlAll(ctx context.Context) ([]*ProjectTaskListStatus, error) {
	var (
		nodes       = []*ProjectTaskListStatus{}
		_spec       = ptlsq.querySpec()
		loadedTypes = [3]bool{
			ptlsq.withProject != nil,
			ptlsq.withTaskListCompletedStatus != nil,
			ptlsq.withTaskListSortStatus != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProjectTaskListStatus{config: ptlsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ptlsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptlsq.withProject; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTaskListStatus)
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

	if query := ptlsq.withTaskListCompletedStatus; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTaskListStatus)
		for i := range nodes {
			fk := nodes[i].TaskListCompletedStatusID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tasklistcompletedstatus.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_list_completed_status_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TaskListCompletedStatus = n
			}
		}
	}

	if query := ptlsq.withTaskListSortStatus; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ProjectTaskListStatus)
		for i := range nodes {
			fk := nodes[i].TaskListSortStatusID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tasklistsortstatus.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "task_list_sort_status_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TaskListSortStatus = n
			}
		}
	}

	return nodes, nil
}

func (ptlsq *ProjectTaskListStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptlsq.querySpec()
	return sqlgraph.CountNodes(ctx, ptlsq.driver, _spec)
}

func (ptlsq *ProjectTaskListStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptlsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ptlsq *ProjectTaskListStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskliststatus.Table,
			Columns: projecttaskliststatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskliststatus.FieldID,
			},
		},
		From:   ptlsq.sql,
		Unique: true,
	}
	if unique := ptlsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ptlsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttaskliststatus.FieldID)
		for i := range fields {
			if fields[i] != projecttaskliststatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptlsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptlsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptlsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptlsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptlsq *ProjectTaskListStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptlsq.driver.Dialect())
	t1 := builder.Table(projecttaskliststatus.Table)
	columns := ptlsq.fields
	if len(columns) == 0 {
		columns = projecttaskliststatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptlsq.sql != nil {
		selector = ptlsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ptlsq.predicates {
		p(selector)
	}
	for _, p := range ptlsq.order {
		p(selector)
	}
	if offset := ptlsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptlsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectTaskListStatusGroupBy is the group-by builder for ProjectTaskListStatus entities.
type ProjectTaskListStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptlsgb *ProjectTaskListStatusGroupBy) Aggregate(fns ...AggregateFunc) *ProjectTaskListStatusGroupBy {
	ptlsgb.fns = append(ptlsgb.fns, fns...)
	return ptlsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptlsgb *ProjectTaskListStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptlsgb.path(ctx)
	if err != nil {
		return err
	}
	ptlsgb.sql = query
	return ptlsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptlsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptlsgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptlsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptlsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptlsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) StringX(ctx context.Context) string {
	v, err := ptlsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptlsgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptlsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptlsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptlsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) IntX(ctx context.Context) int {
	v, err := ptlsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptlsgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptlsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptlsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptlsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptlsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptlsgb.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptlsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptlsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptlsgb *ProjectTaskListStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptlsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptlsgb *ProjectTaskListStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptlsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptlsgb *ProjectTaskListStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptlsgb.fields {
		if !projecttaskliststatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptlsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptlsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptlsgb *ProjectTaskListStatusGroupBy) sqlQuery() *sql.Selector {
	selector := ptlsgb.sql.Select()
	aggregation := make([]string, 0, len(ptlsgb.fns))
	for _, fn := range ptlsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ptlsgb.fields)+len(ptlsgb.fns))
		for _, f := range ptlsgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ptlsgb.fields...)...)
}

// ProjectTaskListStatusSelect is the builder for selecting fields of ProjectTaskListStatus entities.
type ProjectTaskListStatusSelect struct {
	*ProjectTaskListStatusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ptlss *ProjectTaskListStatusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ptlss.prepareQuery(ctx); err != nil {
		return err
	}
	ptlss.sql = ptlss.ProjectTaskListStatusQuery.sqlQuery(ctx)
	return ptlss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ptlss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ptlss.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ptlss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) StringsX(ctx context.Context) []string {
	v, err := ptlss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptlss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) StringX(ctx context.Context) string {
	v, err := ptlss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ptlss.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ptlss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) IntsX(ctx context.Context) []int {
	v, err := ptlss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptlss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) IntX(ctx context.Context) int {
	v, err := ptlss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptlss.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ptlss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ptlss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptlss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := ptlss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ptlss.fields) > 1 {
		return nil, errors.New("ent: ProjectTaskListStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ptlss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := ptlss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ptlss *ProjectTaskListStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptlss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projecttaskliststatus.Label}
	default:
		err = fmt.Errorf("ent: ProjectTaskListStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptlss *ProjectTaskListStatusSelect) BoolX(ctx context.Context) bool {
	v, err := ptlss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptlss *ProjectTaskListStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ptlss.sql.Query()
	if err := ptlss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}