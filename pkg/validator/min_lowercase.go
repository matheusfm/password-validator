package validator

import (
	"fmt"
	"regexp"
)

var regexLowercase = regexp.MustCompile("[a-z]")

func MinLowercase(min int) Validation {
	return func(pwd string) error {
		if matches := regexLowercase.FindAllStringIndex(pwd, -1); len(matches) < min {
			return fmt.Errorf("must be no less than %v lowercase letter", min)
		}
		return nil
	}
}
