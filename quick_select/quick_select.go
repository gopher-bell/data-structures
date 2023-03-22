package main

import "fmt"

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] >= pivot {
			i = i + 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func quickSelect(nums []int, k int) int {
	var l, r int
	r = len(nums) - 1
	for {
		if l == r {
			return nums[l]
		}
		pivotIndex := partition(nums, l, r)
		if k-1 == pivotIndex {
			return nums[pivotIndex]
		} else if k-1 < pivotIndex {
			r = pivotIndex - 1
		} else {
			l = pivotIndex + 1
		}
	}
	return 0
}

/*
{1, 2, 3, 4, 5, 6}
*/

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	res := quickSelect(s, 1)
	fmt.Println(res)
}
