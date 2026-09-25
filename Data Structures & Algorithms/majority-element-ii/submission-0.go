func majorityElement(nums []int) []int {
	target := len(nums) / 3
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	res := []int{}
	for k, v := range frequency {
		if v > target {
			res = append(res, k)
		}
	}
	return res
}
