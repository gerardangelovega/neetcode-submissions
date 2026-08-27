type NumMatrix struct {
	mat [][]int
	pre [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	mat := make([][]int, m)
	pre := make([][]int, m)

	for i, row := range matrix {
		mat[i] = make([]int, n)
		pre[i] = make([]int, n)
		copy(mat[i], row) 
		copy(pre[i], row) 
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ { pre[i][j] = pre[i][j-1] + pre[i][j] }
	}

	return NumMatrix {
		mat: mat,
		pre: pre,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			sum = sum + this.pre[i][col2]
		} else {
			sum = sum + (this.pre[i][col2] - this.pre[i][col1-1])
		}
	}
	return sum
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
