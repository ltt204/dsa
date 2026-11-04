// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1
//
// Idea: We will torn this problem in two:
// 1. Insert/Remove O(1)
// 2. GetRandom O(1)
// ---
// 1. Insert/Remove O(1): Hashmap will be the best choice for this.
// Insert: hmap[key]=val
// Delete: delete(hmap, key)
// 2. GetRandom O(1): Array/slice ofcourse.
// ran := Int.Rand(0, len(arr)-1)
// return arr[ran]
// ----
// Combine two of them, we have a problem about datastructure management.
// with insert and remove as O(1) but combine array and map.
// ## Insert
// with array, we can append consecutively. and also with map.

// ## Remove
// This method is trickier. As usual, to properly delete a value within
// an array, we have to iterate to that specific index/value and overwrite
// it with the rest of the array.

// But we have another way, since we are using a map, we can store the
// pair with {element as key - index as value}. This help us to quicky get
// the index of the element we want to delete.

// After that, we can swap that element with the last index and shrink the
// slice 1 unit along with remove the key --> O(1) delete

// ## Example
// So let visualize the idea first. We have this input array [1,2,4,2]
// --- Insert 1
// arr: [1,]
// map: [{1 - 0}, ]
// --- Insert 4
// arr: [1,2,4]
// map: [{1-0}, {2-1}, {4-2}]
// --- Insert 2
// arr: [1,2]
// map: [{1-0}, {2-1}]

// --- Remove 2
// arr: [1,2,4]
// map: [{1-0}, {2-1}, {4-2}]
// => From map get the index of 2: {2-1} -> 1
// => Swap 2 with the last index, which has value 4: [1,4,2]
// => Update new index for 4 in map: [{1-0}, {2-1}, {4-1}]
// => shrink the array by 1: arr [1,4]
// => Delete the key 2 in map: [{1-0}, {4-1}]

package insertdeletegetrandom1

import "math/rand/v2"

type RandomizedSet struct {
	arr  []int
	hmap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hmap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hmap[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.hmap[val] = len(this.arr) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	valIndex, ok := this.hmap[val]
	if !ok {
		return false
	}

	// Swap index of val with the last index
	n := len(this.arr) - 1
	last := this.arr[n]
	this.arr[n] = this.arr[valIndex]
	this.arr[valIndex] = last

	// Update the map with the new index
	this.hmap[this.arr[valIndex]] = valIndex

	// Shrink the array
	this.arr = this.arr[:n]
	//delete key val in has hmap
	delete(this.hmap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	randInt := rand.IntN(len(this.arr))

	return this.arr[randInt]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
