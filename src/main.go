package main

import (
	"fmt"
	"tutorial"
)

func main() {
	var x uint64 = 9000000
	// fmt.Printf("%b\n", byte(x))
	// fmt.Printf("%b\n", x>>8)
	// fmt.Printf("%b\n", byte(x>>8))
	// fmt.Printf("%b\n", byte(1010>>(1*8)))
	// fmt.Printf("%b\n", 101>>(1*8))
	// fmt.Printf("%b\n", byte(1010>>2))
	fmt.Printf("%d\n", tutorial.PopCountBitClear(x))
}
