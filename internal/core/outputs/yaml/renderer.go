package yaml

import (
	"bytes"
	"fmt"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
	"gopkg.in/yaml.v3"
)

type Renderer struct {
}

func (p *Renderer) Render(source *v1_0.ERPN) ([]byte, error) {
	var buf bytes.Buffer

	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	err := encoder.Encode(source)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal YAML: %w", err)
	}

	return buf.Bytes(), nil
}

func (p *Renderer) SupportedFormats() []string {
	return []string{"yaml", "yml", "x-yaml"}
}
