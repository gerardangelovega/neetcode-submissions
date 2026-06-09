func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	arr := [2]int{}

	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		diff := target - num
		if j, ok := m[diff]; ok && j != i {
			arr[0], arr[1] = i, j
			break
		}
	}

	return arr[:]
}
