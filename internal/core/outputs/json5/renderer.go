package json5

import (
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type Renderer struct {
}

func (p *Renderer) Render(source *v1_0.ERPN) ([]byte, error) {
	return json5.MarshalIndent(source, "", "  ")
}

func (p *Renderer) SupportedFormats() []string {
	return []string{"json5"}
}
