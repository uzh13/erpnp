package std

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/uzh13/erpnp/internal/core/inputs"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
	"github.com/uzh13/erpnp/internal/core/outputs"
)

type StdAdapter struct {
	loader   inputs.Loader
	exporter outputs.Exporter
}

func New(loader inputs.Loader) *StdAdapter {
	return &StdAdapter{
		loader: loader,
	}
}

func (s *StdAdapter) Parse(source, ext string) (*v1_0.ERPN, error) {
	format := s.loader.DetectTypeByExtention(ext)
	if format == inputs.FileTypeUnknown {
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}

	result, err := s.loader.Parse([]byte(source), format)
	if err != nil {
		return nil, fmt.Errorf("failed to parse text: %w", err)
	}

	errs := result.Validate()
	if len(errs) > 0 {
		return nil, fmt.Errorf(
			"failed to validate data: %s",
			strings.Join(lo.Map(errs, func(err error, _ int) string { return fmt.Sprintf("\n%s", err.Error()) }), ""),
		)
	}

	return result, nil
}

func (s *StdAdapter) Show(source *v1_0.ERPN, ext string) error {
	result, err := s.exporter.Export(source, ext)
	if err != nil {
		return fmt.Errorf("failed to export data: %w", err)
	}

	fmt.Print(string(result))

	return nil
}
