import (
	"slices"
)
func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	triplets := make([][]int, 0)
	d := make(map[string]bool)

	for i := 0; i < n - 2; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] == -(nums[j] + nums[k]) {
					str := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[k])
	
					if d[str] {
						continue
					}
	
					d[str] = true
					triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return triplets
}