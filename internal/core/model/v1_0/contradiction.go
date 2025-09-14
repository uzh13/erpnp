package v1_0

// Contradiction represents why we have to create this object
type Contradiction struct {
	Fundamental *Fundamental `json:"fundamental,omitempty" yaml:"fundamental,omitempty" toml:"fundamental,omitempty" json5:"fundamental,omitempty"`
	Space       *Space       `json:"space,omitempty" yaml:"space,omitempty" toml:"space,omitempty" json5:"space,omitempty"`
	Tradeoff    []string     `json:"tradeoff,omitempty" yaml:"tradeoff,omitempty" toml:"tradeoff,omitempty" json5:"tradeoff,omitempty"`
	Resources   []string     `json:"resources,omitempty" yaml:"resources,omitempty" toml:"resources,omitempty" json5:"resources,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (c *Contradiction) UnmarshalJSON(data []byte) error {
	type Alias Contradiction
	aux := (*Alias)(c)

	return unmarshalWithExtra(data, aux, &c.Extra, "fundamental", "space", "tradeoff", "resources")
}

func (c *Contradiction) MarshalJSON() ([]byte, error) {
	type Alias Contradiction

	return marshalWithExtra((*Alias)(c), c.Extra)
}
