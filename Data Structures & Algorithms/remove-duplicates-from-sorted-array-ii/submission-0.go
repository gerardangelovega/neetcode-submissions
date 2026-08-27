func removeDuplicates(nums []int) int {
    l, r, n := 2, 2, len(nums)
	for r < n {
		if nums[l-2] != nums[r] {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	fmt.Println(nums)
	return l
}