package digits

import (
	"fmt"
	"unicode"
)

type (
	InvalidLengthError struct{}
	InvalidDigitError  struct {
		input rune
	}
)

func (e InvalidLengthError) Error() string {
	return "jmbg must be exactly 13 digits long"
}

func (e InvalidDigitError) Error() string {
	return fmt.Sprintf(
		"invalid character %d, jmbg must be comprised only of numeric characters",
		e.input,
	)
}

func New(jmbg string) ([13]uint, error) {
	runes := []rune(jmbg)
	if len(runes) != 13 {
		return [13]uint{}, InvalidLengthError{}
	}

	digits := make([]uint, 13)
	for i, r := range runes {
		if !unicode.IsDigit(r) {
			return [13]uint{}, InvalidDigitError{input: r}
		}

		digits[i] = uint(r - '0')
	}

	return [13]uint(digits), nil
}
