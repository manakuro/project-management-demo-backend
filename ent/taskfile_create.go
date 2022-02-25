// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfile"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFileCreate is the builder for creating a TaskFile entity.
type TaskFileCreate struct {
	config
	mutation *TaskFileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProjectID sets the "project_id" field.
func (tfc *TaskFileCreate) SetProjectID(u ulid.ID) *TaskFileCreate {
	tfc.mutation.SetProjectID(u)
	return tfc
}

// SetTaskID sets the "task_id" field.
func (tfc *TaskFileCreate) SetTaskID(u ulid.ID) *TaskFileCreate {
	tfc.mutation.SetTaskID(u)
	return tfc
}

// SetTaskFeedID sets the "task_feed_id" field.
func (tfc *TaskFileCreate) SetTaskFeedID(u ulid.ID) *TaskFileCreate {
	tfc.mutation.SetTaskFeedID(u)
	return tfc
}

// SetFileTypeID sets the "file_type_id" field.
func (tfc *TaskFileCreate) SetFileTypeID(u ulid.ID) *TaskFileCreate {
	tfc.mutation.SetFileTypeID(u)
	return tfc
}

// SetName sets the "name" field.
func (tfc *TaskFileCreate) SetName(s string) *TaskFileCreate {
	tfc.mutation.SetName(s)
	return tfc
}

// SetSrc sets the "src" field.
func (tfc *TaskFileCreate) SetSrc(s string) *TaskFileCreate {
	tfc.mutation.SetSrc(s)
	return tfc
}

// SetAttached sets the "attached" field.
func (tfc *TaskFileCreate) SetAttached(b bool) *TaskFileCreate {
	tfc.mutation.SetAttached(b)
	return tfc
}

// SetNillableAttached sets the "attached" field if the given value is not nil.
func (tfc *TaskFileCreate) SetNillableAttached(b *bool) *TaskFileCreate {
	if b != nil {
		tfc.SetAttached(*b)
	}
	return tfc
}

// SetCreatedAt sets the "created_at" field.
func (tfc *TaskFileCreate) SetCreatedAt(t time.Time) *TaskFileCreate {
	tfc.mutation.SetCreatedAt(t)
	return tfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tfc *TaskFileCreate) SetNillableCreatedAt(t *time.Time) *TaskFileCreate {
	if t != nil {
		tfc.SetCreatedAt(*t)
	}
	return tfc
}

// SetUpdatedAt sets the "updated_at" field.
func (tfc *TaskFileCreate) SetUpdatedAt(t time.Time) *TaskFileCreate {
	tfc.mutation.SetUpdatedAt(t)
	return tfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tfc *TaskFileCreate) SetNillableUpdatedAt(t *time.Time) *TaskFileCreate {
	if t != nil {
		tfc.SetUpdatedAt(*t)
	}
	return tfc
}

// SetID sets the "id" field.
func (tfc *TaskFileCreate) SetID(u ulid.ID) *TaskFileCreate {
	tfc.mutation.SetID(u)
	return tfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tfc *TaskFileCreate) SetNillableID(u *ulid.ID) *TaskFileCreate {
	if u != nil {
		tfc.SetID(*u)
	}
	return tfc
}

// SetProject sets the "project" edge to the Project entity.
func (tfc *TaskFileCreate) SetProject(p *Project) *TaskFileCreate {
	return tfc.SetProjectID(p.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (tfc *TaskFileCreate) SetTask(t *Task) *TaskFileCreate {
	return tfc.SetTaskID(t.ID)
}

// SetTaskFeed sets the "taskFeed" edge to the TaskFeed entity.
func (tfc *TaskFileCreate) SetTaskFeed(t *TaskFeed) *TaskFileCreate {
	return tfc.SetTaskFeedID(t.ID)
}

// SetFileType sets the "fileType" edge to the FileType entity.
func (tfc *TaskFileCreate) SetFileType(f *FileType) *TaskFileCreate {
	return tfc.SetFileTypeID(f.ID)
}

// Mutation returns the TaskFileMutation object of the builder.
func (tfc *TaskFileCreate) Mutation() *TaskFileMutation {
	return tfc.mutation
}

// Save creates the TaskFile in the database.
func (tfc *TaskFileCreate) Save(ctx context.Context) (*TaskFile, error) {
	var (
		err  error
		node *TaskFile
	)
	tfc.defaults()
	if len(tfc.hooks) == 0 {
		if err = tfc.check(); err != nil {
			return nil, err
		}
		node, err = tfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tfc.check(); err != nil {
				return nil, err
			}
			tfc.mutation = mutation
			if node, err = tfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tfc.hooks) - 1; i >= 0; i-- {
			if tfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tfc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tfc *TaskFileCreate) SaveX(ctx context.Context) *TaskFile {
	v, err := tfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfc *TaskFileCreate) Exec(ctx context.Context) error {
	_, err := tfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfc *TaskFileCreate) ExecX(ctx context.Context) {
	if err := tfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tfc *TaskFileCreate) defaults() {
	if _, ok := tfc.mutation.Attached(); !ok {
		v := taskfile.DefaultAttached
		tfc.mutation.SetAttached(v)
	}
	if _, ok := tfc.mutation.CreatedAt(); !ok {
		v := taskfile.DefaultCreatedAt()
		tfc.mutation.SetCreatedAt(v)
	}
	if _, ok := tfc.mutation.UpdatedAt(); !ok {
		v := taskfile.DefaultUpdatedAt()
		tfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tfc.mutation.ID(); !ok {
		v := taskfile.DefaultID()
		tfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tfc *TaskFileCreate) check() error {
	if _, ok := tfc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "project_id"`)}
	}
	if _, ok := tfc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "task_id"`)}
	}
	if _, ok := tfc.mutation.TaskFeedID(); !ok {
		return &ValidationError{Name: "task_feed_id", err: errors.New(`ent: missing required field "task_feed_id"`)}
	}
	if _, ok := tfc.mutation.FileTypeID(); !ok {
		return &ValidationError{Name: "file_type_id", err: errors.New(`ent: missing required field "file_type_id"`)}
	}
	if _, ok := tfc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tfc.mutation.Name(); ok {
		if err := taskfile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tfc.mutation.Src(); !ok {
		return &ValidationError{Name: "src", err: errors.New(`ent: missing required field "src"`)}
	}
	if v, ok := tfc.mutation.Src(); ok {
		if err := taskfile.SrcValidator(v); err != nil {
			return &ValidationError{Name: "src", err: fmt.Errorf(`ent: validator failed for field "src": %w`, err)}
		}
	}
	if _, ok := tfc.mutation.Attached(); !ok {
		return &ValidationError{Name: "attached", err: errors.New(`ent: missing required field "attached"`)}
	}
	if _, ok := tfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tfc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New("ent: missing required edge \"project\"")}
	}
	if _, ok := tfc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required edge \"task\"")}
	}
	if _, ok := tfc.mutation.TaskFeedID(); !ok {
		return &ValidationError{Name: "taskFeed", err: errors.New("ent: missing required edge \"taskFeed\"")}
	}
	if _, ok := tfc.mutation.FileTypeID(); !ok {
		return &ValidationError{Name: "fileType", err: errors.New("ent: missing required edge \"fileType\"")}
	}
	return nil
}

