package tutorial

import (
	"testing"
)

var x uint64 = 90000000

func BenchmarkPopCountMoveOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountMoveOne(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitClear(x)
	}
}
