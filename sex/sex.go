package sex

const (
	MALE   = 0
	FEMALE = 1
)

type (
	Name struct {
		Cyrillic string `json:"cyrillic"`
		Latin    string `json:"latin"`
	}

	Sex struct {
		Code int  `json:"code"`
		Name Name `json:"name"`
	}
)

func New(digits [13]uint) Sex {
	sex := digits[9]*100 + digits[10]*10 + digits[11]
	if sex <= 499 {
		return Sex{
			Code: MALE,
			Name: Name{
				Cyrillic: "мушки",
				Latin:    "muški",
			},
		}
	}

	return Sex{
		Code: FEMALE,
		Name: Name{
			Cyrillic: "женски",
			Latin:    "ženski",
		},
	}
}
