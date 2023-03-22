package main

import "fmt"

func binarySearch(nums []int, a, b, k int) int {
	if a > b {
		return -1
	}

	m := (a + b) / 2
	if nums[m] == k {
		return m
	} else if nums[m] > k {
		return binarySearch(nums, a, m-1, k)
	} else {
		return binarySearch(nums, m+1, b, k)
	}
}

func main() {
	s := []int{1, 3, 4, 5, 7, 9, 10}
	res := binarySearch(s, 0, len(s)-1, -1)
	fmt.Println(res)
}
