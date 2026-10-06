import "slices"

func permute(nums []int) [][]int {
	n := len(nums)

	var permutate func(int) [][]int
	permutate = func(i int) [][]int {
		if i == n {
			return [][]int{{}}
		}
		permutations := permutate(i + 1)
		var built [][]int
		for _, permutation := range permutations {
			for j := range (len(permutation) + 1) {
				pcopy := slices.Clone(permutation)
				pcopy = slices.Insert(pcopy, j, nums[i])
				built = append(built, pcopy)
			}
		}
		return built
	}
	
	return permutate(0)
}

func factorial(x int) int {
	total := 1
	for i := range x {
		total = total * (i + 1)
	}
	return total
}