func rotate(nums []int, k int) {
	n := len(nums)
	for range k {
		tmp := nums[n-1]
		nums[n-1] = 0
		for i := n-1; i-1 >= 0; i-- {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
		nums[0] = tmp
	}
}
