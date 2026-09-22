// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements
//schoosc.
// nums = [-6,2,3,-4,6,-2,1,2,2], k = 3 --> [2,3,6]Sort the array first, we will have
//
//	[-6,-4,-2,1,2,2,2,3,6]
//
// We can use a map for this. which slightly
// better.
// With slices we need to sort two distinguish
// times for two slices/arrays
// Let's see an example with map:
//  1. We dont have to sort the nums array .
//  2. for the count, we use an hashmap and check, if the element is already in the map --> increase the value. If not, create new key
//  3. tracking for the maximum frequency also.
//  4. Initialize another slice/array with length of the max freq --> counts := DistributedArray
//  5. Iterate throw hashmap -->
//
// counts[map.Key] = map.val. Basically we will have
// this:
//
// index/row represents frequency:	     0     1    2    3    4
// col represents the element:		[  [2,3]  [2]  [3]  [4]  [6]]
//
// From it we can iterate backward and get the kth most frequency element
package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	maxFreq := 1

	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]] = freqMap[nums[i]] + 1

		if freqMap[nums[i]] > maxFreq {
			maxFreq = freqMap[nums[i]]
		}
	}

	// Not gonna use row 0
	buckets := make([][]int, maxFreq+1)
	for key, val := range freqMap {
		buckets[val] = append(buckets[val], key)
	}

	// Iterate backward the buckets and get the exact amount of k elements.
	var result []int
	for row := len(buckets) - 1; row > 0; row-- {
		if len(result) >= k {
			break
		}
		result = append(result, buckets[row]...)
	}

	return result[0:k]
}

func main() {
	nums := []int{-6, -4, -2, 1, 2, 2, 2, 3, 6, 6}
	k := 4
	// nums := []int{1, 1, 1, 2, 2, 3}
	// k := 2

	fmt.Printf("Result: %v\n\n",
		topKFrequent(
			nums,
			k,
		),
	)
}
