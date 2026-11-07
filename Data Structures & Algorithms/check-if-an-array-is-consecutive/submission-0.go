func isConsecutive(nums []int) bool {
	least, most := math.MaxInt, math.MinInt
	for _, num := range nums { least, most = min(least, num), max(most, num) }
	return most - least == len(nums) - 1
}
