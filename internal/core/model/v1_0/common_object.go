package v1_0

// CommonObject represents a common object for various reality flashes
type CommonObject struct {
	Name     *string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	Type     *string `json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty" json5:"type,omitempty"`
	Describe *string `json:"describe,omitempty" yaml:"describe,omitempty" toml:"describe,omitempty" json5:"describe,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (c *CommonObject) UnmarshalJSON(data []byte) error {
	type Alias CommonObject
	aux := (*Alias)(c)
	return unmarshalWithExtra(data, aux, &c.Extra, "name", "type", "describe")
}

func (c *CommonObject) MarshalJSON() ([]byte, error) {
	type Alias CommonObject
	return marshalWithExtra((*Alias)(c), c.Extra)
}
