import "slices"

func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, factorial(n))

	var permutate func(int) [][]int
	permutate = func(i int) [][]int {
		if i == n {
			return [][]int{{}}
		}
		permutations := slices.Clone(permutate(i + 1))
		res = res[:0]
		for _, permutation := range permutations {
			for j := range (len(permutation) + 1) {
				pcopy := slices.Clone(permutation)
				pcopy = slices.Insert(pcopy, j, nums[i])
				res = append(res, pcopy)
			}
		}
		return res
	}
	return permutate(0)
}

func factorial(x int) int {
	total := 1
	for i := range total {
		total = total * (i + 1)
	}
	return total
}