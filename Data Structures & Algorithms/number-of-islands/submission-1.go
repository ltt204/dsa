// 200. Number Of Islands
// https://leetcode.com/problems/number-of-islands
//
// DFS
//
// Note: We dont have to use the visited map. If we already visit it? switch it to 0

package numberofislands

func numIslands(grid [][]byte) int {
	var result int
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= rows || r < 0 ||
			c >= cols || c < 0 ||
			grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				dfs(r, c)
				result++
			}
		}
	}

	return result
}
