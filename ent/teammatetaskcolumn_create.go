// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskcolumn"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetaskcolumn"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskColumnCreate is the builder for creating a TeammateTaskColumn entity.
type TeammateTaskColumnCreate struct {
	config
	mutation *TeammateTaskColumnMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTeammateID sets the "teammate_id" field.
func (ttcc *TeammateTaskColumnCreate) SetTeammateID(u ulid.ID) *TeammateTaskColumnCreate {
	ttcc.mutation.SetTeammateID(u)
	return ttcc
}

// SetTaskColumnID sets the "task_column_id" field.
func (ttcc *TeammateTaskColumnCreate) SetTaskColumnID(u ulid.ID) *TeammateTaskColumnCreate {
	ttcc.mutation.SetTaskColumnID(u)
	return ttcc
}

// SetWorkspaceID sets the "workspace_id" field.
func (ttcc *TeammateTaskColumnCreate) SetWorkspaceID(u ulid.ID) *TeammateTaskColumnCreate {
	ttcc.mutation.SetWorkspaceID(u)
	return ttcc
}

// SetWidth sets the "width" field.
func (ttcc *TeammateTaskColumnCreate) SetWidth(s string) *TeammateTaskColumnCreate {
	ttcc.mutation.SetWidth(s)
	return ttcc
}

// SetDisabled sets the "disabled" field.
func (ttcc *TeammateTaskColumnCreate) SetDisabled(b bool) *TeammateTaskColumnCreate {
	ttcc.mutation.SetDisabled(b)
	return ttcc
}

// SetCustomizable sets the "customizable" field.
func (ttcc *TeammateTaskColumnCreate) SetCustomizable(b bool) *TeammateTaskColumnCreate {
	ttcc.mutation.SetCustomizable(b)
	return ttcc
}

// SetOrder sets the "order" field.
func (ttcc *TeammateTaskColumnCreate) SetOrder(i int) *TeammateTaskColumnCreate {
	ttcc.mutation.SetOrder(i)
	return ttcc
}

// SetCreatedAt sets the "created_at" field.
func (ttcc *TeammateTaskColumnCreate) SetCreatedAt(t time.Time) *TeammateTaskColumnCreate {
	ttcc.mutation.SetCreatedAt(t)
	return ttcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttcc *TeammateTaskColumnCreate) SetNillableCreatedAt(t *time.Time) *TeammateTaskColumnCreate {
	if t != nil {
		ttcc.SetCreatedAt(*t)
	}
	return ttcc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttcc *TeammateTaskColumnCreate) SetUpdatedAt(t time.Time) *TeammateTaskColumnCreate {
	ttcc.mutation.SetUpdatedAt(t)
	return ttcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttcc *TeammateTaskColumnCreate) SetNillableUpdatedAt(t *time.Time) *TeammateTaskColumnCreate {
	if t != nil {
		ttcc.SetUpdatedAt(*t)
	}
	return ttcc
}

// SetID sets the "id" field.
func (ttcc *TeammateTaskColumnCreate) SetID(u ulid.ID) *TeammateTaskColumnCreate {
	ttcc.mutation.SetID(u)
	return ttcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttcc *TeammateTaskColumnCreate) SetNillableID(u *ulid.ID) *TeammateTaskColumnCreate {
	if u != nil {
		ttcc.SetID(*u)
	}
	return ttcc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttcc *TeammateTaskColumnCreate) SetTeammate(t *Teammate) *TeammateTaskColumnCreate {
	return ttcc.SetTeammateID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (ttcc *TeammateTaskColumnCreate) SetWorkspace(w *Workspace) *TeammateTaskColumnCreate {
	return ttcc.SetWorkspaceID(w.ID)
}

// SetTaskColumn sets the "taskColumn" edge to the TaskColumn entity.
func (ttcc *TeammateTaskColumnCreate) SetTaskColumn(t *TaskColumn) *TeammateTaskColumnCreate {
	return ttcc.SetTaskColumnID(t.ID)
}

// Mutation returns the TeammateTaskColumnMutation object of the builder.
func (ttcc *TeammateTaskColumnCreate) Mutation() *TeammateTaskColumnMutation {
	return ttcc.mutation
}

// Save creates the TeammateTaskColumn in the database.
func (ttcc *TeammateTaskColumnCreate) Save(ctx context.Context) (*TeammateTaskColumn, error) {
	var (
		err  error
		node *TeammateTaskColumn
	)
	ttcc.defaults()
	if len(ttcc.hooks) == 0 {
		if err = ttcc.check(); err != nil {
			return nil, err
		}
		node, err = ttcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttcc.check(); err != nil {
				return nil, err
			}
			ttcc.mutation = mutation
			if node, err = ttcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttcc.hooks) - 1; i >= 0; i-- {
			if ttcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttcc *TeammateTaskColumnCreate) SaveX(ctx context.Context) *TeammateTaskColumn {
	v, err := ttcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcc *TeammateTaskColumnCreate) Exec(ctx context.Context) error {
	_, err := ttcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcc *TeammateTaskColumnCreate) ExecX(ctx context.Context) {
	if err := ttcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttcc *TeammateTaskColumnCreate) defaults() {
	if _, ok := ttcc.mutation.CreatedAt(); !ok {
		v := teammatetaskcolumn.DefaultCreatedAt()
		ttcc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttcc.mutation.UpdatedAt(); !ok {
		v := teammatetaskcolumn.DefaultUpdatedAt()
		ttcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttcc.mutation.ID(); !ok {
		v := teammatetaskcolumn.DefaultID()
		ttcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttcc *TeammateTaskColumnCreate) check() error {
	if _, ok := ttcc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "TeammateTaskColumn.teammate_id"`)}
	}
	if _, ok := ttcc.mutation.TaskColumnID(); !ok {
		return &ValidationError{Name: "task_column_id", err: errors.New(`ent: missing required field "TeammateTaskColumn.task_column_id"`)}
	}
	if _, ok := ttcc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "TeammateTaskColumn.workspace_id"`)}
	}
	if _, ok := ttcc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "TeammateTaskColumn.width"`)}
	}
	if v, ok := ttcc.mutation.Width(); ok {
		if err := teammatetaskcolumn.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "TeammateTaskColumn.width": %w`, err)}
		}
	}
	if _, ok := ttcc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New(`ent: missing required field "TeammateTaskColumn.disabled"`)}
	}
	if _, ok := ttcc.mutation.Customizable(); !ok {
		return &ValidationError{Name: "customizable", err: errors.New(`ent: missing required field "TeammateTaskColumn.customizable"`)}
	}
	if _, ok := ttcc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "TeammateTaskColumn.order"`)}
	}
	if _, ok := ttcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TeammateTaskColumn.created_at"`)}
	}
	if _, ok := ttcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TeammateTaskColumn.updated_at"`)}
	}
	if _, ok := ttcc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "TeammateTaskColumn.teammate"`)}
	}
	if _, ok := ttcc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "TeammateTaskColumn.workspace"`)}
	}
	if _, ok := ttcc.mutation.TaskColumnID(); !ok {
		return &ValidationError{Name: "taskColumn", err: errors.New(`ent: missing required edge "TeammateTaskColumn.taskColumn"`)}
	}
	return nil
}

