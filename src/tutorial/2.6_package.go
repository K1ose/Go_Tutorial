package tutorial

import "fmt"

var a = b + c // a 第三个初始化, 为 3
var b = fct() // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1

func fct() int {
	return c + 1
}

// pc[i] is the population count of i.
/* pc[i]为八位二进制数i所含1的个数，即pc[0]~pc[255] 为0~255所含1的个数，例如pc[10]为10的二进制1010所含1的个数，即为2 */
var pc [256]byte

func init() {
	for i := range pc {
		/* 原本这里是pc[i] = pc[i/2] + byte(i&1)，个人感觉改为下述代码好一些，这也表明，推导辅助表格的方式为：对于n位的二进制数，先得到前n-1位二进制数中所含1的个数，再判断最后一位是不是1，是的话就加上*/
		pc[i] = pc[i>>1] + byte(i&1)
		// fmt.Printf("%d: %b ", i, i)
		// fmt.Printf("%d\n", pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	/* byte()每次会取最低八位 */
	return int(pc[byte(x>>(0*8))] + /* 最低八位 */
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]) /* 最高八位 */
}

func PrintCount() {
	fmt.Println("hello")
}
