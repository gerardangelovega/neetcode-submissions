func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	prefix := map[int]int{ 0: 1 }

	for _, num := range nums {
		sum = sum + num
		diff := sum - k
		res = res + prefix[diff]
		prefix[sum]++
	}

	return res
}
