// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetasktabstatus"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskTabStatusCreate is the builder for creating a TeammateTaskTabStatus entity.
type TeammateTaskTabStatusCreate struct {
	config
	mutation *TeammateTaskTabStatusMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWorkspaceID sets the "workspace_id" field.
func (tttsc *TeammateTaskTabStatusCreate) SetWorkspaceID(u ulid.ID) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetWorkspaceID(u)
	return tttsc
}

// SetTeammateID sets the "teammate_id" field.
func (tttsc *TeammateTaskTabStatusCreate) SetTeammateID(u ulid.ID) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetTeammateID(u)
	return tttsc
}

// SetStatusCode sets the "status_code" field.
func (tttsc *TeammateTaskTabStatusCreate) SetStatusCode(tc teammatetasktabstatus.StatusCode) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetStatusCode(tc)
	return tttsc
}

// SetNillableStatusCode sets the "status_code" field if the given value is not nil.
func (tttsc *TeammateTaskTabStatusCreate) SetNillableStatusCode(tc *teammatetasktabstatus.StatusCode) *TeammateTaskTabStatusCreate {
	if tc != nil {
		tttsc.SetStatusCode(*tc)
	}
	return tttsc
}

// SetCreatedAt sets the "created_at" field.
func (tttsc *TeammateTaskTabStatusCreate) SetCreatedAt(t time.Time) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetCreatedAt(t)
	return tttsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tttsc *TeammateTaskTabStatusCreate) SetNillableCreatedAt(t *time.Time) *TeammateTaskTabStatusCreate {
	if t != nil {
		tttsc.SetCreatedAt(*t)
	}
	return tttsc
}

// SetUpdatedAt sets the "updated_at" field.
func (tttsc *TeammateTaskTabStatusCreate) SetUpdatedAt(t time.Time) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetUpdatedAt(t)
	return tttsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tttsc *TeammateTaskTabStatusCreate) SetNillableUpdatedAt(t *time.Time) *TeammateTaskTabStatusCreate {
	if t != nil {
		tttsc.SetUpdatedAt(*t)
	}
	return tttsc
}

