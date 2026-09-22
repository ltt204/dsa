package dsa

type IntHeap []int

// 2. Implement sort.Interface (Len, Less, Swap) for your type
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap logic
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 3. Implement heap.Interface (Push, Pop)
// Note: These must use pointer receivers because they modify the slice's length
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