func (tfc *TaskFileCreate) sqlSave(ctx context.Context) (*TaskFile, error) {
	_node, _spec := tfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tfc.driver, _spec); err != nil {
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

func (tfc *TaskFileCreate) createSpec() (*TaskFile, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskFile{config: tfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: taskfile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfile.FieldID,
			},
		}
	)
	_spec.OnConflict = tfc.conflict
	if id, ok := tfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tfc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tfc.mutation.Src(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldSrc,
		})
		_node.Src = value
	}
	if value, ok := tfc.mutation.Attached(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: taskfile.FieldAttached,
		})
		_node.Attached = value
	}
	if value, ok := tfc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskfile.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tfc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskfile.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tfc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfile.ProjectTable,
			Columns: []string{taskfile.ProjectColumn},
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
	if nodes := tfc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfile.TaskTable,
			Columns: []string{taskfile.TaskColumn},
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
	if nodes := tfc.mutation.TaskFeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfile.TaskFeedTable,
			Columns: []string{taskfile.TaskFeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskfeed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskFeedID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tfc.mutation.FileTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfile.FileTypeTable,
			Columns: []string{taskfile.FileTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: filetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FileTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskFile.Create().
//		SetProjectID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskFileUpsert) {
//			SetProjectID(v+v).
//		}).
//		Exec(ctx)
//
func (tfc *TaskFileCreate) OnConflict(opts ...sql.ConflictOption) *TaskFileUpsertOne {
	tfc.conflict = opts
	return &TaskFileUpsertOne{
		create: tfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tfc *TaskFileCreate) OnConflictColumns(columns ...string) *TaskFileUpsertOne {
	tfc.conflict = append(tfc.conflict, sql.ConflictColumns(columns...))
	return &TaskFileUpsertOne{
		create: tfc,
	}
}

type (
	// TaskFileUpsertOne is the builder for "upsert"-ing
	//  one TaskFile node.
	TaskFileUpsertOne struct {
		create *TaskFileCreate
	}

	// TaskFileUpsert is the "OnConflict" setter.
	TaskFileUpsert struct {
		*sql.UpdateSet
	}
)

// SetProjectID sets the "project_id" field.
func (u *TaskFileUpsert) SetProjectID(v ulid.ID) *TaskFileUpsert {
	u.Set(taskfile.FieldProjectID, v)
	return u
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateProjectID() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldProjectID)
	return u
}

// SetTaskID sets the "task_id" field.
func (u *TaskFileUpsert) SetTaskID(v ulid.ID) *TaskFileUpsert {
	u.Set(taskfile.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateTaskID() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldTaskID)
	return u
}

// SetTaskFeedID sets the "task_feed_id" field.
func (u *TaskFileUpsert) SetTaskFeedID(v ulid.ID) *TaskFileUpsert {
	u.Set(taskfile.FieldTaskFeedID, v)
	return u
}

// UpdateTaskFeedID sets the "task_feed_id" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateTaskFeedID() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldTaskFeedID)
	return u
}

// SetFileTypeID sets the "file_type_id" field.
func (u *TaskFileUpsert) SetFileTypeID(v ulid.ID) *TaskFileUpsert {
	u.Set(taskfile.FieldFileTypeID, v)
	return u
}

// UpdateFileTypeID sets the "file_type_id" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateFileTypeID() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldFileTypeID)
	return u
}

