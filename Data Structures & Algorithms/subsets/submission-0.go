// LeetCode 78. Subsets
// https://leetcode.com/problems/subsets/
//
// Decision at each element: include it or not. Track the choice in `path`,
// recurse, then undo it (backtrack) before trying the other branch.
package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int
	countBacktrack := 0

	var backtrack func(start int)
	backtrack = func(start int) {
		countBacktrack++
		temp := append([]int{}, path...)
		result = append(result, temp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
