package main

import (
	"bufio"
	"os"
	"strconv"
)

var num int = 1

func sendEven(ch chan int) {
	defer close(ch)
	for num <= 100 {
		if num%2 == 0 {
			ch <- num
			num++
		}
	}
}

func sendOdd(ch chan int) {
	for num <= 100 {
		if num%2 != 0 {
			ch <- num
			num++
		}
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	ch := make(chan int)
	go sendEven(ch)
	go sendOdd(ch)

	for v := range ch {
		writer.WriteString(strconv.Itoa(v) + "\n")
	}
}
