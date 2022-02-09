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
	"project-management-demo-backend/ent/task"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskCreate is the builder for creating a ProjectTask entity.
type ProjectTaskCreate struct {
	config
	mutation *ProjectTaskMutation
	hooks    []Hook
}

// SetProjectID sets the "project_id" field.
func (ptc *ProjectTaskCreate) SetProjectID(u ulid.ID) *ProjectTaskCreate {
	ptc.mutation.SetProjectID(u)
	return ptc
}

// SetTaskID sets the "task_id" field.
func (ptc *ProjectTaskCreate) SetTaskID(u ulid.ID) *ProjectTaskCreate {
	ptc.mutation.SetTaskID(u)
	return ptc
}

// SetProjectTaskSectionID sets the "project_task_section_id" field.
func (ptc *ProjectTaskCreate) SetProjectTaskSectionID(u ulid.ID) *ProjectTaskCreate {
	ptc.mutation.SetProjectTaskSectionID(u)
	return ptc
}

// SetCreatedAt sets the "created_at" field.
func (ptc *ProjectTaskCreate) SetCreatedAt(t time.Time) *ProjectTaskCreate {
	ptc.mutation.SetCreatedAt(t)
	return ptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableCreatedAt(t *time.Time) *ProjectTaskCreate {
	if t != nil {
		ptc.SetCreatedAt(*t)
	}
	return ptc
}

// SetUpdatedAt sets the "updated_at" field.
func (ptc *ProjectTaskCreate) SetUpdatedAt(t time.Time) *ProjectTaskCreate {
	ptc.mutation.SetUpdatedAt(t)
	return ptc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableUpdatedAt(t *time.Time) *ProjectTaskCreate {
	if t != nil {
		ptc.SetUpdatedAt(*t)
	}
	return ptc
}

// SetID sets the "id" field.
func (ptc *ProjectTaskCreate) SetID(u ulid.ID) *ProjectTaskCreate {
	ptc.mutation.SetID(u)
	return ptc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ptc *ProjectTaskCreate) SetNillableID(u *ulid.ID) *ProjectTaskCreate {
	if u != nil {
		ptc.SetID(*u)
	}
	return ptc
}

// SetProject sets the "project" edge to the Project entity.
func (ptc *ProjectTaskCreate) SetProject(p *Project) *ProjectTaskCreate {
	return ptc.SetProjectID(p.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ptc *ProjectTaskCreate) SetTask(t *Task) *ProjectTaskCreate {
	return ptc.SetTaskID(t.ID)
}

// SetProjectTaskSection sets the "projectTaskSection" edge to the ProjectTaskSection entity.
func (ptc *ProjectTaskCreate) SetProjectTaskSection(p *ProjectTaskSection) *ProjectTaskCreate {
	return ptc.SetProjectTaskSectionID(p.ID)
}

// Mutation returns the ProjectTaskMutation object of the builder.
func (ptc *ProjectTaskCreate) Mutation() *ProjectTaskMutation {
	return ptc.mutation
}

// Save creates the ProjectTask in the database.
func (ptc *ProjectTaskCreate) Save(ctx context.Context) (*ProjectTask, error) {
	var (
		err  error
		node *ProjectTask
	)
	ptc.defaults()
	if len(ptc.hooks) == 0 {
		if err = ptc.check(); err != nil {
			return nil, err
		}
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptc.check(); err != nil {
				return nil, err
			}
			ptc.mutation = mutation
			if node, err = ptc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			if ptc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *ProjectTaskCreate) SaveX(ctx context.Context) *ProjectTask {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *ProjectTaskCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *ProjectTaskCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptc *ProjectTaskCreate) defaults() {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		v := projecttask.DefaultCreatedAt()
		ptc.mutation.SetCreatedAt(v)
	}
	if _, ok := ptc.mutation.UpdatedAt(); !ok {
		v := projecttask.DefaultUpdatedAt()
		ptc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ptc.mutation.ID(); !ok {
		v := projecttask.DefaultID()
		ptc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *ProjectTaskCreate) check() error {
	if _, ok := ptc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "project_id"`)}
	}
	if _, ok := ptc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "task_id"`)}
	}
	if _, ok := ptc.mutation.ProjectTaskSectionID(); !ok {
		return &ValidationError{Name: "project_task_section_id", err: errors.New(`ent: missing required field "project_task_section_id"`)}
	}
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ptc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ptc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New("ent: missing required edge \"project\"")}
	}
	if _, ok := ptc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required edge \"task\"")}
	}
	if _, ok := ptc.mutation.ProjectTaskSectionID(); !ok {
		return &ValidationError{Name: "projectTaskSection", err: errors.New("ent: missing required edge \"projectTaskSection\"")}
	}
	return nil
}

func (ptc *ProjectTaskCreate) sqlSave(ctx context.Context) (*ProjectTask, error) {
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
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

func (ptc *ProjectTaskCreate) createSpec() (*ProjectTask, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectTask{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projecttask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttask.FieldID,
			},
		}
	)
	if id, ok := ptc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ptc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttask.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ptc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttask.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ptc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttask.ProjectTable,
			Columns: []string{projecttask.ProjectColumn},
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
	if nodes := ptc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttask.TaskTable,
			Columns: []string{projecttask.TaskColumn},
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
	if nodes := ptc.mutation.ProjectTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttask.ProjectTaskSectionTable,
			Columns: []string{projecttask.ProjectTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttasksection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectTaskSectionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectTaskCreateBulk is the builder for creating many ProjectTask entities in bulk.
type ProjectTaskCreateBulk struct {
	config
	builders []*ProjectTaskCreate
}

// Save creates the ProjectTask entities in the database.
func (ptcb *ProjectTaskCreateBulk) Save(ctx context.Context) ([]*ProjectTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*ProjectTask, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *ProjectTaskCreateBulk) SaveX(ctx context.Context) []*ProjectTask {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *ProjectTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *ProjectTaskCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
