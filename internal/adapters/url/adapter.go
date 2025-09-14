package url

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/samber/lo"
	"github.com/uzh13/erpnp/internal/core/inputs"
	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

type UrlAdapter struct {
	loader inputs.Loader
}

func New(loader inputs.Loader) *UrlAdapter {
	return &UrlAdapter{
		loader: loader,
	}
}

func (s *UrlAdapter) Parse(link string) (*v1_0.ERPN, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, errors.New("failed to get data")
	}
	defer resp.Body.Close()

	var ext string
	ct := resp.Header.Get("Content-Type")
	if ct == "" {
		ext = strings.ToLower(path.Ext(link))
	} else {
		mime := strings.ToLower(strings.Split(ct, ";")[0])
		ext = strings.TrimPrefix(mime, "application/")
		ext = strings.TrimPrefix(mime, "text/")
	}

	format := s.loader.DetectTypeByExtention(ext)
	if format == inputs.FileTypeUnknown {
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}

	source, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close link: %w", err)
	}

	result, err := s.loader.Parse(source, format)
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
