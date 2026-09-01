func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	n := newInterval
	inserted := false
	for _, interval := range intervals {
		if n[1] < interval[0] && !inserted { 
			res = append(res, n)
			res = append(res, interval)
			inserted = true
			continue
		}
		if n[0] > interval[1] && !inserted {
			res = append(res, interval)
			continue
		}
		if n[0] <= interval[1] && !inserted { 
			n = []int{ min(interval[0], n[0]), max(interval[1], n[1]) }
			continue
		}
		res = append(res, interval)
	}
	if !inserted {
		res = append(res, n)
	}
	return res
}

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
// A  A  A
//    B  B  B  B

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
//                A  A
//                         A   A