package v1_0

// Fundamental represents construction of contradiction
type Fundamental struct {
	Thesis     *string `json:"thesis,omitempty" yaml:"thesis,omitempty" toml:"thesis,omitempty" json5:"thesis,omitempty"`
	Antithesis *string `json:"antithesis,omitempty" yaml:"antithesis,omitempty" toml:"antithesis,omitempty" json5:"antithesis,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

func (f *Fundamental) UnmarshalJSON(data []byte) error {
	type Alias Fundamental
	aux := (*Alias)(f)
	return unmarshalWithExtra(data, aux, &f.Extra, "thesis", "antithesis")
}

func (f *Fundamental) MarshalJSON() ([]byte, error) {
	type Alias Fundamental
	return marshalWithExtra((*Alias)(f), f.Extra)
}
