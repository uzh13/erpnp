package v1_0

// Essence represents the main description
type Essence struct {
	Contradiction *Contradiction `json:"contradiction,omitempty" yaml:"contradiction,omitempty" toml:"contradiction,omitempty" json5:"contradiction,omitempty"`
	Synthesis     *Synthesis     `json:"synthesis,omitempty" yaml:"synthesis,omitempty" toml:"synthesis,omitempty" json5:"synthesis,omitempty"`
	Realization   *Realization   `json:"realization,omitempty" yaml:"realization,omitempty" toml:"realization,omitempty" json5:"realization,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (e *Essence) UnmarshalJSON(data []byte) error {
	type Alias Essence
	aux := (*Alias)(e)

	return unmarshalWithExtra(data, aux, &e.Extra, "contradiction", "synthesis", "realization")
}

func (e *Essence) MarshalJSON() ([]byte, error) {
	type Alias Essence

	return marshalWithExtra((*Alias)(e), e.Extra)
}
