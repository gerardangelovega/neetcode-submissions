// true is increasing, false is decreasing
func isMonotonic(nums []int) bool {
	trend := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] { continue }

		if trend == 0 && nums[i-1] < nums[i] {
			trend = 1
			continue
		}
		if trend == 0 && nums[i-1] > nums[i] {
			trend = -1
			continue
		}

		if trend == 1 && nums[i-1] > nums[i] { return false }
		if trend == -1 && nums[i-1] < nums[i] { return false }
	}
	return true
}

// 1 2 3 4 5 6
// 1 1 1 1 1 1

// 6 5 4 3 2 1
// 1 1 1 1 1 1

// 6 5 7 3 2 1
// 1 1 2 1 1 1