// SetID sets the "id" field.
func (tttsc *TeammateTaskTabStatusCreate) SetID(u ulid.ID) *TeammateTaskTabStatusCreate {
	tttsc.mutation.SetID(u)
	return tttsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tttsc *TeammateTaskTabStatusCreate) SetNillableID(u *ulid.ID) *TeammateTaskTabStatusCreate {
	if u != nil {
		tttsc.SetID(*u)
	}
	return tttsc
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (tttsc *TeammateTaskTabStatusCreate) SetWorkspace(w *Workspace) *TeammateTaskTabStatusCreate {
	return tttsc.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tttsc *TeammateTaskTabStatusCreate) SetTeammate(t *Teammate) *TeammateTaskTabStatusCreate {
	return tttsc.SetTeammateID(t.ID)
}

// Mutation returns the TeammateTaskTabStatusMutation object of the builder.
func (tttsc *TeammateTaskTabStatusCreate) Mutation() *TeammateTaskTabStatusMutation {
	return tttsc.mutation
}

// Save creates the TeammateTaskTabStatus in the database.
func (tttsc *TeammateTaskTabStatusCreate) Save(ctx context.Context) (*TeammateTaskTabStatus, error) {
	var (
		err  error
		node *TeammateTaskTabStatus
	)
	tttsc.defaults()
	if len(tttsc.hooks) == 0 {
		if err = tttsc.check(); err != nil {
			return nil, err
		}
		node, err = tttsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskTabStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tttsc.check(); err != nil {
				return nil, err
			}
			tttsc.mutation = mutation
			if node, err = tttsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tttsc.hooks) - 1; i >= 0; i-- {
			if tttsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tttsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tttsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tttsc *TeammateTaskTabStatusCreate) SaveX(ctx context.Context) *TeammateTaskTabStatus {
	v, err := tttsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tttsc *TeammateTaskTabStatusCreate) Exec(ctx context.Context) error {
	_, err := tttsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tttsc *TeammateTaskTabStatusCreate) ExecX(ctx context.Context) {
	if err := tttsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tttsc *TeammateTaskTabStatusCreate) defaults() {
	if _, ok := tttsc.mutation.StatusCode(); !ok {
		v := teammatetasktabstatus.DefaultStatusCode
		tttsc.mutation.SetStatusCode(v)
	}
	if _, ok := tttsc.mutation.CreatedAt(); !ok {
		v := teammatetasktabstatus.DefaultCreatedAt()
		tttsc.mutation.SetCreatedAt(v)
	}
	if _, ok := tttsc.mutation.UpdatedAt(); !ok {
		v := teammatetasktabstatus.DefaultUpdatedAt()
		tttsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tttsc.mutation.ID(); !ok {
		v := teammatetasktabstatus.DefaultID()
		tttsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tttsc *TeammateTaskTabStatusCreate) check() error {
	if _, ok := tttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "TeammateTaskTabStatus.workspace_id"`)}
	}
	if _, ok := tttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "TeammateTaskTabStatus.teammate_id"`)}
	}
	if _, ok := tttsc.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`ent: missing required field "TeammateTaskTabStatus.status_code"`)}
	}
	if v, ok := tttsc.mutation.StatusCode(); ok {
		if err := teammatetasktabstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf(`ent: validator failed for field "TeammateTaskTabStatus.status_code": %w`, err)}
		}
	}
	if _, ok := tttsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TeammateTaskTabStatus.created_at"`)}
	}
	if _, ok := tttsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TeammateTaskTabStatus.updated_at"`)}
	}
	if _, ok := tttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "TeammateTaskTabStatus.workspace"`)}
	}
	if _, ok := tttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "TeammateTaskTabStatus.teammate"`)}
	}
	return nil
}

