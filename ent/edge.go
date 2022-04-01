// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (c *Color) ProjectBaseColors(ctx context.Context) ([]*ProjectBaseColor, error) {
	result, err := c.Edges.ProjectBaseColorsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryProjectBaseColors().All(ctx)
	}
	return result, err
}

func (c *Color) ProjectLightColors(ctx context.Context) ([]*ProjectLightColor, error) {
	result, err := c.Edges.ProjectLightColorsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryProjectLightColors().All(ctx)
	}
	return result, err
}

func (c *Color) TaskPriorities(ctx context.Context) ([]*TaskPriority, error) {
	result, err := c.Edges.TaskPrioritiesOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryTaskPriorities().All(ctx)
	}
	return result, err
}

func (c *Color) Tags(ctx context.Context) ([]*Tag, error) {
	result, err := c.Edges.TagsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryTags().All(ctx)
	}
	return result, err
}

func (dt *DeletedTask) Task(ctx context.Context) (*Task, error) {
	result, err := dt.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = dt.QueryTask().Only(ctx)
	}
	return result, err
}

func (dt *DeletedTask) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := dt.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = dt.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (fp *FavoriteProject) Project(ctx context.Context) (*Project, error) {
	result, err := fp.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = fp.QueryProject().Only(ctx)
	}
	return result, err
}

func (fp *FavoriteProject) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := fp.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = fp.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (fw *FavoriteWorkspace) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := fw.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = fw.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (fw *FavoriteWorkspace) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := fw.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = fw.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (ft *FileType) TaskFiles(ctx context.Context) ([]*TaskFile, error) {
	result, err := ft.Edges.TaskFilesOrErr()
	if IsNotLoaded(err) {
		result, err = ft.QueryTaskFiles().All(ctx)
	}
	return result, err
}

func (i *Icon) ProjectIcons(ctx context.Context) ([]*ProjectIcon, error) {
	result, err := i.Edges.ProjectIconsOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryProjectIcons().All(ctx)
	}
	return result, err
}

func (pr *Project) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := pr.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (pr *Project) ProjectBaseColor(ctx context.Context) (*ProjectBaseColor, error) {
	result, err := pr.Edges.ProjectBaseColorOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectBaseColor().Only(ctx)
	}
	return result, err
}

func (pr *Project) ProjectLightColor(ctx context.Context) (*ProjectLightColor, error) {
	result, err := pr.Edges.ProjectLightColorOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectLightColor().Only(ctx)
	}
	return result, err
}

func (pr *Project) ProjectIcon(ctx context.Context) (*ProjectIcon, error) {
	result, err := pr.Edges.ProjectIconOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectIcon().Only(ctx)
	}
	return result, err
}

func (pr *Project) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := pr.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (pr *Project) ProjectTeammates(ctx context.Context) ([]*ProjectTeammate, error) {
	result, err := pr.Edges.ProjectTeammatesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectTeammates().All(ctx)
	}
	return result, err
}

func (pr *Project) FavoriteProjects(ctx context.Context) ([]*FavoriteProject, error) {
	result, err := pr.Edges.FavoriteProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryFavoriteProjects().All(ctx)
	}
	return result, err
}

func (pr *Project) ProjectTaskColumns(ctx context.Context) ([]*ProjectTaskColumn, error) {
	result, err := pr.Edges.ProjectTaskColumnsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectTaskColumns().All(ctx)
	}
	return result, err
}

func (pr *Project) ProjectTaskListStatuses(ctx context.Context) ([]*ProjectTaskListStatus, error) {
	result, err := pr.Edges.ProjectTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectTaskListStatuses().All(ctx)
	}
	return result, err
}

func (pr *Project) ProjectTaskSections(ctx context.Context) ([]*ProjectTaskSection, error) {
	result, err := pr.Edges.ProjectTaskSectionsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectTaskSections().All(ctx)
	}
	return result, err
}

func (pr *Project) ProjectTasks(ctx context.Context) ([]*ProjectTask, error) {
	result, err := pr.Edges.ProjectTasksOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryProjectTasks().All(ctx)
	}
	return result, err
}

func (pr *Project) TaskFiles(ctx context.Context) ([]*TaskFile, error) {
	result, err := pr.Edges.TaskFilesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryTaskFiles().All(ctx)
	}
	return result, err
}

