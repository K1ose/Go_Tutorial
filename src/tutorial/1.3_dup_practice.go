package tutorial

import (
	"bufio"
	"os"
)

func Dup_practice() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {

	}
}

func CountLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
