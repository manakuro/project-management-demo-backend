// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IconUpdate is the builder for updating Icon entities.
type IconUpdate struct {
	config
	hooks    []Hook
	mutation *IconMutation
}

// Where appends a list predicates to the IconUpdate builder.
func (iu *IconUpdate) Where(ps ...predicate.Icon) *IconUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *IconUpdate) SetName(s string) *IconUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetIcon sets the "icon" field.
func (iu *IconUpdate) SetIcon(s string) *IconUpdate {
	iu.mutation.SetIcon(s)
	return iu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (iu *IconUpdate) SetProjectID(id ulid.ID) *IconUpdate {
	iu.mutation.SetProjectID(id)
	return iu
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (iu *IconUpdate) SetNillableProjectID(id *ulid.ID) *IconUpdate {
	if id != nil {
		iu = iu.SetProjectID(*id)
	}
	return iu
}

// SetProject sets the "project" edge to the Project entity.
func (iu *IconUpdate) SetProject(p *Project) *IconUpdate {
	return iu.SetProjectID(p.ID)
}

// Mutation returns the IconMutation object of the builder.
func (iu *IconUpdate) Mutation() *IconMutation {
	return iu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (iu *IconUpdate) ClearProject() *IconUpdate {
	iu.mutation.ClearProject()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IconUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IconMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IconUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IconUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IconUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IconUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := icon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := iu.mutation.Icon(); ok {
		if err := icon.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	return nil
}

func (iu *IconUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   icon.Table,
			Columns: icon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: icon.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icon.FieldName,
		})
	}
	if value, ok := iu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icon.FieldIcon,
		})
	}
	if iu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   icon.ProjectTable,
			Columns: []string{icon.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   icon.ProjectTable,
			Columns: []string{icon.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IconUpdateOne is the builder for updating a single Icon entity.
type IconUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IconMutation
}

// SetName sets the "name" field.
func (iuo *IconUpdateOne) SetName(s string) *IconUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetIcon sets the "icon" field.
func (iuo *IconUpdateOne) SetIcon(s string) *IconUpdateOne {
	iuo.mutation.SetIcon(s)
	return iuo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (iuo *IconUpdateOne) SetProjectID(id ulid.ID) *IconUpdateOne {
	iuo.mutation.SetProjectID(id)
	return iuo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (iuo *IconUpdateOne) SetNillableProjectID(id *ulid.ID) *IconUpdateOne {
	if id != nil {
		iuo = iuo.SetProjectID(*id)
	}
	return iuo
}

// SetProject sets the "project" edge to the Project entity.
func (iuo *IconUpdateOne) SetProject(p *Project) *IconUpdateOne {
	return iuo.SetProjectID(p.ID)
}

// Mutation returns the IconMutation object of the builder.
func (iuo *IconUpdateOne) Mutation() *IconMutation {
	return iuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (iuo *IconUpdateOne) ClearProject() *IconUpdateOne {
	iuo.mutation.ClearProject()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IconUpdateOne) Select(field string, fields ...string) *IconUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Icon entity.
func (iuo *IconUpdateOne) Save(ctx context.Context) (*Icon, error) {
	var (
		err  error
		node *Icon
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IconMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IconUpdateOne) SaveX(ctx context.Context) *Icon {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IconUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IconUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IconUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := icon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := iuo.mutation.Icon(); ok {
		if err := icon.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	return nil
}

func (iuo *IconUpdateOne) sqlSave(ctx context.Context) (_node *Icon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   icon.Table,
			Columns: icon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: icon.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Icon.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, icon.FieldID)
		for _, f := range fields {
			if !icon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != icon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icon.FieldName,
		})
	}
	if value, ok := iuo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icon.FieldIcon,
		})
	}
	if iuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   icon.ProjectTable,
			Columns: []string{icon.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   icon.ProjectTable,
			Columns: []string{icon.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Icon{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
