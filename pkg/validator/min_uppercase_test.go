package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinUppercase(t *testing.T) {
	validator := New(MinUppercase(1))

	assert.NoError(t, validator.IsValid("Foobar"))
	assert.NoError(t, validator.IsValid("fooBar"))
	assert.NoError(t, validator.IsValid("foobaR"))
	assert.NoError(t, validator.IsValid("FooBar"))
	assert.NoError(t, validator.IsValid("foo4baR"))
	assert.NoError(t, validator.IsValid("1234567890A"))
	assert.NoError(t, validator.IsValid("A1234567890"))
	assert.NoError(t, validator.IsValid("!@#$%*()_+=A"))
	assert.NoError(t, validator.IsValid("A!@#$%*()_+="))

	assert.Error(t, validator.IsValid("1234567890"))
	assert.Error(t, validator.IsValid("foobar"))
	assert.Error(t, validator.IsValid("!@#$%*()_+="))
	assert.Error(t, validator.IsValid(""))
}
