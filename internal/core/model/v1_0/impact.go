package v1_0

import (
	"fmt"
)

type Impact float64

func (s Impact) Validate() []error {
	err := []error{fmt.Errorf("must be between 0 and 100: %f", s)}

	if s < 0 || s > 100 {
		return err
	}

	return nil
}
