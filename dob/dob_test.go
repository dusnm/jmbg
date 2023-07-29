package dob

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestShouldParseDate(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	testCases := []struct {
		name     string
		digits   [13]uint
		expected time.Time
	}{
		{
			name:   "the first year digit is 0",
			digits: [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 1, 8},
			expected: func() time.Time {
				value, _ := time.Parse("02-01-2006", "01-01-2000")

				return value
			}(),
		},
		{
			name:   "the first year digit is 9",
			digits: [13]uint{0, 1, 0, 1, 9, 9, 7, 8, 0, 5, 0, 1, 2},
			expected: func() time.Time {
				value, _ := time.Parse("02-01-2006", "01-01-1997")

				return value
			}(),
		},
	}

	// act
	for _, testCase := range testCases {
		t.Log(testCase.name)

		date, err := New(testCase.digits)

		// assert
		assert.NoError(err)
		assert.Equal(testCase.expected, date)
	}
}

func TestShouldFailWithInvalidYearDigitError(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	digits := [13]uint{0, 1, 0, 1, 4, 9, 7, 8, 0, 5, 0, 1, 2}

	// act
	date, err := New(digits)

	// assert
	assert.Error(err)
	assert.ErrorIs(err, InvalidYearDigitError{})
	assert.Equal(time.Time{}, date)
}
