// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFileUpdate is the builder for updating TaskFile entities.
type TaskFileUpdate struct {
	config
	hooks    []Hook
	mutation *TaskFileMutation
}

// Where appends a list predicates to the TaskFileUpdate builder.
func (tfu *TaskFileUpdate) Where(ps ...predicate.TaskFile) *TaskFileUpdate {
	tfu.mutation.Where(ps...)
	return tfu
}

// SetProjectID sets the "project_id" field.
func (tfu *TaskFileUpdate) SetProjectID(u ulid.ID) *TaskFileUpdate {
	tfu.mutation.SetProjectID(u)
	return tfu
}

// SetTaskID sets the "task_id" field.
func (tfu *TaskFileUpdate) SetTaskID(u ulid.ID) *TaskFileUpdate {
	tfu.mutation.SetTaskID(u)
	return tfu
}

// SetTaskFeedID sets the "task_feed_id" field.
func (tfu *TaskFileUpdate) SetTaskFeedID(u ulid.ID) *TaskFileUpdate {
	tfu.mutation.SetTaskFeedID(u)
	return tfu
}

// SetFileTypeID sets the "file_type_id" field.
func (tfu *TaskFileUpdate) SetFileTypeID(u ulid.ID) *TaskFileUpdate {
	tfu.mutation.SetFileTypeID(u)
	return tfu
}

// SetName sets the "name" field.
func (tfu *TaskFileUpdate) SetName(s string) *TaskFileUpdate {
	tfu.mutation.SetName(s)
	return tfu
}

// SetSrc sets the "src" field.
func (tfu *TaskFileUpdate) SetSrc(s string) *TaskFileUpdate {
	tfu.mutation.SetSrc(s)
	return tfu
}

// SetAttached sets the "attached" field.
func (tfu *TaskFileUpdate) SetAttached(b bool) *TaskFileUpdate {
	tfu.mutation.SetAttached(b)
	return tfu
}

// SetNillableAttached sets the "attached" field if the given value is not nil.
func (tfu *TaskFileUpdate) SetNillableAttached(b *bool) *TaskFileUpdate {
	if b != nil {
		tfu.SetAttached(*b)
	}
	return tfu
}

// SetProject sets the "project" edge to the Project entity.
func (tfu *TaskFileUpdate) SetProject(p *Project) *TaskFileUpdate {
	return tfu.SetProjectID(p.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (tfu *TaskFileUpdate) SetTask(t *Task) *TaskFileUpdate {
	return tfu.SetTaskID(t.ID)
}

// SetTaskFeed sets the "taskFeed" edge to the TaskFeed entity.
func (tfu *TaskFileUpdate) SetTaskFeed(t *TaskFeed) *TaskFileUpdate {
	return tfu.SetTaskFeedID(t.ID)
}

// SetFileType sets the "fileType" edge to the FileType entity.
func (tfu *TaskFileUpdate) SetFileType(f *FileType) *TaskFileUpdate {
	return tfu.SetFileTypeID(f.ID)
}

// Mutation returns the TaskFileMutation object of the builder.
func (tfu *TaskFileUpdate) Mutation() *TaskFileMutation {
	return tfu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tfu *TaskFileUpdate) ClearProject() *TaskFileUpdate {
	tfu.mutation.ClearProject()
	return tfu
}

// ClearTask clears the "task" edge to the Task entity.
func (tfu *TaskFileUpdate) ClearTask() *TaskFileUpdate {
	tfu.mutation.ClearTask()
	return tfu
}

// ClearTaskFeed clears the "taskFeed" edge to the TaskFeed entity.
func (tfu *TaskFileUpdate) ClearTaskFeed() *TaskFileUpdate {
	tfu.mutation.ClearTaskFeed()
	return tfu
}

// ClearFileType clears the "fileType" edge to the FileType entity.
func (tfu *TaskFileUpdate) ClearFileType() *TaskFileUpdate {
	tfu.mutation.ClearFileType()
	return tfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tfu *TaskFileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tfu.hooks) == 0 {
		if err = tfu.check(); err != nil {
			return 0, err
		}
		affected, err = tfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tfu.check(); err != nil {
				return 0, err
			}
			tfu.mutation = mutation
			affected, err = tfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tfu.hooks) - 1; i >= 0; i-- {
			if tfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tfu *TaskFileUpdate) SaveX(ctx context.Context) int {
	affected, err := tfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tfu *TaskFileUpdate) Exec(ctx context.Context) error {
	_, err := tfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfu *TaskFileUpdate) ExecX(ctx context.Context) {
	if err := tfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tfu *TaskFileUpdate) check() error {
	if v, ok := tfu.mutation.Name(); ok {
		if err := taskfile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TaskFile.name": %w`, err)}
		}
	}
	if v, ok := tfu.mutation.Src(); ok {
		if err := taskfile.SrcValidator(v); err != nil {
			return &ValidationError{Name: "src", err: fmt.Errorf(`ent: validator failed for field "TaskFile.src": %w`, err)}
		}
	}
	if _, ok := tfu.mutation.ProjectID(); tfu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.project"`)
	}
	if _, ok := tfu.mutation.TaskID(); tfu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.task"`)
	}
	if _, ok := tfu.mutation.TaskFeedID(); tfu.mutation.TaskFeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.taskFeed"`)
	}
	if _, ok := tfu.mutation.FileTypeID(); tfu.mutation.FileTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.fileType"`)
	}
	return nil
}

