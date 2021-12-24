package testuserprofile

import (
	"encoding/json"
	"fmt"
	"io"
	"project-management-demo-backend/pkg/util/conversion"
)

// TestUserProfile of profile
type TestUserProfile struct {
	Address string `json:"address"`
}

// UnmarshalGQL implements the graphql.Unmarshaller interface.
func (t *TestUserProfile) UnmarshalGQL(v interface{}) error {
	return t.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface.
func (t TestUserProfile) MarshalGQL(w io.Writer) {
	val, _ := json.Marshal(t)
	_, _ = io.WriteString(w, string(val))
}

// Scan implements the Scanner interface.
func (t *TestUserProfile) Scan(src interface{}) error {

	switch s := src.(type) {
	case map[string]interface{}:
		var val TestUserProfile
		err := conversion.MapToStruct(s, &val)
		if err != nil {
			return fmt.Errorf("testuserprofile: invalid data format %v", s)
		}
		*t = val
	default:
		return fmt.Errorf("testuserprofile: expected map data %v", s)
	}

	return nil
}
