// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/teammatetaskliststatus"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskListCompletedStatusCreate is the builder for creating a TaskListCompletedStatus entity.
type TaskListCompletedStatusCreate struct {
	config
	mutation *TaskListCompletedStatusMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tlcsc *TaskListCompletedStatusCreate) SetName(s string) *TaskListCompletedStatusCreate {
	tlcsc.mutation.SetName(s)
	return tlcsc
}

// SetStatusCode sets the "status_code" field.
func (tlcsc *TaskListCompletedStatusCreate) SetStatusCode(tc tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusCreate {
	tlcsc.mutation.SetStatusCode(tc)
	return tlcsc
}

// SetCreatedAt sets the "created_at" field.
func (tlcsc *TaskListCompletedStatusCreate) SetCreatedAt(t time.Time) *TaskListCompletedStatusCreate {
	tlcsc.mutation.SetCreatedAt(t)
	return tlcsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tlcsc *TaskListCompletedStatusCreate) SetNillableCreatedAt(t *time.Time) *TaskListCompletedStatusCreate {
	if t != nil {
		tlcsc.SetCreatedAt(*t)
	}
	return tlcsc
}

// SetUpdatedAt sets the "updated_at" field.
func (tlcsc *TaskListCompletedStatusCreate) SetUpdatedAt(t time.Time) *TaskListCompletedStatusCreate {
	tlcsc.mutation.SetUpdatedAt(t)
	return tlcsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tlcsc *TaskListCompletedStatusCreate) SetNillableUpdatedAt(t *time.Time) *TaskListCompletedStatusCreate {
	if t != nil {
		tlcsc.SetUpdatedAt(*t)
	}
	return tlcsc
}

// SetID sets the "id" field.
func (tlcsc *TaskListCompletedStatusCreate) SetID(u ulid.ID) *TaskListCompletedStatusCreate {
	tlcsc.mutation.SetID(u)
	return tlcsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tlcsc *TaskListCompletedStatusCreate) SetNillableID(u *ulid.ID) *TaskListCompletedStatusCreate {
	if u != nil {
		tlcsc.SetID(*u)
	}
	return tlcsc
}

// AddTeammateTaskListStatuseIDs adds the "teammateTaskListStatuses" edge to the TeammateTaskListStatus entity by IDs.
func (tlcsc *TaskListCompletedStatusCreate) AddTeammateTaskListStatuseIDs(ids ...ulid.ID) *TaskListCompletedStatusCreate {
	tlcsc.mutation.AddTeammateTaskListStatuseIDs(ids...)
	return tlcsc
}

// AddTeammateTaskListStatuses adds the "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (tlcsc *TaskListCompletedStatusCreate) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListCompletedStatusCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlcsc.AddTeammateTaskListStatuseIDs(ids...)
}

// AddProjectTaskListStatuseIDs adds the "projectTaskListStatuses" edge to the ProjectTaskListStatus entity by IDs.
func (tlcsc *TaskListCompletedStatusCreate) AddProjectTaskListStatuseIDs(ids ...ulid.ID) *TaskListCompletedStatusCreate {
	tlcsc.mutation.AddProjectTaskListStatuseIDs(ids...)
	return tlcsc
}

// AddProjectTaskListStatuses adds the "projectTaskListStatuses" edges to the ProjectTaskListStatus entity.
func (tlcsc *TaskListCompletedStatusCreate) AddProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListCompletedStatusCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlcsc.AddProjectTaskListStatuseIDs(ids...)
}

// Mutation returns the TaskListCompletedStatusMutation object of the builder.
func (tlcsc *TaskListCompletedStatusCreate) Mutation() *TaskListCompletedStatusMutation {
	return tlcsc.mutation
}

