package checksum

type (
	Checksum struct {
		Value uint
		Valid bool
	}
)

func New(digits [13]uint) Checksum {
	// Assign a letter to each digit like:
	// abcdefg hi jkl m
	//
	// The cheksum K is calculated using the following formula:
	// m = 11 - ((7(a+g) + 6(b+h) + 5(c+i) + 4(d+j) + 3(e+k) + 2(f+l)) mod 11)
	//
	// if m is between 1 and 9 => K = m
	// if m is 10 or 11 => K = 0
	csum := digits[12]
	calculated_csum := 11 - ((7*(digits[0]+digits[6]) +
		6*(digits[1]+digits[7]) +
		5*(digits[2]+digits[8]) +
		4*(digits[3]+digits[9]) +
		3*(digits[4]+digits[10]) +
		2*(digits[5]+digits[11])) % 11)

	if calculated_csum == 10 || calculated_csum == 11 {
		calculated_csum = 0
	}

	if calculated_csum != csum {
		return Checksum{
			Value: csum,
			Valid: false,
		}
	}

	return Checksum{
		Value: csum,
		Valid: true,
	}
}
