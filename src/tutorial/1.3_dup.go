package tutorial

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dup1() {
	counts := make(map[string]int)      /* 内置函数 make 创建空 map */
	input := bufio.NewScanner(os.Stdin) /* Scanner 类型是bufio包最有用的特性之一，它读取输入并将其拆成行或单词；这里的 input 是 bufio.Scanner 类型的变量 */
	for input.Scan() {
		/* 每次调用 input.Scan()，即读入下一行，并移除行末的换行符，读取的内容可以调用 input.Text() 得到 */
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// fmt.Println(input.Text())
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
