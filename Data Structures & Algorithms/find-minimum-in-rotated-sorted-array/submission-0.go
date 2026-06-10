func findMin(nums []int) int {
	min_num := nums[0]
	for _, num := range nums {
		min_num = min(min_num, num)	
	}
	return min_num
}
