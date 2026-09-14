// LeetCode 78. Subsets
// https://leetcode.com/problems/permutations/
//
// Decision at each element: include it or not. Track the choice in `path`,
// recurse, then undo it (backtrack) before trying the other branch.

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	visited := make([]bool, len(nums)+1)

	var backtrack func()

	backtrack = func() {

		if len(path) == len(nums) {
			temp := append([]int{}, path...)
			result = append(result, temp)
			return
		}

		for i, num := range nums {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, num)

			backtrack()

			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()

	return result
}
