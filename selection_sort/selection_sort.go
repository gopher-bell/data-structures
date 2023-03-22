package main

import "fmt"

func selectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func main() {
	s := []int{1, 5, 2, 4, 3, 6, 9, 7, 8, 0}
	selectionSort(s)
	fmt.Println(s)
}
