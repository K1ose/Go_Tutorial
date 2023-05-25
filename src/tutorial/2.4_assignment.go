package tutorial

import "fmt"

func Assignment() {
	var b bool
	var p = &b

	*p = false /* 指针赋值 */

	fmt.Println(gcd(50, 100))
	fmt.Println(fib(5))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
