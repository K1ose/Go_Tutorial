package tutorial

import "fmt"

func IntergerSample() {
	/* % 取模运算符的符号和被取模数的符号总是一致的 */
	fmt.Printf("5 mod -2 = %d\n", 5%(-2))   // 1
	fmt.Printf("-5 mod -2 = %d\n", -5%(-2)) //-1
	fmt.Printf("-5 mod 2 = %d\n", -5%2)     //-1
	fmt.Printf("5 mod 2 = %d\n", 5%2)       // 1

	/* 一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边的bit位是1的话，那么最终结果可能是负的，例如int8的例子 */
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	/* bit位操作运算符，前面4个操作运算符并不区分是有符号还是无符号数： */
	fmt.Printf("%b\n", 1&0)
	fmt.Printf("%b\n", 1|0)
	fmt.Printf("%b\n", 1^0)
	fmt.Printf("%b\n", 1&^1)
	// 位操作运算符&^用于按位置零（AND NOT）：如果对应y中bit位为1的话，表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。
	/*
		1&^0 = 1
		0&^1 = 0
		1&^1 = 0
		0&^0 = 0
	*/

	var x uint8 = 1<<1 | 1<<5
	/* 0000 0010 | 0010 0000 = 0010 0010*/
	var y uint8 = 1<<1 | 1<<2
	/* 0000 0010 | 0000 0100 = 0000 0110*/

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	/* 0010 0010 & 0000 0110 = 0000 0010 */
	fmt.Printf("%08b\n", x&y) // "00000010", the intersection {1}

	/* 0010 0010 | 0000 0110 = 0010 0110 */
	fmt.Printf("%08b\n", x|y) // "00100110", the union {1, 2, 5}

	/* 0010 0010 ^ 0000 0110 = 0010 0100 */
	fmt.Printf("%08b\n", x^y) // "00100100", the symmetric difference {2, 5}

	/* 0010 0010 &^ 0000 0110 = 0010 0000 */
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		/* x = 0010 0010 */
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	f := 3.141 // a float64
	g := int(f)
	fmt.Println(f, g) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	/* 通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。 */
	o := 0666                          // 以0开始的八进制格式
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	p := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", p)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	/* 字符使用%c参数打印，或者是用%q参数打印带单引号的字符 */
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
