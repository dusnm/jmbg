package jmbg

import (
	"time"

	"github.com/dusnm/jmbg/checksum"
	"github.com/dusnm/jmbg/digits"
	"github.com/dusnm/jmbg/dob"
	"github.com/dusnm/jmbg/region"
	"github.com/dusnm/jmbg/sex"
)

type (
	JMBG struct {
		Number      string
		Digits      [13]uint
		DateOfBirth time.Time
		Region      region.Region
		Sex         sex.Sex
		Checksum    checksum.Checksum
	}
)

func New(jmbg string) (JMBG, error) {
	digits, err := digits.New(jmbg)
	if err != nil {
		return JMBG{}, err
	}

	date, err := dob.New(digits)
	if err != nil {
		return JMBG{}, err
	}

	region, err := region.New(digits)
	if err != nil {
		return JMBG{}, err
	}

	return JMBG{
		Number:      jmbg,
		Digits:      digits,
		DateOfBirth: date,
		Region:      region,
		Sex:         sex.New(digits),
		Checksum:    checksum.New(digits),
	}, nil
}
