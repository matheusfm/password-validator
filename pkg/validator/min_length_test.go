package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLength(t *testing.T) {
	validator := New(MinLength(9))

	assert.NoError(t, validator.IsValid("123456789"))
	assert.NoError(t, validator.IsValid("0123456789"))
	assert.NoError(t, validator.IsValid("!@#$%&*(+"))
	assert.NoError(t, validator.IsValid("abcdefghi"))

	assert.Error(t, validator.IsValid("12345678"))
	assert.Error(t, validator.IsValid("!@#$%&*("))
	assert.Error(t, validator.IsValid("abcdefgh"))
	assert.Error(t, validator.IsValid(""))
}
