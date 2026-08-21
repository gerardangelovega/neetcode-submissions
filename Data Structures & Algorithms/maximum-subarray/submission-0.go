func maxSubArray(nums []int) int {
	res := math.MinInt
	sum := 0
	
	for _, n := range nums {
		if sum < 0 {
			sum = 0
		}	
		sum = sum + n
		res = max(res, sum)
	}

	return res
}