// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskSectionCreate is the builder for creating a ProjectTaskSection entity.
type ProjectTaskSectionCreate struct {
	config
	mutation *ProjectTaskSectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProjectID sets the "project_id" field.
func (ptsc *ProjectTaskSectionCreate) SetProjectID(u ulid.ID) *ProjectTaskSectionCreate {
	ptsc.mutation.SetProjectID(u)
	return ptsc
}

// SetName sets the "name" field.
func (ptsc *ProjectTaskSectionCreate) SetName(s string) *ProjectTaskSectionCreate {
	ptsc.mutation.SetName(s)
	return ptsc
}

// SetCreatedAt sets the "created_at" field.
func (ptsc *ProjectTaskSectionCreate) SetCreatedAt(t time.Time) *ProjectTaskSectionCreate {
	ptsc.mutation.SetCreatedAt(t)
	return ptsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptsc *ProjectTaskSectionCreate) SetNillableCreatedAt(t *time.Time) *ProjectTaskSectionCreate {
	if t != nil {
		ptsc.SetCreatedAt(*t)
	}
	return ptsc
}

// SetUpdatedAt sets the "updated_at" field.
func (ptsc *ProjectTaskSectionCreate) SetUpdatedAt(t time.Time) *ProjectTaskSectionCreate {
	ptsc.mutation.SetUpdatedAt(t)
	return ptsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ptsc *ProjectTaskSectionCreate) SetNillableUpdatedAt(t *time.Time) *ProjectTaskSectionCreate {
	if t != nil {
		ptsc.SetUpdatedAt(*t)
	}
	return ptsc
}

// SetID sets the "id" field.
func (ptsc *ProjectTaskSectionCreate) SetID(u ulid.ID) *ProjectTaskSectionCreate {
	ptsc.mutation.SetID(u)
	return ptsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ptsc *ProjectTaskSectionCreate) SetNillableID(u *ulid.ID) *ProjectTaskSectionCreate {
	if u != nil {
		ptsc.SetID(*u)
	}
	return ptsc
}

// SetProject sets the "project" edge to the Project entity.
func (ptsc *ProjectTaskSectionCreate) SetProject(p *Project) *ProjectTaskSectionCreate {
	return ptsc.SetProjectID(p.ID)
}

// AddProjectTaskIDs adds the "projectTasks" edge to the ProjectTask entity by IDs.
func (ptsc *ProjectTaskSectionCreate) AddProjectTaskIDs(ids ...ulid.ID) *ProjectTaskSectionCreate {
	ptsc.mutation.AddProjectTaskIDs(ids...)
	return ptsc
}

// AddProjectTasks adds the "projectTasks" edges to the ProjectTask entity.
func (ptsc *ProjectTaskSectionCreate) AddProjectTasks(p ...*ProjectTask) *ProjectTaskSectionCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptsc.AddProjectTaskIDs(ids...)
}

// Mutation returns the ProjectTaskSectionMutation object of the builder.
func (ptsc *ProjectTaskSectionCreate) Mutation() *ProjectTaskSectionMutation {
	return ptsc.mutation
}

