package v1_0

import "fmt"

// ContentObject represents the main object for notation
type ContentObject struct {
	Link    *string          `json:"link,omitempty" yaml:"link,omitempty" toml:"link,omitempty" json5:"link,omitempty"`
	Name    *string          `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	Version *SemVer          `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty" json5:"version,omitempty"`
	Essence *Essence         `json:"essence,omitempty" yaml:"essence,omitempty" toml:"essence,omitempty" json5:"essence,omitempty"`
	Timing  *Timing          `json:"timing,omitempty" yaml:"timing,omitempty" toml:"timing,omitempty" json5:"timing,omitempty"`
	Content []*ContentObject `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty" json5:"content,omitempty"`
	History []*ContentObject `json:"history,omitempty" yaml:"history,omitempty" toml:"history,omitempty" json5:"history,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`

	Node *Node
}

func (c *ContentObject) UnmarshalJSON(data []byte) error {
	type Alias ContentObject
	aux := (*Alias)(c)
	return unmarshalWithExtra(data, aux, &c.Extra, "link", "name", "version", "essence", "timing", "content", "history")
}

func (c *ContentObject) MarshalJSON() ([]byte, error) {
	type Alias ContentObject
	return marshalWithExtra((*Alias)(c), c.Extra)
}

func (s *ContentObject) Validate() []error {
	var errors, errs []error

	// If we have link, it does not matter what else we have
	if s.Link != nil {
		return errors
	}

	if s.Version != nil {
		errs = s.Version.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("version: %s", err))
		}
	}

	if s.Timing != nil {
		errs = s.Timing.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("timing: %s", err))
		}
	}

	for i, content := range s.Content {
		if content == nil {
			continue
		}

		errs = content.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("content[%d]: %s", i, err))
		}
	}

	return errors
}
