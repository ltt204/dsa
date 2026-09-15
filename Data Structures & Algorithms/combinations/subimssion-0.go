// 77. Combinations
// https://leetcode.com/problems/combinations

func combine(n int, k int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int)

	backtrack = func(start int) {
		if len(path) == k {
			temp := append([]int{}, path...)
			result = append(result, temp)

			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)

	return result
}