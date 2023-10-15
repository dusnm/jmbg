package region

import (
	"fmt"
)

const (
	INT uint = 70
	BGD uint = 71
	SUM uint = 72
	NIS uint = 73
	JZM uint = 74
	ZJC uint = 75
	PDN uint = 76
	PDK uint = 77
	KRA uint = 78
	UZC uint = 79
	NSD uint = 80
	SOM uint = 81
	SUB uint = 82
	KIK uint = 84
	ZRE uint = 85
	PAN uint = 86
	VRS uint = 87
	RUM uint = 88
	SRM uint = 89
	PRI uint = 91
	KSM uint = 92
	PEC uint = 93
	DJK uint = 94
	PRZ uint = 95
	GNJ uint = 96
)

type (
	Name struct {
		Cyrillic string `json:"cyrillic"`
		Latin    string `json:"latin"`
	}

	Region struct {
		Code uint `json:"code"`
		Name Name `json:"name"`
	}

	InvalidCodeError struct {
		input uint
	}
)

func (e InvalidCodeError) Error() string {
	return fmt.Sprintf(
		"invalid region code %d, no region was found using it",
		e.input,
	)
}

func New(digits [13]uint) (Region, error) {
	code := digits[7]*10 + digits[8]
	regions := map[uint][]string{
		INT: {
			"Српски држављани регистровани у иностранству",
			"Srpski državljani registrovani u inostranstvu",
		},
		BGD: {
			"Београд",
			"Beograd",
		},
		SUM: {
			"Регион Шумадије и Поморавља",
			"Region Šumadije i Pomoravlja",
		},
		NIS: {
			"Ниш",
			"Niš",
		},
		JZM: {
			"Регион Јужне Мораве",
			"Region Južne Morave",
		},
		ZJC: {
			"Зајечар",
			"Zaječar",
		},
		PDN: {
			"Регион Подунавља",
			"Region Podunavlja",
		},
		PDK: {
			"Регион Подриња и Колубаре",
			"Region Podrinja i Kolubare",
		},
		KRA: {
			"Краљево",
			"Kraljevo",
		},
		UZC: {
			"Ужице",
			"Užice",
		},
		NSD: {
			"Нови Сад",
			"Novi Sad",
		},
		SOM: {
			"Сомбор",
			"Sombor",
		},
		SUB: {
			"Суботица",
			"Subotica",
		},
		KIK: {
			"Кикинда",
			"Kikinda",
		},
		ZRE: {
			"Зрењанин",
			"Zrenjanin",
		},
		PAN: {
			"Панчево",
			"Pančevo",
		},
		VRS: {
			"Вршац",
			"Vršac",
		},
		RUM: {
			"Рума",
			"Ruma",
		},
		SRM: {
			"Сремска Митровица",
			"Sremska Mitrovica",
		},
		PRI: {
			"Приштина",
			"Priština",
		},
		KSM: {
			"Косовска Митровица",
			"Kosovska Mitrovica",
		},
		PEC: {
			"Пећ",
			"Peć",
		},
		DJK: {
			"Ђаковица",
			"Đakovica",
		},
		PRZ: {
			"Призрен",
			"Prizren",
		},
		GNJ: {
			"Гњилане",
			"Gnjilane",
		},
	}

	region, exists := regions[code]
	if !exists {
		return Region{}, InvalidCodeError{input: code}
	}

	return Region{
		Code: code,
		Name: Name{
			Cyrillic: region[0],
			Latin:    region[1],
		},
	}, nil
}
