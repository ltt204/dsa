// Leetcode 39. Combination Sum
// https://leetcode.com/problems/combination-sum/

package combinationsum

func combinationSum(candidates []int, target int) [][]int {
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
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i)

			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)

	return result
}
