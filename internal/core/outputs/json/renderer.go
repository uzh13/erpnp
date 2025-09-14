package json

import (
	"encoding/json"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type Renderer struct {
}

func (p *Renderer) Render(source *v1_0.ERPN) ([]byte, error) {
	return json.MarshalIndent(source, "", "  ")
}

func (p *Renderer) SupportedFormats() []string {
	return []string{"json"}
}
