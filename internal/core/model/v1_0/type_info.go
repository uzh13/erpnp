package v1_0

import "fmt"

// TypeInfo represents the type information
type TypeInfo struct {
	Name    string `json:"name" yaml:"name" toml:"name" json5:"name"`
	Version SemVer `json:"version" yaml:"version" toml:"version" json5:"version"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (t *TypeInfo) UnmarshalJSON(data []byte) error {
	type Alias TypeInfo
	aux := (*Alias)(t)
	return unmarshalWithExtra(data, aux, &t.Extra, "name", "version")
}

func (t *TypeInfo) MarshalJSON() ([]byte, error) {
	type Alias TypeInfo
	return marshalWithExtra((*Alias)(t), t.Extra)
}

func (s TypeInfo) Validate() []error {
	errors := make([]error, 0, 3)
	// Validate required fields
	if s.Name == "" {
		errors = append(errors, fmt.Errorf("name is required"))
	}
	if s.Version == "" {
		errors = append(errors, fmt.Errorf("version is required"))
	}

	errs := s.Version.Validate()
	for _, err := range errs {
		errors = append(errors, fmt.Errorf("version: %s", err))
	}

	return errors
}
