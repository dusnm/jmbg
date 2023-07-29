package sex

const (
	MALE   = 0
	FEMALE = 1
)

type (
	Name struct {
		Cyrilic string
		Latin   string
	}

	Sex struct {
		Code int
		Name Name
	}
)

func New(digits [13]uint) Sex {
	sex := digits[9]*100 + digits[10]*10 + digits[11]
	if sex <= 499 {
		return Sex{
			Code: MALE,
			Name: Name{
				Cyrilic: "мушки",
				Latin:   "muški",
			},
		}
	}

	return Sex{
		Code: FEMALE,
		Name: Name{
			Cyrilic: "женски",
			Latin:   "ženski",
		},
	}
}