func (ttcc *TeammateTaskColumnCreate) sqlSave(ctx context.Context) (*TeammateTaskColumn, error) {
	_node, _spec := ttcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*ulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ttcc *TeammateTaskColumnCreate) createSpec() (*TeammateTaskColumn, *sqlgraph.CreateSpec) {
	var (
		_node = &TeammateTaskColumn{config: ttcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teammatetaskcolumn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetaskcolumn.FieldID,
			},
		}
	)
	_spec.OnConflict = ttcc.conflict
	if id, ok := ttcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ttcc.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammatetaskcolumn.FieldWidth,
		})
		_node.Width = value
	}
	if value, ok := ttcc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetaskcolumn.FieldDisabled,
		})
		_node.Disabled = value
	}
	if value, ok := ttcc.mutation.Customizable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetaskcolumn.FieldCustomizable,
		})
		_node.Customizable = value
	}
	if value, ok := ttcc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: teammatetaskcolumn.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := ttcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetaskcolumn.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ttcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetaskcolumn.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ttcc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetaskcolumn.TeammateTable,
			Columns: []string{teammatetaskcolumn.TeammateColumn},
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
		_node.TeammateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttcc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetaskcolumn.WorkspaceTable,
			Columns: []string{teammatetaskcolumn.WorkspaceColumn},
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
		_node.WorkspaceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttcc.mutation.TaskColumnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetaskcolumn.TaskColumnTable,
			Columns: []string{teammatetaskcolumn.TaskColumnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskcolumn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskColumnID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskColumn.Create().
//		SetTeammateID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskColumnUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (ttcc *TeammateTaskColumnCreate) OnConflict(opts ...sql.ConflictOption) *TeammateTaskColumnUpsertOne {
	ttcc.conflict = opts
	return &TeammateTaskColumnUpsertOne{
		create: ttcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskColumn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttcc *TeammateTaskColumnCreate) OnConflictColumns(columns ...string) *TeammateTaskColumnUpsertOne {
	ttcc.conflict = append(ttcc.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskColumnUpsertOne{
		create: ttcc,
	}
}

type (
	// TeammateTaskColumnUpsertOne is the builder for "upsert"-ing
	//  one TeammateTaskColumn node.
	TeammateTaskColumnUpsertOne struct {
		create *TeammateTaskColumnCreate
	}

	// TeammateTaskColumnUpsert is the "OnConflict" setter.
	TeammateTaskColumnUpsert struct {
		*sql.UpdateSet
	}
)

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskColumnUpsert) SetTeammateID(v ulid.ID) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateTeammateID() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldTeammateID)
	return u
}