func (pbc *ProjectBaseColor) Projects(ctx context.Context) ([]*Project, error) {
	result, err := pbc.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = pbc.QueryProjects().All(ctx)
	}
	return result, err
}

func (pbc *ProjectBaseColor) Color(ctx context.Context) (*Color, error) {
	result, err := pbc.Edges.ColorOrErr()
	if IsNotLoaded(err) {
		result, err = pbc.QueryColor().Only(ctx)
	}
	return result, err
}

func (pi *ProjectIcon) Projects(ctx context.Context) ([]*Project, error) {
	result, err := pi.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = pi.QueryProjects().All(ctx)
	}
	return result, err
}

func (pi *ProjectIcon) Icon(ctx context.Context) (*Icon, error) {
	result, err := pi.Edges.IconOrErr()
	if IsNotLoaded(err) {
		result, err = pi.QueryIcon().Only(ctx)
	}
	return result, err
}

func (plc *ProjectLightColor) Projects(ctx context.Context) ([]*Project, error) {
	result, err := plc.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = plc.QueryProjects().All(ctx)
	}
	return result, err
}

func (plc *ProjectLightColor) Color(ctx context.Context) (*Color, error) {
	result, err := plc.Edges.ColorOrErr()
	if IsNotLoaded(err) {
		result, err = plc.QueryColor().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTask) Project(ctx context.Context) (*Project, error) {
	result, err := pt.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryProject().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTask) Task(ctx context.Context) (*Task, error) {
	result, err := pt.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryTask().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTask) ProjectTaskSection(ctx context.Context) (*ProjectTaskSection, error) {
	result, err := pt.Edges.ProjectTaskSectionOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryProjectTaskSection().Only(ctx)
	}
	return result, err
}

func (ptc *ProjectTaskColumn) Project(ctx context.Context) (*Project, error) {
	result, err := ptc.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = ptc.QueryProject().Only(ctx)
	}
	return result, err
}

func (ptc *ProjectTaskColumn) TaskColumn(ctx context.Context) (*TaskColumn, error) {
	result, err := ptc.Edges.TaskColumnOrErr()
	if IsNotLoaded(err) {
		result, err = ptc.QueryTaskColumn().Only(ctx)
	}
	return result, err
}

func (ptls *ProjectTaskListStatus) Project(ctx context.Context) (*Project, error) {
	result, err := ptls.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = ptls.QueryProject().Only(ctx)
	}
	return result, err
}

func (ptls *ProjectTaskListStatus) TaskListCompletedStatus(ctx context.Context) (*TaskListCompletedStatus, error) {
	result, err := ptls.Edges.TaskListCompletedStatusOrErr()
	if IsNotLoaded(err) {
		result, err = ptls.QueryTaskListCompletedStatus().Only(ctx)
	}
	return result, err
}

func (ptls *ProjectTaskListStatus) TaskListSortStatus(ctx context.Context) (*TaskListSortStatus, error) {
	result, err := ptls.Edges.TaskListSortStatusOrErr()
	if IsNotLoaded(err) {
		result, err = ptls.QueryTaskListSortStatus().Only(ctx)
	}
	return result, err
}

func (pts *ProjectTaskSection) Project(ctx context.Context) (*Project, error) {
	result, err := pts.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = pts.QueryProject().Only(ctx)
	}
	return result, err
}

func (pts *ProjectTaskSection) ProjectTasks(ctx context.Context) ([]*ProjectTask, error) {
	result, err := pts.Edges.ProjectTasksOrErr()
	if IsNotLoaded(err) {
		result, err = pts.QueryProjectTasks().All(ctx)
	}
	return result, err
}

func (pt *ProjectTeammate) Project(ctx context.Context) (*Project, error) {
	result, err := pt.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryProject().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTeammate) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := pt.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (t *Tag) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := t.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (t *Tag) Color(ctx context.Context) (*Color, error) {
	result, err := t.Edges.ColorOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryColor().Only(ctx)
	}
	return result, err
}

func (t *Tag) TaskTags(ctx context.Context) ([]*TaskTag, error) {
	result, err := t.Edges.TaskTagsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskTags().All(ctx)
	}
	return result, err
}

