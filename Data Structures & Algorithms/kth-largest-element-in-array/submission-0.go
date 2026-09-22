// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array
// 
// Using a min heap --> push and auto sorting
// mean heap ~ complete binary tree --> cannot treat that heap like a slice
// => Pop consecutively til the largest element is the kth largest of the original heap. 
package kthlargestelementinarray

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	intHeap := &IntHeap{}

	for _, val := range nums {
		heap.Push(intHeap, val)
	}

	kthLength := intHeap.Len() - k + 1
	for intHeap.Len() > kthLength {
		// fmt.Printf("%d ", heap.Pop(intHeap))
		heap.Pop(intHeap)
	}

	return (*intHeap)[0]
}
