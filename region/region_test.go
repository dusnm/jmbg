package region

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldFailWhenNoRegionFoundForCode(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	digits := [13]uint{0, 1, 0, 1, 0, 0, 0, 2, 0, 5, 0, 1, 8}
	var code uint = 20

	// act
	region, err := New(digits)

	// assert
	assert.Error(err)
	assert.ErrorIs(err, InvalidCodeError{input: code})
	assert.Equal(region, Region{})
}

func TestShouldPassWhenRegionFoundForCode(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	digits := [13]uint{0, 1, 0, 1, 0, 0, 0, 7, 1, 5, 0, 1, 8}
	code := BGD

	// act
	region, err := New(digits)

	// assert
	assert.NoError(err)
	assert.Equal(
		region,
		Region{
			Code: code,
			Name: Name{
				Cyrilic: "Београд",
				Latin:   "Beograd",
			},
		},
	)
}