func (t *Task) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := t.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammate().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Task) TaskPriority(ctx context.Context) (*TaskPriority, error) {
	result, err := t.Edges.TaskPriorityOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskPriority().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Task) ParentTask(ctx context.Context) (*Task, error) {
	result, err := t.Edges.ParentTaskOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParentTask().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Task) SubTasks(ctx context.Context) ([]*Task, error) {
	result, err := t.Edges.SubTasksOrErr()
	if IsNotLoaded(err) {
		result, err = t.QuerySubTasks().All(ctx)
	}
	return result, err
}

func (t *Task) TeammateTasks(ctx context.Context) ([]*TeammateTask, error) {
	result, err := t.Edges.TeammateTasksOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTasks().All(ctx)
	}
	return result, err
}

func (t *Task) ProjectTasks(ctx context.Context) ([]*ProjectTask, error) {
	result, err := t.Edges.ProjectTasksOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProjectTasks().All(ctx)
	}
	return result, err
}

func (t *Task) TaskLikes(ctx context.Context) ([]*TaskLike, error) {
	result, err := t.Edges.TaskLikesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskLikes().All(ctx)
	}
	return result, err
}

func (t *Task) TaskTags(ctx context.Context) ([]*TaskTag, error) {
	result, err := t.Edges.TaskTagsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskTags().All(ctx)
	}
	return result, err
}

func (t *Task) TaskCollaborators(ctx context.Context) ([]*TaskCollaborator, error) {
	result, err := t.Edges.TaskCollaboratorsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskCollaborators().All(ctx)
	}
	return result, err
}

func (t *Task) TaskFeeds(ctx context.Context) ([]*TaskFeed, error) {
	result, err := t.Edges.TaskFeedsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskFeeds().All(ctx)
	}
	return result, err
}

func (t *Task) TaskFeedLikes(ctx context.Context) ([]*TaskFeedLike, error) {
	result, err := t.Edges.TaskFeedLikesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskFeedLikes().All(ctx)
	}
	return result, err
}

func (t *Task) TaskFiles(ctx context.Context) ([]*TaskFile, error) {
	result, err := t.Edges.TaskFilesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskFiles().All(ctx)
	}
	return result, err
}

func (t *Task) DeletedTasksRef(ctx context.Context) ([]*DeletedTask, error) {
	result, err := t.Edges.DeletedTasksRefOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryDeletedTasksRef().All(ctx)
	}
	return result, err
}

func (tc *TaskCollaborator) Task(ctx context.Context) (*Task, error) {
	result, err := tc.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tc.QueryTask().Only(ctx)
	}
	return result, err
}

func (tc *TaskCollaborator) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tc.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tc.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tc *TaskColumn) TeammateTaskColumns(ctx context.Context) ([]*TeammateTaskColumn, error) {
	result, err := tc.Edges.TeammateTaskColumnsOrErr()
	if IsNotLoaded(err) {
		result, err = tc.QueryTeammateTaskColumns().All(ctx)
	}
	return result, err
}

func (tc *TaskColumn) ProjectTaskColumns(ctx context.Context) ([]*ProjectTaskColumn, error) {
	result, err := tc.Edges.ProjectTaskColumnsOrErr()
	if IsNotLoaded(err) {
		result, err = tc.QueryProjectTaskColumns().All(ctx)
	}
	return result, err
}

func (tf *TaskFeed) Task(ctx context.Context) (*Task, error) {
	result, err := tf.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTask().Only(ctx)
	}
	return result, err
}

func (tf *TaskFeed) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tf.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tf *TaskFeed) TaskFeedLikes(ctx context.Context) ([]*TaskFeedLike, error) {
	result, err := tf.Edges.TaskFeedLikesOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTaskFeedLikes().All(ctx)
	}
	return result, err
}

func (tf *TaskFeed) TaskFiles(ctx context.Context) ([]*TaskFile, error) {
	result, err := tf.Edges.TaskFilesOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTaskFiles().All(ctx)
	}
	return result, err
}

func (tfl *TaskFeedLike) Task(ctx context.Context) (*Task, error) {
	result, err := tfl.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tfl.QueryTask().Only(ctx)
	}
	return result, err
}

