package tutorial

import "fmt"

func StringSample() {
	s := "hello, world"
	fmt.Printf("%s\n", s)

	c := s[len(s)] // panic: index out of range
	fmt.Printf("%d\n", c)

	fmt.Println("len:", len(s[:5]), s[:5]) // "hello"
	fmt.Println("len:", len(s[7:]), s[7:]) // "world"
	fmt.Println("len:", len(s[:]), s[:])   // "hello, world"

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	str := "left foot"
	t := str
	str += ", right foot"

	fmt.Println(str) // "left foot, right foot"
	fmt.Println(t)   // "left foot"

	/* str[0] = 'L' // compile error: cannot assign to s[0] */
	/* TODO */
}
