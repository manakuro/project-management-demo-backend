// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/archivedworkspaceactivity"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedWorkspaceActivityCreate is the builder for creating a ArchivedWorkspaceActivity entity.
type ArchivedWorkspaceActivityCreate struct {
	config
	mutation *ArchivedWorkspaceActivityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetActivityTypeID sets the "activity_type_id" field.
func (awac *ArchivedWorkspaceActivityCreate) SetActivityTypeID(u ulid.ID) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetActivityTypeID(u)
	return awac
}

// SetWorkspaceID sets the "workspace_id" field.
func (awac *ArchivedWorkspaceActivityCreate) SetWorkspaceID(u ulid.ID) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetWorkspaceID(u)
	return awac
}

// SetProjectID sets the "project_id" field.
func (awac *ArchivedWorkspaceActivityCreate) SetProjectID(u ulid.ID) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetProjectID(u)
	return awac
}

// SetTeammateID sets the "teammate_id" field.
func (awac *ArchivedWorkspaceActivityCreate) SetTeammateID(u ulid.ID) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetTeammateID(u)
	return awac
}

// SetCreatedAt sets the "created_at" field.
func (awac *ArchivedWorkspaceActivityCreate) SetCreatedAt(t time.Time) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetCreatedAt(t)
	return awac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (awac *ArchivedWorkspaceActivityCreate) SetNillableCreatedAt(t *time.Time) *ArchivedWorkspaceActivityCreate {
	if t != nil {
		awac.SetCreatedAt(*t)
	}
	return awac
}

// SetUpdatedAt sets the "updated_at" field.
func (awac *ArchivedWorkspaceActivityCreate) SetUpdatedAt(t time.Time) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetUpdatedAt(t)
	return awac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (awac *ArchivedWorkspaceActivityCreate) SetNillableUpdatedAt(t *time.Time) *ArchivedWorkspaceActivityCreate {
	if t != nil {
		awac.SetUpdatedAt(*t)
	}
	return awac
}

