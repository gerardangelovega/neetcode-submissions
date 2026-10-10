func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, 0, n)

	mid := [2]int{ math.MaxInt, 0 }
	for i, num := range nums {
		if abs(num) < abs(mid[0]) {
			mid[0] = num
			mid[1] = i
		}
	}

	l, r := mid[1], mid[1]
	for l >= 0 && r < n {
		ll := nums[l] * nums[l]
		rr := nums[r] * nums[r]

		if l == r {
			res = append(res, ll)
			l--
			r++
		} else if ll < rr {
			res = append(res, ll)
			l--
		} else {
			res = append(res, rr)
			r++
		}
	}
	for l >= 0 {
		ll := nums[l] * nums[l]
		res = append(res, ll)
		l--
	}
	for r < n {
		rr := nums[r] * nums[r]
		res = append(res, rr)
		r++
	}
	return res
}

func abs(x int) int {
	if x < 0 { x = x * -1 }
	return x
}
