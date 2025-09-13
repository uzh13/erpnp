package v1_0

import (
	"encoding/json"
	"time"
)

// ERPN represents the root evolving resolutive process notation object
type ERPN struct {
	Type    TypeInfo       `json:"type" yaml:"type" toml:"type" json5:"type"`
	Content *ContentObject `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty" json5:"content,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// TypeInfo represents the type information
type TypeInfo struct {
	Name    string `json:"name" yaml:"name" toml:"name" json5:"name"`
	Version string `json:"version" yaml:"version" toml:"version" json5:"version"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// ContentObject represents the main object for notation
type ContentObject struct {
	Link    *string          `json:"link,omitempty" yaml:"link,omitempty" toml:"link,omitempty" json5:"link,omitempty"`
	Name    *string          `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	Version *string          `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty" json5:"version,omitempty"`
	Essence *Essence         `json:"essence,omitempty" yaml:"essence,omitempty" toml:"essence,omitempty" json5:"essence,omitempty"`
	Timing  *Timing          `json:"timing,omitempty" yaml:"timing,omitempty" toml:"timing,omitempty" json5:"timing,omitempty"`
	Content []*ContentObject `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty" json5:"content,omitempty"`
	History []*ContentObject `json:"history,omitempty" yaml:"history,omitempty" toml:"history,omitempty" json5:"history,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Essence represents the main description
type Essence struct {
	Contradiction *Contradiction `json:"contradiction,omitempty" yaml:"contradiction,omitempty" toml:"contradiction,omitempty" json5:"contradiction,omitempty"`
	Synthesis     *Synthesis     `json:"synthesis,omitempty" yaml:"synthesis,omitempty" toml:"synthesis,omitempty" json5:"synthesis,omitempty"`
	Realization   *Realization   `json:"realization,omitempty" yaml:"realization,omitempty" toml:"realization,omitempty" json5:"realization,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Contradiction represents why we have to create this object
type Contradiction struct {
	Fundamental *Fundamental `json:"fundamental,omitempty" yaml:"fundamental,omitempty" toml:"fundamental,omitempty" json5:"fundamental,omitempty"`
	Space       *Space       `json:"space,omitempty" yaml:"space,omitempty" toml:"space,omitempty" json5:"space,omitempty"`
	Tradeoff    []string     `json:"tradeoff,omitempty" yaml:"tradeoff,omitempty" toml:"tradeoff,omitempty" json5:"tradeoff,omitempty"`
	Resources   []string     `json:"resources,omitempty" yaml:"resources,omitempty" toml:"resources,omitempty" json5:"resources,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Fundamental represents construction of contradiction
type Fundamental struct {
	Thesis     *string `json:"thesis,omitempty" yaml:"thesis,omitempty" toml:"thesis,omitempty" json5:"thesis,omitempty"`
	Antithesis *string `json:"antithesis,omitempty" yaml:"antithesis,omitempty" toml:"antithesis,omitempty" json5:"antithesis,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Space represents where is the problem and where can we act
type Space struct {
	Space       *string  `json:"space,omitempty" yaml:"space,omitempty" toml:"space,omitempty" json5:"space,omitempty"`
	Limitations []string `json:"limitations,omitempty" yaml:"limitations,omitempty" toml:"limitations,omitempty" json5:"limitations,omitempty"`
	Actors      []string `json:"actors,omitempty" yaml:"actors,omitempty" toml:"actors,omitempty" json5:"actors,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Synthesis represents how to solve main problem
type Synthesis struct {
	Fundamental *string  `json:"fundamental,omitempty" yaml:"fundamental,omitempty" toml:"fundamental,omitempty" json5:"fundamental,omitempty"`
	Resources   []string `json:"resources,omitempty" yaml:"resources,omitempty" toml:"resources,omitempty" json5:"resources,omitempty"`
	Advantages  []string `json:"advantages,omitempty" yaml:"advantages,omitempty" toml:"advantages,omitempty" json5:"advantages,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

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

// CommonObject represents common object for various reality flashes
type CommonObject struct {
	Name     *string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	Type     *string `json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty" json5:"type,omitempty"`
	Describe *string `json:"describe,omitempty" yaml:"describe,omitempty" toml:"describe,omitempty" json5:"describe,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Timing represents time gap of object activity with authors and source of changing
type Timing struct {
	Start         *time.Time `json:"start,omitempty" yaml:"start,omitempty" toml:"start,omitempty" json5:"start,omitempty"`
	End           *time.Time `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty" json5:"end,omitempty"`
	ImpactPercent *float64   `json:"impactPercent,omitempty" yaml:"impactPercent,omitempty" toml:"impactPercent,omitempty" json5:"impactPercent,omitempty"`
	Authors       []*Person  `json:"authors,omitempty" yaml:"authors,omitempty" toml:"authors,omitempty" json5:"authors,omitempty"`
	Source        []string   `json:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty" json5:"source,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// Person represents person description
type Person struct {
	Name          *string  `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" json5:"name,omitempty"`
	ImpactPercent *float64 `json:"impactPercent,omitempty" yaml:"impactPercent,omitempty" toml:"impactPercent,omitempty" json5:"impactPercent,omitempty"`
	Position      *string  `json:"position,omitempty" yaml:"position,omitempty" toml:"position,omitempty" json5:"position,omitempty"`
	Comment       *string  `json:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty" json5:"comment,omitempty"`

	// Extra holds any additional fields not defined in the schema
	Extra map[string]interface{} `json:"-" yaml:"-" toml:"-" json5:"-"`
}

// UnmarshalJSON implements custom JSON unmarshaling to preserve unknown fields
func (e *ERPN) UnmarshalJSON(data []byte) error {
	type Alias ERPN
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	// First unmarshal to get known fields
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Unmarshal to map to get all fields
	var all map[string]interface{}
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	// Remove known fields and store the rest in Extra
	delete(all, "type")
	delete(all, "content")

	if len(all) > 0 {
		e.Extra = all
	}

	return nil
}

// MarshalJSON implements custom JSON marshaling to include unknown fields
func (e *ERPN) MarshalJSON() ([]byte, error) {
	type Alias ERPN
	aux := (*Alias)(e)

	// Marshal known fields first
	known, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}

	if len(e.Extra) == 0 {
		return known, nil
	}

	// Merge with extra fields
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}

	for k, v := range e.Extra {
		knownMap[k] = v
	}

	return json.Marshal(knownMap)
}
