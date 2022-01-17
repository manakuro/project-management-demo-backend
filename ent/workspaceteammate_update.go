// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceTeammateUpdate is the builder for updating WorkspaceTeammate entities.
type WorkspaceTeammateUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceTeammateMutation
}

// Where appends a list predicates to the WorkspaceTeammateUpdate builder.
func (wtu *WorkspaceTeammateUpdate) Where(ps ...predicate.WorkspaceTeammate) *WorkspaceTeammateUpdate {
	wtu.mutation.Where(ps...)
	return wtu
}

// SetWorkspaceID sets the "workspace_id" field.
func (wtu *WorkspaceTeammateUpdate) SetWorkspaceID(u ulid.ID) *WorkspaceTeammateUpdate {
	wtu.mutation.SetWorkspaceID(u)
	return wtu
}

// SetTeammateID sets the "teammate_id" field.
func (wtu *WorkspaceTeammateUpdate) SetTeammateID(u ulid.ID) *WorkspaceTeammateUpdate {
	wtu.mutation.SetTeammateID(u)
	return wtu
}

// SetRole sets the "role" field.
func (wtu *WorkspaceTeammateUpdate) SetRole(s string) *WorkspaceTeammateUpdate {
	wtu.mutation.SetRole(s)
	return wtu
}

// SetIsOwner sets the "is_owner" field.
func (wtu *WorkspaceTeammateUpdate) SetIsOwner(b bool) *WorkspaceTeammateUpdate {
	wtu.mutation.SetIsOwner(b)
	return wtu
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wtu *WorkspaceTeammateUpdate) SetWorkspace(w *Workspace) *WorkspaceTeammateUpdate {
	return wtu.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wtu *WorkspaceTeammateUpdate) SetTeammate(t *Teammate) *WorkspaceTeammateUpdate {
	return wtu.SetTeammateID(t.ID)
}

// Mutation returns the WorkspaceTeammateMutation object of the builder.
func (wtu *WorkspaceTeammateUpdate) Mutation() *WorkspaceTeammateMutation {
	return wtu.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wtu *WorkspaceTeammateUpdate) ClearWorkspace() *WorkspaceTeammateUpdate {
	wtu.mutation.ClearWorkspace()
	return wtu
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (wtu *WorkspaceTeammateUpdate) ClearTeammate() *WorkspaceTeammateUpdate {
	wtu.mutation.ClearTeammate()
	return wtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wtu *WorkspaceTeammateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wtu.hooks) == 0 {
		if err = wtu.check(); err != nil {
			return 0, err
		}
		affected, err = wtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceTeammateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wtu.check(); err != nil {
				return 0, err
			}
			wtu.mutation = mutation
			affected, err = wtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wtu.hooks) - 1; i >= 0; i-- {
			if wtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wtu *WorkspaceTeammateUpdate) SaveX(ctx context.Context) int {
	affected, err := wtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wtu *WorkspaceTeammateUpdate) Exec(ctx context.Context) error {
	_, err := wtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtu *WorkspaceTeammateUpdate) ExecX(ctx context.Context) {
	if err := wtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wtu *WorkspaceTeammateUpdate) check() error {
	if v, ok := wtu.mutation.Role(); ok {
		if err := workspaceteammate.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := wtu.mutation.WorkspaceID(); wtu.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	if _, ok := wtu.mutation.TeammateID(); wtu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (wtu *WorkspaceTeammateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspaceteammate.Table,
			Columns: workspaceteammate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspaceteammate.FieldID,
			},
		},
	}
	if ps := wtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wtu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceteammate.FieldRole,
		})
	}
	if value, ok := wtu.mutation.IsOwner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workspaceteammate.FieldIsOwner,
		})
	}
	if wtu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.WorkspaceTable,
			Columns: []string{workspaceteammate.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.WorkspaceTable,
			Columns: []string{workspaceteammate.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wtu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.TeammateTable,
			Columns: []string{workspaceteammate.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.TeammateTable,
			Columns: []string{workspaceteammate.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceteammate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceTeammateUpdateOne is the builder for updating a single WorkspaceTeammate entity.
type WorkspaceTeammateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkspaceTeammateMutation
}

// SetWorkspaceID sets the "workspace_id" field.
func (wtuo *WorkspaceTeammateUpdateOne) SetWorkspaceID(u ulid.ID) *WorkspaceTeammateUpdateOne {
	wtuo.mutation.SetWorkspaceID(u)
	return wtuo
}

// SetTeammateID sets the "teammate_id" field.
func (wtuo *WorkspaceTeammateUpdateOne) SetTeammateID(u ulid.ID) *WorkspaceTeammateUpdateOne {
	wtuo.mutation.SetTeammateID(u)
	return wtuo
}

// SetRole sets the "role" field.
func (wtuo *WorkspaceTeammateUpdateOne) SetRole(s string) *WorkspaceTeammateUpdateOne {
	wtuo.mutation.SetRole(s)
	return wtuo
}

// SetIsOwner sets the "is_owner" field.
func (wtuo *WorkspaceTeammateUpdateOne) SetIsOwner(b bool) *WorkspaceTeammateUpdateOne {
	wtuo.mutation.SetIsOwner(b)
	return wtuo
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wtuo *WorkspaceTeammateUpdateOne) SetWorkspace(w *Workspace) *WorkspaceTeammateUpdateOne {
	return wtuo.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wtuo *WorkspaceTeammateUpdateOne) SetTeammate(t *Teammate) *WorkspaceTeammateUpdateOne {
	return wtuo.SetTeammateID(t.ID)
}

// Mutation returns the WorkspaceTeammateMutation object of the builder.
func (wtuo *WorkspaceTeammateUpdateOne) Mutation() *WorkspaceTeammateMutation {
	return wtuo.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wtuo *WorkspaceTeammateUpdateOne) ClearWorkspace() *WorkspaceTeammateUpdateOne {
	wtuo.mutation.ClearWorkspace()
	return wtuo
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (wtuo *WorkspaceTeammateUpdateOne) ClearTeammate() *WorkspaceTeammateUpdateOne {
	wtuo.mutation.ClearTeammate()
	return wtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wtuo *WorkspaceTeammateUpdateOne) Select(field string, fields ...string) *WorkspaceTeammateUpdateOne {
	wtuo.fields = append([]string{field}, fields...)
	return wtuo
}

// Save executes the query and returns the updated WorkspaceTeammate entity.
func (wtuo *WorkspaceTeammateUpdateOne) Save(ctx context.Context) (*WorkspaceTeammate, error) {
	var (
		err  error
		node *WorkspaceTeammate
	)
	if len(wtuo.hooks) == 0 {
		if err = wtuo.check(); err != nil {
			return nil, err
		}
		node, err = wtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceTeammateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wtuo.check(); err != nil {
				return nil, err
			}
			wtuo.mutation = mutation
			node, err = wtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wtuo.hooks) - 1; i >= 0; i-- {
			if wtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wtuo *WorkspaceTeammateUpdateOne) SaveX(ctx context.Context) *WorkspaceTeammate {
	node, err := wtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wtuo *WorkspaceTeammateUpdateOne) Exec(ctx context.Context) error {
	_, err := wtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtuo *WorkspaceTeammateUpdateOne) ExecX(ctx context.Context) {
	if err := wtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wtuo *WorkspaceTeammateUpdateOne) check() error {
	if v, ok := wtuo.mutation.Role(); ok {
		if err := workspaceteammate.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := wtuo.mutation.WorkspaceID(); wtuo.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	if _, ok := wtuo.mutation.TeammateID(); wtuo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (wtuo *WorkspaceTeammateUpdateOne) sqlSave(ctx context.Context) (_node *WorkspaceTeammate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspaceteammate.Table,
			Columns: workspaceteammate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspaceteammate.FieldID,
			},
		},
	}
	id, ok := wtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkspaceTeammate.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspaceteammate.FieldID)
		for _, f := range fields {
			if !workspaceteammate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspaceteammate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wtuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceteammate.FieldRole,
		})
	}
	if value, ok := wtuo.mutation.IsOwner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workspaceteammate.FieldIsOwner,
		})
	}
	if wtuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.WorkspaceTable,
			Columns: []string{workspaceteammate.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.WorkspaceTable,
			Columns: []string{workspaceteammate.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wtuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.TeammateTable,
			Columns: []string{workspaceteammate.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wtuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceteammate.TeammateTable,
			Columns: []string{workspaceteammate.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkspaceTeammate{config: wtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceteammate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
