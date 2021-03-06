// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/archivedtaskactivity"
	"project-management-demo-backend/ent/archivedtaskactivitytask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedTaskActivityQuery is the builder for querying ArchivedTaskActivity entities.
type ArchivedTaskActivityQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ArchivedTaskActivity
	// eager-loading edges.
	withTeammate                  *TeammateQuery
	withActivityType              *ActivityTypeQuery
	withWorkspace                 *WorkspaceQuery
	withArchivedTaskActivityTasks *ArchivedTaskActivityTaskQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArchivedTaskActivityQuery builder.
func (ataq *ArchivedTaskActivityQuery) Where(ps ...predicate.ArchivedTaskActivity) *ArchivedTaskActivityQuery {
	ataq.predicates = append(ataq.predicates, ps...)
	return ataq
}

// Limit adds a limit step to the query.
func (ataq *ArchivedTaskActivityQuery) Limit(limit int) *ArchivedTaskActivityQuery {
	ataq.limit = &limit
	return ataq
}

// Offset adds an offset step to the query.
func (ataq *ArchivedTaskActivityQuery) Offset(offset int) *ArchivedTaskActivityQuery {
	ataq.offset = &offset
	return ataq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ataq *ArchivedTaskActivityQuery) Unique(unique bool) *ArchivedTaskActivityQuery {
	ataq.unique = &unique
	return ataq
}

// Order adds an order step to the query.
func (ataq *ArchivedTaskActivityQuery) Order(o ...OrderFunc) *ArchivedTaskActivityQuery {
	ataq.order = append(ataq.order, o...)
	return ataq
}

