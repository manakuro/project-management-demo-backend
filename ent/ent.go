// Code generated by entc, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projecticon"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/tag"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskcollaborator"
	"project-management-demo-backend/ent/taskcolumn"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfeedlike"
	"project-management-demo-backend/ent/taskfile"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/tasklistsortstatus"
	"project-management-demo-backend/ent/taskpriority"
	"project-management-demo-backend/ent/tasksection"
	"project-management-demo-backend/ent/tasktag"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetaskcolumn"
	"project-management-demo-backend/ent/teammatetaskliststatus"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/teammatetasktabstatus"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ent aliases to avoid import conflicts in user's code.
type (
	Op         = ent.Op
	Hook       = ent.Hook
	Value      = ent.Value
	Query      = ent.Query
	Policy     = ent.Policy
	Mutator    = ent.Mutator
	Mutation   = ent.Mutation
	MutateFunc = ent.MutateFunc
)

// OrderFunc applies an ordering on the sql selector.
type OrderFunc func(*sql.Selector)

// columnChecker returns a function indicates if the column exists in the given column.
func columnChecker(table string) func(string) error {
	checks := map[string]func(string) bool{
		color.Table:                   color.ValidColumn,
		deletedtask.Table:             deletedtask.ValidColumn,
		favoriteproject.Table:         favoriteproject.ValidColumn,
		favoriteworkspace.Table:       favoriteworkspace.ValidColumn,
		filetype.Table:                filetype.ValidColumn,
		icon.Table:                    icon.ValidColumn,
		project.Table:                 project.ValidColumn,
		projectbasecolor.Table:        projectbasecolor.ValidColumn,
		projecticon.Table:             projecticon.ValidColumn,
		projectlightcolor.Table:       projectlightcolor.ValidColumn,
		projecttask.Table:             projecttask.ValidColumn,
		projecttaskcolumn.Table:       projecttaskcolumn.ValidColumn,
		projecttaskliststatus.Table:   projecttaskliststatus.ValidColumn,
		projecttasksection.Table:      projecttasksection.ValidColumn,
		projectteammate.Table:         projectteammate.ValidColumn,
		tag.Table:                     tag.ValidColumn,
		task.Table:                    task.ValidColumn,
		taskcollaborator.Table:        taskcollaborator.ValidColumn,
		taskcolumn.Table:              taskcolumn.ValidColumn,
		taskfeed.Table:                taskfeed.ValidColumn,
		taskfeedlike.Table:            taskfeedlike.ValidColumn,
		taskfile.Table:                taskfile.ValidColumn,
		tasklike.Table:                tasklike.ValidColumn,
		tasklistcompletedstatus.Table: tasklistcompletedstatus.ValidColumn,
		tasklistsortstatus.Table:      tasklistsortstatus.ValidColumn,
		taskpriority.Table:            taskpriority.ValidColumn,
		tasksection.Table:             tasksection.ValidColumn,
		tasktag.Table:                 tasktag.ValidColumn,
		teammate.Table:                teammate.ValidColumn,
		teammatetask.Table:            teammatetask.ValidColumn,
		teammatetaskcolumn.Table:      teammatetaskcolumn.ValidColumn,
		teammatetaskliststatus.Table:  teammatetaskliststatus.ValidColumn,
		teammatetasksection.Table:     teammatetasksection.ValidColumn,
		teammatetasktabstatus.Table:   teammatetasktabstatus.ValidColumn,
		testtodo.Table:                testtodo.ValidColumn,
		testuser.Table:                testuser.ValidColumn,
		workspace.Table:               workspace.ValidColumn,
		workspaceteammate.Table:       workspaceteammate.ValidColumn,
	}
	check, ok := checks[table]
	if !ok {
		return func(string) error {
			return fmt.Errorf("unknown table %q", table)
		}
	}
	return func(column string) error {
		if !check(column) {
			return fmt.Errorf("unknown column %q for table %q", column, table)
		}
		return nil
	}
}

// Asc applies the given fields in ASC order.
func Asc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Asc(s.C(f)))
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Desc(s.C(f)))
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
type AggregateFunc func(*sql.Selector) string

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.As(fn(s), end)
	}
}

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.Count("*")
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Max(s.C(field))
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Avg(s.C(field))
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Min(s.C(field))
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Sum(s.C(field))
	}
}

// ValidationError returns when validating a field or edge fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "ent: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "ent: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "ent: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "ent: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}
