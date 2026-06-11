func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			count = 0
		}
		res = max(res, count)
	}	

	return res
}
