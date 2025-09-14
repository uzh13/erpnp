package v1_0

// Synthesis represents how to solve main problem
type Synthesis struct {
	Fundamental *string  `json:"fundamental,omitempty" yaml:"fundamental,omitempty" toml:"fundamental,omitempty" json5:"fundamental,omitempty"`
	Resources   []string `json:"resources,omitempty" yaml:"resources,omitempty" toml:"resources,omitempty" json5:"resources,omitempty"`
	Advantages  []string `json:"advantages,omitempty" yaml:"advantages,omitempty" toml:"advantages,omitempty" json5:"advantages,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (s *Synthesis) UnmarshalJSON(data []byte) error {
	type Alias Synthesis
	aux := (*Alias)(s)
	return unmarshalWithExtra(data, aux, &s.Extra, "fundamental", "resources", "advantages")
}

func (s *Synthesis) MarshalJSON() ([]byte, error) {
	type Alias Synthesis
	return marshalWithExtra((*Alias)(s), s.Extra)
}
