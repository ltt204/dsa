func searchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := len(matrix) - 1
	potentialRow := 0

	for  l <= r {
		m := l + (r - l) / 2

		if target >= matrix[m][0] && target <= matrix[m][len(matrix[0]) - 1] {
			potentialRow = m
			break
		}

		if target > matrix[m][len(matrix[0]) - 1] {
			l = m + 1
		} else if target < matrix[m][0] {
			r = m - 1
		}

	}

	l = 0
	r = len(matrix[0]) - 1
	for l <= r {
		m := l + (r - l) / 2

		if target == matrix[potentialRow][m] {
			return true
		}

		if target > matrix[potentialRow][m] {
			l = m + 1
		} else if target < matrix[potentialRow][m] {
			r = m - 1
		}
	}
	return false
}
