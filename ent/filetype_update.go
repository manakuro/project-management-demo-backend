// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetName sets the "name" field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// SetTypeCode sets the "type_code" field.
func (ftu *FileTypeUpdate) SetTypeCode(fc filetype.TypeCode) *FileTypeUpdate {
	ftu.mutation.SetTypeCode(fc)
	return ftu
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftu *FileTypeUpdate) Mutation() *FileTypeMutation {
	return ftu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftu.hooks) == 0 {
		if err = ftu.check(); err != nil {
			return 0, err
		}
		affected, err = ftu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftu.check(); err != nil {
				return 0, err
			}
			ftu.mutation = mutation
			affected, err = ftu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftu.hooks) - 1; i >= 0; i-- {
			if ftu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FileTypeUpdate) check() error {
	if v, ok := ftu.mutation.Name(); ok {
		if err := filetype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ftu.mutation.TypeCode(); ok {
		if err := filetype.TypeCodeValidator(v); err != nil {
			return &ValidationError{Name: "type_code", err: fmt.Errorf("ent: validator failed for field \"type_code\": %w", err)}
		}
	}
	return nil
}

func (ftu *FileTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: filetype.FieldID,
			},
		},
	}
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filetype.FieldName,
		})
	}
	if value, ok := ftu.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: filetype.FieldTypeCode,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileTypeMutation
}

// SetName sets the "name" field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// SetTypeCode sets the "type_code" field.
func (ftuo *FileTypeUpdateOne) SetTypeCode(fc filetype.TypeCode) *FileTypeUpdateOne {
	ftuo.mutation.SetTypeCode(fc)
	return ftuo
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftuo *FileTypeUpdateOne) Mutation() *FileTypeMutation {
	return ftuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FileTypeUpdateOne) Select(field string, fields ...string) *FileTypeUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FileType entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	var (
		err  error
		node *FileType
	)
	if len(ftuo.hooks) == 0 {
		if err = ftuo.check(); err != nil {
			return nil, err
		}
		node, err = ftuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftuo.check(); err != nil {
				return nil, err
			}
			ftuo.mutation = mutation
			node, err = ftuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftuo.hooks) - 1; i >= 0; i-- {
			if ftuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FileTypeUpdateOne) check() error {
	if v, ok := ftuo.mutation.Name(); ok {
		if err := filetype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ftuo.mutation.TypeCode(); ok {
		if err := filetype.TypeCodeValidator(v); err != nil {
			return &ValidationError{Name: "type_code", err: fmt.Errorf("ent: validator failed for field \"type_code\": %w", err)}
		}
	}
	return nil
}

func (ftuo *FileTypeUpdateOne) sqlSave(ctx context.Context) (_node *FileType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: filetype.FieldID,
			},
		},
	}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FileType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for _, f := range fields {
			if !filetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filetype.FieldName,
		})
	}
	if value, ok := ftuo.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: filetype.FieldTypeCode,
		})
	}
	_node = &FileType{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