func (tfu *TaskFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskfile.Table,
			Columns: taskfile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfile.FieldID,
			},
		},
	}
	if ps := tfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tfu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldName,
		})
	}
	if value, ok := tfu.mutation.Src(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldSrc,
		})
	}
	if value, ok := tfu.mutation.Attached(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: taskfile.FieldAttached,
		})
	}
	if tfu.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfu.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfu.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfu.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfu.mutation.TaskFeedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfu.mutation.TaskFeedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfu.mutation.FileTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfu.mutation.FileTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskFileUpdateOne is the builder for updating a single TaskFile entity.
type TaskFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskFileMutation
}

// SetProjectID sets the "project_id" field.
func (tfuo *TaskFileUpdateOne) SetProjectID(u ulid.ID) *TaskFileUpdateOne {
	tfuo.mutation.SetProjectID(u)
	return tfuo
}

// SetTaskID sets the "task_id" field.
func (tfuo *TaskFileUpdateOne) SetTaskID(u ulid.ID) *TaskFileUpdateOne {
	tfuo.mutation.SetTaskID(u)
	return tfuo
}

// SetTaskFeedID sets the "task_feed_id" field.
func (tfuo *TaskFileUpdateOne) SetTaskFeedID(u ulid.ID) *TaskFileUpdateOne {
	tfuo.mutation.SetTaskFeedID(u)
	return tfuo
}

// SetFileTypeID sets the "file_type_id" field.
func (tfuo *TaskFileUpdateOne) SetFileTypeID(u ulid.ID) *TaskFileUpdateOne {
	tfuo.mutation.SetFileTypeID(u)
	return tfuo
}

// SetName sets the "name" field.
func (tfuo *TaskFileUpdateOne) SetName(s string) *TaskFileUpdateOne {
	tfuo.mutation.SetName(s)
	return tfuo
}

// SetSrc sets the "src" field.
func (tfuo *TaskFileUpdateOne) SetSrc(s string) *TaskFileUpdateOne {
	tfuo.mutation.SetSrc(s)
	return tfuo
}

// SetAttached sets the "attached" field.
func (tfuo *TaskFileUpdateOne) SetAttached(b bool) *TaskFileUpdateOne {
	tfuo.mutation.SetAttached(b)
	return tfuo
}

// SetNillableAttached sets the "attached" field if the given value is not nil.
func (tfuo *TaskFileUpdateOne) SetNillableAttached(b *bool) *TaskFileUpdateOne {
	if b != nil {
		tfuo.SetAttached(*b)
	}
	return tfuo
}

