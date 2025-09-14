package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/uzh13/erpnp/internal/core/inputs"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type FileAdapter struct {
	loader inputs.Loader
}

func New(loader inputs.Loader) *FileAdapter {
	return &FileAdapter{
		loader: loader,
	}
}

func (f *FileAdapter) Parse(filename string) (*v1_0.ERPN, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	format := f.loader.DetectTypeByExtention(ext)
	if format == inputs.FileTypeUnknown {
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	result, err := f.loader.Parse(bytes, format)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
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

func (f *FileAdapter) Save(filename string, source []byte) (int, error) {
	file, err := os.Create(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	size, err := file.Write(source)
	if err != nil {
		return 0, fmt.Errorf("failed to write data: %w", err)
	}

	return size, nil
}
