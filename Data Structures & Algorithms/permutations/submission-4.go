import "slices"

func permute(nums []int) [][]int {
	n := len(nums)

	var permutate func(int) [][]int
	permutate = func(i int) [][]int {
		if i == n {
			return [][]int{{}}
		}
		permutations := permutate(i + 1)
		res := [][]int{}
		for _, permutation := range permutations {
			for j := range len(permutation) + 1 {
				new_permutation := slices.Clone(permutation)
				new_permutation = slices.Insert(new_permutation, j, nums[i])
				res = append(res, new_permutation)
			}
		}
		return res
	}

	return permutate(0)
}

// {}
// {3}
// {2, 3} {3, 2}
// {1, 2, 3} {2, 1, 3} {2, 3, 1} {1, 3, 2} {3, 1, 2}, {3, 2, 1}
