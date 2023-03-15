package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(k int) {
	h.array = append(h.array, k)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	extracted := h.array[0]
	l := len(h.array) - 1

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	h := &MaxHeap{}
	s := []int{100, 300, 200, 400, 500, 600, 900, 700, 800}

	for _, v := range s {
		h.Insert(v)
	}

	for range s {
		h.Extract()
	}

	fmt.Println(h)
}
