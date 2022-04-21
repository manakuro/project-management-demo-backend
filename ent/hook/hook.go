// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
)

// The ActivityTypeFunc type is an adapter to allow the use of ordinary
// function as ActivityType mutator.
type ActivityTypeFunc func(context.Context, *ent.ActivityTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ActivityTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivityTypeMutation", m)
	}
	return f(ctx, mv)
}

// The ArchivedTaskActivityFunc type is an adapter to allow the use of ordinary
// function as ArchivedTaskActivity mutator.
type ArchivedTaskActivityFunc func(context.Context, *ent.ArchivedTaskActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArchivedTaskActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ArchivedTaskActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArchivedTaskActivityMutation", m)
	}
	return f(ctx, mv)
}

// The ArchivedTaskActivityTaskFunc type is an adapter to allow the use of ordinary
// function as ArchivedTaskActivityTask mutator.
type ArchivedTaskActivityTaskFunc func(context.Context, *ent.ArchivedTaskActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArchivedTaskActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ArchivedTaskActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArchivedTaskActivityTaskMutation", m)
	}
	return f(ctx, mv)
}

// The ArchivedWorkspaceActivityFunc type is an adapter to allow the use of ordinary
// function as ArchivedWorkspaceActivity mutator.
type ArchivedWorkspaceActivityFunc func(context.Context, *ent.ArchivedWorkspaceActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArchivedWorkspaceActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ArchivedWorkspaceActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArchivedWorkspaceActivityMutation", m)
	}
	return f(ctx, mv)
}

// The ArchivedWorkspaceActivityTaskFunc type is an adapter to allow the use of ordinary
// function as ArchivedWorkspaceActivityTask mutator.
type ArchivedWorkspaceActivityTaskFunc func(context.Context, *ent.ArchivedWorkspaceActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArchivedWorkspaceActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ArchivedWorkspaceActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArchivedWorkspaceActivityTaskMutation", m)
	}
	return f(ctx, mv)
}

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

// The DeletedProjectTaskFunc type is an adapter to allow the use of ordinary
// function as DeletedProjectTask mutator.
type DeletedProjectTaskFunc func(context.Context, *ent.DeletedProjectTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeletedProjectTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeletedProjectTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeletedProjectTaskMutation", m)
	}
	return f(ctx, mv)
}

// The DeletedTaskFunc type is an adapter to allow the use of ordinary
// function as DeletedTask mutator.
type DeletedTaskFunc func(context.Context, *ent.DeletedTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeletedTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeletedTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeletedTaskMutation", m)
	}
	return f(ctx, mv)
}

// The DeletedTaskActivityTaskFunc type is an adapter to allow the use of ordinary
// function as DeletedTaskActivityTask mutator.
type DeletedTaskActivityTaskFunc func(context.Context, *ent.DeletedTaskActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeletedTaskActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeletedTaskActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeletedTaskActivityTaskMutation", m)
	}
	return f(ctx, mv)
}

// The DeletedTeammateTaskFunc type is an adapter to allow the use of ordinary
// function as DeletedTeammateTask mutator.
type DeletedTeammateTaskFunc func(context.Context, *ent.DeletedTeammateTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeletedTeammateTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeletedTeammateTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeletedTeammateTaskMutation", m)
	}
	return f(ctx, mv)
}

