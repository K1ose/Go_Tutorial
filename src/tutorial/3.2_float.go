package tutorial

import (
	"fmt"
	"math"
)

func FloatSample() {
	/* 因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差 */
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	/* 很小或很大的数最好用科学计数法书写，通过e或E来指定指数部分 */
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数

	/* 打印e的幂，打印精度是小数点后三个小数精度和8个字符宽度 */
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	/* NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1) */
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	/* 函数math.IsNaN用于测试一个数是否是非数NaN，math.NaN则返回非数对应的值。虽然可以用math.NaN来表示一个非法的结果，但是测试一个结果是否是非数NaN则是充满风险的，因为NaN和任何数都是不相等的 */
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

}

/* 如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败 */
/* func compute() (value float64, ok bool) {
	// ...
	if failed {
		return 0, false
	}
	return result, true
} */
