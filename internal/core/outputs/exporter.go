package outputs

import (
	"fmt"
	"strings"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
	"github.com/uzh13/erpnp/internal/core/outputs/json"
	"github.com/uzh13/erpnp/internal/core/outputs/json5"
	"github.com/uzh13/erpnp/internal/core/outputs/toml"
	"github.com/uzh13/erpnp/internal/core/outputs/yaml"
)

type FileType int

const (
	FileTypeUnknown FileType = iota
	FileTypeJSON
	FileTypeYAML
	FileTypeTOML
	FileTypeJSON5
)

type Renderer interface {
	Render(source *v1_0.ERPN) ([]byte, error)
	SupportedFormats() []string
}

type Exporter struct {
	renderers map[FileType]Renderer
}

func NewExporter() *Exporter {
	return &Exporter{
		renderers: map[FileType]Renderer{
			FileTypeJSON:  new(json.Renderer),
			FileTypeJSON5: new(json5.Renderer),
			FileTypeYAML:  new(yaml.Renderer),
			FileTypeTOML:  new(toml.Renderer),
		},
	}
}

func (l *Exporter) SupportedFormats() []string {
	result := make([]string, 0, len(l.renderers))
	for _, parser := range l.renderers {
		result = append(result, parser.SupportedFormats()...)
	}

	return result
}

func (l *Exporter) DetectTypeByExtension(ext string) FileType {
	ext = strings.ToLower(ext)

	for fileType, parser := range l.renderers {
		for _, format := range parser.SupportedFormats() {
			if format != ext {
				continue
			}

			return fileType
		}
	}

	return FileTypeUnknown
}

func (l *Exporter) Export(source *v1_0.ERPN, ext string) ([]byte, error) {
	fileType := l.DetectTypeByExtension(ext)
	if fileType == FileTypeUnknown {
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	renderer, ok := l.renderers[fileType]
	if !ok {
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}

	result, err := renderer.Render(source)
	if err != nil {
		return nil, fmt.Errorf("failed to render data: %w", err)
	}

	return result, nil
}
