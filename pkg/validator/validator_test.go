package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid_When_ThereIsNoValidation(t *testing.T) {
	validator := New()
	assert.Error(t, validator.IsValid("foobar"))
}
