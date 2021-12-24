package testuserprofile

import (
	"encoding/json"
	"fmt"
	"io"
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
	case string:
		var val TestUserProfile
		err := json.Unmarshal([]byte(s), &val)
		if err != nil {
			return fmt.Errorf("testuserprofile: invalid json string %v", s)
		}
		*t = val
	case []byte:
		var val TestUserProfile
		err := json.Unmarshal(s, &val)
		if err != nil {
			return fmt.Errorf("testuserprofile: invalid json string %v", s)
		}
		*t = val
	case map[string]interface{}:
		var val TestUserProfile
		err := MapToStruct(s, &val)
		if err != nil {
			return fmt.Errorf("testuserprofile: invalid data format %v", s)
		}
		*t = val
	default:
		return fmt.Errorf("testuserprofile: expected a json string %v", s)
	}

	return nil
}

// MapToStruct converts map to struct.
func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

//// Value implements the driver Valuer interface.
//func (t *TestUserProfile) Value() (driver.Value, error) {
//	val, err := json.Marshal(t)
//	if err != nil {
//		return nil, fmt.Errorf("testuserprofile: expected a json object %v", t)
//	}
//
//	return string(val), nil
//}
