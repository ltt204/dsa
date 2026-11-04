func searchMatrix(matrix [][]int, target int) bool {
	var l, r int

	for  row := 0; row < len(matrix); row ++ {
		l = 0
		r = len(matrix[0]) - 1

		if target > matrix[row][r] {
			continue
		}

		for l <= r {
			m := l + (r - l) / 2

			if target == matrix[row][m] {
				return true
			}

			if target > matrix[row][m] {
				l = m + 1
			} else if target < matrix[row][m] {
				r = m - 1
			}
		}
	}

	return false
}
