package json5

import (
	"encoding/json"
	"fmt"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type Parser struct {
}

func (p *Parser) Parse(source []byte) (*v1_0.ERPN, error) {
	var erpn v1_0.ERPN

	err := json.Unmarshal(source, &erpn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON5: %w", err)
	}

	return &erpn, nil
}

func (p *Parser) SupportedFormats() []string {
	return []string{"json5"}
}
