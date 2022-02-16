package model

import (
	"project-management-demo-backend/ent/schema/editor"
)

// EditorDescription is an editor description
type EditorDescription = editor.Description
type EditorDescriptionContent = editor.DescriptionContent // revive:disable-line:exported
type EditorDescriptionContentContent = editor.Content     // revive:disable-line:exported
type EditorDescriptionContentContentAttrs = editor.Attrs  // revive:disable-line:exported

// EditorDescriptionInput is an editor description input
type EditorDescriptionInput = editor.Description
type EditorDescriptionContentInput = editor.DescriptionContent // revive:disable-line:exported
type EditorDescriptionContentContentInput = editor.Content     // revive:disable-line:exported
type EditorDescriptionContentContentAttrsInput = editor.Attrs  // revive:disable-line:exported

// DefaultEditorDescription returns default json value for editor.
func DefaultEditorDescription() EditorDescription {
	d := EditorDescription{}
	return d.DefaultValue()
}
