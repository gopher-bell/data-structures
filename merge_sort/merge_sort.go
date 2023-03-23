package main

import "fmt"

func merge(first, second []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	first := mergeSort(nums[:len(nums)/2])
	second := mergeSort(nums[len(nums)/2:])
	return merge(first, second)
}

func main() {
	s := []int{1, 3, 5, 7, 2, 4, 6, 8, 9, 0}
	res := mergeSort(s)
	fmt.Println(res)
}
