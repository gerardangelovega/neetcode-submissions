func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	res := math.MaxInt

	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		}

		mid := left + (right - left) / 2

		if nums[mid] < res {
			res = nums[mid]
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}