// SetTaskColumnID sets the "task_column_id" field.
func (u *TeammateTaskColumnUpsert) SetTaskColumnID(v ulid.ID) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldTaskColumnID, v)
	return u
}

// UpdateTaskColumnID sets the "task_column_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateTaskColumnID() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldTaskColumnID)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskColumnUpsert) SetWorkspaceID(v ulid.ID) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateWorkspaceID() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldWorkspaceID)
	return u
}

// SetWidth sets the "width" field.
func (u *TeammateTaskColumnUpsert) SetWidth(v string) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldWidth, v)
	return u
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateWidth() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldWidth)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *TeammateTaskColumnUpsert) SetDisabled(v bool) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateDisabled() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldDisabled)
	return u
}

// SetCustomizable sets the "customizable" field.
func (u *TeammateTaskColumnUpsert) SetCustomizable(v bool) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldCustomizable, v)
	return u
}

// UpdateCustomizable sets the "customizable" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateCustomizable() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldCustomizable)
	return u
}

// SetOrder sets the "order" field.
func (u *TeammateTaskColumnUpsert) SetOrder(v int) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateOrder() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *TeammateTaskColumnUpsert) AddOrder(v int) *TeammateTaskColumnUpsert {
	u.Add(teammatetaskcolumn.FieldOrder, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskColumnUpsert) SetCreatedAt(v time.Time) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateCreatedAt() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskColumnUpsert) SetUpdatedAt(v time.Time) *TeammateTaskColumnUpsert {
	u.Set(teammatetaskcolumn.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsert) UpdateUpdatedAt() *TeammateTaskColumnUpsert {
	u.SetExcluded(teammatetaskcolumn.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TeammateTaskColumn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetaskcolumn.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskColumnUpsertOne) UpdateNewValues() *TeammateTaskColumnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(teammatetaskcolumn.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(teammatetaskcolumn.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(teammatetaskcolumn.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TeammateTaskColumn.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TeammateTaskColumnUpsertOne) Ignore() *TeammateTaskColumnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskColumnUpsertOne) DoNothing() *TeammateTaskColumnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskColumnCreate.OnConflict
// documentation for more info.
func (u *TeammateTaskColumnUpsertOne) Update(set func(*TeammateTaskColumnUpsert)) *TeammateTaskColumnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskColumnUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskColumnUpsertOne) SetTeammateID(v ulid.ID) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateTeammateID() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateTeammateID()
	})
}

// SetTaskColumnID sets the "task_column_id" field.
func (u *TeammateTaskColumnUpsertOne) SetTaskColumnID(v ulid.ID) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetTaskColumnID(v)
	})
}

// UpdateTaskColumnID sets the "task_column_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateTaskColumnID() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateTaskColumnID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskColumnUpsertOne) SetWorkspaceID(v ulid.ID) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateWorkspaceID() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetWidth sets the "width" field.
func (u *TeammateTaskColumnUpsertOne) SetWidth(v string) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateWidth() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateWidth()
	})
}

// SetDisabled sets the "disabled" field.
func (u *TeammateTaskColumnUpsertOne) SetDisabled(v bool) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateDisabled() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateDisabled()
	})
}

// SetCustomizable sets the "customizable" field.
func (u *TeammateTaskColumnUpsertOne) SetCustomizable(v bool) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetCustomizable(v)
	})
}

// UpdateCustomizable sets the "customizable" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateCustomizable() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateCustomizable()
	})
}

