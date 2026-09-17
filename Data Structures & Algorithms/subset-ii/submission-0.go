// 90. Subsets II
// https://leetcode.com/problems/subsets-ii
//
// Decision: Always append result. First sort the input array and skip if the current index is greater that `start`
// and the current element is equal to the previous one.

package subsetii

import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	var result [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := append([]int{}, path...)
		result = append(result, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
