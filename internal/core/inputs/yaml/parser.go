package yaml

import (
	"fmt"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
	"gopkg.in/yaml.v3"
)

type Parser struct {
}

func (p *Parser) Parse(source []byte) (*v1_0.ERPN, error) {
	var erpn v1_0.ERPN

	err := yaml.Unmarshal(source, &erpn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &erpn, nil
}

func (p *Parser) SupportedFormats() []string {
	return []string{"yaml", "yml", "x-yaml"}
}