// QueryTeammate chains the current query on the "teammate" edge.
func (ataq *ArchivedTaskActivityQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: ataq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ataq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ataq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivity.Table, archivedtaskactivity.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, archivedtaskactivity.TeammateTable, archivedtaskactivity.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(ataq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryActivityType chains the current query on the "activityType" edge.
func (ataq *ArchivedTaskActivityQuery) QueryActivityType() *ActivityTypeQuery {
	query := &ActivityTypeQuery{config: ataq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ataq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ataq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivity.Table, archivedtaskactivity.FieldID, selector),
			sqlgraph.To(activitytype.Table, activitytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, archivedtaskactivity.ActivityTypeTable, archivedtaskactivity.ActivityTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ataq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkspace chains the current query on the "workspace" edge.
func (ataq *ArchivedTaskActivityQuery) QueryWorkspace() *WorkspaceQuery {
	query := &WorkspaceQuery{config: ataq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ataq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ataq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivity.Table, archivedtaskactivity.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, archivedtaskactivity.WorkspaceTable, archivedtaskactivity.WorkspaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ataq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArchivedTaskActivityTasks chains the current query on the "archivedTaskActivityTasks" edge.
func (ataq *ArchivedTaskActivityQuery) QueryArchivedTaskActivityTasks() *ArchivedTaskActivityTaskQuery {
	query := &ArchivedTaskActivityTaskQuery{config: ataq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ataq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ataq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(archivedtaskactivity.Table, archivedtaskactivity.FieldID, selector),
			sqlgraph.To(archivedtaskactivitytask.Table, archivedtaskactivitytask.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, archivedtaskactivity.ArchivedTaskActivityTasksTable, archivedtaskactivity.ArchivedTaskActivityTasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(ataq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArchivedTaskActivity entity from the query.
// Returns a *NotFoundError when no ArchivedTaskActivity was found.
func (ataq *ArchivedTaskActivityQuery) First(ctx context.Context) (*ArchivedTaskActivity, error) {
	nodes, err := ataq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{archivedtaskactivity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) FirstX(ctx context.Context) *ArchivedTaskActivity {
	node, err := ataq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArchivedTaskActivity ID from the query.
// Returns a *NotFoundError when no ArchivedTaskActivity ID was found.
func (ataq *ArchivedTaskActivityQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ataq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{archivedtaskactivity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := ataq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArchivedTaskActivity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArchivedTaskActivity entity is found.
// Returns a *NotFoundError when no ArchivedTaskActivity entities are found.
func (ataq *ArchivedTaskActivityQuery) Only(ctx context.Context) (*ArchivedTaskActivity, error) {
	nodes, err := ataq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{archivedtaskactivity.Label}
	default:
		return nil, &NotSingularError{archivedtaskactivity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) OnlyX(ctx context.Context) *ArchivedTaskActivity {
	node, err := ataq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArchivedTaskActivity ID in the query.
// Returns a *NotSingularError when more than one ArchivedTaskActivity ID is found.
// Returns a *NotFoundError when no entities are found.
func (ataq *ArchivedTaskActivityQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = ataq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = &NotSingularError{archivedtaskactivity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := ataq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArchivedTaskActivities.
func (ataq *ArchivedTaskActivityQuery) All(ctx context.Context) ([]*ArchivedTaskActivity, error) {
	if err := ataq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ataq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) AllX(ctx context.Context) []*ArchivedTaskActivity {
	nodes, err := ataq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArchivedTaskActivity IDs.
func (ataq *ArchivedTaskActivityQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := ataq.Select(archivedtaskactivity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := ataq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ataq *ArchivedTaskActivityQuery) Count(ctx context.Context) (int, error) {
	if err := ataq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ataq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) CountX(ctx context.Context) int {
	count, err := ataq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ataq *ArchivedTaskActivityQuery) Exist(ctx context.Context) (bool, error) {
	if err := ataq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ataq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ataq *ArchivedTaskActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := ataq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArchivedTaskActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ataq *ArchivedTaskActivityQuery) Clone() *ArchivedTaskActivityQuery {
	if ataq == nil {
		return nil
	}
	return &ArchivedTaskActivityQuery{
		config:                        ataq.config,
		limit:                         ataq.limit,
		offset:                        ataq.offset,
		order:                         append([]OrderFunc{}, ataq.order...),
		predicates:                    append([]predicate.ArchivedTaskActivity{}, ataq.predicates...),
		withTeammate:                  ataq.withTeammate.Clone(),
		withActivityType:              ataq.withActivityType.Clone(),
		withWorkspace:                 ataq.withWorkspace.Clone(),
		withArchivedTaskActivityTasks: ataq.withArchivedTaskActivityTasks.Clone(),
		// clone intermediate query.
		sql:    ataq.sql.Clone(),
		path:   ataq.path,
		unique: ataq.unique,
	}
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (ataq *ArchivedTaskActivityQuery) WithTeammate(opts ...func(*TeammateQuery)) *ArchivedTaskActivityQuery {
	query := &TeammateQuery{config: ataq.config}
	for _, opt := range opts {
		opt(query)
	}
	ataq.withTeammate = query
	return ataq
}

// WithActivityType tells the query-builder to eager-load the nodes that are connected to
// the "activityType" edge. The optional arguments are used to configure the query builder of the edge.
func (ataq *ArchivedTaskActivityQuery) WithActivityType(opts ...func(*ActivityTypeQuery)) *ArchivedTaskActivityQuery {
	query := &ActivityTypeQuery{config: ataq.config}
	for _, opt := range opts {
		opt(query)
	}
	ataq.withActivityType = query
	return ataq
}

// WithWorkspace tells the query-builder to eager-load the nodes that are connected to
// the "workspace" edge. The optional arguments are used to configure the query builder of the edge.
func (ataq *ArchivedTaskActivityQuery) WithWorkspace(opts ...func(*WorkspaceQuery)) *ArchivedTaskActivityQuery {
	query := &WorkspaceQuery{config: ataq.config}
	for _, opt := range opts {
		opt(query)
	}
	ataq.withWorkspace = query
	return ataq
}

// WithArchivedTaskActivityTasks tells the query-builder to eager-load the nodes that are connected to
// the "archivedTaskActivityTasks" edge. The optional arguments are used to configure the query builder of the edge.
func (ataq *ArchivedTaskActivityQuery) WithArchivedTaskActivityTasks(opts ...func(*ArchivedTaskActivityTaskQuery)) *ArchivedTaskActivityQuery {
	query := &ArchivedTaskActivityTaskQuery{config: ataq.config}
	for _, opt := range opts {
		opt(query)
	}
	ataq.withArchivedTaskActivityTasks = query
	return ataq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ActivityTypeID ulid.ID `json:"activity_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArchivedTaskActivity.Query().
//		GroupBy(archivedtaskactivity.FieldActivityTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ataq *ArchivedTaskActivityQuery) GroupBy(field string, fields ...string) *ArchivedTaskActivityGroupBy {
	group := &ArchivedTaskActivityGroupBy{config: ataq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ataq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ataq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ActivityTypeID ulid.ID `json:"activity_type_id,omitempty"`
//	}
//
//	client.ArchivedTaskActivity.Query().
//		Select(archivedtaskactivity.FieldActivityTypeID).
//		Scan(ctx, &v)
//
func (ataq *ArchivedTaskActivityQuery) Select(fields ...string) *ArchivedTaskActivitySelect {
	ataq.fields = append(ataq.fields, fields...)
	return &ArchivedTaskActivitySelect{ArchivedTaskActivityQuery: ataq}
}

func (ataq *ArchivedTaskActivityQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ataq.fields {
		if !archivedtaskactivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ataq.path != nil {
		prev, err := ataq.path(ctx)
		if err != nil {
			return err
		}
		ataq.sql = prev
	}
	return nil
}

func (ataq *ArchivedTaskActivityQuery) sqlAll(ctx context.Context) ([]*ArchivedTaskActivity, error) {
	var (
		nodes       = []*ArchivedTaskActivity{}
		_spec       = ataq.querySpec()
		loadedTypes = [4]bool{
			ataq.withTeammate != nil,
			ataq.withActivityType != nil,
			ataq.withWorkspace != nil,
			ataq.withArchivedTaskActivityTasks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ArchivedTaskActivity{config: ataq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ataq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ataq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ArchivedTaskActivity)
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

	if query := ataq.withActivityType; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ArchivedTaskActivity)
		for i := range nodes {
			fk := nodes[i].ActivityTypeID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(activitytype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "activity_type_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ActivityType = n
			}
		}
	}

	if query := ataq.withWorkspace; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*ArchivedTaskActivity)
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

	if query := ataq.withArchivedTaskActivityTasks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*ArchivedTaskActivity)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ArchivedTaskActivityTasks = []*ArchivedTaskActivityTask{}
		}
		query.Where(predicate.ArchivedTaskActivityTask(func(s *sql.Selector) {
			s.Where(sql.InValues(archivedtaskactivity.ArchivedTaskActivityTasksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ArchivedTaskActivityID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "archived_task_activity_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ArchivedTaskActivityTasks = append(node.Edges.ArchivedTaskActivityTasks, n)
		}
	}

	return nodes, nil
}

func (ataq *ArchivedTaskActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ataq.querySpec()
	_spec.Node.Columns = ataq.fields
	if len(ataq.fields) > 0 {
		_spec.Unique = ataq.unique != nil && *ataq.unique
	}
	return sqlgraph.CountNodes(ctx, ataq.driver, _spec)
}

func (ataq *ArchivedTaskActivityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ataq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ataq *ArchivedTaskActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedtaskactivity.Table,
			Columns: archivedtaskactivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedtaskactivity.FieldID,
			},
		},
		From:   ataq.sql,
		Unique: true,
	}
	if unique := ataq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ataq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, archivedtaskactivity.FieldID)
		for i := range fields {
			if fields[i] != archivedtaskactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ataq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ataq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ataq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ataq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ataq *ArchivedTaskActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ataq.driver.Dialect())
	t1 := builder.Table(archivedtaskactivity.Table)
	columns := ataq.fields
	if len(columns) == 0 {
		columns = archivedtaskactivity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ataq.sql != nil {
		selector = ataq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ataq.unique != nil && *ataq.unique {
		selector.Distinct()
	}
	for _, p := range ataq.predicates {
		p(selector)
	}
	for _, p := range ataq.order {
		p(selector)
	}
	if offset := ataq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ataq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArchivedTaskActivityGroupBy is the group-by builder for ArchivedTaskActivity entities.
type ArchivedTaskActivityGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atagb *ArchivedTaskActivityGroupBy) Aggregate(fns ...AggregateFunc) *ArchivedTaskActivityGroupBy {
	atagb.fns = append(atagb.fns, fns...)
	return atagb
}

// Scan applies the group-by query and scans the result into the given value.
func (atagb *ArchivedTaskActivityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := atagb.path(ctx)
	if err != nil {
		return err
	}
	atagb.sql = query
	return atagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := atagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(atagb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := atagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) StringsX(ctx context.Context) []string {
	v, err := atagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) StringX(ctx context.Context) string {
	v, err := atagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(atagb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := atagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) IntsX(ctx context.Context) []int {
	v, err := atagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) IntX(ctx context.Context) int {
	v, err := atagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(atagb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := atagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := atagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) Float64X(ctx context.Context) float64 {
	v, err := atagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(atagb.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivityGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := atagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := atagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atagb *ArchivedTaskActivityGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivityGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atagb *ArchivedTaskActivityGroupBy) BoolX(ctx context.Context) bool {
	v, err := atagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atagb *ArchivedTaskActivityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range atagb.fields {
		if !archivedtaskactivity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atagb *ArchivedTaskActivityGroupBy) sqlQuery() *sql.Selector {
	selector := atagb.sql.Select()
	aggregation := make([]string, 0, len(atagb.fns))
	for _, fn := range atagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(atagb.fields)+len(atagb.fns))
		for _, f := range atagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(atagb.fields...)...)
}

// ArchivedTaskActivitySelect is the builder for selecting fields of ArchivedTaskActivity entities.
type ArchivedTaskActivitySelect struct {
	*ArchivedTaskActivityQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (atas *ArchivedTaskActivitySelect) Scan(ctx context.Context, v interface{}) error {
	if err := atas.prepareQuery(ctx); err != nil {
		return err
	}
	atas.sql = atas.ArchivedTaskActivityQuery.sqlQuery(ctx)
	return atas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) ScanX(ctx context.Context, v interface{}) {
	if err := atas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Strings(ctx context.Context) ([]string, error) {
	if len(atas.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivitySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := atas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) StringsX(ctx context.Context) []string {
	v, err := atas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivitySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) StringX(ctx context.Context) string {
	v, err := atas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Ints(ctx context.Context) ([]int, error) {
	if len(atas.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivitySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := atas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) IntsX(ctx context.Context) []int {
	v, err := atas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivitySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) IntX(ctx context.Context) int {
	v, err := atas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(atas.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivitySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := atas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) Float64sX(ctx context.Context) []float64 {
	v, err := atas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivitySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) Float64X(ctx context.Context) float64 {
	v, err := atas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(atas.fields) > 1 {
		return nil, errors.New("ent: ArchivedTaskActivitySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := atas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) BoolsX(ctx context.Context) []bool {
	v, err := atas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (atas *ArchivedTaskActivitySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{archivedtaskactivity.Label}
	default:
		err = fmt.Errorf("ent: ArchivedTaskActivitySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atas *ArchivedTaskActivitySelect) BoolX(ctx context.Context) bool {
	v, err := atas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atas *ArchivedTaskActivitySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := atas.sql.Query()
	if err := atas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
