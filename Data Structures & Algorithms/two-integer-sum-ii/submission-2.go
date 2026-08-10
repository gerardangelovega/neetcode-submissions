func twoSum(numbers []int, target int) []int {
	for i, n1 := range numbers {
		for j, n2 := range numbers {
			if i == j {
				continue
			}
			if n1 + n2 == target {
				return []int{i+1, j+1}
			}
		}
	}
	return []int{0,0}
}
