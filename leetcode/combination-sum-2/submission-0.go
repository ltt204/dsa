// Leetcode 40. Combination Sum II
// https://leetcode.com/problems/combination-sum-ii

package combinationsum2

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var result [][]int
	var path []int
	var sum int

	var backtrack func(start int)

	backtrack = func(start int) {
		if sum > target {
			return
		}

		if sum == target {
			temp := append([]int{}, path...)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i + 1)

			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)

	return result
}
