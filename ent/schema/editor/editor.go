package editor

import (
	"encoding/json"
	"fmt"
	"io"
	"project-management-demo-backend/pkg/util/conversion"
)

// Attrs is a custom attributes
type Attrs struct {
	MentionID   string `json:"mentionId"`
	MentionType string `json:"mentionType"`
}

// Content is a content
type Content struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Attrs Attrs  `json:"attrs"`
}

// DescriptionContent is a content
type DescriptionContent struct {
	Type    string    `json:"type"`
	Content []Content `json:"content"`
}

// Description is a json type of description
type Description struct {
	Type    string               `json:"type"`
	Content []DescriptionContent `json:"content"`
}

// UnmarshalGQL implements the graphql.Unmarshaller interface.
func (d *Description) UnmarshalGQL(v interface{}) error {
	return d.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface.
func (d Description) MarshalGQL(w io.Writer) {
	val, _ := json.Marshal(d)
	_, _ = io.WriteString(w, string(val))
}

// Scan implements the Scanner interface.
func (d *Description) Scan(src interface{}) error {
	switch s := src.(type) {
	case map[string]interface{}:
		var val Description
		err := conversion.MapToStruct(s, &val)
		if err != nil {
			return fmt.Errorf("editor: invalid data  format %v", s)
		}
		*d = val
	default:
		return fmt.Errorf("editor: exptect map data %v", s)
	}

	return nil
}