func (tttsc *TeammateTaskTabStatusCreate) sqlSave(ctx context.Context) (*TeammateTaskTabStatus, error) {
	_node, _spec := tttsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tttsc.driver, _spec); err != nil {
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

func (tttsc *TeammateTaskTabStatusCreate) createSpec() (*TeammateTaskTabStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &TeammateTaskTabStatus{config: tttsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teammatetasktabstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasktabstatus.FieldID,
			},
		}
	)
	_spec.OnConflict = tttsc.conflict
	if id, ok := tttsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tttsc.mutation.StatusCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: teammatetasktabstatus.FieldStatusCode,
		})
		_node.StatusCode = value
	}
	if value, ok := tttsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasktabstatus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tttsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasktabstatus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tttsc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasktabstatus.WorkspaceTable,
			Columns: []string{teammatetasktabstatus.WorkspaceColumn},
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
	if nodes := tttsc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasktabstatus.TeammateTable,
			Columns: []string{teammatetasktabstatus.TeammateColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskTabStatus.Create().
//		SetWorkspaceID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskTabStatusUpsert) {
//			SetWorkspaceID(v+v).
//		}).
//		Exec(ctx)
//
func (tttsc *TeammateTaskTabStatusCreate) OnConflict(opts ...sql.ConflictOption) *TeammateTaskTabStatusUpsertOne {
	tttsc.conflict = opts
	return &TeammateTaskTabStatusUpsertOne{
		create: tttsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskTabStatus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tttsc *TeammateTaskTabStatusCreate) OnConflictColumns(columns ...string) *TeammateTaskTabStatusUpsertOne {
	tttsc.conflict = append(tttsc.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskTabStatusUpsertOne{
		create: tttsc,
	}
}

type (
	// TeammateTaskTabStatusUpsertOne is the builder for "upsert"-ing
	//  one TeammateTaskTabStatus node.
	TeammateTaskTabStatusUpsertOne struct {
		create *TeammateTaskTabStatusCreate
	}

	// TeammateTaskTabStatusUpsert is the "OnConflict" setter.
	TeammateTaskTabStatusUpsert struct {
		*sql.UpdateSet
	}
)

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskTabStatusUpsert) SetWorkspaceID(v ulid.ID) *TeammateTaskTabStatusUpsert {
	u.Set(teammatetasktabstatus.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsert) UpdateWorkspaceID() *TeammateTaskTabStatusUpsert {
	u.SetExcluded(teammatetasktabstatus.FieldWorkspaceID)
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskTabStatusUpsert) SetTeammateID(v ulid.ID) *TeammateTaskTabStatusUpsert {
	u.Set(teammatetasktabstatus.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsert) UpdateTeammateID() *TeammateTaskTabStatusUpsert {
	u.SetExcluded(teammatetasktabstatus.FieldTeammateID)
	return u
}

// SetStatusCode sets the "status_code" field.
func (u *TeammateTaskTabStatusUpsert) SetStatusCode(v teammatetasktabstatus.StatusCode) *TeammateTaskTabStatusUpsert {
	u.Set(teammatetasktabstatus.FieldStatusCode, v)
	return u
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsert) UpdateStatusCode() *TeammateTaskTabStatusUpsert {
	u.SetExcluded(teammatetasktabstatus.FieldStatusCode)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskTabStatusUpsert) SetCreatedAt(v time.Time) *TeammateTaskTabStatusUpsert {
	u.Set(teammatetasktabstatus.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsert) UpdateCreatedAt() *TeammateTaskTabStatusUpsert {
	u.SetExcluded(teammatetasktabstatus.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskTabStatusUpsert) SetUpdatedAt(v time.Time) *TeammateTaskTabStatusUpsert {
	u.Set(teammatetasktabstatus.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsert) UpdateUpdatedAt() *TeammateTaskTabStatusUpsert {
	u.SetExcluded(teammatetasktabstatus.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TeammateTaskTabStatus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetasktabstatus.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskTabStatusUpsertOne) UpdateNewValues() *TeammateTaskTabStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(teammatetasktabstatus.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(teammatetasktabstatus.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(teammatetasktabstatus.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TeammateTaskTabStatus.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TeammateTaskTabStatusUpsertOne) Ignore() *TeammateTaskTabStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskTabStatusUpsertOne) DoNothing() *TeammateTaskTabStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskTabStatusCreate.OnConflict
// documentation for more info.
func (u *TeammateTaskTabStatusUpsertOne) Update(set func(*TeammateTaskTabStatusUpsert)) *TeammateTaskTabStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskTabStatusUpsert{UpdateSet: update})
	}))
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskTabStatusUpsertOne) SetWorkspaceID(v ulid.ID) *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertOne) UpdateWorkspaceID() *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskTabStatusUpsertOne) SetTeammateID(v ulid.ID) *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertOne) UpdateTeammateID() *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateTeammateID()
	})
}