// SetName sets the "name" field.
func (u *TaskFileUpsert) SetName(v string) *TaskFileUpsert {
	u.Set(taskfile.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateName() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldName)
	return u
}

// SetSrc sets the "src" field.
func (u *TaskFileUpsert) SetSrc(v string) *TaskFileUpsert {
	u.Set(taskfile.FieldSrc, v)
	return u
}

// UpdateSrc sets the "src" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateSrc() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldSrc)
	return u
}

// SetAttached sets the "attached" field.
func (u *TaskFileUpsert) SetAttached(v bool) *TaskFileUpsert {
	u.Set(taskfile.FieldAttached, v)
	return u
}

// UpdateAttached sets the "attached" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateAttached() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldAttached)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskFileUpsert) SetCreatedAt(v time.Time) *TaskFileUpsert {
	u.Set(taskfile.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateCreatedAt() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskFileUpsert) SetUpdatedAt(v time.Time) *TaskFileUpsert {
	u.Set(taskfile.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskFileUpsert) UpdateUpdatedAt() *TaskFileUpsert {
	u.SetExcluded(taskfile.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TaskFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskfile.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskFileUpsertOne) UpdateNewValues() *TaskFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(taskfile.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TaskFile.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TaskFileUpsertOne) Ignore() *TaskFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskFileUpsertOne) DoNothing() *TaskFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskFileCreate.OnConflict
// documentation for more info.
func (u *TaskFileUpsertOne) Update(set func(*TaskFileUpsert)) *TaskFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetProjectID sets the "project_id" field.
func (u *TaskFileUpsertOne) SetProjectID(v ulid.ID) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateProjectID() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateProjectID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *TaskFileUpsertOne) SetTaskID(v ulid.ID) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateTaskID() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateTaskID()
	})
}

// SetTaskFeedID sets the "task_feed_id" field.
func (u *TaskFileUpsertOne) SetTaskFeedID(v ulid.ID) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetTaskFeedID(v)
	})
}

// UpdateTaskFeedID sets the "task_feed_id" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateTaskFeedID() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateTaskFeedID()
	})
}

// SetFileTypeID sets the "file_type_id" field.
func (u *TaskFileUpsertOne) SetFileTypeID(v ulid.ID) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetFileTypeID(v)
	})
}

// UpdateFileTypeID sets the "file_type_id" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateFileTypeID() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateFileTypeID()
	})
}

// SetName sets the "name" field.
func (u *TaskFileUpsertOne) SetName(v string) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateName() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateName()
	})
}

// SetSrc sets the "src" field.
func (u *TaskFileUpsertOne) SetSrc(v string) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetSrc(v)
	})
}

// UpdateSrc sets the "src" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateSrc() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateSrc()
	})
}

// SetAttached sets the "attached" field.
func (u *TaskFileUpsertOne) SetAttached(v bool) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetAttached(v)
	})
}

