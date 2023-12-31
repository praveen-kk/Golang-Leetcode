package main

import "fmt"

//MaxHeap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

//Insert adds an element to the heap

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract return the largest key, and remove it from the heap
func (h *MaxHeap) Extract() int {

	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted

}

//maxHeapifyUp will heapify from bottom top

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)

	}

}

// maxHeapifyUp will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	//loop while index has at least one child
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)

	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			//When left child is the only child
			childToCompare = 1
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		//when left child is the larger child
		//when right child is the larger child
		//compare the array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 390, 53, 5, 99}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)

	}

}
