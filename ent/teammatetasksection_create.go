// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/deletedteammatetask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskSectionCreate is the builder for creating a TeammateTaskSection entity.
type TeammateTaskSectionCreate struct {
	config
	mutation *TeammateTaskSectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTeammateID sets the "teammate_id" field.
func (ttsc *TeammateTaskSectionCreate) SetTeammateID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetTeammateID(u)
	return ttsc
}

// SetWorkspaceID sets the "workspace_id" field.
func (ttsc *TeammateTaskSectionCreate) SetWorkspaceID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetWorkspaceID(u)
	return ttsc
}

// SetName sets the "name" field.
func (ttsc *TeammateTaskSectionCreate) SetName(s string) *TeammateTaskSectionCreate {
	ttsc.mutation.SetName(s)
	return ttsc
}

// SetAssigned sets the "assigned" field.
func (ttsc *TeammateTaskSectionCreate) SetAssigned(b bool) *TeammateTaskSectionCreate {
	ttsc.mutation.SetAssigned(b)
	return ttsc
}

// SetCreatedAt sets the "created_at" field.
func (ttsc *TeammateTaskSectionCreate) SetCreatedAt(t time.Time) *TeammateTaskSectionCreate {
	ttsc.mutation.SetCreatedAt(t)
	return ttsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableCreatedAt(t *time.Time) *TeammateTaskSectionCreate {
	if t != nil {
		ttsc.SetCreatedAt(*t)
	}
	return ttsc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttsc *TeammateTaskSectionCreate) SetUpdatedAt(t time.Time) *TeammateTaskSectionCreate {
	ttsc.mutation.SetUpdatedAt(t)
	return ttsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableUpdatedAt(t *time.Time) *TeammateTaskSectionCreate {
	if t != nil {
		ttsc.SetUpdatedAt(*t)
	}
	return ttsc
}

// SetID sets the "id" field.
func (ttsc *TeammateTaskSectionCreate) SetID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetID(u)
	return ttsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableID(u *ulid.ID) *TeammateTaskSectionCreate {
	if u != nil {
		ttsc.SetID(*u)
	}
	return ttsc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttsc *TeammateTaskSectionCreate) SetTeammate(t *Teammate) *TeammateTaskSectionCreate {
	return ttsc.SetTeammateID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (ttsc *TeammateTaskSectionCreate) SetWorkspace(w *Workspace) *TeammateTaskSectionCreate {
	return ttsc.SetWorkspaceID(w.ID)
}

// AddTeammateTaskIDs adds the "teammateTasks" edge to the TeammateTask entity by IDs.
func (ttsc *TeammateTaskSectionCreate) AddTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.AddTeammateTaskIDs(ids...)
	return ttsc
}

// AddTeammateTasks adds the "teammateTasks" edges to the TeammateTask entity.
func (ttsc *TeammateTaskSectionCreate) AddTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsc.AddTeammateTaskIDs(ids...)
}

// AddDeletedTeammateTaskIDs adds the "deletedTeammateTasks" edge to the DeletedTeammateTask entity by IDs.
func (ttsc *TeammateTaskSectionCreate) AddDeletedTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.AddDeletedTeammateTaskIDs(ids...)
	return ttsc
}

// AddDeletedTeammateTasks adds the "deletedTeammateTasks" edges to the DeletedTeammateTask entity.
func (ttsc *TeammateTaskSectionCreate) AddDeletedTeammateTasks(d ...*DeletedTeammateTask) *TeammateTaskSectionCreate {
	ids := make([]ulid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ttsc.AddDeletedTeammateTaskIDs(ids...)
}

// Mutation returns the TeammateTaskSectionMutation object of the builder.
func (ttsc *TeammateTaskSectionCreate) Mutation() *TeammateTaskSectionMutation {
	return ttsc.mutation
}

// Save creates the TeammateTaskSection in the database.
func (ttsc *TeammateTaskSectionCreate) Save(ctx context.Context) (*TeammateTaskSection, error) {
	var (
		err  error
		node *TeammateTaskSection
	)
	ttsc.defaults()
	if len(ttsc.hooks) == 0 {
		if err = ttsc.check(); err != nil {
			return nil, err
		}
		node, err = ttsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttsc.check(); err != nil {
				return nil, err
			}
			ttsc.mutation = mutation
			if node, err = ttsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttsc.hooks) - 1; i >= 0; i-- {
			if ttsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttsc *TeammateTaskSectionCreate) SaveX(ctx context.Context) *TeammateTaskSection {
	v, err := ttsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttsc *TeammateTaskSectionCreate) Exec(ctx context.Context) error {
	_, err := ttsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttsc *TeammateTaskSectionCreate) ExecX(ctx context.Context) {
	if err := ttsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttsc *TeammateTaskSectionCreate) defaults() {
	if _, ok := ttsc.mutation.CreatedAt(); !ok {
		v := teammatetasksection.DefaultCreatedAt()
		ttsc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttsc.mutation.UpdatedAt(); !ok {
		v := teammatetasksection.DefaultUpdatedAt()
		ttsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttsc.mutation.ID(); !ok {
		v := teammatetasksection.DefaultID()
		ttsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttsc *TeammateTaskSectionCreate) check() error {
	if _, ok := ttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "TeammateTaskSection.teammate_id"`)}
	}
	if _, ok := ttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "TeammateTaskSection.workspace_id"`)}
	}
	if _, ok := ttsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TeammateTaskSection.name"`)}
	}
	if v, ok := ttsc.mutation.Name(); ok {
		if err := teammatetasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TeammateTaskSection.name": %w`, err)}
		}
	}
	if _, ok := ttsc.mutation.Assigned(); !ok {
		return &ValidationError{Name: "assigned", err: errors.New(`ent: missing required field "TeammateTaskSection.assigned"`)}
	}
	if _, ok := ttsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TeammateTaskSection.created_at"`)}
	}
	if _, ok := ttsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TeammateTaskSection.updated_at"`)}
	}
	if _, ok := ttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "TeammateTaskSection.teammate"`)}
	}
	if _, ok := ttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "TeammateTaskSection.workspace"`)}
	}
	return nil
}

