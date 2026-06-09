import (
	"slices"
)
func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	triplets := make([][]int, 0)

	for i := 0; i < n - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			}
			if sum > 0 {
				k--
				continue
			}
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}

	return triplets
}