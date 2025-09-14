package inputs

import (
	"fmt"
	"io"
	"strings"

	"github.com/uzh13/erpnp/internal/core/inputs/json"
	"github.com/uzh13/erpnp/internal/core/inputs/json5"
	"github.com/uzh13/erpnp/internal/core/inputs/toml"
	"github.com/uzh13/erpnp/internal/core/inputs/yaml"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type FileType int

const (
	FileTypeUnknown FileType = iota
	FileTypeJSON
	FileTypeYAML
	FileTypeTOML
	FileTypeJSON5
)

type Parser interface {
	Parse(source []byte) (*v1_0.ERPN, error)
	SupportedFormats() []string
}

type Loader struct {
	Parsers map[FileType]Parser
}

func NewLoader() *Loader {
	return &Loader{
		Parsers: map[FileType]Parser{
			FileTypeJSON:  new(json.Parser),
			FileTypeJSON5: new(json5.Parser),
			FileTypeYAML:  new(yaml.Parser),
			FileTypeTOML:  new(toml.Parser),
		},
	}
}

func (l *Loader) SupportedFormats() []string {
	result := make([]string, 0, len(l.Parsers))
	for _, parser := range l.Parsers {
		result = append(result, parser.SupportedFormats()...)
	}

	return result
}

func (l *Loader) DetectTypeByExtention(ext string) FileType {
	ext = strings.ToLower(ext)

	for fileType, parser := range l.Parsers {
		for _, format := range parser.SupportedFormats() {
			if format != ext {
				continue
			}

			return fileType
		}
	}

	return FileTypeUnknown
}

func (l *Loader) Parse(data []byte, format FileType) (*v1_0.ERPN, error) {
	parser, ok := l.Parsers[format]
	if !ok {
		return nil, fmt.Errorf("unsupported format: %d", format)
	}

	result, err := parser.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse data: %w", err)
	}

	return result, nil
}

func (l *Loader) ParseFromReader(r io.Reader, format FileType) (*v1_0.ERPN, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	return l.Parse(data, format)
}
