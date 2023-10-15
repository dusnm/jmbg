package jmbg

import (
	"testing"
	"time"

	"github.com/dusnm/jmbg/checksum"
	"github.com/dusnm/jmbg/region"
	"github.com/dusnm/jmbg/sex"
	"github.com/stretchr/testify/require"
)

func TestShouldParseJmbg(t *testing.T) {
	t.Parallel()

	// arrange
	assert := require.New(t)
	number := "0101000805018"
	expected := JMBG{
		Number: number,
		Digits: [13]uint{0, 1, 0, 1, 0, 0, 0, 8, 0, 5, 0, 1, 8},
		DateOfBirth: func() time.Time {
			value, _ := time.Parse(
				"01-02-2006",
				"01-01-2000",
			)

			return value
		}(),
		Region: region.Region{
			Code: region.NSD,
			Name: region.Name{
				Cyrillic: "Нови Сад",
				Latin:    "Novi Sad",
			},
		},
		Sex: sex.Sex{
			Code: sex.FEMALE,
			Name: sex.Name{
				Cyrillic: "женски",
				Latin:    "ženski",
			},
		},
		Checksum: checksum.Checksum{
			Value: 8,
			Valid: true,
		},
	}

	// act
	jmbg, err := New(number)

	// assert
	assert.NoError(err)
	assert.Equal(expected, jmbg)
}
