func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 1 {
		return []int{nums[0]}
	}

	res := []int{}
	max_val, max_idx := -1, -1

	l, r := 0, k-1
	for r < n {
		if max_idx < l {
			tmp_val, tmp_idx := nums[l], l
			for i := l; i <= r; i++ {
				if nums[i] > tmp_val {
					tmp_val, tmp_idx = nums[i], i
				}
			}
			max_val, max_idx = tmp_val, tmp_idx
		} else if nums[r] > max_val {
			max_val, max_idx = nums[r], r
		}
		res = append(res, max_val)
		l++
		r++
	}
	
	return res
}

// max will contain both the value and index
// for every sliding of the window, we check if max is still within the window
// if not, then we recalculate the max of the current window and update

// checking can be done through l <= max_index <= r

// 1 2 1 0 4 2 6 k = 3

