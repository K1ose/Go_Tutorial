package main

import (
	"fmt"
	"os"
	"strconv"
	"weightconv"
)

func main() {
	var k float64 = 0
	if len(os.Args[1:]) == 0 {
		fmt.Scanln(&k)
	} else {
		for _, arg := range os.Args[1:] {
			if os.Args[1:] != nil {
				var err error
				k, err = strconv.ParseFloat(arg, 64) // ParseFloat() 将string类型转为对应的float类型
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					os.Exit(1)
				}
			}
		}
	}
	j := weightconv.Kgtojin(weightconv.Kilo(k))
	p := weightconv.Kgtopound(weightconv.Kilo(k))
	fmt.Printf("%s == %s == %s", weightconv.Kilo(k).String(), j.String(), p.String())
}
