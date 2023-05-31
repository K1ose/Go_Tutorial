package tempconv

import "fmt"

type Kelvin float64
type Fahrenheit float64
type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() {
	fmt.Printf("%g°C ", c)
}

func (f Fahrenheit) String() {
	fmt.Printf("%g°F ", f)
}

func (k Kelvin) String() {
	fmt.Printf("%gK ", k)
}
