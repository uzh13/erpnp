package v1_0

import (
	"encoding/json"
	"fmt"
)

// ERPN represents the root evolving resolutive process notation object
type ERPN struct {
	Type    TypeInfo       `json:"type" yaml:"type" toml:"type" json5:"type"`
	Content *ContentObject `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty" json5:"content,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`

	Tree *Tree
}

// UnmarshalJSON implements custom JSON unmarshaling to preserve unknown fields
func (e *ERPN) UnmarshalJSON(data []byte) error {
	type Alias ERPN
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	// First unmarshal to get known fields
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Unmarshal to map to get all fields
	var all map[string]interface{}
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	// Remove known fields and store the rest in Extra
	delete(all, "type")
	delete(all, "content")

	if len(all) > 0 {
		e.Extra = all
	}

	return nil
}

// MarshalJSON implements custom JSON marshaling to include unknown fields
func (e *ERPN) MarshalJSON() ([]byte, error) {
	type Alias ERPN
	aux := (*Alias)(e)

	// Marshal known fields first
	known, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}

	if len(e.Extra) == 0 {
		return known, nil
	}

	// Merge with extra fields
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}

	for k, v := range e.Extra {
		knownMap[k] = v
	}

	return json.Marshal(knownMap)
}

func (s *ERPN) Validate() []error {
	var errors, errs []error

	errs = s.Type.Validate()
	for _, err := range errs {
		errors = append(errors, fmt.Errorf("type: %s", err))
	}

	if s.Content != nil {
		errs = s.Content.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("content: %s", err))
		}
	}

	return errors
}
