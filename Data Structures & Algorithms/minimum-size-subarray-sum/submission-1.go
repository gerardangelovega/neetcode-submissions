func minSubArrayLen(target int, nums []int) int {
	l, r, n := 0, 0, len(nums)
	res, sum := math.MaxInt, 0
	_ = n

	for l <= r {
		if sum >= target {
			res = min(res, r-l) 
			sum = sum - nums[l]
			l++
		} else {
			if r >= n { break }
			sum = sum + nums[r]
			r++
		}
		fmt.Println(l,r, sum)
	}

	if res == math.MaxInt { return 0 }
	return res
}
