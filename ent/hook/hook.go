// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
)

// The ColorFunc type is an adapter to allow the use of ordinary
// function as Color mutator.
type ColorFunc func(context.Context, *ent.ColorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ColorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ColorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ColorMutation", m)
	}
	return f(ctx, mv)
}

// The FavoriteProjectFunc type is an adapter to allow the use of ordinary
// function as FavoriteProject mutator.
type FavoriteProjectFunc func(context.Context, *ent.FavoriteProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FavoriteProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FavoriteProjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FavoriteProjectMutation", m)
	}
	return f(ctx, mv)
}

// The FavoriteWorkspaceFunc type is an adapter to allow the use of ordinary
// function as FavoriteWorkspace mutator.
type FavoriteWorkspaceFunc func(context.Context, *ent.FavoriteWorkspaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FavoriteWorkspaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FavoriteWorkspaceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FavoriteWorkspaceMutation", m)
	}
	return f(ctx, mv)
}

// The IconFunc type is an adapter to allow the use of ordinary
// function as Icon mutator.
type IconFunc func(context.Context, *ent.IconMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IconFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IconMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IconMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *ent.ProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectBaseColorFunc type is an adapter to allow the use of ordinary
// function as ProjectBaseColor mutator.
type ProjectBaseColorFunc func(context.Context, *ent.ProjectBaseColorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectBaseColorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectBaseColorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectBaseColorMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectIconFunc type is an adapter to allow the use of ordinary
// function as ProjectIcon mutator.
type ProjectIconFunc func(context.Context, *ent.ProjectIconMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectIconFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectIconMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectIconMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectLightColorFunc type is an adapter to allow the use of ordinary
// function as ProjectLightColor mutator.
type ProjectLightColorFunc func(context.Context, *ent.ProjectLightColorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectLightColorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectLightColorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectLightColorMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectTaskColumnFunc type is an adapter to allow the use of ordinary
// function as ProjectTaskColumn mutator.
type ProjectTaskColumnFunc func(context.Context, *ent.ProjectTaskColumnMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectTaskColumnFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectTaskColumnMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectTaskColumnMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectTeammateFunc type is an adapter to allow the use of ordinary
// function as ProjectTeammate mutator.
type ProjectTeammateFunc func(context.Context, *ent.ProjectTeammateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectTeammateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectTeammateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectTeammateMutation", m)
	}
	return f(ctx, mv)
}

// The TaskColumnFunc type is an adapter to allow the use of ordinary
// function as TaskColumn mutator.
type TaskColumnFunc func(context.Context, *ent.TaskColumnMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskColumnFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskColumnMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskColumnMutation", m)
	}
	return f(ctx, mv)
}

// The TaskListCompletedStatusFunc type is an adapter to allow the use of ordinary
// function as TaskListCompletedStatus mutator.
type TaskListCompletedStatusFunc func(context.Context, *ent.TaskListCompletedStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskListCompletedStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskListCompletedStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskListCompletedStatusMutation", m)
	}
	return f(ctx, mv)
}

// The TaskListSortStatusFunc type is an adapter to allow the use of ordinary
// function as TaskListSortStatus mutator.
type TaskListSortStatusFunc func(context.Context, *ent.TaskListSortStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskListSortStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskListSortStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskListSortStatusMutation", m)
	}
	return f(ctx, mv)
}

// The TaskSectionFunc type is an adapter to allow the use of ordinary
// function as TaskSection mutator.
type TaskSectionFunc func(context.Context, *ent.TaskSectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskSectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskSectionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskSectionMutation", m)
	}
	return f(ctx, mv)
}

// The TeammateFunc type is an adapter to allow the use of ordinary
// function as Teammate mutator.
type TeammateFunc func(context.Context, *ent.TeammateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateMutation", m)
	}
	return f(ctx, mv)
}

// The TeammateTabStatusFunc type is an adapter to allow the use of ordinary
// function as TeammateTabStatus mutator.
type TeammateTabStatusFunc func(context.Context, *ent.TeammateTabStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTabStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTabStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTabStatusMutation", m)
	}
	return f(ctx, mv)
}

// The TeammateTaskColumnFunc type is an adapter to allow the use of ordinary
// function as TeammateTaskColumn mutator.
type TeammateTaskColumnFunc func(context.Context, *ent.TeammateTaskColumnMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTaskColumnFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTaskColumnMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTaskColumnMutation", m)
	}
	return f(ctx, mv)
}

// The TestTodoFunc type is an adapter to allow the use of ordinary
// function as TestTodo mutator.
type TestTodoFunc func(context.Context, *ent.TestTodoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestTodoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TestTodoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestTodoMutation", m)
	}
	return f(ctx, mv)
}

// The TestUserFunc type is an adapter to allow the use of ordinary
// function as TestUser mutator.
type TestUserFunc func(context.Context, *ent.TestUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TestUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestUserMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceFunc type is an adapter to allow the use of ordinary
// function as Workspace mutator.
type WorkspaceFunc func(context.Context, *ent.WorkspaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkspaceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkspaceMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceTeammateFunc type is an adapter to allow the use of ordinary
// function as WorkspaceTeammate mutator.
type WorkspaceTeammateFunc func(context.Context, *ent.WorkspaceTeammateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceTeammateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkspaceTeammateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkspaceTeammateMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
