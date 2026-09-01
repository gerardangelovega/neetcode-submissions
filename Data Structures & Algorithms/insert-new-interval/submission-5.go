func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	n := newInterval
	for _, interval := range intervals {
		if n[0] != -1 && n[1] != -1 {
			if n[1] < interval[0] { 
				res = append(res, n)
				res = append(res, interval)
				n = []int{-1, -1}	
				continue
			}
			if n[0] > interval[1] {
				res = append(res, interval)
				continue
			}
			if n[0] <= interval[1] { 
				n = []int{ min(interval[0], n[0]), max(interval[1], n[1]) }
				continue
			}
		}
		res = append(res, interval)
	}
	if n[0] != -1 && n[1] != -1 { 
		res = append(res, n)
	}
	return res
}