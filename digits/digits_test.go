package digits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldConvertJmbgStringToDigits(t *testing.T) {
	// arrange
	assert := require.New(t)
	jmbg := "0101000805018"
	expected := [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 1, 8}

	// act
	actual, err := New(jmbg)

	// assert
	assert.NoError(err)
	assert.Equal(expected, actual)
}

func TestShouldFailWithInvalidLength(t *testing.T) {
	// arrange
	assert := require.New(t)
	jmbg := "010100080501"

	// act
	actual, err := New(jmbg)

	// assert
	assert.Error(err)
	assert.ErrorIs(err, InvalidLengthError{})
	assert.Equal([13]uint{}, actual)
}

func TestShouldFailWithInvalidDigit(t *testing.T) {
	// arrange
	assert := require.New(t)
	jmbg := "01010008界0501"

	// act
	actual, err := New(jmbg)

	// assert
	assert.Error(err)
	assert.ErrorIs(err, InvalidDigitError{input: '界'})
	assert.Equal([13]uint{}, actual)
}
