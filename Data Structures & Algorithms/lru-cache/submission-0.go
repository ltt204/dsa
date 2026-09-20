// 146. LRU Cache
// https://leetcode.com/problems/lru-cache
//
// Case: Out of Capacity
// ==> Check the least recently usage
// --> need 2nd datastructure, why?

// When we get a specific key, we likely "used" that key, therefor the
// order/priority is changed. The map itself cannot handle this logic so we have to
// use another datastructure for this.

// When a key is accessed, we have to reorder the stored keys, but usually,
// reorder would coupling to sorting. In this case using the array is not a good
// idea since the problem required that all method should produce O(1) time
// complexity.

// We might think another way, like the Insert/Remove/GetRandom O(1) problem,
// we used a map with pair constitute with {val - array_index} and the array for
// storing values. We not iterating for removing the val, we have O(1) lookup for
// values' index by virtue of the map, and we can also have O(1) removal by
// swapping that specific val to the last index and then shrink the array by 1. The
// question is, can we reuse that design for this LRU cache?

// ## get(key int)
// Easily O(1) for get by using map and array. But this is the LRU cache, we have
// to update the prioty of the array and the new index within the map also.

// for example we have:
//     - array: [1,3,4,5] --> value
//     - map: [{1-0}, {3-1}, {4-2}, {5-3}] --> key

// when we get(4), we have to:
//     1. update array: [1,2,5,4] - with the assumpton that the left most one is
// LRU and the right most one is MRU
//     2. Update new index for the map: [{1-0}, {3-1}, {4-3}, {5-2}]

// ## put (key int, val int)
// append to the end of the array and add new key-index to the map.

// when out of capacity, we remove the Left most value within the array, also in
// the map and start to append like usual.
// ---

// I have not seen no problem with array. Let's try it first.

// I saw the problem now. The inferiority of the array is the amount of
// attributes it can store. when we put a pair key-value obey the requirement, I
// simply misunderstood that the array value within the array will equal with the
// key - like [1,1] [4,4]. but if we have the [2,4] pair the swapping logic will be
// derailed since after we swapped we wont able to access back to the map. We
// only store the key with index, and val is at that index. but when we swap, the
// val is no longer stay at the index and also we have no clue how to get back the
// key paired with that value also.
// ---

// Back to doubly linked list again.

// With doubly Linked-list. The map will now storing the key paired with the Node.
// which mean when we swap the nodes within the list, we dont have to update
// the map since the value is a pointer.

// Still preserve the advantage of the map, we quickly get the node with the key.
// but instead of swap, we remove and reinsert it at the end of the list.

// To achieve O(1), the Insert and Remove method of the linkedlist should
// produce O(1) also. To make this able, we have to store the head and tail, with
// akin to the idea of using array - we storeing the lru and mru respectively left
// mode value and right most value. Insert and Remove will compute based on
// those two node only.

// The out of capacity is still easy to maintain, by just modify the lru node

package lrucache

type Node struct {
	Left  *Node
	Right *Node
	Val   int
	Key   int
}

func NewNode() *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Val:   0,
		Key:   0,
	}
}

type LRUCache struct {
	hmap map[int]*Node
	// Tracking for the least recently used
	lru *Node
	// Tracking for the most recently used
	mru *Node

	capacity int
}

func Constructor(capacity int) LRUCache {
	lru := NewNode()
	mru := NewNode()

	lru.Right = mru
	mru.Left = lru

	return LRUCache{
		hmap:     make(map[int]*Node),
		lru:      lru,
		mru:      mru,
		capacity: capacity,
	}
}

func (this *LRUCache) Insert(key, val int) {
	new := &Node{Key: key, Val: val}

	new.Right, new.Left = this.mru, this.mru.Left
	this.mru.Left.Right, this.mru.Left = new, new

	this.hmap[key] = new
}

func (this *LRUCache) Remove(key int) {
	node := this.hmap[key]

	right, left := node.Right, node.Left

	left.Right = node.Right
	right.Left = node.Left

	delete(this.hmap, key)
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.hmap[key]; !ok {
		return -1
	}

	val := this.hmap[key].Val
	this.Remove(key)
	this.Insert(key, val)

	node := this.hmap[key]
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.hmap[key]; ok {
		this.Remove(key)
		this.Insert(key, value)

		return
	}

	this.Insert(key, value)

	if len(this.hmap) > this.capacity {
		this.Remove(this.lru.Right.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
