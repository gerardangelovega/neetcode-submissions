func findErrorNums(nums []int) []int {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	res := []int{}
	for k, v := range m {
		if len(v) > 1 { res = append(res, k) }
	}
	for i := range len(nums) {
		if _, e := m[i+1]; !e { res = append(res, i+1) }
	}
	return res
}


// if (v[1] + 1) != k {
// 	res = append(res, v[1]+1)
// } else {
// 	res = append(res, v[0]+1)
// }