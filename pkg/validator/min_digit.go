package validator

import (
	"fmt"
	"regexp"
)

var regexDigit = regexp.MustCompile("\\d")

func MinDigits(min int) Validation {
	return func(pwd string) error {
		if matches := regexDigit.FindAllStringIndex(pwd, -1); len(matches) < min {
			return fmt.Errorf("must be no less than %v digits", min)
		}
		return nil
	}
}
