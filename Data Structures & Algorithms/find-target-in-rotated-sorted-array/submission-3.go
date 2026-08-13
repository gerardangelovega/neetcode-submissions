func search(nums []int, target int) int {
	n := len(nums)
	l, r, m := 0, n-1, (n-1)/2

	for l <= r {
		m = (l + r) / 2
		if target == nums[m] {
			return m
		}

		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1	
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
// 1,2,3,4,5,6
// 5,6,1,2,3,4
// 3,4,5,6,1,2