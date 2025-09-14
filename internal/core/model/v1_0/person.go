package v1_0

import "fmt"

// Person represents person description
type Person struct {
	Name          *string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	ImpactPercent *Impact `json:"impactPercent,omitempty" yaml:"impactPercent,omitempty" toml:"impactPercent,omitempty" json5:"impactPercent,omitempty"`
	Position      *string `json:"position,omitempty" yaml:"position,omitempty" toml:"position,omitempty" json5:"position,omitempty"`
	Comment       *string `json:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty" json5:"comment,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (p *Person) UnmarshalJSON(data []byte) error {
	type Alias Person
	aux := (*Alias)(p)
	return unmarshalWithExtra(data, aux, &p.Extra, "name", "impactPercent", "position", "comment")
}

func (p *Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return marshalWithExtra((*Alias)(p), p.Extra)
}

func (s *Person) Validate() []error {
	var errors, errs []error

	if s.ImpactPercent != nil {
		errs = s.ImpactPercent.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("impactPercent: %s", err))
		}
	}

	return errors
}
