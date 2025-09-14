package toml

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type Parser struct {
}

func (p *Parser) Parse(source []byte) (*v1_0.ERPN, error) {
	var erpn v1_0.ERPN

	err := toml.Unmarshal(source, &erpn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TOML: %w", err)
	}

	return &erpn, nil
}

func (p *Parser) SupportedFormats() []string {
	return []string{"toml"}
}
