func twoSum(nums []int, target int) []int {
	arr := [2]int{}

	outer:
	for i := range len(nums) {
		for j := range len(nums) {
			if i == j {
				continue
			}	
			if nums[i] + nums[j] == target {
				arr[0], arr[1] = i, j
				break outer
			}
		}
	}
	
	return arr[:]
}