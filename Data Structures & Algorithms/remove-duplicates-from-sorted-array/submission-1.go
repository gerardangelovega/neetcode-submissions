func removeDuplicates(nums []int) int {
	l, r, n := 0, 0, len(nums)
	for r < n {
		nums[l]	= nums[r] // replace the value at l with the found unique number
		for r < n && nums[r] == nums[l] { r++ } // loops until a unique number is found
		l++ // moves l forward
	}
	return l
}
