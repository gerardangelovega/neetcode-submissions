func searchMatrix(matrix [][]int, target int) bool {
	n_row, n_col := len(matrix), len(matrix[0])
	left, right := 0, n_row * n_col - 1

	fmt.Println(n_row, n_col, left, right)

	for left <= right {
		mid := (left + right) / 2
		mid_row := mid / n_col
		mid_col := mid % n_col

		fmt.Println(mid, mid_row, mid_col, matrix[mid_row][mid_col])

		if target == matrix[mid_row][mid_col] {
			fmt.Println("Target Found")
			return true
		}
		if target > matrix[mid_row][mid_col] {
			fmt.Println("Target Greater")
			left = mid + 1
			continue
		}
		if target < matrix[mid_row][mid_col] {
			fmt.Println("Target Lesser")
			right = mid - 1
			continue
		}
	}

	return false
}
