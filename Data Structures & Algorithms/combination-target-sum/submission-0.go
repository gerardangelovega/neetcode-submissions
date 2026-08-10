func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(int, int)
	dfs = func(i, t int) {
		if t == 0 {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}
		if i >= len(nums) || t < 0 {
			return
		}
		subset = append(subset, nums[i])
		dfs(i, t-nums[i])
		subset = subset[:len(subset)-1]	
		dfs(i+1, t)
	}

	dfs(0, target)

	return res
}
