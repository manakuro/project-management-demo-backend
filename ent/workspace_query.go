// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceQuery is the builder for querying Workspace entities.
type WorkspaceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Workspace
	// eager-loading edges.
	withTeammate           *TeammateQuery
	withProjects           *ProjectQuery
	withWorkspaceTeammates *WorkspaceTeammateQuery
	withFavoriteWorkspaces *FavoriteWorkspaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkspaceQuery builder.
func (wq *WorkspaceQuery) Where(ps ...predicate.Workspace) *WorkspaceQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WorkspaceQuery) Limit(limit int) *WorkspaceQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WorkspaceQuery) Offset(offset int) *WorkspaceQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkspaceQuery) Unique(unique bool) *WorkspaceQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *WorkspaceQuery) Order(o ...OrderFunc) *WorkspaceQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryTeammate chains the current query on the "teammate" edge.
func (wq *WorkspaceQuery) QueryTeammate() *TeammateQuery {
	query := &TeammateQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, selector),
			sqlgraph.To(teammate.Table, teammate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspace.TeammateTable, workspace.TeammateColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjects chains the current query on the "projects" edge.
func (wq *WorkspaceQuery) QueryProjects() *ProjectQuery {
	query := &ProjectQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.ProjectsTable, workspace.ProjectsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkspaceTeammates chains the current query on the "workspace_teammates" edge.
func (wq *WorkspaceQuery) QueryWorkspaceTeammates() *WorkspaceTeammateQuery {
	query := &WorkspaceTeammateQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, selector),
			sqlgraph.To(workspaceteammate.Table, workspaceteammate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.WorkspaceTeammatesTable, workspace.WorkspaceTeammatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFavoriteWorkspaces chains the current query on the "favorite_workspaces" edge.
func (wq *WorkspaceQuery) QueryFavoriteWorkspaces() *FavoriteWorkspaceQuery {
	query := &FavoriteWorkspaceQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, selector),
			sqlgraph.To(favoriteworkspace.Table, favoriteworkspace.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.FavoriteWorkspacesTable, workspace.FavoriteWorkspacesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Workspace entity from the query.
// Returns a *NotFoundError when no Workspace was found.
func (wq *WorkspaceQuery) First(ctx context.Context) (*Workspace, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workspace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkspaceQuery) FirstX(ctx context.Context) *Workspace {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Workspace ID from the query.
// Returns a *NotFoundError when no Workspace ID was found.
func (wq *WorkspaceQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workspace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkspaceQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Workspace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Workspace entity is not found.
// Returns a *NotFoundError when no Workspace entities are found.
func (wq *WorkspaceQuery) Only(ctx context.Context) (*Workspace, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workspace.Label}
	default:
		return nil, &NotSingularError{workspace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkspaceQuery) OnlyX(ctx context.Context) *Workspace {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Workspace ID in the query.
// Returns a *NotSingularError when exactly one Workspace ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkspaceQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = &NotSingularError{workspace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkspaceQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workspaces.
func (wq *WorkspaceQuery) All(ctx context.Context) ([]*Workspace, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkspaceQuery) AllX(ctx context.Context) []*Workspace {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Workspace IDs.
func (wq *WorkspaceQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := wq.Select(workspace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkspaceQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkspaceQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkspaceQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkspaceQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkspaceQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkspaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkspaceQuery) Clone() *WorkspaceQuery {
	if wq == nil {
		return nil
	}
	return &WorkspaceQuery{
		config:                 wq.config,
		limit:                  wq.limit,
		offset:                 wq.offset,
		order:                  append([]OrderFunc{}, wq.order...),
		predicates:             append([]predicate.Workspace{}, wq.predicates...),
		withTeammate:           wq.withTeammate.Clone(),
		withProjects:           wq.withProjects.Clone(),
		withWorkspaceTeammates: wq.withWorkspaceTeammates.Clone(),
		withFavoriteWorkspaces: wq.withFavoriteWorkspaces.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithTeammate tells the query-builder to eager-load the nodes that are connected to
// the "teammate" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkspaceQuery) WithTeammate(opts ...func(*TeammateQuery)) *WorkspaceQuery {
	query := &TeammateQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withTeammate = query
	return wq
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkspaceQuery) WithProjects(opts ...func(*ProjectQuery)) *WorkspaceQuery {
	query := &ProjectQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withProjects = query
	return wq
}

// WithWorkspaceTeammates tells the query-builder to eager-load the nodes that are connected to
// the "workspace_teammates" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkspaceQuery) WithWorkspaceTeammates(opts ...func(*WorkspaceTeammateQuery)) *WorkspaceQuery {
	query := &WorkspaceTeammateQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkspaceTeammates = query
	return wq
}

// WithFavoriteWorkspaces tells the query-builder to eager-load the nodes that are connected to
// the "favorite_workspaces" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkspaceQuery) WithFavoriteWorkspaces(opts ...func(*FavoriteWorkspaceQuery)) *WorkspaceQuery {
	query := &FavoriteWorkspaceQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withFavoriteWorkspaces = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy ulid.ID `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Workspace.Query().
//		GroupBy(workspace.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WorkspaceQuery) GroupBy(field string, fields ...string) *WorkspaceGroupBy {
	group := &WorkspaceGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy ulid.ID `json:"created_by,omitempty"`
//	}
//
//	client.Workspace.Query().
//		Select(workspace.FieldCreatedBy).
//		Scan(ctx, &v)
//
func (wq *WorkspaceQuery) Select(fields ...string) *WorkspaceSelect {
	wq.fields = append(wq.fields, fields...)
	return &WorkspaceSelect{WorkspaceQuery: wq}
}

func (wq *WorkspaceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !workspace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkspaceQuery) sqlAll(ctx context.Context) ([]*Workspace, error) {
	var (
		nodes       = []*Workspace{}
		_spec       = wq.querySpec()
		loadedTypes = [4]bool{
			wq.withTeammate != nil,
			wq.withProjects != nil,
			wq.withWorkspaceTeammates != nil,
			wq.withFavoriteWorkspaces != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Workspace{config: wq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wq.withTeammate; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*Workspace)
		for i := range nodes {
			fk := nodes[i].CreatedBy
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
				return nil, fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Teammate = n
			}
		}
	}

	if query := wq.withProjects; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*Workspace)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Projects = []*Project{}
		}
		query.Where(predicate.Project(func(s *sql.Selector) {
			s.Where(sql.InValues(workspace.ProjectsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.WorkspaceID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workspace_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Projects = append(node.Edges.Projects, n)
		}
	}

	if query := wq.withWorkspaceTeammates; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*Workspace)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.WorkspaceTeammates = []*WorkspaceTeammate{}
		}
		query.Where(predicate.WorkspaceTeammate(func(s *sql.Selector) {
			s.Where(sql.InValues(workspace.WorkspaceTeammatesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.WorkspaceID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workspace_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.WorkspaceTeammates = append(node.Edges.WorkspaceTeammates, n)
		}
	}

	if query := wq.withFavoriteWorkspaces; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[ulid.ID]*Workspace)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FavoriteWorkspaces = []*FavoriteWorkspace{}
		}
		query.Where(predicate.FavoriteWorkspace(func(s *sql.Selector) {
			s.Where(sql.InValues(workspace.FavoriteWorkspacesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.WorkspaceID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workspace_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.FavoriteWorkspaces = append(node.Edges.FavoriteWorkspaces, n)
		}
	}

	return nodes, nil
}

func (wq *WorkspaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkspaceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wq *WorkspaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspace.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspace.FieldID)
		for i := range fields {
			if fields[i] != workspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkspaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(workspace.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = workspace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkspaceGroupBy is the group-by builder for Workspace entities.
type WorkspaceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkspaceGroupBy) Aggregate(fns ...AggregateFunc) *WorkspaceGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *WorkspaceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkspaceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkspaceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkspaceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WorkspaceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WorkspaceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *WorkspaceGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *WorkspaceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !workspace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WorkspaceGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// WorkspaceSelect is the builder for selecting fields of Workspace entities.
type WorkspaceSelect struct {
	*WorkspaceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkspaceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.WorkspaceQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *WorkspaceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkspaceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *WorkspaceSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *WorkspaceSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkspaceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *WorkspaceSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *WorkspaceSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkspaceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *WorkspaceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *WorkspaceSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WorkspaceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *WorkspaceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ws *WorkspaceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspace.Label}
	default:
		err = fmt.Errorf("ent: WorkspaceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *WorkspaceSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *WorkspaceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
