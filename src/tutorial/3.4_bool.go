package tutorial

import "fmt"

func BoolSample() {
	/* 布尔值可以和&&（AND）和||（OR）操作符结合，并且有短路行为：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，因此下面的表达式总是安全的 */
	var s string = ""
	if s != "" && s[0] == 'h' {
		fmt.Printf("%s\n", "yes")
	}

	/* &&的优先级比||高 */
	var c byte = 'd'
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		fmt.Printf("%s\n", "yesyes")
	}

	/* 布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换 */
	i := 0
	var b bool = true
	if b {
		i = 1
	} else {
		fmt.Printf("%d\n", i)
	}

	fmt.Printf("%d\n", btoi(b))
	if itob(1) {
		fmt.Printf("%s\n", "True!")
	}
}

/* 如果需要经常做类似的转换，将布尔值转换为0或1，包装成一个函数会更方便 */
// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

/* 数字到布尔型的逆转换则非常简单，不过为了保持对称，我们也可以包装一个函数 */
// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }
