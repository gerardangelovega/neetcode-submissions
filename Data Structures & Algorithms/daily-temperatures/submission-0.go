func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i, t1 := range temperatures {
		temp := 0
		for _, t2 := range temperatures[i+1:] {
			temp++
			if t2 > t1 {
				res[i] = temp
				break;
			}
		}
	}

	return res
}