func (ttsc *TeammateTaskSectionCreate) sqlSave(ctx context.Context) (*TeammateTaskSection, error) {
	_node, _spec := ttsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttsc.driver, _spec); err != nil {
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

func (ttsc *TeammateTaskSectionCreate) createSpec() (*TeammateTaskSection, *sqlgraph.CreateSpec) {
	var (
		_node = &TeammateTaskSection{config: ttsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teammatetasksection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasksection.FieldID,
			},
		}
	)
	_spec.OnConflict = ttsc.conflict
	if id, ok := ttsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ttsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammatetasksection.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ttsc.mutation.Assigned(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetasksection.FieldAssigned,
		})
		_node.Assigned = value
	}
	if value, ok := ttsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasksection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ttsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasksection.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ttsc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if nodes := ttsc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if nodes := ttsc.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttsc.mutation.DeletedTeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.DeletedTeammateTasksTable,
			Columns: []string{teammatetasksection.DeletedTeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: deletedteammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskSection.Create().
//		SetTeammateID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskSectionUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (ttsc *TeammateTaskSectionCreate) OnConflict(opts ...sql.ConflictOption) *TeammateTaskSectionUpsertOne {
	ttsc.conflict = opts
	return &TeammateTaskSectionUpsertOne{
		create: ttsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskSection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttsc *TeammateTaskSectionCreate) OnConflictColumns(columns ...string) *TeammateTaskSectionUpsertOne {
	ttsc.conflict = append(ttsc.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskSectionUpsertOne{
		create: ttsc,
	}
}

type (
	// TeammateTaskSectionUpsertOne is the builder for "upsert"-ing
	//  one TeammateTaskSection node.
	TeammateTaskSectionUpsertOne struct {
		create *TeammateTaskSectionCreate
	}

	// TeammateTaskSectionUpsert is the "OnConflict" setter.
	TeammateTaskSectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskSectionUpsert) SetTeammateID(v ulid.ID) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateTeammateID() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldTeammateID)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskSectionUpsert) SetWorkspaceID(v ulid.ID) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateWorkspaceID() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldWorkspaceID)
	return u
}

// SetName sets the "name" field.
func (u *TeammateTaskSectionUpsert) SetName(v string) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateName() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldName)
	return u
}

// SetAssigned sets the "assigned" field.
func (u *TeammateTaskSectionUpsert) SetAssigned(v bool) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldAssigned, v)
	return u
}

// UpdateAssigned sets the "assigned" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateAssigned() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldAssigned)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskSectionUpsert) SetCreatedAt(v time.Time) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateCreatedAt() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskSectionUpsert) SetUpdatedAt(v time.Time) *TeammateTaskSectionUpsert {
	u.Set(teammatetasksection.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsert) UpdateUpdatedAt() *TeammateTaskSectionUpsert {
	u.SetExcluded(teammatetasksection.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TeammateTaskSection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetasksection.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskSectionUpsertOne) UpdateNewValues() *TeammateTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(teammatetasksection.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(teammatetasksection.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(teammatetasksection.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TeammateTaskSection.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TeammateTaskSectionUpsertOne) Ignore() *TeammateTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskSectionUpsertOne) DoNothing() *TeammateTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskSectionCreate.OnConflict
// documentation for more info.
func (u *TeammateTaskSectionUpsertOne) Update(set func(*TeammateTaskSectionUpsert)) *TeammateTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskSectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskSectionUpsertOne) SetTeammateID(v ulid.ID) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateTeammateID() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateTeammateID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskSectionUpsertOne) SetWorkspaceID(v ulid.ID) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateWorkspaceID() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetName sets the "name" field.
func (u *TeammateTaskSectionUpsertOne) SetName(v string) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateName() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateName()
	})
}