// Save creates the ProjectTaskSection in the database.
func (ptsc *ProjectTaskSectionCreate) Save(ctx context.Context) (*ProjectTaskSection, error) {
	var (
		err  error
		node *ProjectTaskSection
	)
	ptsc.defaults()
	if len(ptsc.hooks) == 0 {
		if err = ptsc.check(); err != nil {
			return nil, err
		}
		node, err = ptsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptsc.check(); err != nil {
				return nil, err
			}
			ptsc.mutation = mutation
			if node, err = ptsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ptsc.hooks) - 1; i >= 0; i-- {
			if ptsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptsc *ProjectTaskSectionCreate) SaveX(ctx context.Context) *ProjectTaskSection {
	v, err := ptsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptsc *ProjectTaskSectionCreate) Exec(ctx context.Context) error {
	_, err := ptsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsc *ProjectTaskSectionCreate) ExecX(ctx context.Context) {
	if err := ptsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptsc *ProjectTaskSectionCreate) defaults() {
	if _, ok := ptsc.mutation.CreatedAt(); !ok {
		v := projecttasksection.DefaultCreatedAt()
		ptsc.mutation.SetCreatedAt(v)
	}
	if _, ok := ptsc.mutation.UpdatedAt(); !ok {
		v := projecttasksection.DefaultUpdatedAt()
		ptsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ptsc.mutation.ID(); !ok {
		v := projecttasksection.DefaultID()
		ptsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptsc *ProjectTaskSectionCreate) check() error {
	if _, ok := ptsc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "ProjectTaskSection.project_id"`)}
	}
	if _, ok := ptsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProjectTaskSection.name"`)}
	}
	if v, ok := ptsc.mutation.Name(); ok {
		if err := projecttasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProjectTaskSection.name": %w`, err)}
		}
	}
	if _, ok := ptsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProjectTaskSection.created_at"`)}
	}
	if _, ok := ptsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProjectTaskSection.updated_at"`)}
	}
	if _, ok := ptsc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "ProjectTaskSection.project"`)}
	}
	return nil
}

func (ptsc *ProjectTaskSectionCreate) sqlSave(ctx context.Context) (*ProjectTaskSection, error) {
	_node, _spec := ptsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptsc.driver, _spec); err != nil {
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

func (ptsc *ProjectTaskSectionCreate) createSpec() (*ProjectTaskSection, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectTaskSection{config: ptsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projecttasksection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttasksection.FieldID,
			},
		}
	)
	_spec.OnConflict = ptsc.conflict
	if id, ok := ptsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ptsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttasksection.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ptsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttasksection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ptsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttasksection.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ptsc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttasksection.ProjectTable,
			Columns: []string{projecttasksection.ProjectColumn},
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
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptsc.mutation.ProjectTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
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
//	client.ProjectTaskSection.Create().
//		SetProjectID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectTaskSectionUpsert) {
//			SetProjectID(v+v).
//		}).
//		Exec(ctx)
//
func (ptsc *ProjectTaskSectionCreate) OnConflict(opts ...sql.ConflictOption) *ProjectTaskSectionUpsertOne {
	ptsc.conflict = opts
	return &ProjectTaskSectionUpsertOne{
		create: ptsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectTaskSection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ptsc *ProjectTaskSectionCreate) OnConflictColumns(columns ...string) *ProjectTaskSectionUpsertOne {
	ptsc.conflict = append(ptsc.conflict, sql.ConflictColumns(columns...))
	return &ProjectTaskSectionUpsertOne{
		create: ptsc,
	}
}

type (
	// ProjectTaskSectionUpsertOne is the builder for "upsert"-ing
	//  one ProjectTaskSection node.
	ProjectTaskSectionUpsertOne struct {
		create *ProjectTaskSectionCreate
	}

	// ProjectTaskSectionUpsert is the "OnConflict" setter.
	ProjectTaskSectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetProjectID sets the "project_id" field.
func (u *ProjectTaskSectionUpsert) SetProjectID(v ulid.ID) *ProjectTaskSectionUpsert {
	u.Set(projecttasksection.FieldProjectID, v)
	return u
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsert) UpdateProjectID() *ProjectTaskSectionUpsert {
	u.SetExcluded(projecttasksection.FieldProjectID)
	return u
}

// SetName sets the "name" field.
func (u *ProjectTaskSectionUpsert) SetName(v string) *ProjectTaskSectionUpsert {
	u.Set(projecttasksection.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsert) UpdateName() *ProjectTaskSectionUpsert {
	u.SetExcluded(projecttasksection.FieldName)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProjectTaskSectionUpsert) SetCreatedAt(v time.Time) *ProjectTaskSectionUpsert {
	u.Set(projecttasksection.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsert) UpdateCreatedAt() *ProjectTaskSectionUpsert {
	u.SetExcluded(projecttasksection.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectTaskSectionUpsert) SetUpdatedAt(v time.Time) *ProjectTaskSectionUpsert {
	u.Set(projecttasksection.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsert) UpdateUpdatedAt() *ProjectTaskSectionUpsert {
	u.SetExcluded(projecttasksection.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ProjectTaskSection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(projecttasksection.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProjectTaskSectionUpsertOne) UpdateNewValues() *ProjectTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(projecttasksection.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(projecttasksection.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(projecttasksection.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ProjectTaskSection.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ProjectTaskSectionUpsertOne) Ignore() *ProjectTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectTaskSectionUpsertOne) DoNothing() *ProjectTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectTaskSectionCreate.OnConflict
// documentation for more info.
func (u *ProjectTaskSectionUpsertOne) Update(set func(*ProjectTaskSectionUpsert)) *ProjectTaskSectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectTaskSectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetProjectID sets the "project_id" field.
func (u *ProjectTaskSectionUpsertOne) SetProjectID(v ulid.ID) *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertOne) UpdateProjectID() *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateProjectID()
	})
}

// SetName sets the "name" field.
func (u *ProjectTaskSectionUpsertOne) SetName(v string) *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertOne) UpdateName() *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ProjectTaskSectionUpsertOne) SetCreatedAt(v time.Time) *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertOne) UpdateCreatedAt() *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectTaskSectionUpsertOne) SetUpdatedAt(v time.Time) *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertOne) UpdateUpdatedAt() *ProjectTaskSectionUpsertOne {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ProjectTaskSectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectTaskSectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectTaskSectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProjectTaskSectionUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ProjectTaskSectionUpsertOne.ID is not supported by MySQL driver. Use ProjectTaskSectionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectTaskSectionUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProjectTaskSectionCreateBulk is the builder for creating many ProjectTaskSection entities in bulk.
type ProjectTaskSectionCreateBulk struct {
	config
	builders []*ProjectTaskSectionCreate
	conflict []sql.ConflictOption
}

// Save creates the ProjectTaskSection entities in the database.
func (ptscb *ProjectTaskSectionCreateBulk) Save(ctx context.Context) ([]*ProjectTaskSection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptscb.builders))
	nodes := make([]*ProjectTaskSection, len(ptscb.builders))
	mutators := make([]Mutator, len(ptscb.builders))
	for i := range ptscb.builders {
		func(i int, root context.Context) {
			builder := ptscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectTaskSectionMutation)
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
					_, err = mutators[i+1].Mutate(root, ptscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ptscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptscb *ProjectTaskSectionCreateBulk) SaveX(ctx context.Context) []*ProjectTaskSection {
	v, err := ptscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptscb *ProjectTaskSectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ptscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptscb *ProjectTaskSectionCreateBulk) ExecX(ctx context.Context) {
	if err := ptscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProjectTaskSection.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectTaskSectionUpsert) {
//			SetProjectID(v+v).
//		}).
//		Exec(ctx)
//
func (ptscb *ProjectTaskSectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProjectTaskSectionUpsertBulk {
	ptscb.conflict = opts
	return &ProjectTaskSectionUpsertBulk{
		create: ptscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectTaskSection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ptscb *ProjectTaskSectionCreateBulk) OnConflictColumns(columns ...string) *ProjectTaskSectionUpsertBulk {
	ptscb.conflict = append(ptscb.conflict, sql.ConflictColumns(columns...))
	return &ProjectTaskSectionUpsertBulk{
		create: ptscb,
	}
}

// ProjectTaskSectionUpsertBulk is the builder for "upsert"-ing
// a bulk of ProjectTaskSection nodes.
type ProjectTaskSectionUpsertBulk struct {
	create *ProjectTaskSectionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProjectTaskSection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(projecttasksection.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProjectTaskSectionUpsertBulk) UpdateNewValues() *ProjectTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(projecttasksection.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(projecttasksection.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(projecttasksection.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProjectTaskSection.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ProjectTaskSectionUpsertBulk) Ignore() *ProjectTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectTaskSectionUpsertBulk) DoNothing() *ProjectTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectTaskSectionCreateBulk.OnConflict
// documentation for more info.
func (u *ProjectTaskSectionUpsertBulk) Update(set func(*ProjectTaskSectionUpsert)) *ProjectTaskSectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectTaskSectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetProjectID sets the "project_id" field.
func (u *ProjectTaskSectionUpsertBulk) SetProjectID(v ulid.ID) *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertBulk) UpdateProjectID() *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateProjectID()
	})
}

// SetName sets the "name" field.
func (u *ProjectTaskSectionUpsertBulk) SetName(v string) *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertBulk) UpdateName() *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ProjectTaskSectionUpsertBulk) SetCreatedAt(v time.Time) *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertBulk) UpdateCreatedAt() *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectTaskSectionUpsertBulk) SetUpdatedAt(v time.Time) *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectTaskSectionUpsertBulk) UpdateUpdatedAt() *ProjectTaskSectionUpsertBulk {
	return u.Update(func(s *ProjectTaskSectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ProjectTaskSectionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProjectTaskSectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectTaskSectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectTaskSectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
