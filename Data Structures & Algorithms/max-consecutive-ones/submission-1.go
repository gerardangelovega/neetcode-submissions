func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count++
			res = max(res, count)
		} else {
			count = 0
		}
	}	

	return res
}