// SetOrder sets the "order" field.
func (u *TeammateTaskColumnUpsertOne) SetOrder(v int) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *TeammateTaskColumnUpsertOne) AddOrder(v int) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateOrder() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateOrder()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskColumnUpsertOne) SetCreatedAt(v time.Time) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateCreatedAt() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskColumnUpsertOne) SetUpdatedAt(v time.Time) *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertOne) UpdateUpdatedAt() *TeammateTaskColumnUpsertOne {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskColumnUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskColumnCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskColumnUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeammateTaskColumnUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TeammateTaskColumnUpsertOne.ID is not supported by MySQL driver. Use TeammateTaskColumnUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeammateTaskColumnUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeammateTaskColumnCreateBulk is the builder for creating many TeammateTaskColumn entities in bulk.
type TeammateTaskColumnCreateBulk struct {
	config
	builders []*TeammateTaskColumnCreate
	conflict []sql.ConflictOption
}

// Save creates the TeammateTaskColumn entities in the database.
func (ttccb *TeammateTaskColumnCreateBulk) Save(ctx context.Context) ([]*TeammateTaskColumn, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttccb.builders))
	nodes := make([]*TeammateTaskColumn, len(ttccb.builders))
	mutators := make([]Mutator, len(ttccb.builders))
	for i := range ttccb.builders {
		func(i int, root context.Context) {
			builder := ttccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeammateTaskColumnMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ttccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ttccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ttccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttccb *TeammateTaskColumnCreateBulk) SaveX(ctx context.Context) []*TeammateTaskColumn {
	v, err := ttccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttccb *TeammateTaskColumnCreateBulk) Exec(ctx context.Context) error {
	_, err := ttccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttccb *TeammateTaskColumnCreateBulk) ExecX(ctx context.Context) {
	if err := ttccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskColumn.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskColumnUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (ttccb *TeammateTaskColumnCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeammateTaskColumnUpsertBulk {
	ttccb.conflict = opts
	return &TeammateTaskColumnUpsertBulk{
		create: ttccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskColumn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttccb *TeammateTaskColumnCreateBulk) OnConflictColumns(columns ...string) *TeammateTaskColumnUpsertBulk {
	ttccb.conflict = append(ttccb.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskColumnUpsertBulk{
		create: ttccb,
	}
}

// TeammateTaskColumnUpsertBulk is the builder for "upsert"-ing
// a bulk of TeammateTaskColumn nodes.
type TeammateTaskColumnUpsertBulk struct {
	create *TeammateTaskColumnCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TeammateTaskColumn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetaskcolumn.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskColumnUpsertBulk) UpdateNewValues() *TeammateTaskColumnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(teammatetaskcolumn.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(teammatetaskcolumn.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(teammatetaskcolumn.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TeammateTaskColumn.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TeammateTaskColumnUpsertBulk) Ignore() *TeammateTaskColumnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskColumnUpsertBulk) DoNothing() *TeammateTaskColumnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskColumnCreateBulk.OnConflict
// documentation for more info.
func (u *TeammateTaskColumnUpsertBulk) Update(set func(*TeammateTaskColumnUpsert)) *TeammateTaskColumnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskColumnUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskColumnUpsertBulk) SetTeammateID(v ulid.ID) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateTeammateID() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateTeammateID()
	})
}

// SetTaskColumnID sets the "task_column_id" field.
func (u *TeammateTaskColumnUpsertBulk) SetTaskColumnID(v ulid.ID) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetTaskColumnID(v)
	})
}

// UpdateTaskColumnID sets the "task_column_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateTaskColumnID() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateTaskColumnID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskColumnUpsertBulk) SetWorkspaceID(v ulid.ID) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateWorkspaceID() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetWidth sets the "width" field.
func (u *TeammateTaskColumnUpsertBulk) SetWidth(v string) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateWidth() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateWidth()
	})
}

// SetDisabled sets the "disabled" field.
func (u *TeammateTaskColumnUpsertBulk) SetDisabled(v bool) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateDisabled() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateDisabled()
	})
}

// SetCustomizable sets the "customizable" field.
func (u *TeammateTaskColumnUpsertBulk) SetCustomizable(v bool) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetCustomizable(v)
	})
}

// UpdateCustomizable sets the "customizable" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateCustomizable() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateCustomizable()
	})
}

// SetOrder sets the "order" field.
func (u *TeammateTaskColumnUpsertBulk) SetOrder(v int) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *TeammateTaskColumnUpsertBulk) AddOrder(v int) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateOrder() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateOrder()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskColumnUpsertBulk) SetCreatedAt(v time.Time) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateCreatedAt() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskColumnUpsertBulk) SetUpdatedAt(v time.Time) *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskColumnUpsertBulk) UpdateUpdatedAt() *TeammateTaskColumnUpsertBulk {
	return u.Update(func(s *TeammateTaskColumnUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskColumnUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TeammateTaskColumnCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskColumnCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskColumnUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