// SetProject sets the "project" edge to the Project entity.
func (tfuo *TaskFileUpdateOne) SetProject(p *Project) *TaskFileUpdateOne {
	return tfuo.SetProjectID(p.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (tfuo *TaskFileUpdateOne) SetTask(t *Task) *TaskFileUpdateOne {
	return tfuo.SetTaskID(t.ID)
}

// SetTaskFeed sets the "taskFeed" edge to the TaskFeed entity.
func (tfuo *TaskFileUpdateOne) SetTaskFeed(t *TaskFeed) *TaskFileUpdateOne {
	return tfuo.SetTaskFeedID(t.ID)
}

// SetFileType sets the "fileType" edge to the FileType entity.
func (tfuo *TaskFileUpdateOne) SetFileType(f *FileType) *TaskFileUpdateOne {
	return tfuo.SetFileTypeID(f.ID)
}

// Mutation returns the TaskFileMutation object of the builder.
func (tfuo *TaskFileUpdateOne) Mutation() *TaskFileMutation {
	return tfuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tfuo *TaskFileUpdateOne) ClearProject() *TaskFileUpdateOne {
	tfuo.mutation.ClearProject()
	return tfuo
}

// ClearTask clears the "task" edge to the Task entity.
func (tfuo *TaskFileUpdateOne) ClearTask() *TaskFileUpdateOne {
	tfuo.mutation.ClearTask()
	return tfuo
}

// ClearTaskFeed clears the "taskFeed" edge to the TaskFeed entity.
func (tfuo *TaskFileUpdateOne) ClearTaskFeed() *TaskFileUpdateOne {
	tfuo.mutation.ClearTaskFeed()
	return tfuo
}

// ClearFileType clears the "fileType" edge to the FileType entity.
func (tfuo *TaskFileUpdateOne) ClearFileType() *TaskFileUpdateOne {
	tfuo.mutation.ClearFileType()
	return tfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tfuo *TaskFileUpdateOne) Select(field string, fields ...string) *TaskFileUpdateOne {
	tfuo.fields = append([]string{field}, fields...)
	return tfuo
}

// Save executes the query and returns the updated TaskFile entity.
func (tfuo *TaskFileUpdateOne) Save(ctx context.Context) (*TaskFile, error) {
	var (
		err  error
		node *TaskFile
	)
	if len(tfuo.hooks) == 0 {
		if err = tfuo.check(); err != nil {
			return nil, err
		}
		node, err = tfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tfuo.check(); err != nil {
				return nil, err
			}
			tfuo.mutation = mutation
			node, err = tfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tfuo.hooks) - 1; i >= 0; i-- {
			if tfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tfuo *TaskFileUpdateOne) SaveX(ctx context.Context) *TaskFile {
	node, err := tfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tfuo *TaskFileUpdateOne) Exec(ctx context.Context) error {
	_, err := tfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfuo *TaskFileUpdateOne) ExecX(ctx context.Context) {
	if err := tfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tfuo *TaskFileUpdateOne) check() error {
	if v, ok := tfuo.mutation.Name(); ok {
		if err := taskfile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TaskFile.name": %w`, err)}
		}
	}
	if v, ok := tfuo.mutation.Src(); ok {
		if err := taskfile.SrcValidator(v); err != nil {
			return &ValidationError{Name: "src", err: fmt.Errorf(`ent: validator failed for field "TaskFile.src": %w`, err)}
		}
	}
	if _, ok := tfuo.mutation.ProjectID(); tfuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.project"`)
	}
	if _, ok := tfuo.mutation.TaskID(); tfuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.task"`)
	}
	if _, ok := tfuo.mutation.TaskFeedID(); tfuo.mutation.TaskFeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.taskFeed"`)
	}
	if _, ok := tfuo.mutation.FileTypeID(); tfuo.mutation.FileTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskFile.fileType"`)
	}
	return nil
}

func (tfuo *TaskFileUpdateOne) sqlSave(ctx context.Context) (_node *TaskFile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskfile.Table,
			Columns: taskfile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfile.FieldID,
			},
		},
	}
	id, ok := tfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskfile.FieldID)
		for _, f := range fields {
			if !taskfile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tfuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldName,
		})
	}
	if value, ok := tfuo.mutation.Src(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskfile.FieldSrc,
		})
	}
	if value, ok := tfuo.mutation.Attached(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: taskfile.FieldAttached,
		})
	}
	if tfuo.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfuo.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfuo.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfuo.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfuo.mutation.TaskFeedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfuo.mutation.TaskFeedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tfuo.mutation.FileTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tfuo.mutation.FileTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TaskFile{config: tfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
