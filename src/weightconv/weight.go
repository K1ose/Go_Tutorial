// Package weightconv performs Kilogram, jin and pound conversions.
package weightconv

import "fmt"

type Kilo float64
type Jin float64
type Pound float64

func (k Kilo) String() string {
	return fmt.Sprintf("%gkg", k)
}

func (j Jin) String() string {
	return fmt.Sprintf("%gjin", j)
}

func (p Pound) String() string {
	return fmt.Sprintf("%gpound(s)", p)
}
