# JMBG

[![License: LGPL v3](https://img.shields.io/badge/License-LGPL_v3-blue.svg)](https://github.com/dusnm/jmbg/blob/main/LICENSE)
![Tests](https://github.com/dusnm/jmbg/actions/workflows/test.yml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/dusnm/jmbg)](https://goreportcard.com/report/github.com/dusnm/jmbg)

A go library to parse the Serbian JMBG (Jedinstveni Matični Broj Građana), or UMCN (Unique Master Citizen Number)

## Explanation
The number that the government of Serbia assigns to its citizens is not random and is made up of the following information:

* A person's date of birth
* Which region of the country they were born in
* The sex they were assigned at birth
* A single digit checksum for error detection

```go
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
```

This library attempts to make it simple to extract this information out of a string representation of this number.

## Basic usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/dusnm/jmbg"
)

func main() {
	number := "0101000805018"

	data, err := jmbg.New(number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

## Licensing
Licensed under the terms of the GNU Lesser General Public License, version 3