// The DeletedWorkspaceActivityTaskFunc type is an adapter to allow the use of ordinary
// function as DeletedWorkspaceActivityTask mutator.
type DeletedWorkspaceActivityTaskFunc func(context.Context, *ent.DeletedWorkspaceActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeletedWorkspaceActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeletedWorkspaceActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeletedWorkspaceActivityTaskMutation", m)
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

// The FileTypeFunc type is an adapter to allow the use of ordinary
// function as FileType mutator.
type FileTypeFunc func(context.Context, *ent.FileTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileTypeMutation", m)
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

// The ProjectTaskFunc type is an adapter to allow the use of ordinary
// function as ProjectTask mutator.
type ProjectTaskFunc func(context.Context, *ent.ProjectTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectTaskMutation", m)
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

// The ProjectTaskListStatusFunc type is an adapter to allow the use of ordinary
// function as ProjectTaskListStatus mutator.
type ProjectTaskListStatusFunc func(context.Context, *ent.ProjectTaskListStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectTaskListStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectTaskListStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectTaskListStatusMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectTaskSectionFunc type is an adapter to allow the use of ordinary
// function as ProjectTaskSection mutator.
type ProjectTaskSectionFunc func(context.Context, *ent.ProjectTaskSectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectTaskSectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectTaskSectionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectTaskSectionMutation", m)
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

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *ent.TagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TagMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TagMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
	}
	return f(ctx, mv)
}

// The TaskActivityFunc type is an adapter to allow the use of ordinary
// function as TaskActivity mutator.
type TaskActivityFunc func(context.Context, *ent.TaskActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskActivityMutation", m)
	}
	return f(ctx, mv)
}

// The TaskActivityTaskFunc type is an adapter to allow the use of ordinary
// function as TaskActivityTask mutator.
type TaskActivityTaskFunc func(context.Context, *ent.TaskActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskActivityTaskMutation", m)
	}
	return f(ctx, mv)
}

// The TaskCollaboratorFunc type is an adapter to allow the use of ordinary
// function as TaskCollaborator mutator.
type TaskCollaboratorFunc func(context.Context, *ent.TaskCollaboratorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskCollaboratorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskCollaboratorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskCollaboratorMutation", m)
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

// The TaskFeedFunc type is an adapter to allow the use of ordinary
// function as TaskFeed mutator.
type TaskFeedFunc func(context.Context, *ent.TaskFeedMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFeedFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskFeedMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskFeedMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFeedLikeFunc type is an adapter to allow the use of ordinary
// function as TaskFeedLike mutator.
type TaskFeedLikeFunc func(context.Context, *ent.TaskFeedLikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFeedLikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskFeedLikeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskFeedLikeMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFileFunc type is an adapter to allow the use of ordinary
// function as TaskFile mutator.
type TaskFileFunc func(context.Context, *ent.TaskFileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskFileMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskFileMutation", m)
	}
	return f(ctx, mv)
}

// The TaskLikeFunc type is an adapter to allow the use of ordinary
// function as TaskLike mutator.
type TaskLikeFunc func(context.Context, *ent.TaskLikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskLikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskLikeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskLikeMutation", m)
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

// The TaskPriorityFunc type is an adapter to allow the use of ordinary
// function as TaskPriority mutator.
type TaskPriorityFunc func(context.Context, *ent.TaskPriorityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskPriorityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskPriorityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskPriorityMutation", m)
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

// The TaskTagFunc type is an adapter to allow the use of ordinary
// function as TaskTag mutator.
type TaskTagFunc func(context.Context, *ent.TaskTagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskTagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskTagMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskTagMutation", m)
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

// The TeammateTaskFunc type is an adapter to allow the use of ordinary
// function as TeammateTask mutator.
type TeammateTaskFunc func(context.Context, *ent.TeammateTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTaskMutation", m)
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

// The TeammateTaskListStatusFunc type is an adapter to allow the use of ordinary
// function as TeammateTaskListStatus mutator.
type TeammateTaskListStatusFunc func(context.Context, *ent.TeammateTaskListStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTaskListStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTaskListStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTaskListStatusMutation", m)
	}
	return f(ctx, mv)
}

// The TeammateTaskSectionFunc type is an adapter to allow the use of ordinary
// function as TeammateTaskSection mutator.
type TeammateTaskSectionFunc func(context.Context, *ent.TeammateTaskSectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTaskSectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTaskSectionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTaskSectionMutation", m)
	}
	return f(ctx, mv)
}

// The TeammateTaskTabStatusFunc type is an adapter to allow the use of ordinary
// function as TeammateTaskTabStatus mutator.
type TeammateTaskTabStatusFunc func(context.Context, *ent.TeammateTaskTabStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeammateTaskTabStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeammateTaskTabStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeammateTaskTabStatusMutation", m)
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

// The WorkspaceActivityFunc type is an adapter to allow the use of ordinary
// function as WorkspaceActivity mutator.
type WorkspaceActivityFunc func(context.Context, *ent.WorkspaceActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkspaceActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkspaceActivityMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceActivityTaskFunc type is an adapter to allow the use of ordinary
// function as WorkspaceActivityTask mutator.
type WorkspaceActivityTaskFunc func(context.Context, *ent.WorkspaceActivityTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceActivityTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkspaceActivityTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkspaceActivityTaskMutation", m)
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
