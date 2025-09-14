package v1_0

// Realization represents how main object was solved
type Realization struct {
	Input             []*CommonObject `json:"input,omitempty" yaml:"input,omitempty" toml:"input,omitempty" json5:"input,omitempty"`
	Output            []*CommonObject `json:"output,omitempty" yaml:"output,omitempty" toml:"output,omitempty" json5:"output,omitempty"`
	Resources         []*CommonObject `json:"resources,omitempty" yaml:"resources,omitempty" toml:"resources,omitempty" json5:"resources,omitempty"`
	Value             []*CommonObject `json:"value,omitempty" yaml:"value,omitempty" toml:"value,omitempty" json5:"value,omitempty"`
	CommonDescription *string         `json:"commonDescription,omitempty" yaml:"commonDescription,omitempty" toml:"commonDescription,omitempty" json5:"commonDescription,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (r *Realization) UnmarshalJSON(data []byte) error {
	type Alias Realization
	aux := (*Alias)(r)
	return unmarshalWithExtra(data, aux, &r.Extra, "input", "output", "resources", "value", "commonDescription")
}

func (r *Realization) MarshalJSON() ([]byte, error) {
	type Alias Realization
	return marshalWithExtra((*Alias)(r), r.Extra)
}
