package toml

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type Renderer struct {
}

func (p *Renderer) Render(source *v1_0.ERPN) ([]byte, error) {
	var buf bytes.Buffer

	encoder := toml.NewEncoder(&buf)

	err := encoder.Encode(source)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TOML: %w", err)
	}

	return buf.Bytes(), nil
}

func (p *Renderer) SupportedFormats() []string {
	return []string{"toml"}
}
