package tutorial

import(
	"os"
	"strconv"	// strconv 提供了一些于string相关的类型转换函数
	"time"
	"fmt"
)


/*  练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。 */
func Echo_arg_practice1() string {
	return os.Args[0]
}

/* 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。*/
func Echo_arg_practice2() string {
	var s string
	for idx, arg := range os.Args[1:] {
		s += strconv.Itoa(idx) + " " + arg + "\n"
	} 
	return s
}


/* 练习 1.3： 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6 节讲解了部分 time 包，11.4 节展示了如何写标准测试程序，以得到系统性的性能评测。） */
func Echo_arg_practice3_1() {
	start_time := time.Now()
	s := Echo_argument1()
	fmt.Print(s + " | ")
	fmt.Printf("%.6fs \n", time.Since(start_time).Seconds())
}

func Echo_arg_practice3_2() {
	start_time := time.Now()
	s := Echo_argument2()
	fmt.Print(s + " | ")
	fmt.Printf("%.6fs \n", time.Since(start_time).Seconds())
}

func Echo_arg_practice3_3() {
	start_time := time.Now()
	s := Echo_argument3()
	fmt.Print(s + " | ")
	fmt.Printf("%.6fs \n", time.Since(start_time).Seconds())
}