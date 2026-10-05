func combine(n int, k int) [][]int {
	combinations, combination := make([][]int, 0, 4), make([]int, 0, 4)

	var generate func(int)
	generate = func(i int) {
		m := len(combination)
		if i > n {
			if m != k {
				return	
			}
			tmp := make([]int, m)
			copy(tmp, combination)
			combinations = append(combinations, tmp)
			return
		}
		combination = append(combination, i)
		generate(i + 1)

		combination = combination[:len(combination)-1]
		generate(i + 1)
	}
	generate(1)

	return combinations
}
