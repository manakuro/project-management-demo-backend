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
