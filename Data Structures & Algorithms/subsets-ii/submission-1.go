func subsetsWithDup(nums []int) [][]int {
	subsets, subset := [][]int{}, []int{}
	dfs(quickSort(nums), 0, &subsets, &subset)
	return subsets
}

func dfs(nums []int, i int, subsets *[][]int, subset *[]int) {
	if i > len(nums)-1 {
		temp := make([]int, len(*subset))
		copy(temp, (*subset))
		*subsets = append(*subsets, temp)
		return
	}
	*subset = append(*subset, nums[i])
	dfs(nums, i+1, subsets, subset)
	*subset = (*subset)[:len(*subset)-1]
	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}
	dfs(nums, i+1, subsets, subset)
}

func quickSort(arr []int) []int {
	var qSort func([]int, int, int) []int
	qSort = func(arr []int, s, e int) []int {
		if (e - s + 1) < 2 { return arr }	
		left := s
		for i := s; i < e; i++ {
			if arr[i] < arr[e] {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}
		arr[e], arr[left] = arr[left], arr[e]
		qSort(arr, s, left-1)
		qSort(arr, left+1, e)
		return arr
	}
	return qSort(arr, 0, len(arr)-1)
}