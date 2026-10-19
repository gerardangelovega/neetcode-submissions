import "slices"
func permute(nums []int) [][]int {
	n := len(nums)
	var permutate func(int) [][]int
	permutate = func(i int) [][]int {
		if i == n { return [][]int{{}} }
		permutations := permutate(i + 1)
		res := [][]int{}
		for _, permutation := range permutations {
			for j := range len(permutation) + 1 {
				res = append(res, slices.Insert(slices.Clone(permutation), j, nums[i]))
			}
		}
		return res
	}
	return permutate(0)
}
