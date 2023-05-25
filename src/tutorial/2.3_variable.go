package tutorial

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/* 包级别声明的变量会在main入口函数执行前完成初始化 */
const con = "hello"

func VariableNormal() {
	/* var 变量名字 类型 = 表达式 */
	var s string // s = ""
	fmt.Println(s)

	var i, j, k int
	fmt.Printf("i=%d j=%d k=%d\n", i, j, k)

	/* 类型推导 */
	var b, f, str = true, 2.3, "four" // bool, float64, string
	fmt.Printf("b=%t f=%.6f str=%s\n", b, f, str)

	fmt.Println(con)

	var exist_filename = "D:\\Klose\\project\\go\\src\\tutorial\\textbook.txt"
	var not_exist_filename = "D:\\Klose\\project\\go\\src\\tutorial\\textbook"

	/*
		fl_e 类型为 *os.File，err_e 为报错信
	*/
	var fl_e, err_e = os.Open(exist_filename)
	fmt.Println(fl_e, err_e)
	var fl_ne, err_ne = os.Open(not_exist_filename)
	fmt.Println(fl_ne, err_ne)

	/* 简短变量声明语句的形式可用于声明和初始化局部变量 */
	simple_statement := "It is simple."
	freq := rand.Float64() * 2.0
	fmt.Printf("%s, %.2f", simple_statement, freq)
	/* 简短变量声明左边的变量可能并不是全部都是刚刚声明的。如果有一些已经在相同的词法域声明过了（§2.7），那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了，但是需要同时声明一个新的变量 */
	new_v, freq := 1.0, 2.0 // 对freq赋值，同时声明一个新的变量
	/* 直接freq:=2.0 会报错 */
	fmt.Printf("%.2f, %.2f", new_v, freq)
}

/***************** Pointer  ****************/
func VariablePointer() {
	var x = 1
	p := &x /* &[变量名] 的表达式将产生一个指向该变量的指针，即 取该变量的内存地址 */
	fmt.Println(p)
	fmt.Println(*p)

	*p = 2 // 更新指针p所指向的变量x的值为2

	var i, j int
	fmt.Println(&i == &i, &i == &j, &i == nil)
	// "true false false"

}

/* 返回函数中局部变量的地址也是安全的 */
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	(*p)++
	return *p
}

func PointerF() {
	var p1 = f()
	var p2 = f()
	fmt.Println(p1 == p2) // "false"

	v := 1
	incr(&v)              // 2
	fmt.Println(incr(&v)) // 3
}

/*
	调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量:

func flag.Bool(name string, value bool, usage string) *bool
- name: 命令行标志参数的名字
- value: 标志参数的默认值
- usage: 该标志参数对应的描述信息
*/
/* n, sep分别是指向对应命令行标志参数变量的指针，必须用*n, *sep 来进行引用*/
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func HandleFlag() {
	flag.Parse() // 必须在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）
	// 非标志参数的普通命令行参数可以通过调用flag.Args()函数来访问，返回值对应一个字符串类型的slice

	fmt.Print(strings.Join(flag.Args(), *sep))
	// fmt.Println("*sep: ", *sep)
	// fmt.Println("n: ", *n)
	if *n {
		fmt.Println("You used '-n'.")
	} else {
		fmt.Println("no no no.")
	}
}

/***************** New *********************/
func NewFunc() {
	p := new(int)
	fmt.Println(p, *p)

	q := new(int)
	fmt.Println(p == q)

	/* 由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型。 */
	new := "hello"
	fmt.Println(new)

	/* 如果new作为函数参数名，则该函数内部无法使用new作为预定义函数 */
}
