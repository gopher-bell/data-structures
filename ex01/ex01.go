package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	writer  = bufio.NewWriter(os.Stdout)
	scanner = bufio.NewScanner(os.Stdin)
)

type tower struct {
	idx    int
	height int
}

func getNumber() int {
	ret := 0
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	runs := getNumber()

	stack := make([]tower, 0, runs)
	for i := 0; i < runs; i++ {
		num := getNumber()
		for len(stack) > 0 && stack[len(stack)-1].height < num {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			writer.WriteString("0 ")
		} else {
			writer.WriteString(strconv.Itoa(stack[len(stack)-1].idx) + " ")
		}

		stack = append(stack, tower{idx: i + 1, height: num})
	}
}
