func canJump(nums []int) bool {
	return jump(nums, 0, make(map[int]bool))
}

func jump(nums []int, i int, cache map[int]bool) bool {
	if i >= len(nums)-1 {
		return true
	}
	if nums[i] == 0 {
		return false
	}
	if v, e := cache[i]; e {
		return v
	}
	res := false
	for j := nums[i]; j > 0; j-- {
		cache[i] = jump(nums, i + j, cache)
		res = res || cache[i]
	}
	return res
}