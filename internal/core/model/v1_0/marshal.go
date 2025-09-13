package v1_0

import (
	"encoding/json"
)

// marshalWithExtra is a helper function to marshal structs with extra fields
func marshalWithExtra(v interface{}, extra map[string]interface{}) ([]byte, error) {
	// Marshal known fields first
	known, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	if len(extra) == 0 {
		return known, nil
	}

	// Merge with extra fields
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}

	for k, v := range extra {
		knownMap[k] = v
	}

	return json.Marshal(knownMap)
}

// unmarshalWithExtra is a helper function to unmarshal structs preserving extra fields
func unmarshalWithExtra(data []byte, v interface{}, extra *map[string]interface{}, knownFields ...string) error {
	// First unmarshal to get known fields
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	// Unmarshal to map to get all fields
	var all map[string]interface{}
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	// Remove known fields
	for _, field := range knownFields {
		delete(all, field)
	}

	if len(all) > 0 {
		*extra = all
	}

	return nil
}

// TypeInfo marshal/unmarshal methods
func (t *TypeInfo) UnmarshalJSON(data []byte) error {
	type Alias TypeInfo
	aux := (*Alias)(t)
	return unmarshalWithExtra(data, aux, &t.Extra, "name", "version")
}

func (t *TypeInfo) MarshalJSON() ([]byte, error) {
	type Alias TypeInfo
	return marshalWithExtra((*Alias)(t), t.Extra)
}

// ContentObject marshal/unmarshal methods
func (c *ContentObject) UnmarshalJSON(data []byte) error {
	type Alias ContentObject
	aux := (*Alias)(c)
	return unmarshalWithExtra(data, aux, &c.Extra, "link", "name", "version", "essence", "timing", "content", "history")
}

func (c *ContentObject) MarshalJSON() ([]byte, error) {
	type Alias ContentObject
	return marshalWithExtra((*Alias)(c), c.Extra)
}

// Essence marshal/unmarshal methods
func (e *Essence) UnmarshalJSON(data []byte) error {
	type Alias Essence
	aux := (*Alias)(e)
	return unmarshalWithExtra(data, aux, &e.Extra, "contradiction", "synthesis", "realization")
}

func (e *Essence) MarshalJSON() ([]byte, error) {
	type Alias Essence
	return marshalWithExtra((*Alias)(e), e.Extra)
}

// Contradiction marshal/unmarshal methods
func (c *Contradiction) UnmarshalJSON(data []byte) error {
	type Alias Contradiction
	aux := (*Alias)(c)
	return unmarshalWithExtra(data, aux, &c.Extra, "fundamental", "space", "tradeoff", "resources")
}

func (c *Contradiction) MarshalJSON() ([]byte, error) {
	type Alias Contradiction
	return marshalWithExtra((*Alias)(c), c.Extra)
}

// Fundamental marshal/unmarshal methods
func (f *Fundamental) UnmarshalJSON(data []byte) error {
	type Alias Fundamental
	aux := (*Alias)(f)
	return unmarshalWithExtra(data, aux, &f.Extra, "thesis", "antithesis")
}

func (f *Fundamental) MarshalJSON() ([]byte, error) {
	type Alias Fundamental
	return marshalWithExtra((*Alias)(f), f.Extra)
}

// Space marshal/unmarshal methods
func (s *Space) UnmarshalJSON(data []byte) error {
	type Alias Space
	aux := (*Alias)(s)
	return unmarshalWithExtra(data, aux, &s.Extra, "space", "limitations", "actors")
}

func (s *Space) MarshalJSON() ([]byte, error) {
	type Alias Space
	return marshalWithExtra((*Alias)(s), s.Extra)
}

// Synthesis marshal/unmarshal methods
func (s *Synthesis) UnmarshalJSON(data []byte) error {
	type Alias Synthesis
	aux := (*Alias)(s)
	return unmarshalWithExtra(data, aux, &s.Extra, "fundamental", "resources", "advantages")
}

func (s *Synthesis) MarshalJSON() ([]byte, error) {
	type Alias Synthesis
	return marshalWithExtra((*Alias)(s), s.Extra)
}

// Realization marshal/unmarshal methods
func (r *Realization) UnmarshalJSON(data []byte) error {
	type Alias Realization
	aux := (*Alias)(r)
	return unmarshalWithExtra(data, aux, &r.Extra, "input", "output", "resources", "value", "commonDescription")
}

func (r *Realization) MarshalJSON() ([]byte, error) {
	type Alias Realization
	return marshalWithExtra((*Alias)(r), r.Extra)
}

// CommonObject marshal/unmarshal methods
func (c *CommonObject) UnmarshalJSON(data []byte) error {
	type Alias CommonObject
	aux := (*Alias)(c)
	return unmarshalWithExtra(data, aux, &c.Extra, "name", "type", "describe")
}

func (c *CommonObject) MarshalJSON() ([]byte, error) {
	type Alias CommonObject
	return marshalWithExtra((*Alias)(c), c.Extra)
}

// Timing marshal/unmarshal methods
func (t *Timing) UnmarshalJSON(data []byte) error {
	type Alias Timing
	aux := (*Alias)(t)
	return unmarshalWithExtra(data, aux, &t.Extra, "start", "end", "impactPercent", "authors", "source")
}

func (t *Timing) MarshalJSON() ([]byte, error) {
	type Alias Timing
	return marshalWithExtra((*Alias)(t), t.Extra)
}

// Person marshal/unmarshal methods
func (p *Person) UnmarshalJSON(data []byte) error {
	type Alias Person
	aux := (*Alias)(p)
	return unmarshalWithExtra(data, aux, &p.Extra, "name", "impactPercent", "position", "comment")
}

func (p *Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return marshalWithExtra((*Alias)(p), p.Extra)
}
