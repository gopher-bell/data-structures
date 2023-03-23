package main

import "fmt"

func partition(nums []int, left, right int) int {
	pivotIndex := right
	j := left
	for i := left; i < right; i++ {
		if nums[i] > nums[pivotIndex] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	nums[j], nums[pivotIndex] = nums[pivotIndex], nums[j]
	return j
}

func quickSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p-1)
		quickSort(nums, p+1, right)
	}
}

func main() {
	s := []int{1, 5, 7, 9, 3, 2, 4, 6, 8, 0}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