// UpdateAttached sets the "attached" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateAttached() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateAttached()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskFileUpsertOne) SetCreatedAt(v time.Time) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateCreatedAt() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskFileUpsertOne) SetUpdatedAt(v time.Time) *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskFileUpsertOne) UpdateUpdatedAt() *TaskFileUpsertOne {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskFileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskFileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskFileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskFileUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TaskFileUpsertOne.ID is not supported by MySQL driver. Use TaskFileUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskFileUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskFileCreateBulk is the builder for creating many TaskFile entities in bulk.
type TaskFileCreateBulk struct {
	config
	builders []*TaskFileCreate
	conflict []sql.ConflictOption
}

// Save creates the TaskFile entities in the database.
func (tfcb *TaskFileCreateBulk) Save(ctx context.Context) ([]*TaskFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tfcb.builders))
	nodes := make([]*TaskFile, len(tfcb.builders))
	mutators := make([]Mutator, len(tfcb.builders))
	for i := range tfcb.builders {
		func(i int, root context.Context) {
			builder := tfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskFileMutation)
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
					_, err = mutators[i+1].Mutate(root, tfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tfcb *TaskFileCreateBulk) SaveX(ctx context.Context) []*TaskFile {
	v, err := tfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfcb *TaskFileCreateBulk) Exec(ctx context.Context) error {
	_, err := tfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfcb *TaskFileCreateBulk) ExecX(ctx context.Context) {
	if err := tfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskFile.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskFileUpsert) {
//			SetProjectID(v+v).
//		}).
//		Exec(ctx)
//
func (tfcb *TaskFileCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskFileUpsertBulk {
	tfcb.conflict = opts
	return &TaskFileUpsertBulk{
		create: tfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tfcb *TaskFileCreateBulk) OnConflictColumns(columns ...string) *TaskFileUpsertBulk {
	tfcb.conflict = append(tfcb.conflict, sql.ConflictColumns(columns...))
	return &TaskFileUpsertBulk{
		create: tfcb,
	}
}

// TaskFileUpsertBulk is the builder for "upsert"-ing
// a bulk of TaskFile nodes.
type TaskFileUpsertBulk struct {
	create *TaskFileCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TaskFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskfile.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskFileUpsertBulk) UpdateNewValues() *TaskFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(taskfile.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TaskFile.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TaskFileUpsertBulk) Ignore() *TaskFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskFileUpsertBulk) DoNothing() *TaskFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskFileCreateBulk.OnConflict
// documentation for more info.
func (u *TaskFileUpsertBulk) Update(set func(*TaskFileUpsert)) *TaskFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetProjectID sets the "project_id" field.
func (u *TaskFileUpsertBulk) SetProjectID(v ulid.ID) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetProjectID(v)
	})
}

// UpdateProjectID sets the "project_id" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateProjectID() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateProjectID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *TaskFileUpsertBulk) SetTaskID(v ulid.ID) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateTaskID() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateTaskID()
	})
}

// SetTaskFeedID sets the "task_feed_id" field.
func (u *TaskFileUpsertBulk) SetTaskFeedID(v ulid.ID) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetTaskFeedID(v)
	})
}

// UpdateTaskFeedID sets the "task_feed_id" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateTaskFeedID() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateTaskFeedID()
	})
}

// SetFileTypeID sets the "file_type_id" field.
func (u *TaskFileUpsertBulk) SetFileTypeID(v ulid.ID) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetFileTypeID(v)
	})
}

// UpdateFileTypeID sets the "file_type_id" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateFileTypeID() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateFileTypeID()
	})
}

// SetName sets the "name" field.
func (u *TaskFileUpsertBulk) SetName(v string) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateName() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateName()
	})
}

// SetSrc sets the "src" field.
func (u *TaskFileUpsertBulk) SetSrc(v string) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetSrc(v)
	})
}

// UpdateSrc sets the "src" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateSrc() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateSrc()
	})
}

// SetAttached sets the "attached" field.
func (u *TaskFileUpsertBulk) SetAttached(v bool) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetAttached(v)
	})
}

// UpdateAttached sets the "attached" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateAttached() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateAttached()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskFileUpsertBulk) SetCreatedAt(v time.Time) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateCreatedAt() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskFileUpsertBulk) SetUpdatedAt(v time.Time) *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskFileUpsertBulk) UpdateUpdatedAt() *TaskFileUpsertBulk {
	return u.Update(func(s *TaskFileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskFileUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskFileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskFileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskFileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
