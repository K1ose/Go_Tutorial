package tutorial

import (
	"fmt"
	"unicode/utf8"
)

func StringSample() {
	s := "hello, world"
	fmt.Printf("%s\n", s)

	// c := s[len(s)] // panic: index out of range
	// fmt.Printf("%d\n", c)

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

	const strUsage = `Go is the greatest program language.`
	fmt.Println(strUsage[:2])
	if hasPrefix(strUsage, "Go") {
		fmt.Println("has prefix: " + "Go")
	}
	if hasSuffix(strUsage, "ge.") {
		fmt.Println("has suffix: " + "ge.")
	}
	if contains(strUsage, "greatest") {
		fmt.Println("contains: " + "greatest")
	}

	/* 字符串包含13个字节，以UTF8形式编码，只对应9个Unicode字符 */
	s_unicode := "Hello, 世界"
	fmt.Println(len(s_unicode))                    // "13"
	fmt.Println(utf8.RuneCountInString(s_unicode)) // "9"

	/* 为了处理这些真实的字符，我们需要一个UTF8解码器。unicode/utf8包提供了该功能 */
	/* 每一次调用DecodeRuneInString函数都返回一个r和长度，r对应字符本身，长度对应r采用UTF8编码后的编码字节数目。长度可以用于更新第i个字符在字符串中的字节索引位置。 */
	for i := 0; i < len(s_unicode); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

/* 前缀包含 */
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

/* 后缀包含 */
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

/* 字串包含 */
func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
