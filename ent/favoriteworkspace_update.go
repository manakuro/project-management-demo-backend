// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteWorkspaceUpdate is the builder for updating FavoriteWorkspace entities.
type FavoriteWorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *FavoriteWorkspaceMutation
}

// Where appends a list predicates to the FavoriteWorkspaceUpdate builder.
func (fwu *FavoriteWorkspaceUpdate) Where(ps ...predicate.FavoriteWorkspace) *FavoriteWorkspaceUpdate {
	fwu.mutation.Where(ps...)
	return fwu
}

// SetWorkspaceID sets the "workspace_id" field.
func (fwu *FavoriteWorkspaceUpdate) SetWorkspaceID(u ulid.ID) *FavoriteWorkspaceUpdate {
	fwu.mutation.SetWorkspaceID(u)
	return fwu
}

// SetTeammateID sets the "teammate_id" field.
func (fwu *FavoriteWorkspaceUpdate) SetTeammateID(u ulid.ID) *FavoriteWorkspaceUpdate {
	fwu.mutation.SetTeammateID(u)
	return fwu
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (fwu *FavoriteWorkspaceUpdate) SetWorkspace(w *Workspace) *FavoriteWorkspaceUpdate {
	return fwu.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (fwu *FavoriteWorkspaceUpdate) SetTeammate(t *Teammate) *FavoriteWorkspaceUpdate {
	return fwu.SetTeammateID(t.ID)
}

// Mutation returns the FavoriteWorkspaceMutation object of the builder.
func (fwu *FavoriteWorkspaceUpdate) Mutation() *FavoriteWorkspaceMutation {
	return fwu.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (fwu *FavoriteWorkspaceUpdate) ClearWorkspace() *FavoriteWorkspaceUpdate {
	fwu.mutation.ClearWorkspace()
	return fwu
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (fwu *FavoriteWorkspaceUpdate) ClearTeammate() *FavoriteWorkspaceUpdate {
	fwu.mutation.ClearTeammate()
	return fwu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fwu *FavoriteWorkspaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fwu.hooks) == 0 {
		if err = fwu.check(); err != nil {
			return 0, err
		}
		affected, err = fwu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteWorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fwu.check(); err != nil {
				return 0, err
			}
			fwu.mutation = mutation
			affected, err = fwu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fwu.hooks) - 1; i >= 0; i-- {
			if fwu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fwu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fwu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fwu *FavoriteWorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := fwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fwu *FavoriteWorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := fwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fwu *FavoriteWorkspaceUpdate) ExecX(ctx context.Context) {
	if err := fwu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fwu *FavoriteWorkspaceUpdate) check() error {
	if _, ok := fwu.mutation.WorkspaceID(); fwu.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FavoriteWorkspace.workspace"`)
	}
	if _, ok := fwu.mutation.TeammateID(); fwu.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FavoriteWorkspace.teammate"`)
	}
	return nil
}

func (fwu *FavoriteWorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favoriteworkspace.Table,
			Columns: favoriteworkspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteworkspace.FieldID,
			},
		},
	}
	if ps := fwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fwu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.WorkspaceTable,
			Columns: []string{favoriteworkspace.WorkspaceColumn},
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
	if nodes := fwu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.WorkspaceTable,
			Columns: []string{favoriteworkspace.WorkspaceColumn},
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
	if fwu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.TeammateTable,
			Columns: []string{favoriteworkspace.TeammateColumn},
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
	if nodes := fwu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.TeammateTable,
			Columns: []string{favoriteworkspace.TeammateColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favoriteworkspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FavoriteWorkspaceUpdateOne is the builder for updating a single FavoriteWorkspace entity.
type FavoriteWorkspaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FavoriteWorkspaceMutation
}

// SetWorkspaceID sets the "workspace_id" field.
func (fwuo *FavoriteWorkspaceUpdateOne) SetWorkspaceID(u ulid.ID) *FavoriteWorkspaceUpdateOne {
	fwuo.mutation.SetWorkspaceID(u)
	return fwuo
}

// SetTeammateID sets the "teammate_id" field.
func (fwuo *FavoriteWorkspaceUpdateOne) SetTeammateID(u ulid.ID) *FavoriteWorkspaceUpdateOne {
	fwuo.mutation.SetTeammateID(u)
	return fwuo
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (fwuo *FavoriteWorkspaceUpdateOne) SetWorkspace(w *Workspace) *FavoriteWorkspaceUpdateOne {
	return fwuo.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (fwuo *FavoriteWorkspaceUpdateOne) SetTeammate(t *Teammate) *FavoriteWorkspaceUpdateOne {
	return fwuo.SetTeammateID(t.ID)
}

// Mutation returns the FavoriteWorkspaceMutation object of the builder.
func (fwuo *FavoriteWorkspaceUpdateOne) Mutation() *FavoriteWorkspaceMutation {
	return fwuo.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (fwuo *FavoriteWorkspaceUpdateOne) ClearWorkspace() *FavoriteWorkspaceUpdateOne {
	fwuo.mutation.ClearWorkspace()
	return fwuo
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (fwuo *FavoriteWorkspaceUpdateOne) ClearTeammate() *FavoriteWorkspaceUpdateOne {
	fwuo.mutation.ClearTeammate()
	return fwuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fwuo *FavoriteWorkspaceUpdateOne) Select(field string, fields ...string) *FavoriteWorkspaceUpdateOne {
	fwuo.fields = append([]string{field}, fields...)
	return fwuo
}

// Save executes the query and returns the updated FavoriteWorkspace entity.
func (fwuo *FavoriteWorkspaceUpdateOne) Save(ctx context.Context) (*FavoriteWorkspace, error) {
	var (
		err  error
		node *FavoriteWorkspace
	)
	if len(fwuo.hooks) == 0 {
		if err = fwuo.check(); err != nil {
			return nil, err
		}
		node, err = fwuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteWorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fwuo.check(); err != nil {
				return nil, err
			}
			fwuo.mutation = mutation
			node, err = fwuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fwuo.hooks) - 1; i >= 0; i-- {
			if fwuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fwuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fwuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fwuo *FavoriteWorkspaceUpdateOne) SaveX(ctx context.Context) *FavoriteWorkspace {
	node, err := fwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fwuo *FavoriteWorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := fwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fwuo *FavoriteWorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := fwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fwuo *FavoriteWorkspaceUpdateOne) check() error {
	if _, ok := fwuo.mutation.WorkspaceID(); fwuo.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FavoriteWorkspace.workspace"`)
	}
	if _, ok := fwuo.mutation.TeammateID(); fwuo.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FavoriteWorkspace.teammate"`)
	}
	return nil
}

func (fwuo *FavoriteWorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *FavoriteWorkspace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favoriteworkspace.Table,
			Columns: favoriteworkspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteworkspace.FieldID,
			},
		},
	}
	id, ok := fwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FavoriteWorkspace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favoriteworkspace.FieldID)
		for _, f := range fields {
			if !favoriteworkspace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != favoriteworkspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fwuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.WorkspaceTable,
			Columns: []string{favoriteworkspace.WorkspaceColumn},
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
	if nodes := fwuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.WorkspaceTable,
			Columns: []string{favoriteworkspace.WorkspaceColumn},
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
	if fwuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.TeammateTable,
			Columns: []string{favoriteworkspace.TeammateColumn},
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
	if nodes := fwuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteworkspace.TeammateTable,
			Columns: []string{favoriteworkspace.TeammateColumn},
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
	_node = &FavoriteWorkspace{config: fwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favoriteworkspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