func (tfl *TaskFeedLike) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tfl.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tfl.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tfl *TaskFeedLike) Feed(ctx context.Context) (*TaskFeed, error) {
	result, err := tfl.Edges.FeedOrErr()
	if IsNotLoaded(err) {
		result, err = tfl.QueryFeed().Only(ctx)
	}
	return result, err
}

func (tf *TaskFile) Project(ctx context.Context) (*Project, error) {
	result, err := tf.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryProject().Only(ctx)
	}
	return result, err
}

func (tf *TaskFile) Task(ctx context.Context) (*Task, error) {
	result, err := tf.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTask().Only(ctx)
	}
	return result, err
}

func (tf *TaskFile) TaskFeed(ctx context.Context) (*TaskFeed, error) {
	result, err := tf.Edges.TaskFeedOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryTaskFeed().Only(ctx)
	}
	return result, err
}

func (tf *TaskFile) FileType(ctx context.Context) (*FileType, error) {
	result, err := tf.Edges.FileTypeOrErr()
	if IsNotLoaded(err) {
		result, err = tf.QueryFileType().Only(ctx)
	}
	return result, err
}

func (tl *TaskLike) Task(ctx context.Context) (*Task, error) {
	result, err := tl.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tl.QueryTask().Only(ctx)
	}
	return result, err
}

func (tl *TaskLike) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tl.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tl.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tl *TaskLike) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := tl.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = tl.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (tlcs *TaskListCompletedStatus) TeammateTaskListStatuses(ctx context.Context) ([]*TeammateTaskListStatus, error) {
	result, err := tlcs.Edges.TeammateTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = tlcs.QueryTeammateTaskListStatuses().All(ctx)
	}
	return result, err
}

func (tlcs *TaskListCompletedStatus) ProjectTaskListStatuses(ctx context.Context) ([]*ProjectTaskListStatus, error) {
	result, err := tlcs.Edges.ProjectTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = tlcs.QueryProjectTaskListStatuses().All(ctx)
	}
	return result, err
}

func (tlss *TaskListSortStatus) TeammateTaskListStatuses(ctx context.Context) ([]*TeammateTaskListStatus, error) {
	result, err := tlss.Edges.TeammateTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = tlss.QueryTeammateTaskListStatuses().All(ctx)
	}
	return result, err
}

func (tlss *TaskListSortStatus) ProjectTaskListStatuses(ctx context.Context) ([]*ProjectTaskListStatus, error) {
	result, err := tlss.Edges.ProjectTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = tlss.QueryProjectTaskListStatuses().All(ctx)
	}
	return result, err
}

func (tp *TaskPriority) Color(ctx context.Context) (*Color, error) {
	result, err := tp.Edges.ColorOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryColor().Only(ctx)
	}
	return result, err
}

func (tp *TaskPriority) Tasks(ctx context.Context) ([]*Task, error) {
	result, err := tp.Edges.TasksOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryTasks().All(ctx)
	}
	return result, err
}

func (tt *TaskTag) Task(ctx context.Context) (*Task, error) {
	result, err := tt.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTask().Only(ctx)
	}
	return result, err
}

func (tt *TaskTag) Tag(ctx context.Context) (*Tag, error) {
	result, err := tt.Edges.TagOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTag().Only(ctx)
	}
	return result, err
}

func (t *Teammate) Workspaces(ctx context.Context) ([]*Workspace, error) {
	result, err := t.Edges.WorkspacesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryWorkspaces().All(ctx)
	}
	return result, err
}

func (t *Teammate) Projects(ctx context.Context) ([]*Project, error) {
	result, err := t.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProjects().All(ctx)
	}
	return result, err
}

func (t *Teammate) ProjectTeammates(ctx context.Context) ([]*ProjectTeammate, error) {
	result, err := t.Edges.ProjectTeammatesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProjectTeammates().All(ctx)
	}
	return result, err
}

func (t *Teammate) WorkspaceTeammates(ctx context.Context) ([]*WorkspaceTeammate, error) {
	result, err := t.Edges.WorkspaceTeammatesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryWorkspaceTeammates().All(ctx)
	}
	return result, err
}

func (t *Teammate) FavoriteProjects(ctx context.Context) ([]*FavoriteProject, error) {
	result, err := t.Edges.FavoriteProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryFavoriteProjects().All(ctx)
	}
	return result, err
}

