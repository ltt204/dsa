// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements
//
// nums = [-6,2,3,-4,6,-2,1,2,2], k = 3 --> [2,3,6]Sort the array first, we will have
//
//	[-6,-4,-2,1,2,2,2,3,6]
//
// After sorting, we count the adjacent element, which will be equivalent since we sorted. The counts will be append into a slice/array, that array will hold a pair {val-freq}
// [{-6,1}, {-4,1}, {-2, 1}, {1,1}, {2,3}, {3,-1}, {6,-1}]
// Sort descreasing again the counts slice
// --> [{2,3},{-6,1}, {-4,1}, {-2, 1}, {1,1}, {3,1}, {6,1}, ]
// the result is the counts[i].Val with i in range(0, k)
package main

import (
	"fmt"
	"slices"
	"sort"
)

func solution(nums []int, k int) []int {
	type Pair struct {
		Val, Freq int
	}
	var counts []Pair

	// Note 1: Use hashmap --> avoid sorting for nums
	slices.Sort(nums)
	fmt.Println(nums)

	count := 1
	for i := 1; i <= len(nums); i++ {
		if i != len(nums) && nums[i-1] == nums[i] {
			count++
		} else {
			counts = append(counts, Pair{Val: nums[i-1], Freq: count})
			count = 1
		}
	}
	fmt.Println(counts)

	sort.Slice(
		counts,
		func(i, j int) bool {
			return counts[i].Freq > counts[j].Freq
		},
	)

	var result []int
	for idx, pair := range counts {
		if idx >= k {
			break
		}
		result = append(result, pair.Val)
	}

	return result
}

func main() {
	// nums := []int{-6, -4, -2, 1, 2, 2, 2, 3, 6, 6}
	// k := 4
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Printf("Result: %v",
		solution(
			nums,
			k,
		),
	)
}
