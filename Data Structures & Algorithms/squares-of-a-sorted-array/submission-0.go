func sortedSquares(nums []int) []int {
	for i, _ := range nums {
		nums[i] = nums[i] * nums[i]
	}	
	return quicksort(nums)
}

func quicksort(nums []int) []int {
	var sort func(int, int)
	sort = func(start, end int) {
		if start >= end { return }

		i, k := start, start
		for i < end {
			if nums[i] < nums[end] {
				nums[i], nums[k] = nums[k], nums[i]
				k++
			}
			i++
		}
		nums[k], nums[end] = nums[end], nums[k]

		sort(start, k-1)
		sort(k+1, end)
	}
	sort(0, len(nums)-1)
	return nums
}