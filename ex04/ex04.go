package main

import "fmt"

func towerBreakers(n int32, m int32) int32 {
	// Write your code here
	res := int32(0)
	tmp := m
	for {
		target := (tmp / m) * m
		res += tmp - target
		tmp = target
		if tmp < m {
			break
		}
		res += 1
		tmp /= m
	}
	res += tmp - 1
	res *= n

	if res%2 == 0 {
		return 2
	}

	return 1
}

func main() {
	res := towerBreakers(1, 4)
	fmt.Println(res)
}
