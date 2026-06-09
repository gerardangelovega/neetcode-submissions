func missingNumber(nums []int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	} 
	
	summation := ((len(nums)) * (len(nums) + 1)) / 2
	return summation - total
}
