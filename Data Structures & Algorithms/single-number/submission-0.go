func singleNumber(nums []int) int {
	res := nums[0]
	for _, n := range nums[1:] {
		res = res ^ n
	}
	return res
}

// 11
// 10
// 11

// 11
// 01
// 10

