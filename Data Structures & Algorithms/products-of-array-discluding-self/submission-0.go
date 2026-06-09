func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = 1
		for j := 0; j < len(nums); j++ {
			if j == i { 
				continue 
			}
			res[i] = res[i] * nums[j]
		}
	}

	return res
}
