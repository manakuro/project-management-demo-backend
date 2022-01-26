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

func (ttc *TeammateTaskColumn) Teammate(ctx context.Context) (*Teammate, error) {
	result, err := ttc.Edges.TeammateOrErr()
	if IsNotLoaded(err) {
		result, err = ttc.QueryTeammate().Only(ctx)
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