func (t *Teammate) FavoriteWorkspaces(ctx context.Context) ([]*FavoriteWorkspace, error) {
	result, err := t.Edges.FavoriteWorkspacesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryFavoriteWorkspaces().All(ctx)
	}
	return result, err
}

func (t *Teammate) TeammateTaskTabStatuses(ctx context.Context) ([]*TeammateTaskTabStatus, error) {
	result, err := t.Edges.TeammateTaskTabStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTaskTabStatuses().All(ctx)
	}
	return result, err
}

func (t *Teammate) TeammateTaskColumns(ctx context.Context) ([]*TeammateTaskColumn, error) {
	result, err := t.Edges.TeammateTaskColumnsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTaskColumns().All(ctx)
	}
	return result, err
}

func (t *Teammate) TeammateTaskListStatuses(ctx context.Context) ([]*TeammateTaskListStatus, error) {
	result, err := t.Edges.TeammateTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTaskListStatuses().All(ctx)
	}
	return result, err
}

func (t *Teammate) TeammateTaskSections(ctx context.Context) ([]*TeammateTaskSection, error) {
	result, err := t.Edges.TeammateTaskSectionsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTaskSections().All(ctx)
	}
	return result, err
}

func (t *Teammate) Tasks(ctx context.Context) ([]*Task, error) {
	result, err := t.Edges.TasksOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTasks().All(ctx)
	}
	return result, err
}

func (t *Teammate) TeammateTasks(ctx context.Context) ([]*TeammateTask, error) {
	result, err := t.Edges.TeammateTasksOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTeammateTasks().All(ctx)
	}
	return result, err
}

func (t *Teammate) TaskLikes(ctx context.Context) ([]*TaskLike, error) {
	result, err := t.Edges.TaskLikesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskLikes().All(ctx)
	}
	return result, err
}

func (t *Teammate) TaskCollaborators(ctx context.Context) ([]*TaskCollaborator, error) {
	result, err := t.Edges.TaskCollaboratorsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskCollaborators().All(ctx)
	}
	return result, err
}

func (t *Teammate) TaskFeeds(ctx context.Context) ([]*TaskFeed, error) {
	result, err := t.Edges.TaskFeedsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskFeeds().All(ctx)
	}
	return result, err
}

func (t *Teammate) TaskFeedLikes(ctx context.Context) ([]*TaskFeedLike, error) {
	result, err := t.Edges.TaskFeedLikesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryTaskFeedLikes().All(ctx)
	}
	return result, err
}

func (tt *TeammateTask) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tt.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tt *TeammateTask) Task(ctx context.Context) (*Task, error) {
	result, err := tt.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTask().Only(ctx)
	}
	return result, err
}

func (tt *TeammateTask) TeammateTaskSection(ctx context.Context) (*TeammateTaskSection, error) {
	result, err := tt.Edges.TeammateTaskSectionOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTeammateTaskSection().Only(ctx)
	}
	return result, err
}

func (tt *TeammateTask) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := tt.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (ttc *TeammateTaskColumn) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := ttc.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = ttc.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (ttc *TeammateTaskColumn) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := ttc.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = ttc.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (ttc *TeammateTaskColumn) TaskColumn(ctx context.Context) (*TaskColumn, error) {
	result, err := ttc.Edges.TaskColumnOrErr()
	if IsNotLoaded(err) {
		result, err = ttc.QueryTaskColumn().Only(ctx)
	}
	return result, err
}

func (ttls *TeammateTaskListStatus) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := ttls.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = ttls.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (ttls *TeammateTaskListStatus) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := ttls.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = ttls.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (ttls *TeammateTaskListStatus) TaskListCompletedStatus(ctx context.Context) (*TaskListCompletedStatus, error) {
	result, err := ttls.Edges.TaskListCompletedStatusOrErr()
	if IsNotLoaded(err) {
		result, err = ttls.QueryTaskListCompletedStatus().Only(ctx)
	}
	return result, err
}

func (ttls *TeammateTaskListStatus) TaskListSortStatus(ctx context.Context) (*TaskListSortStatus, error) {
	result, err := ttls.Edges.TaskListSortStatusOrErr()
	if IsNotLoaded(err) {
		result, err = ttls.QueryTaskListSortStatus().Only(ctx)
	}
	return result, err
}

