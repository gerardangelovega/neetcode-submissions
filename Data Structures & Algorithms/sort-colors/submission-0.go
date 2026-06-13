func sortColors(nums []int) {
	colors := make([]int, 3)

	for _, num := range nums {
		colors[num]++
	}

	i := 0
	for c, color := range colors {
		for j := 0; j < color; j++ {
			nums[i] = c
			i++
		}
	}
}
