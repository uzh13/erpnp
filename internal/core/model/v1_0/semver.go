package v1_0

import (
	"fmt"
)

type SemVer string

func (s SemVer) Validate() []error {
	// check empty must be on the upper layer
	if s == "" {
		return nil
	}

	failResult := []error{fmt.Errorf("semver is not valid: %s", s)}

	parts := make([]string, 0, 3)
	current := ""

	for _, r := range s {
		if r == '.' {
			if current == "" {
				return failResult
			}
			parts = append(parts, current)
			current = ""
			continue
		}

		if r >= '0' && r <= '9' {
			current += string(r)
			continue
		}

		return failResult
	}

	if current != "" {
		parts = append(parts, current)
	}

	if len(parts) != 3 {
		return failResult
	}

	return nil
}
