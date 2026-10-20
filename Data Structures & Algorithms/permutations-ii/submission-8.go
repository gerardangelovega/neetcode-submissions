import "slices"
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	slices.Sort(nums)

	var permutate func(int)
	permutate = func(i int) {
		if i == n {
			res = append(res, slices.Clone(nums))
			return;
		}

		// swap
		for j := i; j < n; j++ {
			if j > i && nums[j] == nums[i] {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			permutate(i + 1)
		}

		// restore
		for j := n-1; j > i; j-- {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	permutate(0)
	return res
}
