package v1_0

import (
	"fmt"
	"testing"
)

func ExampleERPN_marshaling() {
	// Create a sample ERPN with extra fields
	erpn := &ERPN{
		Type: TypeInfo{
			Name:    "Test Process",
			Version: "1.0.0",
			Extra: map[string]interface{}{
				"customTypeField": "custom value",
			},
		},
		Content: &ContentObject{
			Name:    stringPtr("Test Content"),
			Version: stringPtr("1.0.0"),
			Extra: map[string]interface{}{
				"customContentField": 42,
			},
		},
		Extra: map[string]interface{}{
			"customRootField": []string{"value1", "value2"},
		},
	}

	// Convert to JSON
	jsonStr, err := erpn.ToString("json")
	if err != nil {
		fmt.Printf("JSON error: %v\n", err)
		return
	}
	fmt.Println("JSON output:")
	fmt.Println(jsonStr)

	// Convert to YAML
	yamlStr, err := erpn.ToString("yaml")
	if err != nil {
		fmt.Printf("YAML error: %v\n", err)
		return
	}
	fmt.Println("\nYAML output:")
	fmt.Println(yamlStr)

	// Parse back from JSON and verify extra fields are preserved
	parsedERPN, err := ParseFromBytes([]byte(jsonStr), "json")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		return
	}

	fmt.Println("\nParsed back - extra fields preserved:")
	fmt.Printf("Root extra: %v\n", parsedERPN.Extra)
	fmt.Printf("Type extra: %v\n", parsedERPN.Type.Extra)
	fmt.Printf("Content extra: %v\n", parsedERPN.Content.Extra)
}

func TestRoundTripMarshaling(t *testing.T) {
	formats := []string{"json", "yaml", "toml"}

	original := &ERPN{
		Type: TypeInfo{
			Name:    "Test Process",
			Version: "1.0.0",
			Extra: map[string]interface{}{
				"customField": "value",
			},
		},
		Extra: map[string]interface{}{
			"rootExtra": 123,
		},
	}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			// Marshal to format
			data, err := original.ToBytes(format)
			if err != nil {
				t.Fatalf("Failed to marshal to %s: %v", format, err)
			}

			// Parse back
			parsed, err := ParseFromBytes(data, format)
			if err != nil {
				t.Fatalf("Failed to parse %s: %v", format, err)
			}

			// Verify basic fields
			if parsed.Type.Name != original.Type.Name {
				t.Errorf("Name mismatch: got %s, want %s", parsed.Type.Name, original.Type.Name)
			}

			// Verify extra fields are preserved (for JSON format)
			if format == "json" {
				if parsed.Extra["rootExtra"] == nil {
					t.Error("Root extra field not preserved")
				}
				if parsed.Type.Extra["customField"] == nil {
					t.Error("Type extra field not preserved")
				}
			}
		})
	}
}

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"test.json", "json"},
		{"test.json5", "json5"},
		{"test.yaml", "yaml"},
		{"test.yml", "yaml"},
		{"test.toml", "toml"},
		{"test.txt", ""},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := DetectFormat(tt.filename)
			if result != tt.expected {
				t.Errorf("DetectFormat(%s) = %s, want %s", tt.filename, result, tt.expected)
			}
		})
	}
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}

// Example showing how to work with the existing blog example
func ExampleWorkingWithExistingFile() {
	// Sample JSON that includes extra fields
	jsonWithExtra := `{
		"type": {
			"name": "Test Process",
			"version": "1.0.0",
			"customTypeField": "should be preserved"
		},
		"content": {
			"name": "Test Content",
			"version": "1.0.0",
			"customContentField": 42
		},
		"customRootField": ["value1", "value2"]
	}`

	// Parse JSON
	erpn, err := ParseFromBytes([]byte(jsonWithExtra), "json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed ERPN with preserved extra fields:\n")
	fmt.Printf("Root extra: %v\n", erpn.Extra)
	fmt.Printf("Type extra: %v\n", erpn.Type.Extra)
	fmt.Printf("Content extra: %v\n", erpn.Content.Extra)

	// Convert to different format while preserving extra fields
	yamlOutput, err := erpn.ToString("yaml")
	if err != nil {
		fmt.Printf("Error converting to YAML: %v\n", err)
		return
	}

	fmt.Println("\nConverted to YAML (extra fields preserved):")
	fmt.Println(yamlOutput)
}
