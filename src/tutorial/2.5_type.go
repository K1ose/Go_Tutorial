package tutorial

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func ctof(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func ftoc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func Type() {
	var c float64 = 1.0
	fmt.Printf("%.2f Celsius = %.2f Fahrenheit", c, ctof(Celsius(c)))
	/* 即使Celsius也是float64类型，但是还是需要显式转换为Celsius类型，因为函数声明的参数是Celsius类型，用ctof(c)会报错 */

	var f Fahrenheit = 3.2
	fmt.Printf("%.2f Fahrenheit = %.2f Celsius", f, ftoc(f))

	cel := ftoc(21.0)
	fmt.Println(cel.String())
}

func (c Celsius) String() string { return fmt.Sprintf("result is %g°C", c) }
