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
		Number      string            `json:"number"`
		Digits      [13]uint          `json:"digits"`
		DateOfBirth time.Time         `json:"date_of_birth"`
		Region      region.Region     `json:"region"`
		Sex         sex.Sex           `json:"sex"`
		Checksum    checksum.Checksum `json:"checksum"`
	}
)

func New(jmbg string) (JMBG, error) {
	d, err := digits.New(jmbg)
	if err != nil {
		return JMBG{}, err
	}

	date, err := dob.New(d)
	if err != nil {
		return JMBG{}, err
	}

	r, err := region.New(d)
	if err != nil {
		return JMBG{}, err
	}

	return JMBG{
		Number:      jmbg,
		Digits:      d,
		DateOfBirth: date,
		Region:      r,
		Sex:         sex.New(d),
		Checksum:    checksum.New(d),
	}, nil
}