func (tts *TeammateTaskSection) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := tts.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = tts.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tts *TeammateTaskSection) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := tts.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = tts.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (tts *TeammateTaskSection) TeammateTasks(ctx context.Context) ([]*TeammateTask, error) {
	result, err := tts.Edges.TeammateTasksOrErr()
	if IsNotLoaded(err) {
		result, err = tts.QueryTeammateTasks().All(ctx)
	}
	return result, err
}

func (ttts *TeammateTaskTabStatus) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := ttts.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = ttts.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (ttts *TeammateTaskTabStatus) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := ttts.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = ttts.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (tt *TestTodo) TestUser(ctx context.Context) (*TestUser, error) {
	result, err := tt.Edges.TestUserOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryTestUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (tt *TestTodo) Parent(ctx context.Context) (*TestTodo, error) {
	result, err := tt.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (tt *TestTodo) Children(ctx context.Context) ([]*TestTodo, error) {
	result, err := tt.Edges.ChildrenOrErr()
	if IsNotLoaded(err) {
		result, err = tt.QueryChildren().All(ctx)
	}
	return result, err
}

func (tu *TestUser) TestTodos(ctx context.Context) ([]*TestTodo, error) {
	result, err := tu.Edges.TestTodosOrErr()
	if IsNotLoaded(err) {
		result, err = tu.QueryTestTodos().All(ctx)
	}
	return result, err
}

func (w *Workspace) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := w.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammate().Only(ctx)
	}
	return result, err
}

func (w *Workspace) Projects(ctx context.Context) ([]*Project, error) {
	result, err := w.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryProjects().All(ctx)
	}
	return result, err
}

func (w *Workspace) WorkspaceTeammates(ctx context.Context) ([]*WorkspaceTeammate, error) {
	result, err := w.Edges.WorkspaceTeammatesOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryWorkspaceTeammates().All(ctx)
	}
	return result, err
}

func (w *Workspace) FavoriteWorkspaces(ctx context.Context) ([]*FavoriteWorkspace, error) {
	result, err := w.Edges.FavoriteWorkspacesOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryFavoriteWorkspaces().All(ctx)
	}
	return result, err
}

func (w *Workspace) TeammateTaskTabStatuses(ctx context.Context) ([]*TeammateTaskTabStatus, error) {
	result, err := w.Edges.TeammateTaskTabStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammateTaskTabStatuses().All(ctx)
	}
	return result, err
}

func (w *Workspace) TeammateTaskListStatuses(ctx context.Context) ([]*TeammateTaskListStatus, error) {
	result, err := w.Edges.TeammateTaskListStatusesOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammateTaskListStatuses().All(ctx)
	}
	return result, err
}

func (w *Workspace) TeammateTaskSections(ctx context.Context) ([]*TeammateTaskSection, error) {
	result, err := w.Edges.TeammateTaskSectionsOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammateTaskSections().All(ctx)
	}
	return result, err
}

func (w *Workspace) TaskLikes(ctx context.Context) ([]*TaskLike, error) {
	result, err := w.Edges.TaskLikesOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTaskLikes().All(ctx)
	}
	return result, err
}

func (w *Workspace) Tags(ctx context.Context) ([]*Tag, error) {
	result, err := w.Edges.TagsOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTags().All(ctx)
	}
	return result, err
}

func (w *Workspace) TeammateTaskColumns(ctx context.Context) ([]*TeammateTaskColumn, error) {
	result, err := w.Edges.TeammateTaskColumnsOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammateTaskColumns().All(ctx)
	}
	return result, err
}

func (w *Workspace) TeammateTasks(ctx context.Context) ([]*TeammateTask, error) {
	result, err := w.Edges.TeammateTasksOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryTeammateTasks().All(ctx)
	}
	return result, err
}

func (w *Workspace) DeletedTasksRef(ctx context.Context) ([]*DeletedTask, error) {
	result, err := w.Edges.DeletedTasksRefOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryDeletedTasksRef().All(ctx)
	}
	return result, err
}

func (wt *WorkspaceTeammate) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := wt.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = wt.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (wt *WorkspaceTeammate) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := wt.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = wt.QueryTeammate().Only(ctx)
	}
	return result, err
}
