package dob

import (
	"fmt"
	"time"
)

type (
	InvalidYearDigitError struct{}
)

func (e InvalidYearDigitError) Error() string {
	return "the first digit of the year can either be a 9 or a 0"
}

func New(digits [13]uint) (time.Time, error) {
	var year uint

	// This is ugly, naive and stupid, but it must be done
	switch digits[4] {
	case 0:
		year = 2000 + digits[5]*10 + digits[6]
	case 9:
		year = 1000 + digits[4]*100 + digits[5]*10 + digits[6]
	default:
		return time.Time{}, InvalidYearDigitError{}
	}

	return time.Parse(
		"02-01-2006",
		fmt.Sprintf(
			"%d%d-%d%d-%d",
			digits[0],
			digits[1],
			digits[2],
			digits[3],
			year,
		),
	)
}
