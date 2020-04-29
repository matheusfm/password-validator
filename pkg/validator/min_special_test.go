package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSpecialChar(t *testing.T) {
	validator := New(MinSpecialChar(1))

	assert.NoError(t, validator.IsValid("foob@r"))
	assert.NoError(t, validator.IsValid("foobar!"))
	assert.NoError(t, validator.IsValid("#foobar"))
	assert.NoError(t, validator.IsValid("FOOBAR%"))
	assert.NoError(t, validator.IsValid("!\"#$%&'()*+,-./:;<=>?@[]^_`{|}~"))

	assert.Error(t, validator.IsValid(""))
	assert.Error(t, validator.IsValid("foo bar"))
	assert.Error(t, validator.IsValid("ç"))
	assert.Error(t, validator.IsValid("foO0báàãâär"))
}
