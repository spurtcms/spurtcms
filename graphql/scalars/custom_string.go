package scalars

import (
	"encoding/json"
)

// Define the maximum length for your string
// const MaxLength int = 1000000000000000000

// CustomString is a custom scalar for strings with length constraints
type CustomString string

func (s *CustomString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = CustomString(str)
	return nil
}

func (s CustomString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}