// Save creates the TaskListCompletedStatus in the database.
func (tlcsc *TaskListCompletedStatusCreate) Save(ctx context.Context) (*TaskListCompletedStatus, error) {
	var (
		err  error
		node *TaskListCompletedStatus
	)
	tlcsc.defaults()
	if len(tlcsc.hooks) == 0 {
		if err = tlcsc.check(); err != nil {
			return nil, err
		}
		node, err = tlcsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskListCompletedStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlcsc.check(); err != nil {
				return nil, err
			}
			tlcsc.mutation = mutation
			if node, err = tlcsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tlcsc.hooks) - 1; i >= 0; i-- {
			if tlcsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlcsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlcsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tlcsc *TaskListCompletedStatusCreate) SaveX(ctx context.Context) *TaskListCompletedStatus {
	v, err := tlcsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlcsc *TaskListCompletedStatusCreate) Exec(ctx context.Context) error {
	_, err := tlcsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcsc *TaskListCompletedStatusCreate) ExecX(ctx context.Context) {
	if err := tlcsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tlcsc *TaskListCompletedStatusCreate) defaults() {
	if _, ok := tlcsc.mutation.CreatedAt(); !ok {
		v := tasklistcompletedstatus.DefaultCreatedAt()
		tlcsc.mutation.SetCreatedAt(v)
	}
	if _, ok := tlcsc.mutation.UpdatedAt(); !ok {
		v := tasklistcompletedstatus.DefaultUpdatedAt()
		tlcsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tlcsc.mutation.ID(); !ok {
		v := tasklistcompletedstatus.DefaultID()
		tlcsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlcsc *TaskListCompletedStatusCreate) check() error {
	if _, ok := tlcsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tlcsc.mutation.Name(); ok {
		if err := tasklistcompletedstatus.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tlcsc.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`ent: missing required field "status_code"`)}
	}
	if v, ok := tlcsc.mutation.StatusCode(); ok {
		if err := tasklistcompletedstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf(`ent: validator failed for field "status_code": %w`, err)}
		}
	}
	if _, ok := tlcsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tlcsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (tlcsc *TaskListCompletedStatusCreate) sqlSave(ctx context.Context) (*TaskListCompletedStatus, error) {
	_node, _spec := tlcsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tlcsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(ulid.ID)
	}
	return _node, nil
}

func (tlcsc *TaskListCompletedStatusCreate) createSpec() (*TaskListCompletedStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskListCompletedStatus{config: tlcsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tasklistcompletedstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistcompletedstatus.FieldID,
			},
		}
	)
	_spec.OnConflict = tlcsc.conflict
	if id, ok := tlcsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tlcsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklistcompletedstatus.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tlcsc.mutation.StatusCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tasklistcompletedstatus.FieldStatusCode,
		})
		_node.StatusCode = value
	}
	if value, ok := tlcsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tasklistcompletedstatus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tlcsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tasklistcompletedstatus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tlcsc.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tlcsc.mutation.ProjectTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
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
//	client.TaskListCompletedStatus.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskListCompletedStatusUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tlcsc *TaskListCompletedStatusCreate) OnConflict(opts ...sql.ConflictOption) *TaskListCompletedStatusUpsertOne {
	tlcsc.conflict = opts
	return &TaskListCompletedStatusUpsertOne{
		create: tlcsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskListCompletedStatus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tlcsc *TaskListCompletedStatusCreate) OnConflictColumns(columns ...string) *TaskListCompletedStatusUpsertOne {
	tlcsc.conflict = append(tlcsc.conflict, sql.ConflictColumns(columns...))
	return &TaskListCompletedStatusUpsertOne{
		create: tlcsc,
	}
}

type (
	// TaskListCompletedStatusUpsertOne is the builder for "upsert"-ing
	//  one TaskListCompletedStatus node.
	TaskListCompletedStatusUpsertOne struct {
		create *TaskListCompletedStatusCreate
	}

	// TaskListCompletedStatusUpsert is the "OnConflict" setter.
	TaskListCompletedStatusUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TaskListCompletedStatusUpsert) SetName(v string) *TaskListCompletedStatusUpsert {
	u.Set(tasklistcompletedstatus.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsert) UpdateName() *TaskListCompletedStatusUpsert {
	u.SetExcluded(tasklistcompletedstatus.FieldName)
	return u
}

// SetStatusCode sets the "status_code" field.
func (u *TaskListCompletedStatusUpsert) SetStatusCode(v tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusUpsert {
	u.Set(tasklistcompletedstatus.FieldStatusCode, v)
	return u
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsert) UpdateStatusCode() *TaskListCompletedStatusUpsert {
	u.SetExcluded(tasklistcompletedstatus.FieldStatusCode)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskListCompletedStatusUpsert) SetCreatedAt(v time.Time) *TaskListCompletedStatusUpsert {
	u.Set(tasklistcompletedstatus.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsert) UpdateCreatedAt() *TaskListCompletedStatusUpsert {
	u.SetExcluded(tasklistcompletedstatus.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskListCompletedStatusUpsert) SetUpdatedAt(v time.Time) *TaskListCompletedStatusUpsert {
	u.Set(tasklistcompletedstatus.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsert) UpdateUpdatedAt() *TaskListCompletedStatusUpsert {
	u.SetExcluded(tasklistcompletedstatus.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TaskListCompletedStatus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tasklistcompletedstatus.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskListCompletedStatusUpsertOne) UpdateNewValues() *TaskListCompletedStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tasklistcompletedstatus.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TaskListCompletedStatus.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TaskListCompletedStatusUpsertOne) Ignore() *TaskListCompletedStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskListCompletedStatusUpsertOne) DoNothing() *TaskListCompletedStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskListCompletedStatusCreate.OnConflict
// documentation for more info.
func (u *TaskListCompletedStatusUpsertOne) Update(set func(*TaskListCompletedStatusUpsert)) *TaskListCompletedStatusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskListCompletedStatusUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TaskListCompletedStatusUpsertOne) SetName(v string) *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertOne) UpdateName() *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateName()
	})
}

