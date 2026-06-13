func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n - 1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
			continue
		} 
		if target < nums[mid] {
			right = mid - 1
			continue
		} 
	}

	return -1
}