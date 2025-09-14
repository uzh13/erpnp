package v1_0

// Space represents where is the problem and where can we act
type Space struct {
	Space       *string  `json:"space,omitempty" yaml:"space,omitempty" toml:"space,omitempty" json5:"space,omitempty"`
	Limitations []string `json:"limitations,omitempty" yaml:"limitations,omitempty" toml:"limitations,omitempty" json5:"limitations,omitempty"`
	Actors      []string `json:"actors,omitempty" yaml:"actors,omitempty" toml:"actors,omitempty" json5:"actors,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (s *Space) UnmarshalJSON(data []byte) error {
	type Alias Space
	aux := (*Alias)(s)
	return unmarshalWithExtra(data, aux, &s.Extra, "space", "limitations", "actors")
}

func (s *Space) MarshalJSON() ([]byte, error) {
	type Alias Space
	return marshalWithExtra((*Alias)(s), s.Extra)
}
