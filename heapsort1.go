package main

import (
	"fmt"
)

/*
what is a binary Tree?
Binary tree is data structure with n number of child/parent objects with index and key , consisting of atleast one child arranged in an array
left index = i*2 +1, right index = i*2+2
traverse through the binary tree through heap sort and arrange from bottom to top of binary tree with the max key

defining structs for the slice, index, key
create a slice to hold the array

*/

type maxHeap struct {
	slice []int
}

func (h *maxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapbottomToUp(len(h.slice) - 1)

}

func (h *maxHeap) maxHeapbottomToUp(index int) {

	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}

// Extract returns the largest key and removes it from heap
func (h *maxHeap) Extract() int {
	extracted := h.slice[0]
	l := len(h.slice) - 1
	if len(h.slice) == 0 {
		return -1
	}
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]
	h.maxHeapifyDown(0)
	return extracted

}

func (h *maxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.slice) - 1
	l, r := left(index), right(index)
	childToComapre := 0

	//loop while index has atleast one child
	for l <= lastIndex {
		//	if only left child is there; if left child > right child ; right > left
		if l == lastIndex {
			childToComapre = l
		} else if h.slice[l] > h.slice[r] {
			childToComapre = l
		} else {
			childToComapre = r
		}
		//	compare the slice
		if h.slice[index] < h.slice[childToComapre] {
			h.swap(index, childToComapre)
			index = childToComapre
			l, r = left(index), right(index)
		} else {
			return
		}

	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]

}
func left(i int) int {

	return i*2 + 1

}
func right(i int) int {
	return i*2 + 2
}

func main() {
	m := &maxHeap{}
	fmt.Println(m)
	heap := []int{1, 3, 88, 5, 9, 76, 987, 34}
	for _, v := range heap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < len(heap); i++ {
		m.Extract()
		fmt.Println(m)
	}

}
