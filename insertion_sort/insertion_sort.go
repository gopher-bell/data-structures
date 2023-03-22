package main

import "fmt"

func insertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

func main() {
	s := []int{1, 5, 2, 4, 3, 6, 9, 7, 8, 0}
	insertionSort(s)
	fmt.Println(s)
}
