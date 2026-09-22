func subsets(nums []int) [][]int {
	subsets, subset := [][]int{}, []int{}
	dfs(nums, 0, &subsets, &subset)
	return subsets
}

func dfs(nums []int, i int, subsets *[][]int, subset *[]int) {
	if i > len(nums)-1 {
		temp := make([]int, len(*subset))
		copy(temp, *subset)
		*subsets = append(*subsets, temp)
		return
	}
	*subset = append(*subset, nums[i])
	dfs(nums, i+1, subsets, subset)
	*subset = (*subset)[:len(*subset)-1]
	dfs(nums, i+1, subsets, subset)
}
