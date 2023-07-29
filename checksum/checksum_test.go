package checksum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChecksumShouldBeCalculatedAndValid(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	testCases := []struct {
		name     string
		digits   [13]uint
		expected uint
	}{
		{
			name:     "1 <= m <= 9, K = m",
			digits:   [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 1, 8},
			expected: 8,
		},
		{
			name:     "m = 10, K = 0",
			digits:   [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "m = 11, K = 0",
			digits:   [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 3, 5, 0, 3, 0},
			expected: 0,
		},
	}

	// act
	for _, testCase := range testCases {
		t.Log(testCase.name)

		csum := New(testCase.digits)

		// assert
		assert.True(csum.Valid)
		assert.Equal(testCase.expected, csum.Value)
	}
}

func TestChecksumShouldBeCalculatedAndInvalid(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	digits := [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 3, 5, 0, 3, 1}
	var expected uint = 1

	// act
	csum := New(digits)

    // assert
	assert.False(csum.Valid)
	assert.Equal(expected, csum.Value)
}
