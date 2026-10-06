import "slices"

func permute(nums []int) [][]int {
	n := len(nums)

	var permutate func(int) [][]int
	permutate = func(i int) [][]int {
		if i == n {
			return [][]int{{}}
		}
		permutations := permutate(i + 1)
		var res [][]int
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