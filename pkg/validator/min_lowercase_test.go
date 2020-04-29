package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLowercase(t *testing.T) {
	validator := New(MinLowercase(1))

	assert.NoError(t, validator.IsValid("foobar"))
	assert.NoError(t, validator.IsValid("fooBAR"))
	assert.NoError(t, validator.IsValid("FOObar"))
	assert.NoError(t, validator.IsValid("FOObAR"))
	assert.NoError(t, validator.IsValid("FOOb4R"))
	assert.NoError(t, validator.IsValid("1234567890a"))
	assert.NoError(t, validator.IsValid("a1234567890"))
	assert.NoError(t, validator.IsValid("!@#$%*()_+=a"))
	assert.NoError(t, validator.IsValid("aA!@#$%*()_+="))

	assert.Error(t, validator.IsValid("1234567890"))
	assert.Error(t, validator.IsValid("FOOBAR"))
	assert.Error(t, validator.IsValid("!@#$%*()_+="))
	assert.Error(t, validator.IsValid(""))
}
