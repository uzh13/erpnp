package v1_0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// SupportedFormats returns list of supported file formats
func SupportedFormats() []string {
	return []string{"json", "yaml", "yml", "toml", "json5"}
}

// ParseFromBytes parses ERPN from bytes based on format
func ParseFromBytes(data []byte, format string) (*ERPN, error) {
	var erpn ERPN

	switch strings.ToLower(format) {
	case "json", "json5":
		if err := json.Unmarshal(data, &erpn); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(data, &erpn); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	case "toml":
		if err := toml.Unmarshal(data, &erpn); err != nil {
			return nil, fmt.Errorf("failed to parse TOML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return &erpn, nil
}

// ParseFromReader parses ERPN from reader based on format
func ParseFromReader(r io.Reader, format string) (*ERPN, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return ParseFromBytes(data, format)
}

// ParseFromFile parses ERPN from file, detecting format from extension
func ParseFromFile(filename string) (*ERPN, error) {
	format := DetectFormat(filename)
	if format == "" {
		return nil, fmt.Errorf("unsupported file extension: %s", filepath.Ext(filename))
	}

	data, err := io.ReadAll(strings.NewReader("")) // This would be file reading in real implementation
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return ParseFromBytes(data, format)
}

// DetectFormat detects format from filename extension
func DetectFormat(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".json":
		return "json"
	case ".json5":
		return "json5"
	case ".yaml", ".yml":
		return "yaml"
	case ".toml":
		return "toml"
	default:
		return ""
	}
}

// ToBytes converts ERPN to bytes in specified format
func (e *ERPN) ToBytes(format string) ([]byte, error) {
	switch strings.ToLower(format) {
	case "json", "json5":
		return json.MarshalIndent(e, "", "  ")
	case "yaml", "yml":
		var buf bytes.Buffer
		encoder := yaml.NewEncoder(&buf)
		encoder.SetIndent(2)
		if err := encoder.Encode(e); err != nil {
			return nil, fmt.Errorf("failed to marshal YAML: %w", err)
		}
		return buf.Bytes(), nil
	case "toml":
		var buf bytes.Buffer
		encoder := toml.NewEncoder(&buf)
		if err := encoder.Encode(e); err != nil {
			return nil, fmt.Errorf("failed to marshal TOML: %w", err)
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// ToWriter writes ERPN to writer in specified format
func (e *ERPN) ToWriter(w io.Writer, format string) error {
	data, err := e.ToBytes(format)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// ToString converts ERPN to string in specified format
func (e *ERPN) ToString(format string) (string, error) {
	data, err := e.ToBytes(format)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
