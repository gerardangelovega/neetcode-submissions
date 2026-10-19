import "slices"
func permuteUnique(nums []int) [][]int {
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

	permutations := permutate(0)
	slices.SortFunc(permutations, func(a, b []int) int {
		return slices.Compare(a, b)
	})

	pn := len(permutations)
	for i := 0; i < pn; {
		if i + 1 >= pn {
			break
		}
		if slices.Compare(permutations[i], permutations[i + 1]) == 0 {
			permutations = slices.Delete(permutations, i, i + 1)
			pn = len(permutations)
			continue
		}
		i++
	}

	return permutations
}