// SetStatusCode sets the "status_code" field.
func (u *TaskListCompletedStatusUpsertOne) SetStatusCode(v tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetStatusCode(v)
	})
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertOne) UpdateStatusCode() *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateStatusCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskListCompletedStatusUpsertOne) SetCreatedAt(v time.Time) *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertOne) UpdateCreatedAt() *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskListCompletedStatusUpsertOne) SetUpdatedAt(v time.Time) *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertOne) UpdateUpdatedAt() *TaskListCompletedStatusUpsertOne {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskListCompletedStatusUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskListCompletedStatusCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskListCompletedStatusUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskListCompletedStatusUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TaskListCompletedStatusUpsertOne.ID is not supported by MySQL driver. Use TaskListCompletedStatusUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskListCompletedStatusUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskListCompletedStatusCreateBulk is the builder for creating many TaskListCompletedStatus entities in bulk.
type TaskListCompletedStatusCreateBulk struct {
	config
	builders []*TaskListCompletedStatusCreate
	conflict []sql.ConflictOption
}

// Save creates the TaskListCompletedStatus entities in the database.
func (tlcscb *TaskListCompletedStatusCreateBulk) Save(ctx context.Context) ([]*TaskListCompletedStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tlcscb.builders))
	nodes := make([]*TaskListCompletedStatus, len(tlcscb.builders))
	mutators := make([]Mutator, len(tlcscb.builders))
	for i := range tlcscb.builders {
		func(i int, root context.Context) {
			builder := tlcscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskListCompletedStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, tlcscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tlcscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tlcscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tlcscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tlcscb *TaskListCompletedStatusCreateBulk) SaveX(ctx context.Context) []*TaskListCompletedStatus {
	v, err := tlcscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlcscb *TaskListCompletedStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := tlcscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcscb *TaskListCompletedStatusCreateBulk) ExecX(ctx context.Context) {
	if err := tlcscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskListCompletedStatus.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskListCompletedStatusUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tlcscb *TaskListCompletedStatusCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskListCompletedStatusUpsertBulk {
	tlcscb.conflict = opts
	return &TaskListCompletedStatusUpsertBulk{
		create: tlcscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskListCompletedStatus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tlcscb *TaskListCompletedStatusCreateBulk) OnConflictColumns(columns ...string) *TaskListCompletedStatusUpsertBulk {
	tlcscb.conflict = append(tlcscb.conflict, sql.ConflictColumns(columns...))
	return &TaskListCompletedStatusUpsertBulk{
		create: tlcscb,
	}
}

// TaskListCompletedStatusUpsertBulk is the builder for "upsert"-ing
// a bulk of TaskListCompletedStatus nodes.
type TaskListCompletedStatusUpsertBulk struct {
	create *TaskListCompletedStatusCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TaskListCompletedStatus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tasklistcompletedstatus.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskListCompletedStatusUpsertBulk) UpdateNewValues() *TaskListCompletedStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tasklistcompletedstatus.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TaskListCompletedStatus.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TaskListCompletedStatusUpsertBulk) Ignore() *TaskListCompletedStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskListCompletedStatusUpsertBulk) DoNothing() *TaskListCompletedStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskListCompletedStatusCreateBulk.OnConflict
// documentation for more info.
func (u *TaskListCompletedStatusUpsertBulk) Update(set func(*TaskListCompletedStatusUpsert)) *TaskListCompletedStatusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskListCompletedStatusUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TaskListCompletedStatusUpsertBulk) SetName(v string) *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertBulk) UpdateName() *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateName()
	})
}

// SetStatusCode sets the "status_code" field.
func (u *TaskListCompletedStatusUpsertBulk) SetStatusCode(v tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetStatusCode(v)
	})
}

// UpdateStatusCode sets the "status_code" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertBulk) UpdateStatusCode() *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateStatusCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskListCompletedStatusUpsertBulk) SetCreatedAt(v time.Time) *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertBulk) UpdateCreatedAt() *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskListCompletedStatusUpsertBulk) SetUpdatedAt(v time.Time) *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskListCompletedStatusUpsertBulk) UpdateUpdatedAt() *TaskListCompletedStatusUpsertBulk {
	return u.Update(func(s *TaskListCompletedStatusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskListCompletedStatusUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskListCompletedStatusCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskListCompletedStatusCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskListCompletedStatusUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