// SetStatusCode sets the "status_code" field.
func (u *TeammateTaskTabStatusUpsertOne) SetStatusCode(v teammatetasktabstatus.StatusCode) *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetStatusCode(v)
	})
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertOne) UpdateStatusCode() *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateStatusCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskTabStatusUpsertOne) SetCreatedAt(v time.Time) *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertOne) UpdateCreatedAt() *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskTabStatusUpsertOne) SetUpdatedAt(v time.Time) *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertOne) UpdateUpdatedAt() *TeammateTaskTabStatusUpsertOne {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskTabStatusUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskTabStatusCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskTabStatusUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeammateTaskTabStatusUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TeammateTaskTabStatusUpsertOne.ID is not supported by MySQL driver. Use TeammateTaskTabStatusUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeammateTaskTabStatusUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeammateTaskTabStatusCreateBulk is the builder for creating many TeammateTaskTabStatus entities in bulk.
type TeammateTaskTabStatusCreateBulk struct {
	config
	builders []*TeammateTaskTabStatusCreate
	conflict []sql.ConflictOption
}

// Save creates the TeammateTaskTabStatus entities in the database.
func (tttscb *TeammateTaskTabStatusCreateBulk) Save(ctx context.Context) ([]*TeammateTaskTabStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tttscb.builders))
	nodes := make([]*TeammateTaskTabStatus, len(tttscb.builders))
	mutators := make([]Mutator, len(tttscb.builders))
	for i := range tttscb.builders {
		func(i int, root context.Context) {
			builder := tttscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeammateTaskTabStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, tttscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tttscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tttscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tttscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tttscb *TeammateTaskTabStatusCreateBulk) SaveX(ctx context.Context) []*TeammateTaskTabStatus {
	v, err := tttscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tttscb *TeammateTaskTabStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := tttscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tttscb *TeammateTaskTabStatusCreateBulk) ExecX(ctx context.Context) {
	if err := tttscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeammateTaskTabStatus.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeammateTaskTabStatusUpsert) {
//			SetWorkspaceID(v+v).
//		}).
//		Exec(ctx)
//
func (tttscb *TeammateTaskTabStatusCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeammateTaskTabStatusUpsertBulk {
	tttscb.conflict = opts
	return &TeammateTaskTabStatusUpsertBulk{
		create: tttscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeammateTaskTabStatus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tttscb *TeammateTaskTabStatusCreateBulk) OnConflictColumns(columns ...string) *TeammateTaskTabStatusUpsertBulk {
	tttscb.conflict = append(tttscb.conflict, sql.ConflictColumns(columns...))
	return &TeammateTaskTabStatusUpsertBulk{
		create: tttscb,
	}
}

// TeammateTaskTabStatusUpsertBulk is the builder for "upsert"-ing
// a bulk of TeammateTaskTabStatus nodes.
type TeammateTaskTabStatusUpsertBulk struct {
	create *TeammateTaskTabStatusCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TeammateTaskTabStatus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teammatetasktabstatus.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TeammateTaskTabStatusUpsertBulk) UpdateNewValues() *TeammateTaskTabStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(teammatetasktabstatus.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(teammatetasktabstatus.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(teammatetasktabstatus.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TeammateTaskTabStatus.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TeammateTaskTabStatusUpsertBulk) Ignore() *TeammateTaskTabStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeammateTaskTabStatusUpsertBulk) DoNothing() *TeammateTaskTabStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeammateTaskTabStatusCreateBulk.OnConflict
// documentation for more info.
func (u *TeammateTaskTabStatusUpsertBulk) Update(set func(*TeammateTaskTabStatusUpsert)) *TeammateTaskTabStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeammateTaskTabStatusUpsert{UpdateSet: update})
	}))
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *TeammateTaskTabStatusUpsertBulk) SetWorkspaceID(v ulid.ID) *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertBulk) UpdateWorkspaceID() *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TeammateTaskTabStatusUpsertBulk) SetTeammateID(v ulid.ID) *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertBulk) UpdateTeammateID() *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateTeammateID()
	})
}

// SetStatusCode sets the "status_code" field.
func (u *TeammateTaskTabStatusUpsertBulk) SetStatusCode(v teammatetasktabstatus.StatusCode) *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetStatusCode(v)
	})
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertBulk) UpdateStatusCode() *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateStatusCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TeammateTaskTabStatusUpsertBulk) SetCreatedAt(v time.Time) *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertBulk) UpdateCreatedAt() *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeammateTaskTabStatusUpsertBulk) SetUpdatedAt(v time.Time) *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeammateTaskTabStatusUpsertBulk) UpdateUpdatedAt() *TeammateTaskTabStatusUpsertBulk {
	return u.Update(func(s *TeammateTaskTabStatusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TeammateTaskTabStatusUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TeammateTaskTabStatusCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeammateTaskTabStatusCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeammateTaskTabStatusUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
