package v1_0

import (
	"fmt"
	"time"
)

// Timing represents time gap of object activity with authors and source of changing
type Timing struct {
	Start         *time.Time `json:"start,omitempty" yaml:"start,omitempty" toml:"start,omitempty" json5:"start,omitempty"`
	End           *time.Time `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty" json5:"end,omitempty"`
	ImpactPercent *Impact    `json:"impactPercent,omitempty" yaml:"impactPercent,omitempty" toml:"impactPercent,omitempty" json5:"impactPercent,omitempty"`
	Authors       []*Person  `json:"authors,omitempty" yaml:"authors,omitempty" toml:"authors,omitempty" json5:"authors,omitempty"`
	Source        []string   `json:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty" json5:"source,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (t *Timing) UnmarshalJSON(data []byte) error {
	type Alias Timing
	aux := (*Alias)(t)
	return unmarshalWithExtra(data, aux, &t.Extra, "start", "end", "impactPercent", "authors", "source")
}

func (t *Timing) MarshalJSON() ([]byte, error) {
	type Alias Timing
	return marshalWithExtra((*Alias)(t), t.Extra)
}

func (s *Timing) Validate() []error {
	var errors, errs []error

	if s.ImpactPercent != nil {
		errs = s.ImpactPercent.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("impactPercent: %s", err))
		}
	}

	for i, author := range s.Authors {
		if author == nil {
			continue
		}

		errs = author.Validate()
		for _, err := range errs {
			errors = append(errors, fmt.Errorf("authors[%d]: %s", i, err))
		}
	}

	return errors
}
