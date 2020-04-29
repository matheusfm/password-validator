package validator

import (
	"fmt"
)

func MinLength(min int) Validation {
	return func(pwd string) error {
		if len(pwd) < min {
			return fmt.Errorf("must be no less than %v characters", min)
		}
		return nil
	}
}
