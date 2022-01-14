// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     ulid.ID  `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string    `json:"type,omitempty"` // edge type.
	Name string    `json:"name,omitempty"` // edge name.
	IDs  []ulid.ID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (c *Color) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Color",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Color); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "color",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Hex); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "hex",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "ProjectBaseColor",
		Name: "project_base_colors",
	}
	err = c.QueryProjectBaseColors().
		Select(projectbasecolor.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (i *Icon) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     i.ID,
		Type:   "Icon",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(i.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Icon); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "icon",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Project",
		Name: "projects",
	}
	err = i.QueryProjects().
		Select(project.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (pr *Project) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pr.ID,
		Type:   "Project",
		Fields: make([]*Field, 10),
		Edges:  make([]*Edge, 5),
	}
	var buf []byte
	if buf, err = json.Marshal(pr.WorkspaceID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "ulid.ID",
		Name:  "workspace_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.ProjectBaseColorID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "ulid.ID",
		Name:  "project_base_color_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.IconID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "ulid.ID",
		Name:  "icon_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "ulid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Description); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "editor.Description",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.DescriptionTitle); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "description_title",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.DueDate); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "due_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workspace",
		Name: "workspace",
	}
	err = pr.QueryWorkspace().
		Select(workspace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "ProjectBaseColor",
		Name: "project_base_color",
	}
	err = pr.QueryProjectBaseColor().
		Select(projectbasecolor.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Icon",
		Name: "icon",
	}
	err = pr.QueryIcon().
		Select(icon.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Teammate",
		Name: "teammate",
	}
	err = pr.QueryTeammate().
		Select(teammate.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "ProjectTeammate",
		Name: "project_teammates",
	}
	err = pr.QueryProjectTeammates().
		Select(projectteammate.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (pbc *ProjectBaseColor) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pbc.ID,
		Type:   "ProjectBaseColor",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(pbc.ColorID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "ulid.ID",
		Name:  "color_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pbc.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pbc.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Project",
		Name: "projects",
	}
	err = pbc.QueryProjects().
		Select(project.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Color",
		Name: "color",
	}
	err = pbc.QueryColor().
		Select(color.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (pt *ProjectTeammate) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pt.ID,
		Type:   "ProjectTeammate",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(pt.ProjectID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "ulid.ID",
		Name:  "project_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pt.TeammateID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "ulid.ID",
		Name:  "teammate_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pt.Role); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "role",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pt.IsOwner); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "bool",
		Name:  "is_owner",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pt.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pt.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Project",
		Name: "project",
	}
	err = pt.QueryProject().
		Select(project.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Teammate",
		Name: "teammate",
	}
	err = pt.QueryTeammate().
		Select(teammate.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (t *Teammate) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Teammate",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(t.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Image); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "image",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Email); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workspace",
		Name: "workspaces",
	}
	err = t.QueryWorkspaces().
		Select(workspace.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Project",
		Name: "projects",
	}
	err = t.QueryProjects().
		Select(project.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "ProjectTeammate",
		Name: "project_teammates",
	}
	err = t.QueryProjectTeammates().
		Select(projectteammate.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (tt *TestTodo) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     tt.ID,
		Type:   "TestTodo",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(tt.TestUserID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "ulid.ID",
		Name:  "test_user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tt.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tt.Status); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "testtodo.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tt.Priority); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "priority",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tt.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tt.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "TestUser",
		Name: "test_user",
	}
	err = tt.QueryTestUser().
		Select(testuser.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (tu *TestUser) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     tu.ID,
		Type:   "TestUser",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(tu.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tu.Age); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "age",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tu.Profile); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "testuserprofile.TestUserProfile",
		Name:  "profile",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tu.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tu.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "TestTodo",
		Name: "test_todos",
	}
	err = tu.QueryTestTodos().
		Select(testtodo.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (w *Workspace) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     w.ID,
		Type:   "Workspace",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(w.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "ulid.ID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.Description); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "editor.Description",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Teammate",
		Name: "teammate",
	}
	err = w.QueryTeammate().
		Select(teammate.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Project",
		Name: "projects",
	}
	err = w.QueryProjects().
		Select(project.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id ulid.ID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, ulid.ID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, ulid.ID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, ulid.ID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id ulid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id ulid.ID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id ulid.ID) (Noder, error) {
	switch table {
	case color.Table:
		n, err := c.Color.Query().
			Where(color.ID(id)).
			CollectFields(ctx, "Color").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case icon.Table:
		n, err := c.Icon.Query().
			Where(icon.ID(id)).
			CollectFields(ctx, "Icon").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case project.Table:
		n, err := c.Project.Query().
			Where(project.ID(id)).
			CollectFields(ctx, "Project").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case projectbasecolor.Table:
		n, err := c.ProjectBaseColor.Query().
			Where(projectbasecolor.ID(id)).
			CollectFields(ctx, "ProjectBaseColor").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case projectteammate.Table:
		n, err := c.ProjectTeammate.Query().
			Where(projectteammate.ID(id)).
			CollectFields(ctx, "ProjectTeammate").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case teammate.Table:
		n, err := c.Teammate.Query().
			Where(teammate.ID(id)).
			CollectFields(ctx, "Teammate").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case testtodo.Table:
		n, err := c.TestTodo.Query().
			Where(testtodo.ID(id)).
			CollectFields(ctx, "TestTodo").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case testuser.Table:
		n, err := c.TestUser.Query().
			Where(testuser.ID(id)).
			CollectFields(ctx, "TestUser").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workspace.Table:
		n, err := c.Workspace.Query().
			Where(workspace.ID(id)).
			CollectFields(ctx, "Workspace").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []ulid.ID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]ulid.ID)
	id2idx := make(map[ulid.ID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []ulid.ID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[ulid.ID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case color.Table:
		nodes, err := c.Color.Query().
			Where(color.IDIn(ids...)).
			CollectFields(ctx, "Color").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case icon.Table:
		nodes, err := c.Icon.Query().
			Where(icon.IDIn(ids...)).
			CollectFields(ctx, "Icon").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case project.Table:
		nodes, err := c.Project.Query().
			Where(project.IDIn(ids...)).
			CollectFields(ctx, "Project").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case projectbasecolor.Table:
		nodes, err := c.ProjectBaseColor.Query().
			Where(projectbasecolor.IDIn(ids...)).
			CollectFields(ctx, "ProjectBaseColor").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case projectteammate.Table:
		nodes, err := c.ProjectTeammate.Query().
			Where(projectteammate.IDIn(ids...)).
			CollectFields(ctx, "ProjectTeammate").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case teammate.Table:
		nodes, err := c.Teammate.Query().
			Where(teammate.IDIn(ids...)).
			CollectFields(ctx, "Teammate").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case testtodo.Table:
		nodes, err := c.TestTodo.Query().
			Where(testtodo.IDIn(ids...)).
			CollectFields(ctx, "TestTodo").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case testuser.Table:
		nodes, err := c.TestUser.Query().
			Where(testuser.IDIn(ids...)).
			CollectFields(ctx, "TestUser").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workspace.Table:
		nodes, err := c.Workspace.Query().
			Where(workspace.IDIn(ids...)).
			CollectFields(ctx, "Workspace").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
