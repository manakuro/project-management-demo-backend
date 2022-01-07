// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (c *Color) Projects(ctx context.Context) (*Project, error) {
	result, err := c.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryProjects().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Icon) Projects(ctx context.Context) (*Project, error) {
	result, err := i.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryProjects().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Project) Workspace(ctx context.Context) (*Workspace, error) {
	result, err := pr.Edges.WorkspaceOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryWorkspace().Only(ctx)
	}
	return result, err
}

func (pr *Project) Color(ctx context.Context) (*Color, error) {
	result, err := pr.Edges.ColorOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryColor().Only(ctx)
	}
	return result, err
}

func (pr *Project) Icon(ctx context.Context) (*Icon, error) {
	result, err := pr.Edges.IconOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryIcon().Only(ctx)
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

func (t *Teammate) Workspaces(ctx context.Context) (*Workspace, error) {
	result, err := t.Edges.WorkspacesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryWorkspaces().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Teammate) Projects(ctx context.Context) (*Project, error) {
	result, err := t.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProjects().Only(ctx)
	}
	return result, MaskNotFound(err)
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
