package main

import "fmt"

type Inch float64
type Feet float64

const (
	InchPerFoot = 12.0
)

func (i Inch) String() string {
	return fmt.Sprintf("%g''", i)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g'", f)
}

func (i Inch) ToFeet() Feet {
	return Feet(i / InchPerFoot)
}

func (f Feet) ToInches() Inch {
	return Inch(f * InchPerFoot)
}
