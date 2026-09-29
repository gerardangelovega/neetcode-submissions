func subsetXORSum(nums []int) int {
	sum := 0
	for _, subset := range subsets(nums) {
		tmp := 0	
		for _, num := range subset {
			tmp = tmp ^ num
		}	
		sum = sum + tmp
	}
	return sum
}

func subsets(nums []int) [][]int {
	subsets, subset := make([][]int, 0), make([]int, 0)
	n := len(nums)

	var generate func(int) 
	generate = func(i int) {
		if i > n-1 {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			subsets = append(subsets, tmp)
			return
		}
		subset = append(subset, nums[i])
		generate(i+1)

		subset = subset[:len(subset)-1]
		generate(i+1)
	}
	generate(0)

	return subsets
}