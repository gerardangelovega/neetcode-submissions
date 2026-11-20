func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums { m[n]++ }
	key, value := math.MinInt, math.MinInt
	for k, v := range m { if v > value { key, value = k, v }}
	return key
}
