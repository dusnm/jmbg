package sex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldParseSex(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	testCases := []struct {
		name     string
		digits   [13]uint
		expected Sex
	}{
		{
			name:   "number indicates a person is male",
			digits: [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 4, 0, 1, 1},
			expected: Sex{
				Code: MALE,
				Name: Name{
					Cyrillic: "мушки",
					Latin:    "muški",
				},
			},
		},
		{
			name:   "number indicates a person is female",
			digits: [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 1, 8},
			expected: Sex{
				Code: FEMALE,
				Name: Name{
					Cyrillic: "женски",
					Latin:    "ženski",
				},
			},
		},
	}

	// act
	for _, testCase := range testCases {
		t.Log(testCase.name)

		sex := New(testCase.digits)

		// assert
		assert.Equal(testCase.expected, sex)
	}
}
