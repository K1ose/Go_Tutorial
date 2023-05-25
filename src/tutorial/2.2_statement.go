package tutorial

import "fmt"

const boiling_point = 212.0

func BoilingF() {
	var f = boiling_point
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g °C\n", f, c)
}

func FtoC() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing point = %g°F or %g°C\n", freezingF, transform(freezingF))
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, transform(boilingF))
}

func transform(f float64) float64 {
	return (f - 32) * 5 / 9
}
