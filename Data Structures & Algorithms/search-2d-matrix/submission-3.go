func searchMatrix(matrix [][]int, target int) bool {
	n_row, n_col := len(matrix), len(matrix[0])
	left, right := 0, n_row * n_col - 1

	for left <= right {
		mid := (left + right) / 2
		ele := matrix[mid / n_col][mid % n_col]

		if target == ele {
			return true
		} else if target > ele {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