// SetAssigned sets the "assigned" field.
func (u *TeammateTaskSectionUpsertOne) SetAssigned(v bool) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetAssigned(v)
	})
}

// UpdateAssigned sets the "assigned" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateAssigned() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateAssigned()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskSectionUpsertOne) SetCreatedAt(v time.Time) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateCreatedAt() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskSectionUpsertOne) SetUpdatedAt(v time.Time) *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertOne) UpdateUpdatedAt() *TeammateTaskSectionUpsertOne {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskSectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskSectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskSectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeammateTaskSectionUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TeammateTaskSectionUpsertOne.ID is not supported by MySQL driver. Use TeammateTaskSectionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeammateTaskSectionUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeammateTaskSectionCreateBulk is the builder for creating many TeammateTaskSection entities in bulk.
type TeammateTaskSectionCreateBulk struct {
	config
	builders []*TeammateTaskSectionCreate
	conflict []sql.ConflictOption
}

// Save creates the TeammateTaskSection entities in the database.
func (ttscb *TeammateTaskSectionCreateBulk) Save(ctx context.Context) ([]*TeammateTaskSection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttscb.builders))
	nodes := make([]*TeammateTaskSection, len(ttscb.builders))
	mutators := make([]Mutator, len(ttscb.builders))
	for i := range ttscb.builders {
		func(i int, root context.Context) {
			builder := ttscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeammateTaskSectionMutation)
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
					_, err = mutators[i+1].Mutate(root, ttscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ttscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ttscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttscb *TeammateTaskSectionCreateBulk) SaveX(ctx context.Context) []*TeammateTaskSection {
	v, err := ttscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttscb *TeammateTaskSectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ttscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttscb *TeammateTaskSectionCreateBulk) ExecX(ctx context.Context) {
	if err := ttscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskSection.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskSectionUpsert) {
//			SetTeammateID(v+v).
//		}).
//		Exec(ctx)
//
func (ttscb *TeammateTaskSectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeammateTaskSectionUpsertBulk {
	ttscb.conflict = opts
	return &TeammateTaskSectionUpsertBulk{
		create: ttscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskSection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttscb *TeammateTaskSectionCreateBulk) OnConflictColumns(columns ...string) *TeammateTaskSectionUpsertBulk {
	ttscb.conflict = append(ttscb.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskSectionUpsertBulk{
		create: ttscb,
	}
}

// TeammateTaskSectionUpsertBulk is the builder for "upsert"-ing
// a bulk of TeammateTaskSection nodes.
type TeammateTaskSectionUpsertBulk struct {
	create *TeammateTaskSectionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TeammateTaskSection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetasksection.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskSectionUpsertBulk) UpdateNewValues() *TeammateTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(teammatetasksection.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(teammatetasksection.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(teammatetasksection.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TeammateTaskSection.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TeammateTaskSectionUpsertBulk) Ignore() *TeammateTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskSectionUpsertBulk) DoNothing() *TeammateTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskSectionCreateBulk.OnConflict
// documentation for more info.
func (u *TeammateTaskSectionUpsertBulk) Update(set func(*TeammateTaskSectionUpsert)) *TeammateTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskSectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskSectionUpsertBulk) SetTeammateID(v ulid.ID) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateTeammateID() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateTeammateID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskSectionUpsertBulk) SetWorkspaceID(v ulid.ID) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateWorkspaceID() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetName sets the "name" field.
func (u *TeammateTaskSectionUpsertBulk) SetName(v string) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateName() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateName()
	})
}

// SetAssigned sets the "assigned" field.
func (u *TeammateTaskSectionUpsertBulk) SetAssigned(v bool) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetAssigned(v)
	})
}

// UpdateAssigned sets the "assigned" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateAssigned() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateAssigned()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskSectionUpsertBulk) SetCreatedAt(v time.Time) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateCreatedAt() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskSectionUpsertBulk) SetUpdatedAt(v time.Time) *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskSectionUpsertBulk) UpdateUpdatedAt() *TeammateTaskSectionUpsertBulk {
	return u.Update(func(s *TeammateTaskSectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskSectionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TeammateTaskSectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskSectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskSectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
