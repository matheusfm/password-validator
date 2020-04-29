package validator

import (
	"fmt"
	"regexp"
)

var regexUppercase = regexp.MustCompile("[A-Z]")

func MinUppercase(min int) Validation {
	return func(pwd string) error {
		if matches := regexUppercase.FindAllStringIndex(pwd, -1); len(matches) < min {
			return fmt.Errorf("must be no less than %v uppercase letter", min)
		}
		return nil
	}
}
