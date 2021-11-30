// Code generated by entc, DO NOT EDIT.

package ent

import (
	"project-management-demo-backend/ent/schema"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	testtodoMixin := schema.TestTodo{}.Mixin()
	testtodoMixinFields0 := testtodoMixin[0].Fields()
	_ = testtodoMixinFields0
	testtodoMixinFields1 := testtodoMixin[1].Fields()
	_ = testtodoMixinFields1
	testtodoMixinFields2 := testtodoMixin[2].Fields()
	_ = testtodoMixinFields2
	testtodoFields := schema.TestTodo{}.Fields()
	_ = testtodoFields
	// testtodoDescName is the schema descriptor for name field.
	testtodoDescName := testtodoMixinFields1[1].Descriptor()
	// testtodo.DefaultName holds the default value on creation for the name field.
	testtodo.DefaultName = testtodoDescName.Default.(string)
	// testtodoDescPriority is the schema descriptor for priority field.
	testtodoDescPriority := testtodoMixinFields1[3].Descriptor()
	// testtodo.DefaultPriority holds the default value on creation for the priority field.
	testtodo.DefaultPriority = testtodoDescPriority.Default.(int)
	// testtodoDescCreatedAt is the schema descriptor for created_at field.
	testtodoDescCreatedAt := testtodoMixinFields2[0].Descriptor()
	// testtodo.DefaultCreatedAt holds the default value on creation for the created_at field.
	testtodo.DefaultCreatedAt = testtodoDescCreatedAt.Default.(func() time.Time)
	// testtodoDescUpdatedAt is the schema descriptor for updated_at field.
	testtodoDescUpdatedAt := testtodoMixinFields2[1].Descriptor()
	// testtodo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	testtodo.DefaultUpdatedAt = testtodoDescUpdatedAt.Default.(func() time.Time)
	// testtodoDescID is the schema descriptor for id field.
	testtodoDescID := testtodoMixinFields0[0].Descriptor()
	// testtodo.DefaultID holds the default value on creation for the id field.
	testtodo.DefaultID = testtodoDescID.Default.(func() ulid.ID)
	testuserMixin := schema.TestUser{}.Mixin()
	testuserMixinFields0 := testuserMixin[0].Fields()
	_ = testuserMixinFields0
	testuserMixinFields1 := testuserMixin[1].Fields()
	_ = testuserMixinFields1
	testuserMixinFields2 := testuserMixin[2].Fields()
	_ = testuserMixinFields2
	testuserFields := schema.TestUser{}.Fields()
	_ = testuserFields
	// testuserDescName is the schema descriptor for name field.
	testuserDescName := testuserMixinFields1[0].Descriptor()
	// testuser.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testuser.NameValidator = testuserDescName.Validators[0].(func(string) error)
	// testuserDescCreatedAt is the schema descriptor for created_at field.
	testuserDescCreatedAt := testuserMixinFields2[0].Descriptor()
	// testuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	testuser.DefaultCreatedAt = testuserDescCreatedAt.Default.(func() time.Time)
	// testuserDescUpdatedAt is the schema descriptor for updated_at field.
	testuserDescUpdatedAt := testuserMixinFields2[1].Descriptor()
	// testuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	testuser.DefaultUpdatedAt = testuserDescUpdatedAt.Default.(func() time.Time)
	// testuserDescID is the schema descriptor for id field.
	testuserDescID := testuserMixinFields0[0].Descriptor()
	// testuser.DefaultID holds the default value on creation for the id field.
	testuser.DefaultID = testuserDescID.Default.(func() ulid.ID)
}
