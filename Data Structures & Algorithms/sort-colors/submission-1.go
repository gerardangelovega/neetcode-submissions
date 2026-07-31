func sortColors(nums []int) {
	colors := [3]int{}

	for _, num := range nums {
		colors[num]++
	}

	i := 0
	for c, color := range colors {
		for range color {
			nums[i] = c
			i++
		}
	}
}