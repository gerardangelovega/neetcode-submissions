func removeElement(nums []int, val int) int {
	n := len(nums)
	swap_index := n - 1

	i := 0
	for i < n {
		if i > swap_index { 
			break
		}
		if nums[i] == val {
			nums[i], nums[swap_index] = nums[swap_index], nums[i]
			swap_index--
			continue
		}
		i++
	}

	fmt.Println(swap_index + 1)

	return swap_index + 1
}
