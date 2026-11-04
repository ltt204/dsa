// 200. Number Of Islands
// https://leetcode.com/problems/number-of-islands
//
// If we hit the "1", from that we trarverse around it to
// find all other connected "1".

// At every indices, we can see it as a root of a tree.
// Therefor we can either use BFS or DFS for finding
// the connected element.

// If we successfully implement the core logic, we have to
// solve a subproblem which is the duplicated visiting.
// Since we start from a "1", we find all other "1" --> we
// already check for "that" other "1" and we do not
// want to check it again.
package numberofislands

func numIslands(grid [][]byte) int {
	type Coord struct {
		X, Y int
	}

	var result int
	visited := make(map[Coord]struct{})
	rows, cols := len(grid), len(grid[0])

	bfs := func(r, c int) {
		// predefine the directions for quickly traversing
		directions := []Coord{
			{X: 1, Y: 0},
			{X: -1, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: -1},
		}

		// Queue for bfs
		var queue []Coord
		queue = append(queue, Coord{X: r, Y: c})

		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]
			for _, direct := range directions {

				row := pop.X + direct.X
				col := pop.Y + direct.Y
				coord := Coord{X: row, Y: col}

				if _, ok := visited[coord]; !ok &&
					row < rows && row >= 0 &&
					col < cols && col >= 0 &&
					grid[row][col] == '1' {
					queue = append(queue, coord)
					visited[coord] = struct{}{}
				}
			}
		}

	}

	for r := range rows {
		for c := range cols {
			coord := Coord{X: r, Y: c}
			if _, ok := visited[coord]; !ok && grid[r][c] == '1' {
				bfs(r, c)
				result++
			}
		}
	}

	return result
}