// SetID sets the "id" field.
func (awac *ArchivedWorkspaceActivityCreate) SetID(u ulid.ID) *ArchivedWorkspaceActivityCreate {
	awac.mutation.SetID(u)
	return awac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (awac *ArchivedWorkspaceActivityCreate) SetNillableID(u *ulid.ID) *ArchivedWorkspaceActivityCreate {
	if u != nil {
		awac.SetID(*u)
	}
	return awac
}

// SetActivityType sets the "activityType" edge to the ActivityType entity.
func (awac *ArchivedWorkspaceActivityCreate) SetActivityType(a *ActivityType) *ArchivedWorkspaceActivityCreate {
	return awac.SetActivityTypeID(a.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (awac *ArchivedWorkspaceActivityCreate) SetWorkspace(w *Workspace) *ArchivedWorkspaceActivityCreate {
	return awac.SetWorkspaceID(w.ID)
}

// SetProject sets the "project" edge to the Project entity.
func (awac *ArchivedWorkspaceActivityCreate) SetProject(p *Project) *ArchivedWorkspaceActivityCreate {
	return awac.SetProjectID(p.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (awac *ArchivedWorkspaceActivityCreate) SetTeammate(t *Teammate) *ArchivedWorkspaceActivityCreate {
	return awac.SetTeammateID(t.ID)
}

// Mutation returns the ArchivedWorkspaceActivityMutation object of the builder.
func (awac *ArchivedWorkspaceActivityCreate) Mutation() *ArchivedWorkspaceActivityMutation {
	return awac.mutation
}

// Save creates the ArchivedWorkspaceActivity in the database.
func (awac *ArchivedWorkspaceActivityCreate) Save(ctx context.Context) (*ArchivedWorkspaceActivity, error) {
	var (
		err  error
		node *ArchivedWorkspaceActivity
	)
	awac.defaults()
	if len(awac.hooks) == 0 {
		if err = awac.check(); err != nil {
			return nil, err
		}
		node, err = awac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedWorkspaceActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awac.check(); err != nil {
				return nil, err
			}
			awac.mutation = mutation
			if node, err = awac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(awac.hooks) - 1; i >= 0; i-- {
			if awac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = awac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (awac *ArchivedWorkspaceActivityCreate) SaveX(ctx context.Context) *ArchivedWorkspaceActivity {
	v, err := awac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (awac *ArchivedWorkspaceActivityCreate) Exec(ctx context.Context) error {
	_, err := awac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awac *ArchivedWorkspaceActivityCreate) ExecX(ctx context.Context) {
	if err := awac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (awac *ArchivedWorkspaceActivityCreate) defaults() {
	if _, ok := awac.mutation.CreatedAt(); !ok {
		v := archivedworkspaceactivity.DefaultCreatedAt()
		awac.mutation.SetCreatedAt(v)
	}
	if _, ok := awac.mutation.UpdatedAt(); !ok {
		v := archivedworkspaceactivity.DefaultUpdatedAt()
		awac.mutation.SetUpdatedAt(v)
	}
	if _, ok := awac.mutation.ID(); !ok {
		v := archivedworkspaceactivity.DefaultID()
		awac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awac *ArchivedWorkspaceActivityCreate) check() error {
	if _, ok := awac.mutation.ActivityTypeID(); !ok {
		return &ValidationError{Name: "activity_type_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.activity_type_id"`)}
	}
	if _, ok := awac.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.workspace_id"`)}
	}
	if _, ok := awac.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.project_id"`)}
	}
	if _, ok := awac.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.teammate_id"`)}
	}
	if _, ok := awac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.created_at"`)}
	}
	if _, ok := awac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivity.updated_at"`)}
	}
	if _, ok := awac.mutation.ActivityTypeID(); !ok {
		return &ValidationError{Name: "activityType", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivity.activityType"`)}
	}
	if _, ok := awac.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivity.workspace"`)}
	}
	if _, ok := awac.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivity.project"`)}
	}
	if _, ok := awac.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivity.teammate"`)}
	}
	return nil
}

func (awac *ArchivedWorkspaceActivityCreate) sqlSave(ctx context.Context) (*ArchivedWorkspaceActivity, error) {
	_node, _spec := awac.createSpec()
	if err := sqlgraph.CreateNode(ctx, awac.driver, _spec); err != nil {
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

func (awac *ArchivedWorkspaceActivityCreate) createSpec() (*ArchivedWorkspaceActivity, *sqlgraph.CreateSpec) {
	var (
		_node = &ArchivedWorkspaceActivity{config: awac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: archivedworkspaceactivity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedworkspaceactivity.FieldID,
			},
		}
	)
	_spec.OnConflict = awac.conflict
	if id, ok := awac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := awac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedworkspaceactivity.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := awac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedworkspaceactivity.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := awac.mutation.ActivityTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ActivityTypeTable,
			Columns: []string{archivedworkspaceactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActivityTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := awac.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.WorkspaceTable,
			Columns: []string{archivedworkspaceactivity.WorkspaceColumn},
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
	if nodes := awac.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ProjectTable,
			Columns: []string{archivedworkspaceactivity.ProjectColumn},
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
	if nodes := awac.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.TeammateTable,
			Columns: []string{archivedworkspaceactivity.TeammateColumn},
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
//	client.ArchivedWorkspaceActivity.Create().
//		SetActivityTypeID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedWorkspaceActivityUpsert) {
//			SetActivityTypeID(v+v).
//		}).
//		Exec(ctx)
//
func (awac *ArchivedWorkspaceActivityCreate) OnConflict(opts ...sql.ConflictOption) *ArchivedWorkspaceActivityUpsertOne {
	awac.conflict = opts
	return &ArchivedWorkspaceActivityUpsertOne{
		create: awac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (awac *ArchivedWorkspaceActivityCreate) OnConflictColumns(columns ...string) *ArchivedWorkspaceActivityUpsertOne {
	awac.conflict = append(awac.conflict, sql.ConflictColumns(columns...))
	return &ArchivedWorkspaceActivityUpsertOne{
		create: awac,
	}
}

type (
	// ArchivedWorkspaceActivityUpsertOne is the builder for "upsert"-ing
	//  one ArchivedWorkspaceActivity node.
	ArchivedWorkspaceActivityUpsertOne struct {
		create *ArchivedWorkspaceActivityCreate
	}

	// ArchivedWorkspaceActivityUpsert is the "OnConflict" setter.
	ArchivedWorkspaceActivityUpsert struct {
		*sql.UpdateSet
	}
)

// SetActivityTypeID sets the "activity_type_id" field.
func (u *ArchivedWorkspaceActivityUpsert) SetActivityTypeID(v ulid.ID) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldActivityTypeID, v)
	return u
}

// UpdateActivityTypeID sets the "activity_type_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateActivityTypeID() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldActivityTypeID)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *ArchivedWorkspaceActivityUpsert) SetWorkspaceID(v ulid.ID) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateWorkspaceID() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldWorkspaceID)
	return u
}

// SetProjectID sets the "project_id" field.
func (u *ArchivedWorkspaceActivityUpsert) SetProjectID(v ulid.ID) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldProjectID, v)
	return u
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateProjectID() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldProjectID)
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *ArchivedWorkspaceActivityUpsert) SetTeammateID(v ulid.ID) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateTeammateID() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldTeammateID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityUpsert) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateCreatedAt() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityUpsert) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityUpsert {
	u.Set(archivedworkspaceactivity.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsert) UpdateUpdatedAt() *ArchivedWorkspaceActivityUpsert {
	u.SetExcluded(archivedworkspaceactivity.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedworkspaceactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateNewValues() *ArchivedWorkspaceActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(archivedworkspaceactivity.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(archivedworkspaceactivity.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(archivedworkspaceactivity.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ArchivedWorkspaceActivity.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ArchivedWorkspaceActivityUpsertOne) Ignore() *ArchivedWorkspaceActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedWorkspaceActivityUpsertOne) DoNothing() *ArchivedWorkspaceActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedWorkspaceActivityCreate.OnConflict
// documentation for more info.
func (u *ArchivedWorkspaceActivityUpsertOne) Update(set func(*ArchivedWorkspaceActivityUpsert)) *ArchivedWorkspaceActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedWorkspaceActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetActivityTypeID sets the "activity_type_id" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetActivityTypeID(v ulid.ID) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetActivityTypeID(v)
	})
}

// UpdateActivityTypeID sets the "activity_type_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateActivityTypeID() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateActivityTypeID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetWorkspaceID(v ulid.ID) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateWorkspaceID() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetProjectID sets the "project_id" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetProjectID(v ulid.ID) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateProjectID() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateProjectID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetTeammateID(v ulid.ID) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateTeammateID() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateCreatedAt() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityUpsertOne) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertOne) UpdateUpdatedAt() *ArchivedWorkspaceActivityUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedWorkspaceActivityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedWorkspaceActivityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArchivedWorkspaceActivityUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArchivedWorkspaceActivityUpsertOne.ID is not supported by MySQL driver. Use ArchivedWorkspaceActivityUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArchivedWorkspaceActivityCreateBulk is the builder for creating many ArchivedWorkspaceActivity entities in bulk.
type ArchivedWorkspaceActivityCreateBulk struct {
	config
	builders []*ArchivedWorkspaceActivityCreate
	conflict []sql.ConflictOption
}

// Save creates the ArchivedWorkspaceActivity entities in the database.
func (awacb *ArchivedWorkspaceActivityCreateBulk) Save(ctx context.Context) ([]*ArchivedWorkspaceActivity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(awacb.builders))
	nodes := make([]*ArchivedWorkspaceActivity, len(awacb.builders))
	mutators := make([]Mutator, len(awacb.builders))
	for i := range awacb.builders {
		func(i int, root context.Context) {
			builder := awacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArchivedWorkspaceActivityMutation)
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
					_, err = mutators[i+1].Mutate(root, awacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = awacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, awacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, awacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (awacb *ArchivedWorkspaceActivityCreateBulk) SaveX(ctx context.Context) []*ArchivedWorkspaceActivity {
	v, err := awacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (awacb *ArchivedWorkspaceActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := awacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awacb *ArchivedWorkspaceActivityCreateBulk) ExecX(ctx context.Context) {
	if err := awacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArchivedWorkspaceActivity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedWorkspaceActivityUpsert) {
//			SetActivityTypeID(v+v).
//		}).
//		Exec(ctx)
//
func (awacb *ArchivedWorkspaceActivityCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArchivedWorkspaceActivityUpsertBulk {
	awacb.conflict = opts
	return &ArchivedWorkspaceActivityUpsertBulk{
		create: awacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (awacb *ArchivedWorkspaceActivityCreateBulk) OnConflictColumns(columns ...string) *ArchivedWorkspaceActivityUpsertBulk {
	awacb.conflict = append(awacb.conflict, sql.ConflictColumns(columns...))
	return &ArchivedWorkspaceActivityUpsertBulk{
		create: awacb,
	}
}

// ArchivedWorkspaceActivityUpsertBulk is the builder for "upsert"-ing
// a bulk of ArchivedWorkspaceActivity nodes.
type ArchivedWorkspaceActivityUpsertBulk struct {
	create *ArchivedWorkspaceActivityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedworkspaceactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateNewValues() *ArchivedWorkspaceActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(archivedworkspaceactivity.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(archivedworkspaceactivity.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(archivedworkspaceactivity.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityUpsertBulk) Ignore() *ArchivedWorkspaceActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedWorkspaceActivityUpsertBulk) DoNothing() *ArchivedWorkspaceActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedWorkspaceActivityCreateBulk.OnConflict
// documentation for more info.
func (u *ArchivedWorkspaceActivityUpsertBulk) Update(set func(*ArchivedWorkspaceActivityUpsert)) *ArchivedWorkspaceActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedWorkspaceActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetActivityTypeID sets the "activity_type_id" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetActivityTypeID(v ulid.ID) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetActivityTypeID(v)
	})
}

// UpdateActivityTypeID sets the "activity_type_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateActivityTypeID() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateActivityTypeID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetWorkspaceID(v ulid.ID) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateWorkspaceID() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetProjectID sets the "project_id" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetProjectID(v ulid.ID) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateProjectID() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateProjectID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetTeammateID(v ulid.ID) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateTeammateID() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateCreatedAt() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityUpsertBulk) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityUpsertBulk) UpdateUpdatedAt() *ArchivedWorkspaceActivityUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedWorkspaceActivityUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArchivedWorkspaceActivityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedWorkspaceActivityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
