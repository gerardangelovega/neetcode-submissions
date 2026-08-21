func maxSubarraySumCircular(nums []int) int {
	minimum, maximum, total, sum := math.MaxInt, math.MinInt, 0, 0

	for _, n := range nums {
		sum = max(sum, 0)
		sum = sum + n
		maximum = max(maximum, sum)
	}
	fmt.Println(maximum)

	for _, n := range nums {
		sum = min(sum, 0)
		sum = sum + n
		minimum = min(minimum, sum)
	}
	fmt.Println(minimum)

	for _, n := range nums {
		total = total + n
	}
	fmt.Println(total)

	if maximum < 0 { return maximum }
	return max(maximum, total - minimum)
}
