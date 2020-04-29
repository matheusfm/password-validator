package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDigits(t *testing.T) {
	validator := New(MinDigits(1))

	assert.NoError(t, validator.IsValid("foo4bar"))
	assert.NoError(t, validator.IsValid("foobar4"))
	assert.NoError(t, validator.IsValid("4foobar"))
	assert.NoError(t, validator.IsValid("4"))
	assert.NoError(t, validator.IsValid("4foo4bar4"))

	assert.Error(t, validator.IsValid("foo.bar"))
	assert.Error(t, validator.IsValid("foobar"))
	assert.Error(t, validator.IsValid(""))
	assert.Error(t, validator.IsValid("!@#$%*()_+="))
}
