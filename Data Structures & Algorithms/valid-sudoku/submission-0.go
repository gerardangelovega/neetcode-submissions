func isValidSudoku(board [][]byte) bool {
	for i, row := range board {
		row_set := make(map[byte]bool)
		col_set := make(map[byte]bool)

		for j := range len(row) {
			if _, exists := row_set[board[i][j]]; !exists {
				if board[i][j] != '.' {
					row_set[board[i][j]] = true
				}
			} else {
				fmt.Println("Duplicate number found in row", i)
				return false
			}

			if _, exists := col_set[board[j][i]]; !exists {
				if board[j][i] != '.' {
					col_set[board[j][i]] = true
				}
			} else {
				fmt.Println("Duplicate number found in column", j)
				return false
			}
		}
	}

	for i := range 3 {
		i = i * 3

		for j := range 3 {
			mat_set := make(map[byte]bool)
			j = j * 3

			for m := range 9 {
				mr := i + m / 3
				mc := j + m % 3
				if _, exists := mat_set[board[mr][mc]]; !exists {
					if board[mr][mc] != '.' {
						mat_set[board[mr][mc]] = true
					}
				} else {
					fmt.Println("Duplicate number found in 3x3", i, j, mr, mc)
					return false
				}
			}
		}
	}
	return true
}
