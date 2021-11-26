// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTestTodo = "TestTodo"
	TypeTestUser = "TestUser"
)

// TestTodoMutation represents an operation that mutates the TestTodo nodes in the graph.
type TestTodoMutation struct {
	config
	op               Op
	typ              string
	id               *ulid.ID
	name             *string
	status           *testtodo.Status
	priority         *int
	addpriority      *int
	created_at       *time.Time
	updated_at       *time.Time
	clearedFields    map[string]struct{}
	test_user        *ulid.ID
	clearedtest_user bool
	done             bool
	oldValue         func(context.Context) (*TestTodo, error)
	predicates       []predicate.TestTodo
}

var _ ent.Mutation = (*TestTodoMutation)(nil)

// testtodoOption allows management of the mutation configuration using functional options.
type testtodoOption func(*TestTodoMutation)

// newTestTodoMutation creates new mutation for the TestTodo entity.
func newTestTodoMutation(c config, op Op, opts ...testtodoOption) *TestTodoMutation {
	m := &TestTodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTestTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTestTodoID sets the ID field of the mutation.
func withTestTodoID(id ulid.ID) testtodoOption {
	return func(m *TestTodoMutation) {
		var (
			err   error
			once  sync.Once
			value *TestTodo
		)
		m.oldValue = func(ctx context.Context) (*TestTodo, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TestTodo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTestTodo sets the old TestTodo of the mutation.
func withTestTodo(node *TestTodo) testtodoOption {
	return func(m *TestTodoMutation) {
		m.oldValue = func(context.Context) (*TestTodo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TestTodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TestTodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of TestTodo entities.
func (m *TestTodoMutation) SetID(id ulid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TestTodoMutation) ID() (id ulid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTestUserID sets the "test_user_id" field.
func (m *TestTodoMutation) SetTestUserID(u ulid.ID) {
	m.test_user = &u
}

// TestUserID returns the value of the "test_user_id" field in the mutation.
func (m *TestTodoMutation) TestUserID() (r ulid.ID, exists bool) {
	v := m.test_user
	if v == nil {
		return
	}
	return *v, true
}

// OldTestUserID returns the old "test_user_id" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldTestUserID(ctx context.Context) (v ulid.ID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTestUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTestUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTestUserID: %w", err)
	}
	return oldValue.TestUserID, nil
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (m *TestTodoMutation) ClearTestUserID() {
	m.test_user = nil
	m.clearedFields[testtodo.FieldTestUserID] = struct{}{}
}

// TestUserIDCleared returns if the "test_user_id" field was cleared in this mutation.
func (m *TestTodoMutation) TestUserIDCleared() bool {
	_, ok := m.clearedFields[testtodo.FieldTestUserID]
	return ok
}

// ResetTestUserID resets all changes to the "test_user_id" field.
func (m *TestTodoMutation) ResetTestUserID() {
	m.test_user = nil
	delete(m.clearedFields, testtodo.FieldTestUserID)
}

// SetName sets the "name" field.
func (m *TestTodoMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TestTodoMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TestTodoMutation) ResetName() {
	m.name = nil
}

// SetStatus sets the "status" field.
func (m *TestTodoMutation) SetStatus(t testtodo.Status) {
	m.status = &t
}

// Status returns the value of the "status" field in the mutation.
func (m *TestTodoMutation) Status() (r testtodo.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldStatus(ctx context.Context) (v testtodo.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *TestTodoMutation) ResetStatus() {
	m.status = nil
}

// SetPriority sets the "priority" field.
func (m *TestTodoMutation) SetPriority(i int) {
	m.priority = &i
	m.addpriority = nil
}

// Priority returns the value of the "priority" field in the mutation.
func (m *TestTodoMutation) Priority() (r int, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldPriority(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// AddPriority adds i to the "priority" field.
func (m *TestTodoMutation) AddPriority(i int) {
	if m.addpriority != nil {
		*m.addpriority += i
	} else {
		m.addpriority = &i
	}
}

// AddedPriority returns the value that was added to the "priority" field in this mutation.
func (m *TestTodoMutation) AddedPriority() (r int, exists bool) {
	v := m.addpriority
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriority resets all changes to the "priority" field.
func (m *TestTodoMutation) ResetPriority() {
	m.priority = nil
	m.addpriority = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TestTodoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TestTodoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TestTodoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TestTodoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TestTodoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the TestTodo entity.
// If the TestTodo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestTodoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TestTodoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// ClearTestUser clears the "test_user" edge to the TestUser entity.
func (m *TestTodoMutation) ClearTestUser() {
	m.clearedtest_user = true
}

// TestUserCleared reports if the "test_user" edge to the TestUser entity was cleared.
func (m *TestTodoMutation) TestUserCleared() bool {
	return m.TestUserIDCleared() || m.clearedtest_user
}

// TestUserIDs returns the "test_user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// TestUserID instead. It exists only for internal usage by the builders.
func (m *TestTodoMutation) TestUserIDs() (ids []ulid.ID) {
	if id := m.test_user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetTestUser resets all changes to the "test_user" edge.
func (m *TestTodoMutation) ResetTestUser() {
	m.test_user = nil
	m.clearedtest_user = false
}

// Where appends a list predicates to the TestTodoMutation builder.
func (m *TestTodoMutation) Where(ps ...predicate.TestTodo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TestTodoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TestTodo).
func (m *TestTodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TestTodoMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.test_user != nil {
		fields = append(fields, testtodo.FieldTestUserID)
	}
	if m.name != nil {
		fields = append(fields, testtodo.FieldName)
	}
	if m.status != nil {
		fields = append(fields, testtodo.FieldStatus)
	}
	if m.priority != nil {
		fields = append(fields, testtodo.FieldPriority)
	}
	if m.created_at != nil {
		fields = append(fields, testtodo.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, testtodo.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TestTodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case testtodo.FieldTestUserID:
		return m.TestUserID()
	case testtodo.FieldName:
		return m.Name()
	case testtodo.FieldStatus:
		return m.Status()
	case testtodo.FieldPriority:
		return m.Priority()
	case testtodo.FieldCreatedAt:
		return m.CreatedAt()
	case testtodo.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TestTodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case testtodo.FieldTestUserID:
		return m.OldTestUserID(ctx)
	case testtodo.FieldName:
		return m.OldName(ctx)
	case testtodo.FieldStatus:
		return m.OldStatus(ctx)
	case testtodo.FieldPriority:
		return m.OldPriority(ctx)
	case testtodo.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case testtodo.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown TestTodo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TestTodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case testtodo.FieldTestUserID:
		v, ok := value.(ulid.ID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTestUserID(v)
		return nil
	case testtodo.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case testtodo.FieldStatus:
		v, ok := value.(testtodo.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case testtodo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	case testtodo.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case testtodo.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown TestTodo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TestTodoMutation) AddedFields() []string {
	var fields []string
	if m.addpriority != nil {
		fields = append(fields, testtodo.FieldPriority)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TestTodoMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case testtodo.FieldPriority:
		return m.AddedPriority()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TestTodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	case testtodo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriority(v)
		return nil
	}
	return fmt.Errorf("unknown TestTodo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TestTodoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(testtodo.FieldTestUserID) {
		fields = append(fields, testtodo.FieldTestUserID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TestTodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TestTodoMutation) ClearField(name string) error {
	switch name {
	case testtodo.FieldTestUserID:
		m.ClearTestUserID()
		return nil
	}
	return fmt.Errorf("unknown TestTodo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TestTodoMutation) ResetField(name string) error {
	switch name {
	case testtodo.FieldTestUserID:
		m.ResetTestUserID()
		return nil
	case testtodo.FieldName:
		m.ResetName()
		return nil
	case testtodo.FieldStatus:
		m.ResetStatus()
		return nil
	case testtodo.FieldPriority:
		m.ResetPriority()
		return nil
	case testtodo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case testtodo.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown TestTodo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TestTodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.test_user != nil {
		edges = append(edges, testtodo.EdgeTestUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TestTodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case testtodo.EdgeTestUser:
		if id := m.test_user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TestTodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TestTodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TestTodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtest_user {
		edges = append(edges, testtodo.EdgeTestUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TestTodoMutation) EdgeCleared(name string) bool {
	switch name {
	case testtodo.EdgeTestUser:
		return m.clearedtest_user
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TestTodoMutation) ClearEdge(name string) error {
	switch name {
	case testtodo.EdgeTestUser:
		m.ClearTestUser()
		return nil
	}
	return fmt.Errorf("unknown TestTodo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TestTodoMutation) ResetEdge(name string) error {
	switch name {
	case testtodo.EdgeTestUser:
		m.ResetTestUser()
		return nil
	}
	return fmt.Errorf("unknown TestTodo edge %s", name)
}

// TestUserMutation represents an operation that mutates the TestUser nodes in the graph.
type TestUserMutation struct {
	config
	op                Op
	typ               string
	id                *ulid.ID
	name              *string
	age               *int
	addage            *int
	created_at        *time.Time
	updated_at        *time.Time
	clearedFields     map[string]struct{}
	test_todos        map[ulid.ID]struct{}
	removedtest_todos map[ulid.ID]struct{}
	clearedtest_todos bool
	done              bool
	oldValue          func(context.Context) (*TestUser, error)
	predicates        []predicate.TestUser
}

var _ ent.Mutation = (*TestUserMutation)(nil)

// testuserOption allows management of the mutation configuration using functional options.
type testuserOption func(*TestUserMutation)

// newTestUserMutation creates new mutation for the TestUser entity.
func newTestUserMutation(c config, op Op, opts ...testuserOption) *TestUserMutation {
	m := &TestUserMutation{
		config:        c,
		op:            op,
		typ:           TypeTestUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTestUserID sets the ID field of the mutation.
func withTestUserID(id ulid.ID) testuserOption {
	return func(m *TestUserMutation) {
		var (
			err   error
			once  sync.Once
			value *TestUser
		)
		m.oldValue = func(ctx context.Context) (*TestUser, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TestUser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTestUser sets the old TestUser of the mutation.
func withTestUser(node *TestUser) testuserOption {
	return func(m *TestUserMutation) {
		m.oldValue = func(context.Context) (*TestUser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TestUserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TestUserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of TestUser entities.
func (m *TestUserMutation) SetID(id ulid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TestUserMutation) ID() (id ulid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *TestUserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TestUserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the TestUser entity.
// If the TestUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestUserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TestUserMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *TestUserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *TestUserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the TestUser entity.
// If the TestUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestUserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *TestUserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *TestUserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *TestUserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TestUserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TestUserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the TestUser entity.
// If the TestUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestUserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TestUserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TestUserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TestUserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the TestUser entity.
// If the TestUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TestUserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TestUserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddTestTodoIDs adds the "test_todos" edge to the TestTodo entity by ids.
func (m *TestUserMutation) AddTestTodoIDs(ids ...ulid.ID) {
	if m.test_todos == nil {
		m.test_todos = make(map[ulid.ID]struct{})
	}
	for i := range ids {
		m.test_todos[ids[i]] = struct{}{}
	}
}

// ClearTestTodos clears the "test_todos" edge to the TestTodo entity.
func (m *TestUserMutation) ClearTestTodos() {
	m.clearedtest_todos = true
}

// TestTodosCleared reports if the "test_todos" edge to the TestTodo entity was cleared.
func (m *TestUserMutation) TestTodosCleared() bool {
	return m.clearedtest_todos
}

// RemoveTestTodoIDs removes the "test_todos" edge to the TestTodo entity by IDs.
func (m *TestUserMutation) RemoveTestTodoIDs(ids ...ulid.ID) {
	if m.removedtest_todos == nil {
		m.removedtest_todos = make(map[ulid.ID]struct{})
	}
	for i := range ids {
		delete(m.test_todos, ids[i])
		m.removedtest_todos[ids[i]] = struct{}{}
	}
}

// RemovedTestTodos returns the removed IDs of the "test_todos" edge to the TestTodo entity.
func (m *TestUserMutation) RemovedTestTodosIDs() (ids []ulid.ID) {
	for id := range m.removedtest_todos {
		ids = append(ids, id)
	}
	return
}

// TestTodosIDs returns the "test_todos" edge IDs in the mutation.
func (m *TestUserMutation) TestTodosIDs() (ids []ulid.ID) {
	for id := range m.test_todos {
		ids = append(ids, id)
	}
	return
}

// ResetTestTodos resets all changes to the "test_todos" edge.
func (m *TestUserMutation) ResetTestTodos() {
	m.test_todos = nil
	m.clearedtest_todos = false
	m.removedtest_todos = nil
}

// Where appends a list predicates to the TestUserMutation builder.
func (m *TestUserMutation) Where(ps ...predicate.TestUser) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TestUserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (TestUser).
func (m *TestUserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TestUserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, testuser.FieldName)
	}
	if m.age != nil {
		fields = append(fields, testuser.FieldAge)
	}
	if m.created_at != nil {
		fields = append(fields, testuser.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, testuser.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TestUserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case testuser.FieldName:
		return m.Name()
	case testuser.FieldAge:
		return m.Age()
	case testuser.FieldCreatedAt:
		return m.CreatedAt()
	case testuser.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TestUserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case testuser.FieldName:
		return m.OldName(ctx)
	case testuser.FieldAge:
		return m.OldAge(ctx)
	case testuser.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case testuser.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown TestUser field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TestUserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case testuser.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case testuser.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case testuser.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case testuser.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown TestUser field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TestUserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, testuser.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TestUserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case testuser.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TestUserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case testuser.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown TestUser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TestUserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TestUserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TestUserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TestUser nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TestUserMutation) ResetField(name string) error {
	switch name {
	case testuser.FieldName:
		m.ResetName()
		return nil
	case testuser.FieldAge:
		m.ResetAge()
		return nil
	case testuser.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case testuser.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown TestUser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TestUserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.test_todos != nil {
		edges = append(edges, testuser.EdgeTestTodos)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TestUserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case testuser.EdgeTestTodos:
		ids := make([]ent.Value, 0, len(m.test_todos))
		for id := range m.test_todos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TestUserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtest_todos != nil {
		edges = append(edges, testuser.EdgeTestTodos)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TestUserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case testuser.EdgeTestTodos:
		ids := make([]ent.Value, 0, len(m.removedtest_todos))
		for id := range m.removedtest_todos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TestUserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtest_todos {
		edges = append(edges, testuser.EdgeTestTodos)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TestUserMutation) EdgeCleared(name string) bool {
	switch name {
	case testuser.EdgeTestTodos:
		return m.clearedtest_todos
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TestUserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown TestUser unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TestUserMutation) ResetEdge(name string) error {
	switch name {
	case testuser.EdgeTestTodos:
		m.ResetTestTodos()
		return nil
	}
	return fmt.Errorf("unknown TestUser edge %s", name)
}
