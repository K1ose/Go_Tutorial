package tutorial

func PopCountLoop(x uint64) int {
	var res int = 0
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func PopCountMoveOne(x uint64) int {
	var res int = 0
	for x > 0 {
		res += int(x & 1)
		x = x >> 1
	}
	return res
}

func PopCountBitClear(x uint64) int {
	var res int = 0
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}
