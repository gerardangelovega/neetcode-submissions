func maxSubArray(nums []int) int {
	res, sum := math.MinInt, 0
	for _, n := range nums {
		sum = max(sum, 0)
		sum = sum + n
		res = max(res, sum)
	}
	return res
}