package validator

import (
	"errors"
)

type Validation func(pwd string) error

type Validator struct {
	validations []Validation
}

func New(validations ...Validation) *Validator {
	return &Validator{validations: validations}
}

func (r *Validator) IsValid(pwd string) error {
	if len(r.validations) == 0 {
		return errors.New("the validator must be no less than 1 validation function")
	}
	for _, validation := range r.validations {
		if err := validation(pwd); err != nil {
			return err
		}
	}
	return nil
}
