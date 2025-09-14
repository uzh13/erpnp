package v1_0

import "encoding/json"

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
