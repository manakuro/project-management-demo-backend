// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/deletedteammatetask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeletedTeammateTaskCreate is the builder for creating a DeletedTeammateTask entity.
type DeletedTeammateTaskCreate struct {
	config
	mutation *DeletedTeammateTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTeammateID sets the "teammate_id" field.
func (dttc *DeletedTeammateTaskCreate) SetTeammateID(u ulid.ID) *DeletedTeammateTaskCreate {
	dttc.mutation.SetTeammateID(u)
	return dttc
}

// SetTaskID sets the "task_id" field.
func (dttc *DeletedTeammateTaskCreate) SetTaskID(u ulid.ID) *DeletedTeammateTaskCreate {
	dttc.mutation.SetTaskID(u)
	return dttc
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (dttc *DeletedTeammateTaskCreate) SetTeammateTaskSectionID(u ulid.ID) *DeletedTeammateTaskCreate {
	dttc.mutation.SetTeammateTaskSectionID(u)
	return dttc
}

// SetWorkspaceID sets the "workspace_id" field.
func (dttc *DeletedTeammateTaskCreate) SetWorkspaceID(u ulid.ID) *DeletedTeammateTaskCreate {
	dttc.mutation.SetWorkspaceID(u)
	return dttc
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (dttc *DeletedTeammateTaskCreate) SetTeammateTaskCreatedAt(t time.Time) *DeletedTeammateTaskCreate {
	dttc.mutation.SetTeammateTaskCreatedAt(t)
	return dttc
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (dttc *DeletedTeammateTaskCreate) SetTeammateTaskUpdatedAt(t time.Time) *DeletedTeammateTaskCreate {
	dttc.mutation.SetTeammateTaskUpdatedAt(t)
	return dttc
}

// SetCreatedAt sets the "created_at" field.
func (dttc *DeletedTeammateTaskCreate) SetCreatedAt(t time.Time) *DeletedTeammateTaskCreate {
	dttc.mutation.SetCreatedAt(t)
	return dttc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dttc *DeletedTeammateTaskCreate) SetNillableCreatedAt(t *time.Time) *DeletedTeammateTaskCreate {
	if t != nil {
		dttc.SetCreatedAt(*t)
	}
	return dttc
}

// SetUpdatedAt sets the "updated_at" field.
func (dttc *DeletedTeammateTaskCreate) SetUpdatedAt(t time.Time) *DeletedTeammateTaskCreate {
	dttc.mutation.SetUpdatedAt(t)
	return dttc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dttc *DeletedTeammateTaskCreate) SetNillableUpdatedAt(t *time.Time) *DeletedTeammateTaskCreate {
	if t != nil {
		dttc.SetUpdatedAt(*t)
	}
	return dttc
}

// SetID sets the "id" field.
func (dttc *DeletedTeammateTaskCreate) SetID(u ulid.ID) *DeletedTeammateTaskCreate {
	dttc.mutation.SetID(u)
	return dttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dttc *DeletedTeammateTaskCreate) SetNillableID(u *ulid.ID) *DeletedTeammateTaskCreate {
	if u != nil {
		dttc.SetID(*u)
	}
	return dttc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (dttc *DeletedTeammateTaskCreate) SetTeammate(t *Teammate) *DeletedTeammateTaskCreate {
	return dttc.SetTeammateID(t.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (dttc *DeletedTeammateTaskCreate) SetTask(t *Task) *DeletedTeammateTaskCreate {
	return dttc.SetTaskID(t.ID)
}

// SetTeammateTaskSection sets the "teammateTaskSection" edge to the TeammateTaskSection entity.
func (dttc *DeletedTeammateTaskCreate) SetTeammateTaskSection(t *TeammateTaskSection) *DeletedTeammateTaskCreate {
	return dttc.SetTeammateTaskSectionID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (dttc *DeletedTeammateTaskCreate) SetWorkspace(w *Workspace) *DeletedTeammateTaskCreate {
	return dttc.SetWorkspaceID(w.ID)
}

// Mutation returns the DeletedTeammateTaskMutation object of the builder.
func (dttc *DeletedTeammateTaskCreate) Mutation() *DeletedTeammateTaskMutation {
	return dttc.mutation
}

// Save creates the DeletedTeammateTask in the database.
func (dttc *DeletedTeammateTaskCreate) Save(ctx context.Context) (*DeletedTeammateTask, error) {
	var (
		err  error
		node *DeletedTeammateTask
	)
	dttc.defaults()
	if len(dttc.hooks) == 0 {
		if err = dttc.check(); err != nil {
			return nil, err
		}
		node, err = dttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTeammateTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dttc.check(); err != nil {
				return nil, err
			}
			dttc.mutation = mutation
			if node, err = dttc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dttc.hooks) - 1; i >= 0; i-- {
			if dttc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dttc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dttc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dttc *DeletedTeammateTaskCreate) SaveX(ctx context.Context) *DeletedTeammateTask {
	v, err := dttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dttc *DeletedTeammateTaskCreate) Exec(ctx context.Context) error {
	_, err := dttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dttc *DeletedTeammateTaskCreate) ExecX(ctx context.Context) {
	if err := dttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dttc *DeletedTeammateTaskCreate) defaults() {
	if _, ok := dttc.mutation.CreatedAt(); !ok {
		v := deletedteammatetask.DefaultCreatedAt()
		dttc.mutation.SetCreatedAt(v)
	}
	if _, ok := dttc.mutation.UpdatedAt(); !ok {
		v := deletedteammatetask.DefaultUpdatedAt()
		dttc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dttc.mutation.ID(); !ok {
		v := deletedteammatetask.DefaultID()
		dttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dttc *DeletedTeammateTaskCreate) check() error {
	if _, ok := dttc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "DeletedTeammateTask.teammate_id"`)}
	}
	if _, ok := dttc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "DeletedTeammateTask.task_id"`)}
	}
	if _, ok := dttc.mutation.TeammateTaskSectionID(); !ok {
		return &ValidationError{Name: "teammate_task_section_id", err: errors.New(`ent: missing required field "DeletedTeammateTask.teammate_task_section_id"`)}
	}
	if _, ok := dttc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "DeletedTeammateTask.workspace_id"`)}
	}
	if _, ok := dttc.mutation.TeammateTaskCreatedAt(); !ok {
		return &ValidationError{Name: "teammate_task_created_at", err: errors.New(`ent: missing required field "DeletedTeammateTask.teammate_task_created_at"`)}
	}
	if _, ok := dttc.mutation.TeammateTaskUpdatedAt(); !ok {
		return &ValidationError{Name: "teammate_task_updated_at", err: errors.New(`ent: missing required field "DeletedTeammateTask.teammate_task_updated_at"`)}
	}
	if _, ok := dttc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeletedTeammateTask.created_at"`)}
	}
	if _, ok := dttc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeletedTeammateTask.updated_at"`)}
	}
	if _, ok := dttc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "DeletedTeammateTask.teammate"`)}
	}
	if _, ok := dttc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "DeletedTeammateTask.task"`)}
	}
	if _, ok := dttc.mutation.TeammateTaskSectionID(); !ok {
		return &ValidationError{Name: "teammateTaskSection", err: errors.New(`ent: missing required edge "DeletedTeammateTask.teammateTaskSection"`)}
	}
	if _, ok := dttc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "DeletedTeammateTask.workspace"`)}
	}
	return nil
}

func (dttc *DeletedTeammateTaskCreate) sqlSave(ctx context.Context) (*DeletedTeammateTask, error) {
	_node, _spec := dttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dttc.driver, _spec); err != nil {
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

func (dttc *DeletedTeammateTaskCreate) createSpec() (*DeletedTeammateTask, *sqlgraph.CreateSpec) {
	var (
		_node = &DeletedTeammateTask{config: dttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deletedteammatetask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedteammatetask.FieldID,
			},
		}
	)
	_spec.OnConflict = dttc.conflict
	if id, ok := dttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dttc.mutation.TeammateTaskCreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskCreatedAt,
		})
		_node.TeammateTaskCreatedAt = value
	}
	if value, ok := dttc.mutation.TeammateTaskUpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskUpdatedAt,
		})
		_node.TeammateTaskUpdatedAt = value
	}
	if value, ok := dttc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dttc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := dttc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTable,
			Columns: []string{deletedteammatetask.TeammateColumn},
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
	if nodes := dttc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TaskTable,
			Columns: []string{deletedteammatetask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dttc.mutation.TeammateTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTaskSectionTable,
			Columns: []string{deletedteammatetask.TeammateTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasksection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeammateTaskSectionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dttc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.WorkspaceTable,
			Columns: []string{deletedteammatetask.WorkspaceColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeletedTeammateTask.Create().
//		SetTeammateID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeletedTeammateTaskUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (dttc *DeletedTeammateTaskCreate) OnConflict(opts ...sql.ConflictOption) *DeletedTeammateTaskUpsertOne {
	dttc.conflict = opts
	return &DeletedTeammateTaskUpsertOne{
		create: dttc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeletedTeammateTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dttc *DeletedTeammateTaskCreate) OnConflictColumns(columns ...string) *DeletedTeammateTaskUpsertOne {
	dttc.conflict = append(dttc.conflict, sql.ConflictColumns(columns...))
	return &DeletedTeammateTaskUpsertOne{
		create: dttc,
	}
}

type (
	// DeletedTeammateTaskUpsertOne is the builder for "upsert"-ing
	//  one DeletedTeammateTask node.
	DeletedTeammateTaskUpsertOne struct {
		create *DeletedTeammateTaskCreate
	}

	// DeletedTeammateTaskUpsert is the "OnConflict" setter.
	DeletedTeammateTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetTeammateID sets the "teammate_id" field.
func (u *DeletedTeammateTaskUpsert) SetTeammateID(v ulid.ID) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateTeammateID() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldTeammateID)
	return u
}

// SetTaskID sets the "task_id" field.
func (u *DeletedTeammateTaskUpsert) SetTaskID(v ulid.ID) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateTaskID() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldTaskID)
	return u
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (u *DeletedTeammateTaskUpsert) SetTeammateTaskSectionID(v ulid.ID) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldTeammateTaskSectionID, v)
	return u
}

// UpdateTeammateTaskSectionID sets the "teammate_task_section_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateTeammateTaskSectionID() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldTeammateTaskSectionID)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTeammateTaskUpsert) SetWorkspaceID(v ulid.ID) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateWorkspaceID() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldWorkspaceID)
	return u
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (u *DeletedTeammateTaskUpsert) SetTeammateTaskCreatedAt(v time.Time) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldTeammateTaskCreatedAt, v)
	return u
}

// UpdateTeammateTaskCreatedAt sets the "teammate_task_created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateTeammateTaskCreatedAt() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldTeammateTaskCreatedAt)
	return u
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (u *DeletedTeammateTaskUpsert) SetTeammateTaskUpdatedAt(v time.Time) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldTeammateTaskUpdatedAt, v)
	return u
}

// UpdateTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateTeammateTaskUpdatedAt() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldTeammateTaskUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTeammateTaskUpsert) SetCreatedAt(v time.Time) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateCreatedAt() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTeammateTaskUpsert) SetUpdatedAt(v time.Time) *DeletedTeammateTaskUpsert {
	u.Set(deletedteammatetask.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsert) UpdateUpdatedAt() *DeletedTeammateTaskUpsert {
	u.SetExcluded(deletedteammatetask.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeletedTeammateTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deletedteammatetask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeletedTeammateTaskUpsertOne) UpdateNewValues() *DeletedTeammateTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deletedteammatetask.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(deletedteammatetask.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(deletedteammatetask.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DeletedTeammateTask.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DeletedTeammateTaskUpsertOne) Ignore() *DeletedTeammateTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeletedTeammateTaskUpsertOne) DoNothing() *DeletedTeammateTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeletedTeammateTaskCreate.OnConflict
// documentation for more info.
func (u *DeletedTeammateTaskUpsertOne) Update(set func(*DeletedTeammateTaskUpsert)) *DeletedTeammateTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeletedTeammateTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *DeletedTeammateTaskUpsertOne) SetTeammateID(v ulid.ID) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateTeammateID() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *DeletedTeammateTaskUpsertOne) SetTaskID(v ulid.ID) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateTaskID() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (u *DeletedTeammateTaskUpsertOne) SetTeammateTaskSectionID(v ulid.ID) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskSectionID(v)
	})
}

// UpdateTeammateTaskSectionID sets the "teammate_task_section_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateTeammateTaskSectionID() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskSectionID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTeammateTaskUpsertOne) SetWorkspaceID(v ulid.ID) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateWorkspaceID() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (u *DeletedTeammateTaskUpsertOne) SetTeammateTaskCreatedAt(v time.Time) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskCreatedAt(v)
	})
}

// UpdateTeammateTaskCreatedAt sets the "teammate_task_created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateTeammateTaskCreatedAt() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskCreatedAt()
	})
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (u *DeletedTeammateTaskUpsertOne) SetTeammateTaskUpdatedAt(v time.Time) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskUpdatedAt(v)
	})
}

// UpdateTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateTeammateTaskUpdatedAt() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTeammateTaskUpsertOne) SetCreatedAt(v time.Time) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateCreatedAt() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTeammateTaskUpsertOne) SetUpdatedAt(v time.Time) *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertOne) UpdateUpdatedAt() *DeletedTeammateTaskUpsertOne {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeletedTeammateTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeletedTeammateTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeletedTeammateTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeletedTeammateTaskUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DeletedTeammateTaskUpsertOne.ID is not supported by MySQL driver. Use DeletedTeammateTaskUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeletedTeammateTaskUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeletedTeammateTaskCreateBulk is the builder for creating many DeletedTeammateTask entities in bulk.
type DeletedTeammateTaskCreateBulk struct {
	config
	builders []*DeletedTeammateTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the DeletedTeammateTask entities in the database.
func (dttcb *DeletedTeammateTaskCreateBulk) Save(ctx context.Context) ([]*DeletedTeammateTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dttcb.builders))
	nodes := make([]*DeletedTeammateTask, len(dttcb.builders))
	mutators := make([]Mutator, len(dttcb.builders))
	for i := range dttcb.builders {
		func(i int, root context.Context) {
			builder := dttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeletedTeammateTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, dttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dttcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dttcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dttcb *DeletedTeammateTaskCreateBulk) SaveX(ctx context.Context) []*DeletedTeammateTask {
	v, err := dttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dttcb *DeletedTeammateTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := dttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dttcb *DeletedTeammateTaskCreateBulk) ExecX(ctx context.Context) {
	if err := dttcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeletedTeammateTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeletedTeammateTaskUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (dttcb *DeletedTeammateTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeletedTeammateTaskUpsertBulk {
	dttcb.conflict = opts
	return &DeletedTeammateTaskUpsertBulk{
		create: dttcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeletedTeammateTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dttcb *DeletedTeammateTaskCreateBulk) OnConflictColumns(columns ...string) *DeletedTeammateTaskUpsertBulk {
	dttcb.conflict = append(dttcb.conflict, sql.ConflictColumns(columns...))
	return &DeletedTeammateTaskUpsertBulk{
		create: dttcb,
	}
}

// DeletedTeammateTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of DeletedTeammateTask nodes.
type DeletedTeammateTaskUpsertBulk struct {
	create *DeletedTeammateTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeletedTeammateTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deletedteammatetask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeletedTeammateTaskUpsertBulk) UpdateNewValues() *DeletedTeammateTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deletedteammatetask.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(deletedteammatetask.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(deletedteammatetask.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeletedTeammateTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DeletedTeammateTaskUpsertBulk) Ignore() *DeletedTeammateTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeletedTeammateTaskUpsertBulk) DoNothing() *DeletedTeammateTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeletedTeammateTaskCreateBulk.OnConflict
// documentation for more info.
func (u *DeletedTeammateTaskUpsertBulk) Update(set func(*DeletedTeammateTaskUpsert)) *DeletedTeammateTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeletedTeammateTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *DeletedTeammateTaskUpsertBulk) SetTeammateID(v ulid.ID) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateTeammateID() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *DeletedTeammateTaskUpsertBulk) SetTaskID(v ulid.ID) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateTaskID() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (u *DeletedTeammateTaskUpsertBulk) SetTeammateTaskSectionID(v ulid.ID) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskSectionID(v)
	})
}

// UpdateTeammateTaskSectionID sets the "teammate_task_section_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateTeammateTaskSectionID() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskSectionID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTeammateTaskUpsertBulk) SetWorkspaceID(v ulid.ID) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateWorkspaceID() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (u *DeletedTeammateTaskUpsertBulk) SetTeammateTaskCreatedAt(v time.Time) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskCreatedAt(v)
	})
}

// UpdateTeammateTaskCreatedAt sets the "teammate_task_created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateTeammateTaskCreatedAt() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskCreatedAt()
	})
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (u *DeletedTeammateTaskUpsertBulk) SetTeammateTaskUpdatedAt(v time.Time) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetTeammateTaskUpdatedAt(v)
	})
}

// UpdateTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateTeammateTaskUpdatedAt() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateTeammateTaskUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTeammateTaskUpsertBulk) SetCreatedAt(v time.Time) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateCreatedAt() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTeammateTaskUpsertBulk) SetUpdatedAt(v time.Time) *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTeammateTaskUpsertBulk) UpdateUpdatedAt() *DeletedTeammateTaskUpsertBulk {
	return u.Update(func(s *DeletedTeammateTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeletedTeammateTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeletedTeammateTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeletedTeammateTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeletedTeammateTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
