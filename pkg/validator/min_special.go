package validator

import (
	"fmt"
	"regexp"
)

//https://owasp.org/www-community/password-special-characters
var regexSpecialChar = regexp.MustCompile("[!\"#$%&'()*+,\\-./:;<=>?@\\[\\]^_`{|}~]")

func MinSpecialChar(min int) Validation {
	return func(pwd string) error {
		if matches := regexSpecialChar.FindAllStringIndex(pwd, -1); len(matches) < min {
			return fmt.Errorf("must be no less than %v special characters", min)
		}
		return nil
	}
}
