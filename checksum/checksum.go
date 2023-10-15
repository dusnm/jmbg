package checksum

type (
	Checksum struct {
		Value uint `json:"value"`
		Valid bool `json:"valid"`
	}
)

func New(digits [13]uint) Checksum {
	// Assign a letter to each digit like:
	// abcdefg hi jkl m
	//
	// The checksum K is calculated using the following formula:
	// m = 11 - ((7(a+g) + 6(b+h) + 5(c+i) + 4(d+j) + 3(e+k) + 2(f+l)) mod 11)
	//
	// if m is between 1 and 9 => K = m
	// if m is 10 or 11 => K = 0
	checksum := digits[12]
	calculatedChecksum := 11 - ((7*(digits[0]+digits[6]) +
		6*(digits[1]+digits[7]) +
		5*(digits[2]+digits[8]) +
		4*(digits[3]+digits[9]) +
		3*(digits[4]+digits[10]) +
		2*(digits[5]+digits[11])) % 11)

	if calculatedChecksum == 10 || calculatedChecksum == 11 {
		calculatedChecksum = 0
	}

	if calculatedChecksum != checksum {
		return Checksum{
			Value: checksum,
			Valid: false,
		}
	}

	return Checksum{
		Value: checksum,
		Valid: true,
	}
}
